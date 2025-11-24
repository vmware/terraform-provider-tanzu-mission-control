// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package authctx

import (
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
	issuerURL := config.VMWCloudEndPoint
	token := config.Token
	proxyConfig := config.TLSConfig

	if config.IsSelfManaged() {
		username := config.SMUsername

		return func() (map[string]string, error) {
			return getSMUserAuthCtx(issuerURL, username, token, proxyConfig)
		}
	}

	return func() (map[string]string, error) {
		return getSaaSUserAuthCtx(issuerURL, token, proxyConfig)
	}
}
