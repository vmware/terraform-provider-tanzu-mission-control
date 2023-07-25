/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package gitrepositoryclusterclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiSubGroup                        = "namespaces"
	apiKind                            = "fluxcd/gitrepositories"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
	queryParamKeyOrgID                 = "fullName.orgID"
)

// New creates a new cluster Flux CD git repository resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster Flux CD git repository resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceService Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceCreate(request *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest) (*gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse, error)

	VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceDelete(fn *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName) error

	VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceGet(fn *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName) (*gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGetGitRepositoryResponse, error)

	VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceUpdate(request *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest) (*gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceCreate creates a Flux CD git repository scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceCreate(request *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest) (*gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.GitRepository.FullName.ClusterName, apiSubGroup, request.GitRepository.FullName.NamespaceName, apiKind).String()
	fluxCDGitRepositoryClusterResponse := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse{}
	err := p.Create(requestURL, request, fluxCDGitRepositoryClusterResponse)

	return fluxCDGitRepositoryClusterResponse, err
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceDelete deletes a Flux CD git repository scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceDelete(fn *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiSubGroup, fn.NamespaceName, apiKind, fn.Name).AppendQueryParams(queryParams).String()

	return p.Delete(requestURL)
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceGet gets a Flux CD git repository scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceGet(fn *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName) (*gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGetGitRepositoryResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiSubGroup, fn.NamespaceName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	fluxCDGitRepositoryClusterResponse := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGetGitRepositoryResponse{}
	err := p.Get(requestURL, fluxCDGitRepositoryClusterResponse)

	return fluxCDGitRepositoryClusterResponse, err
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceUpdate updates overwrite a Flux CD git repository scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceUpdate(request *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest) (*gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.GitRepository.FullName.ClusterName, apiSubGroup, request.GitRepository.FullName.NamespaceName, apiKind, request.GitRepository.FullName.Name).String()
	fluxCDGitRepositoryClusterResponse := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse{}
	err := p.Update(requestURL, request, fluxCDGitRepositoryClusterResponse)

	return fluxCDGitRepositoryClusterResponse, err
}
