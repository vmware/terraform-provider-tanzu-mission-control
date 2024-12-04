// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package releaseclusterclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/cluster"
)

const (
	apiVersionAndGroup                            = "v1alpha1/clusters"
	apiKind                                       = "fluxcd/helm"
	queryParamKeyFullNameManagementClusterName    = "fullName.managementClusterName"
	queryParamKeyFullNameProvisionerName          = "fullName.provisionerName"
	queryParamKeyOrgID                            = "fullName.orgID"
	queryParamKeySearchScopeManagementClusterName = "searchScope.managementClusterName"
	queryParamKeySearchScopeProvisionerName       = "searchScope.provisionerName"
)

// New creates a new cluster Flux CD helm feature resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster Flux CD helm feature resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for VmwareTanzuManageV1alpha1ClusterHelmResourceService Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClusterHelmResourceServiceCreate(request *helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmRequest) (*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmResponse, error)

	VmwareTanzuManageV1alpha1ClusterHelmResourceServiceDelete(fn *helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmFullName) error

	VmwareTanzuManageV1alpha1ClusterHelmResourceServiceList(rp *helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmRequestParameters) (*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClusterHelmResourceServiceCreate creates a Flux CD helm feature scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterHelmResourceServiceCreate(request *helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmRequest) (*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Helm.FullName.ClusterName, apiKind).String()
	helmClusterResponse := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmResponse{}
	err := p.Create(requestURL, request, helmClusterResponse)

	return helmClusterResponse, err
}

/*
VmwareTanzuManageV1alpha1ClusterHelmResourceServiceDelete deletes a Flux CD helm feature scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterHelmResourceServiceDelete(fn *helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyFullNameManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyFullNameProvisionerName, fn.ProvisionerName)
	}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind).AppendQueryParams(queryParams).String()

	return p.Delete(requestURL)
}

/*
VmwareTanzuManageV1alpha1ClusterHelmResourceServiceList lists Flux CD helm scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterHelmResourceServiceList(rp *helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmRequestParameters) (*helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse, error) {
	queryParams := url.Values{}

	if rp.SearchScope.ManagementClusterName != "" {
		queryParams.Add(queryParamKeySearchScopeManagementClusterName, rp.SearchScope.ManagementClusterName)
	}

	if rp.SearchScope.ProvisionerName != "" {
		queryParams.Add(queryParamKeySearchScopeProvisionerName, rp.SearchScope.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, rp.SearchScope.ClusterName, apiKind).AppendQueryParams(queryParams).String()
	helmClusterResponse := &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse{}
	err := p.Get(requestURL, helmClusterResponse)

	return helmClusterResponse, err
}
