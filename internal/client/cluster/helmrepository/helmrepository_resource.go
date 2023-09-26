/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmrepositoryclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	helmchartsorgmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrepository"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiSubGroup                        = "namespaces"
	apiKind                            = "fluxcd/helm/repositories"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
	queryParamKeyOrgID                 = "fullName.orgID"
)

// New creates a new cluster Flux CD helm charts resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster Flux CD helm charts resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for VmwareTanzuManageV1alpha1ClusterFluxcdHelmChartsResourceService Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClusterFluxcdHelmRepositoryResourceServiceGet(fn *helmchartsorgmodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryFullName) (*helmchartsorgmodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryGetResponse, error)

	VmwareTanzuManageV1alpha1ClusterFluxcdHelmRepositoryResourceServiceList(request *helmchartsorgmodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySearchScope) (*helmchartsorgmodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryListResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdHelmRepositoryResourceServiceGet gets a Flux CD helm charts scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdHelmRepositoryResourceServiceGet(fn *helmchartsorgmodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryFullName) (*helmchartsorgmodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryGetResponse, error) {
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
	helmchartResponse := &helmchartsorgmodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryGetResponse{}
	err := p.Get(requestURL, helmchartResponse)

	return helmchartResponse, err
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdHelmRepositoryResourceServiceList updates overwrite a Flux CD helm charts scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdHelmRepositoryResourceServiceList(fn *helmchartsorgmodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySearchScope) (*helmchartsorgmodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryListResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiSubGroup, fn.NamespaceName, apiKind).AppendQueryParams(queryParams).String()
	helmchartResponse := &helmchartsorgmodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryListResponse{}
	err := p.Get(requestURL, helmchartResponse)

	return helmchartResponse, err
}
