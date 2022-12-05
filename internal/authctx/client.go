/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package authctx

import (
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
)

const (
	ServerEndpointEnvVar             = "TMC_ENDPOINT"
	VMWCloudEndpointEnvVar           = "VMW_CLOUD_ENDPOINT"
	VMWCloudAPITokenEnvVar           = "VMW_CLOUD_API_TOKEN"
	InsecureAllowUnverifiedSSLEnvVar = "INSECURE_ALLOW_UNVERIFIED_SSL"
	ClientAuthCertFileEnvVar         = "CLIENT_AUTH_CERT_FILE"
	ClientAuthKeyFileEnvVar          = "CLIENT_AUTH_KEY_FILE"
	CAFileEnvVar                     = "CA_FILE"
	ClientAuthCertEnvVar             = "CLIENT_AUTH_CERT"
	ClientAuthKeyEnvVar              = "CLIENT_AUTH_KEY"
	CACertEnvVar                     = "CA_CERT"
)

type TanzuContext struct {
	ServerEndpoint   string
	Token            string
	VMWCloudEndPoint string
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

func (cfg *TanzuContext) SetupWithDefaultTransport() (err error) {
	cfg.TMCConnection = client.NewHTTPClientWithDefaultTransport()
	return setup(cfg)
}

func setup(cfg *TanzuContext) (err error) {
	md, err := getUserAuthCtx(cfg)
	if err != nil {
		return errors.Wrap(err, "unable to get user context")
	}

	cfg.TMCConnection.WithHost(cfg.ServerEndpoint)
	cfg.TMCConnection.Headers.Set("Host", cfg.ServerEndpoint)

	for key, value := range md {
		cfg.TMCConnection.Headers.Set(key, value)
	}

	return nil
}
