// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmreleaseclusterclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	releaseclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/cluster"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiSubGroup                        = "namespaces"
	apiKind                            = "fluxcd/helm/releases"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
	queryParamKeyOrgID                 = "fullName.orgID"
)

// New creates a new cluster Flux CD helm release resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster Flux CD helm release resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for VmwareTanzuManageV1alpha1ClusterReleaseResourceService Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceCreate(request *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest) (*releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse, error)

	VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceDelete(fn *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName) error

	VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceGet(fn *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName) (*releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGetResponse, error)

	VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceUpdate(request *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest) (*releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceCreate creates a Flux CD helm release scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceCreate(request *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest) (*releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Release.FullName.ClusterName, apiSubGroup, request.Release.FullName.NamespaceName, apiKind).String()
	releaseClusterResponse := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse{}
	err := p.Create(requestURL, request, releaseClusterResponse)

	return releaseClusterResponse, err
}

/*
VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceDelete deletes a Flux CD helm release scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceDelete(fn *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName) error {
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
VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceGet gets a Flux CD helm release scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceGet(fn *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName) (*releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGetResponse, error) {
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
	releaseClusterResponse := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGetResponse{}
	err := p.Get(requestURL, releaseClusterResponse)

	return releaseClusterResponse, err
}

/*
VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceUpdate updates overwrite a Flux CD helm release scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterReleaseResourceServiceUpdate(request *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest) (*releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Release.FullName.ClusterName, apiSubGroup, request.Release.FullName.NamespaceName, apiKind, request.Release.FullName.Name).String()
	releaseClusterResponse := &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse{}
	err := p.Update(requestURL, request, releaseClusterResponse)

	return releaseClusterResponse, err
}
