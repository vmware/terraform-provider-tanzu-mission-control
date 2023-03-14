/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package gitrepositoryclustergroupclient

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	gitrepositoryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
)

const (
	apiVersionAndGroup = "v1alpha1/clustergroups"
	apiSubGroup        = "namespaces"
	apiKind            = "fluxcd/gitrepositories"
)

// New creates a new cluster Flux CD git repository resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster group Flux CD git repository resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceService Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceCreate(request *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest) (*gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse, error)

	VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceDelete(fn *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName) error

	VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceGet(fn *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName) (*gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGetGitRepositoryResponse, error)

	VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceUpdate(request *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest) (*gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceCreate creates a Flux CD git repository scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceCreate(request *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest) (*gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.GitRepository.FullName.ClusterGroupName, apiSubGroup, request.GitRepository.FullName.NamespaceName, apiKind).String()
	fluxCDGitRepositoryClusterGroupResponse := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse{}
	err := p.Create(requestURL, request, fluxCDGitRepositoryClusterGroupResponse)

	return fluxCDGitRepositoryClusterGroupResponse, err
}

/*
VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceDelete deletes a Flux CD git repository scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceDelete(fn *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName) error {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiSubGroup, fn.NamespaceName, apiKind, fn.Name).String()

	return p.Delete(requestURL)
}

/*
VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceGet gets a Flux CD git repository scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceGet(fn *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName) (*gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGetGitRepositoryResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiSubGroup, fn.NamespaceName, apiKind, fn.Name).String()
	fluxCDGitRepositoryClusterGroupResponse := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGetGitRepositoryResponse{}
	err := p.Get(requestURL, fluxCDGitRepositoryClusterGroupResponse)

	return fluxCDGitRepositoryClusterGroupResponse, err
}

/*
VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceUpdate updates overwrite a Flux CD git repository scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceUpdate(request *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest) (*gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.GitRepository.FullName.ClusterGroupName, apiSubGroup, request.GitRepository.FullName.NamespaceName, apiKind, request.GitRepository.FullName.Name).String()
	fluxCDGitRepositoryClusterGroupResponse := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse{}
	err := p.Update(requestURL, request, fluxCDGitRepositoryClusterGroupResponse)

	return fluxCDGitRepositoryClusterGroupResponse, err
}
