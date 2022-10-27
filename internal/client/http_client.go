/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package client

import (
	"net/http"
	"runtime"

	clusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster"
	iamclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/iam_policy"
	policyclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/policy"
	clustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup"
	iamclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/iam_policy"
	policyclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/policy"
	namespaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/namespace"
	iamnamespaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/namespace/iam_policy"
	nodepoolclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/nodepool"
	iamorganizationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/organization/iam_policy"
	policyorganizationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/organization/policy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	workspaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/workspace"
	iamworkspaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/workspace/iam_policy"
	policyworkspaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/workspace/policy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
)

// NewHTTPClient creates a new tanzu mission control HTTP client.
func NewHTTPClient(config *helper.TLSConfig) (*TanzuMissionControl, error) {
	httpClient, err := transport.NewClient(config)
	if err != nil {
		return nil, err
	}

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Connection", "keep-alive")
	headers.Set("x-client-name", "tmc-terraform-provider")
	headers.Set("x-client-platform", runtime.GOOS)
	headers.Set("x-client-version", "1.0.2")

	httpClient.AddHeaders(headers)

	return &TanzuMissionControl{
		Client:                            httpClient,
		ClusterResourceService:            clusterclient.New(httpClient),
		WorkspaceResourceService:          workspaceclient.New(httpClient),
		NamespaceResourceService:          namespaceclient.New(httpClient),
		ClusterGroupResourceService:       clustergroupclient.New(httpClient),
		NodePoolResourceService:           nodepoolclient.New(httpClient),
		OrganizationIAMResourceService:    iamorganizationclient.New(httpClient),
		ClusterGroupIAMResourceService:    iamclustergroupclient.New(httpClient),
		ClusterIAMResourceService:         iamclusterclient.New(httpClient),
		WorkspaceIAMResourceService:       iamworkspaceclient.New(httpClient),
		NamespaceIAMResourceService:       iamnamespaceclient.New(httpClient),
		ClusterPolicyResourceService:      policyclusterclient.New(httpClient),
		ClusterGroupPolicyResourceService: policyclustergroupclient.New(httpClient),
		WorkspacePolicyResourceService:    policyworkspaceclient.New(httpClient),
		OrganizationPolicyResourceService: policyorganizationclient.New(httpClient),
	}, nil
}

// TanzuMissionControl is a client for tanzu mission control.
type TanzuMissionControl struct {
	*transport.Client
	ClusterResourceService            clusterclient.ClientService
	WorkspaceResourceService          workspaceclient.ClientService
	NamespaceResourceService          namespaceclient.ClientService
	ClusterGroupResourceService       clustergroupclient.ClientService
	NodePoolResourceService           nodepoolclient.ClientService
	OrganizationIAMResourceService    iamorganizationclient.ClientService
	ClusterGroupIAMResourceService    iamclustergroupclient.ClientService
	ClusterIAMResourceService         iamclusterclient.ClientService
	WorkspaceIAMResourceService       iamworkspaceclient.ClientService
	NamespaceIAMResourceService       iamnamespaceclient.ClientService
	ClusterPolicyResourceService      policyclusterclient.ClientService
	ClusterGroupPolicyResourceService policyclustergroupclient.ClientService
	WorkspacePolicyResourceService    policyworkspaceclient.ClientService
	OrganizationPolicyResourceService policyorganizationclient.ClientService
}
