/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package client

import (
	"net/http"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/transport"

	clusterclient "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/cluster"
)

// NewHTTPClient creates a new  tanzu mission control HTTP client.
func NewHTTPClient() *TanzuMissionControl {
	httpClient := transport.NewClient()

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Connection", "keep-alive")

	cfg := transport.DefaultTransportConfig().AddHeaders(headers)

	return &TanzuMissionControl{
		Config:                 cfg,
		Transport:              httpClient,
		ClusterResourceService: clusterclient.New(httpClient, cfg),
	}
}

// TanzuMissionControl is a client for  tanzu mission control.
type TanzuMissionControl struct {
	*transport.Config
	Transport              *transport.Client
	ClusterResourceService clusterclient.ClientService
}
