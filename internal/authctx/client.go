/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package authctx

import (
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
)

const (
	ServerEndpointEnvVar   = "TMC_ENDPOINT"
	VMWCloudEndpointEnvVar = "VMW_CLOUD_ENDPOINT"
	VMWCloudAPITokenEnvVar = "VMW_CLOUD_API_TOKEN"
)

type TanzuContext struct {
	ServerEndpoint   string
	Token            string
	VMWCloudEndPoint string
	TMCConnection    *client.TanzuMissionControl
	TLSConfig        *helper.TLSConfig
}

func (cfg *TanzuContext) Setup() (err error) {
	cfg.TMCConnection, err = client.NewHTTPClient(cfg.TLSConfig)
	if err != nil {
		return
	}

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
