/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package client

import (
	"net/http"
	"runtime"

	clusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster"
	continuousdeliveryclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/continuousdelivery"
	gitrepositoryclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/gitrepository"
	iamclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/iam_policy"
	kustomizationclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/kustomization"
	policyclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/policy"
	sourcesecretclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/cluster/sourcesecret"
	clustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup"
	continuousdeliveryclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/continuousdelivery"
	gitrepositoryclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/gitrepository"
	iamclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/iam_policy"
	kustomizationclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/kustomization"
	policyclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/policy"
	sourcesecretclustergroupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clustergroup/sourcesecret"
	credentialclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/credential"
	eksclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/ekscluster"
	eksnodepoolclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/ekscluster/nodepool"
	integrationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/integration"
	namespaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/namespace"
	iamnamespaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/namespace/iam_policy"
	nodepoolclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/nodepool"
	iamorganizationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/organization/iam_policy"
	policyorganizationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/organization/policy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	workspaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/workspace"
	iamworkspaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/workspace/iam_policy"
	policyworkspaceclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/workspace/policy"
)

// NewHTTPClient creates a new tanzu mission control HTTP client.
func NewHTTPClient(config *proxy.TLSConfig) (*TanzuMissionControl, error) {
	httpClient, err := transport.NewClient(config)
	if err != nil {
		return nil, err
	}

	return newHTTPClient(httpClient), nil
}

// NewTestHTTPClientWithDefaultTransport is intended primarily for testing only, as httpmock requires a default transport object be used in order to intercept and mock traffic.
func NewTestHTTPClientWithDefaultTransport() *TanzuMissionControl {
	return newHTTPClient(transport.NewClientWithDefaultTransport())
}

func newHTTPClient(httpClient *transport.Client) *TanzuMissionControl {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Connection", "keep-alive")
	headers.Set("x-client-name", "tmc-terraform-provider")
	headers.Set("x-client-platform", runtime.GOOS)
	headers.Set("x-client-version", "1.0.2")

	httpClient.AddHeaders(headers)

	return &TanzuMissionControl{
		Client:                                        httpClient,
		ClusterResourceService:                        clusterclient.New(httpClient),
		EKSClusterResourceService:                     eksclusterclient.New(httpClient),
		EKSNodePoolResourceService:                    eksnodepoolclient.New(httpClient),
		WorkspaceResourceService:                      workspaceclient.New(httpClient),
		NamespaceResourceService:                      namespaceclient.New(httpClient),
		ClusterGroupResourceService:                   clustergroupclient.New(httpClient),
		NodePoolResourceService:                       nodepoolclient.New(httpClient),
		OrganizationIAMResourceService:                iamorganizationclient.New(httpClient),
		ClusterGroupIAMResourceService:                iamclustergroupclient.New(httpClient),
		ClusterIAMResourceService:                     iamclusterclient.New(httpClient),
		WorkspaceIAMResourceService:                   iamworkspaceclient.New(httpClient),
		NamespaceIAMResourceService:                   iamnamespaceclient.New(httpClient),
		ClusterPolicyResourceService:                  policyclusterclient.New(httpClient),
		ClusterGroupPolicyResourceService:             policyclustergroupclient.New(httpClient),
		WorkspacePolicyResourceService:                policyworkspaceclient.New(httpClient),
		OrganizationPolicyResourceService:             policyorganizationclient.New(httpClient),
		CredentialResourceService:                     credentialclient.New(httpClient),
		IntegrationResourceService:                    integrationclient.New(httpClient),
		ClusterContinuousDeliveryResourceService:      continuousdeliveryclusterclient.New(httpClient),
		ClusterGitRepositoryResourceService:           gitrepositoryclusterclient.New(httpClient),
		ClusterKustomizationResourceService:           kustomizationclusterclient.New(httpClient),
		ClusterGroupContinuousDeliveryResourceService: continuousdeliveryclustergroupclient.New(httpClient),
		ClusterGroupGitRepositoryResourceService:      gitrepositoryclustergroupclient.New(httpClient),
		ClusterGroupKustomizationResourceService:      kustomizationclustergroupclient.New(httpClient),
		ClusterSourcesecretResourceService:            sourcesecretclusterclient.New(httpClient),
		ClusterGroupSourcesecretResourceService:       sourcesecretclustergroupclient.New(httpClient),
	}
}

// TanzuMissionControl is a client for tanzu mission control.
type TanzuMissionControl struct {
	*transport.Client
	ClusterResourceService                        clusterclient.ClientService
	EKSClusterResourceService                     eksclusterclient.ClientService
	EKSNodePoolResourceService                    eksnodepoolclient.ClientService
	WorkspaceResourceService                      workspaceclient.ClientService
	NamespaceResourceService                      namespaceclient.ClientService
	ClusterGroupResourceService                   clustergroupclient.ClientService
	NodePoolResourceService                       nodepoolclient.ClientService
	OrganizationIAMResourceService                iamorganizationclient.ClientService
	ClusterGroupIAMResourceService                iamclustergroupclient.ClientService
	ClusterIAMResourceService                     iamclusterclient.ClientService
	WorkspaceIAMResourceService                   iamworkspaceclient.ClientService
	NamespaceIAMResourceService                   iamnamespaceclient.ClientService
	ClusterPolicyResourceService                  policyclusterclient.ClientService
	ClusterGroupPolicyResourceService             policyclustergroupclient.ClientService
	WorkspacePolicyResourceService                policyworkspaceclient.ClientService
	OrganizationPolicyResourceService             policyorganizationclient.ClientService
	CredentialResourceService                     credentialclient.ClientService
	IntegrationResourceService                    integrationclient.ClientService
	ClusterContinuousDeliveryResourceService      continuousdeliveryclusterclient.ClientService
	ClusterGitRepositoryResourceService           gitrepositoryclusterclient.ClientService
	ClusterKustomizationResourceService           kustomizationclusterclient.ClientService
	ClusterGroupContinuousDeliveryResourceService continuousdeliveryclustergroupclient.ClientService
	ClusterGroupGitRepositoryResourceService      gitrepositoryclustergroupclient.ClientService
	ClusterGroupKustomizationResourceService      kustomizationclustergroupclient.ClientService
	ClusterSourcesecretResourceService            sourcesecretclusterclient.ClientService
	ClusterGroupSourcesecretResourceService       sourcesecretclustergroupclient.ClientService
}
