// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package authctx

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	oidcapi "go.pinniped.dev/generated/latest/apis/supervisor/oidc"
	"go.pinniped.dev/pkg/oidcclient/pkce"
	"go.pinniped.dev/pkg/oidcclient/state"
	"golang.org/x/oauth2"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
)

const (
	tokenEndpointSuffix         = "oauth2/token"
	authorizationEndpointSuffix = "oauth2/authorize"
	redirectURL                 = "http://127.0.0.1/callback"
	pinnipedCLIClientID         = oidcapi.ClientIDPinnipedCLI
	contextTimeout              = 60 * time.Second
	extraIDToken                = "id_token"
	federationDomainPath        = "provider/pinniped"

	mdKeyAuthToken    = "Authorization"
	authTokenPrefix   = "Bearer "
	mdKeyAuthIDToken  = "grpc-metadata-x-user-id"
	mdKeyRefreshToken = "grpc-metadata-x-refresh-token"
)

type smSession struct {
	sharedOauthConfig             *oauth2.Config
	tlsConfig                     *tls.Config
	issuerURL, username, password string
	pkceCodePair                  pkce.Code
	stateVal                      state.State
}

// todo: proxy support is not added for the self-managed flow. Add it when there is a requirement.
func getSMUserAuthCtx(pinnipedURL, uName, password string, config *proxy.TLSConfig) (metadata map[string]string, err error) {
	if pinnipedURL == "" || uName == "" || password == "" {
		return nil, errors.New("Invalid auth configuration for self_managed")
	}

	tlsConfig, err := proxy.GetConnectorTLSConfig(config)
	if err != nil {
		return nil, err
	}

	session, err := initSession(pinnipedURL, uName, password, tlsConfig)
	if err != nil {
		return nil, err
	}

	expectedRedirectURL, err := url.Parse(session.sharedOauthConfig.RedirectURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse expected redirect URL %s", session.sharedOauthConfig.RedirectURL)
	}

	actualRedirectURL, err := session.initiateAuthorizeRequestUnamePwd()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to initiate authorize request with issuer %s", session.issuerURL)
	}

	// Check that the redirect was to the expected location.
	if actualRedirectURL.Scheme != expectedRedirectURL.Scheme ||
		actualRedirectURL.Host != expectedRedirectURL.Host || actualRedirectURL.Path != expectedRedirectURL.Path {
		return nil, fmt.Errorf("error getting authorization: redirected to the wrong location: %s",
			actualRedirectURL.String())
	}

	// validate the state param to detect and prevent CSRF attacks.
	if err := session.stateVal.Validate(actualRedirectURL.Query().Get("state")); err != nil {
		return nil, errors.Wrap(err, "failed to validate state")
	}

	// Get the auth code or return the error from the server.
	authCode := actualRedirectURL.Query().Get("code")
	if authCode == "" {
		// Check for error response parameters. See https://openid.net/specs/openid-connect-core-1_0.html#AuthError.
		requiredErrorCode := actualRedirectURL.Query().Get("error")

		optionalErrorDescription := actualRedirectURL.Query().Get("error_description")
		if optionalErrorDescription == "" {
			return nil, fmt.Errorf("login failed with code %q", requiredErrorCode)
		}

		return nil, fmt.Errorf("login failed with code %q: %s", requiredErrorCode, optionalErrorDescription)
	}

	customClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
			Proxy:           http.ProxyFromEnvironment,
		},
	}
	// Exchange the authorization code for access, ID, and refresh tokens and perform required
	// validations on the returned ID token.
	ctxWithValue := context.WithValue(context.Background(), oauth2.HTTPClient, customClient)

	tokenCtx, tokenCtxCancelFunc := context.WithTimeout(ctxWithValue, contextTimeout)
	defer tokenCtxCancelFunc()

	token, err := session.sharedOauthConfig.Exchange(tokenCtx, authCode, session.pkceCodePair.Verifier())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to exchange auth code for oauth tokens")
	}

	extraFields := map[string]interface{}{extraIDToken: token.Extra(extraIDToken).(string)}
	token = &oauth2.Token{
		AccessToken:  token.AccessToken,
		Expiry:       token.Expiry,
		RefreshToken: token.RefreshToken,
	}

	token = token.WithExtra(extraFields)

	return getSMHeaders(token), nil
}

// todo: if slowness is experienced, then we can avoid re-initialising same values again.
func initSession(pinnipedURL, uName, password string, config *tls.Config) (*smSession, error) {
	// TMC Local Pinniped sample endpoint:
	// https://pinniped-supervisor.*******.com/provider/pinniped
	u := url.URL{
		Scheme: "https",
		Host:   pinnipedURL,
		Path:   federationDomainPath,
	}

	issuerURL := u.String()

	sharedOauthConfig := &oauth2.Config{
		RedirectURL:  redirectURL,
		ClientID:     pinnipedCLIClientID,
		ClientSecret: "",
		Scopes:       []string{"openid", "offline_access", "username", "groups"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("%s/%s", issuerURL, authorizationEndpointSuffix),
			TokenURL: fmt.Sprintf("%s/%s", issuerURL, tokenEndpointSuffix),
		},
	}

	var err error

	session := &smSession{
		sharedOauthConfig: sharedOauthConfig,
		tlsConfig:         config,
		issuerURL:         issuerURL,
		username:          uName,
		password:          password,
	}

	if session.pkceCodePair, err = pkce.Generate(); err != nil {
		return nil, errors.Wrapf(err, "failed to generate pkce code pair generator")
	}

	if session.stateVal, err = state.Generate(); err != nil {
		return nil, errors.Wrap(err, "failed to generate state parameter")
	}

	return session, nil
}

func getSMHeaders(token *oauth2.Token) map[string]string {
	headers := map[string]string{mdKeyAuthToken: authTokenPrefix + " " + token.AccessToken}
	headers[mdKeyRefreshToken] = token.RefreshToken

	if idTok := getIDTokenFromTokenSource(*token); idTok != "" {
		headers[mdKeyAuthIDToken] = idTok
	}

	return headers
}

func getIDTokenFromTokenSource(token oauth2.Token) string {
	idTok := ""

	extraTok := token.Extra(extraIDToken)
	if extraTok != nil {
		idTok = extraTok.(string)
	}

	return idTok
}

func (s *smSession) initiateAuthorizeRequestUnamePwd() (*url.URL, error) {
	authCodeURL := s.getAuthCodeURL()

	// Send an authorize request.
	authCtx, authorizeCtxCancelFunc := context.WithTimeout(context.Background(), contextTimeout)
	defer authorizeCtxCancelFunc()

	authReq, err := http.NewRequestWithContext(authCtx, http.MethodGet, authCodeURL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not build authorize request: %w", err)
	}

	authReq.Header.Set(oidcapi.AuthorizeUsernameHeaderName, s.username)
	authReq.Header.Set(oidcapi.AuthorizePasswordHeaderName, s.password)

	redirected := false
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: s.tlsConfig,
			Proxy:           http.ProxyFromEnvironment,
		},
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			redirected = true
			return http.ErrUseLastResponse
		},
	}

	authRes, err := httpClient.Do(authReq)
	if err != nil {
		return nil, fmt.Errorf("authorization response error: %w", err)
	}

	_ = authRes.Body.Close()

	if !redirected {
		return nil, fmt.Errorf("error getting authorization: expected to be redirected, but response status was %s", authRes.Status)
	}

	rawLocation := authRes.Header.Get("Location")

	location, err := url.Parse(rawLocation)
	if err != nil {
		// This shouldn't be possible in practice because httpClient.Do() already parses the Location header.
		return nil, fmt.Errorf("error getting authorization: could not parse redirect location %s: %w", rawLocation, err)
	}

	return location, nil
}

func (s *smSession) getAuthCodeURL() string {
	opts := []oauth2.AuthCodeOption{
		s.pkceCodePair.Challenge(),
		s.pkceCodePair.Method(),
	}

	return s.sharedOauthConfig.AuthCodeURL(s.stateVal.String(), opts...)
}

func refreshSMUserAuthCtx(config *TanzuContext) {
	md, _ := getSMUserAuthCtx(config.VMWCloudEndPoint, config.SMUsername, config.Token, config.TLSConfig)
	for key, value := range md {
		config.TMCConnection.Headers.Set(key, value)
	}
}
