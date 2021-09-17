/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package client

import (
	"net/http"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/transport"

	clusterclient "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/cluster"
	clustergroupclient "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/clustergroup"
	namespaceclient "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/namespace"
	workspaceclient "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/workspace"
)

// NewHTTPClient creates a new  tanzu mission control HTTP client.
func NewHTTPClient() *TanzuMissionControl {
	httpClient := transport.NewClient()

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Connection", "keep-alive")

	httpClient.AddHeaders(headers)

	return &TanzuMissionControl{
		Client:                      httpClient,
		ClusterResourceService:      clusterclient.New(httpClient),
		WorkspaceResourceService:    workspaceclient.New(httpClient),
		NamespaceResourceService:    namespaceclient.New(httpClient),
		ClusterGroupResourceService: clustergroupclient.New(httpClient),
	}
}

// TanzuMissionControl is a client for  tanzu mission control.
type TanzuMissionControl struct {
	*transport.Client
	ClusterResourceService      clusterclient.ClientService
	WorkspaceResourceService    workspaceclient.ClientService
	NamespaceResourceService    namespaceclient.ClientService
	ClusterGroupResourceService clustergroupclient.ClientService
}
