/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package client

import (
	"net/http"
	"runtime"

	clusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster"
	clustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup"
	namespaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/namespace"
	nodepoolclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/nodepool"
	policyclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/policy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	workspaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/workspace"
)

// NewHTTPClient creates a new  tanzu mission control HTTP client.
func NewHTTPClient() *TanzuMissionControl {
	httpClient := transport.NewClient()

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Connection", "keep-alive")
	headers.Set("x-client-name", "tmc-terraform-provider")
	headers.Set("x-client-platform", runtime.GOOS)
	headers.Set("x-client-version", "1.0.2")

	httpClient.AddHeaders(headers)

	return &TanzuMissionControl{
		Client:                      httpClient,
		ClusterResourceService:      clusterclient.New(httpClient),
		WorkspaceResourceService:    workspaceclient.New(httpClient),
		NamespaceResourceService:    namespaceclient.New(httpClient),
		ClusterGroupResourceService: clustergroupclient.New(httpClient),
		NodePoolResourceService:     nodepoolclient.New(httpClient),
		PolicyResourceService:       policyclient.New(httpClient),
	}
}

// TanzuMissionControl is a client for  tanzu mission control.
type TanzuMissionControl struct {
	*transport.Client
	ClusterResourceService      clusterclient.ClientService
	WorkspaceResourceService    workspaceclient.ClientService
	NamespaceResourceService    namespaceclient.ClientService
	ClusterGroupResourceService clustergroupclient.ClientService
	NodePoolResourceService     nodepoolclient.ClientService
	PolicyResourceService       policyclient.ClientService
}
