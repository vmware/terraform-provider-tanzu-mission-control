// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package authctx

import (
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
)

const (
	ServerEndpointEnvVar = "TMC_ENDPOINT"
	ProjectIDEnvVar      = "PROJECT_ID"

	// TMC SaaS env variables.
	VMWCloudEndpointEnvVar = "VMW_CLOUD_ENDPOINT"
	VMWCloudAPITokenEnvVar = "VMW_CLOUD_API_TOKEN"

	// TMC self managed env variables.
	OIDCIssuerEndpointEnvVar = "OIDC_ISSUER"
	TmcSMUsernameEnvVar      = "TMC_SM_USERNAME"
	TmcSMPasswordEnvVar      = "TMC_SM_PASSWORD"

	// Proxy config values.
	InsecureAllowUnverifiedSSLEnvVar = "INSECURE_ALLOW_UNVERIFIED_SSL"
	ClientAuthCertFileEnvVar         = "CLIENT_AUTH_CERT_FILE"
	ClientAuthKeyFileEnvVar          = "CLIENT_AUTH_KEY_FILE"
	CAFileEnvVar                     = "CA_FILE"
	ClientAuthCertEnvVar             = "CLIENT_AUTH_CERT"
	ClientAuthKeyEnvVar              = "CLIENT_AUTH_KEY"
	CACertEnvVar                     = "CA_CERT"
)

type TanzuContext struct {
	SelfManaged      bool
	ServerEndpoint   string
	ProjectID        string
	SMUsername       string
	Token            string // selfmanaged password is stored here
	VMWCloudEndPoint string // selfmanaged odic issuer is stored here
	TMCConnection    *client.TanzuMissionControl
	TLSConfig        *proxy.TLSConfig

	// Cache for self-managed OIDC tokens
	smTokenCache *TokenCache
}

// TokenCache manages the lifecycle of the self-managed OIDC tokens.
type TokenCache struct {
	mu            sync.Mutex
	cachedHeaders map[string]string
	expiry        time.Time
}

// NewTokenCache creates a new TokenCache instance.
func NewTokenCache() *TokenCache {
	return &TokenCache{}
}

// GetToken returns the cached headers if they are still valid (with a 1-minute buffer),
// or fetches new headers using the provided fetch function.
func (tc *TokenCache) GetToken(fetch func() (map[string]string, time.Time, error)) (map[string]string, error) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	if tc.cachedHeaders != nil && time.Now().Add(1*time.Minute).Before(tc.expiry) {
		return tc.cachedHeaders, nil
	}

	headers, expiry, err := fetch()
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh token")
	}

	tc.cachedHeaders = headers
	tc.expiry = expiry

	return headers, nil
}

// Update explicitly updates the cached headers and expiry.
func (tc *TokenCache) Update(headers map[string]string, expiry time.Time) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	tc.cachedHeaders = headers
	tc.expiry = expiry
}

func (cfg *TanzuContext) Setup() (err error) {
	cfg.TMCConnection, err = client.NewHTTPClient(cfg.TLSConfig)
	if err != nil {
		return err
	}

	return setup(cfg)
}

func (cfg *TanzuContext) IsSelfManaged() bool {
	return cfg.SelfManaged
}

// The default transport is needed for mocking. The http mocking library used in testing
// can only intercept calls if they're made with the default transport.
func (cfg *TanzuContext) SetupWithDefaultTransportForTesting() (err error) {
	cfg.TMCConnection = client.NewTestHTTPClientWithDefaultTransport()
	return setup(cfg)
}

func setup(cfg *TanzuContext) (err error) {
	if cfg.IsSelfManaged() {
		cfg.smTokenCache = NewTokenCache()
	}

	fetchAuthHeaders := getUserAuthCtxHeaders(cfg)

	md, err := fetchAuthHeaders()
	if err != nil {
		return errors.Wrap(err, "unable to get user context")
	}

	cfg.TMCConnection.WithHost(cfg.ServerEndpoint)
	cfg.TMCConnection.Headers.Set("Host", cfg.ServerEndpoint)

	if cfg.ProjectID != "" {
		cfg.TMCConnection.Headers.Set("X-Project-Id", cfg.ProjectID)
	}

	if cfg.IsSelfManaged() {
		// We need to add this only for self-managed flow because the SaaS token has a longer ttl.
		cfg.TMCConnection.WithRefreshAuthCtx(fetchAuthHeaders)
	}

	for key, value := range md {
		cfg.TMCConnection.Headers.Set(key, value)
	}

	return nil
}

func getUserAuthCtxHeaders(config *TanzuContext) func() (map[string]string, error) {
	if config.IsSelfManaged() {
		return func() (map[string]string, error) {
			// For compatibility considerations.
			if config.smTokenCache == nil {
				headers, _, err := getSMUserAuthCtx(config.VMWCloudEndPoint, config.SMUsername, config.Token, config.TLSConfig)
				if err != nil {
					return nil, errors.Wrap(err, "failed to get self-managed user auth context headers")
				}

				return headers, nil
			}

			return config.smTokenCache.GetToken(func() (map[string]string, time.Time, error) {
				return getSMUserAuthCtx(config.VMWCloudEndPoint, config.SMUsername, config.Token, config.TLSConfig)
			})
		}
	}

	issuerURL := config.VMWCloudEndPoint
	token := config.Token
	proxyConfig := config.TLSConfig

	return func() (map[string]string, error) {
		return getSaaSUserAuthCtx(issuerURL, token, proxyConfig)
	}
}
