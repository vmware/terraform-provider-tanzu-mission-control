/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package authctx

import (
	"github.com/pkg/errors"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client"
)

const (
	ServerEndpointEnvVar = "TMC_ENDPOINT"
	CSPTokenEnvVar       = "TMC_CSP_TOKEN"
)

type TanzuContext struct {
	ServerEndpoint string
	Token          string
	CSPEndPoint    string
	TMCConnection  *client.TanzuMissionControl
}

func (cfg *TanzuContext) Setup() error {
	cfg.TMCConnection = client.NewHTTPClient()

	md, err := getUserAuthCtx(cfg)
	if err != nil {
		return errors.Wrap(err, "while getting user ctx")
	}

	cfg.TMCConnection.WithHost(cfg.ServerEndpoint)
	cfg.TMCConnection.Headers.Set("Host", cfg.ServerEndpoint)

	for key, value := range md {
		cfg.TMCConnection.Headers.Set(key, value)
	}

	return nil
}
