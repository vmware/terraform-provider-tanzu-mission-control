// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmfeatureclustergroupclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	helmclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"
)

const (
	apiVersionAndGroup = "v1alpha1/clustergroups"
	apiKind            = "fluxcd/helm"
	queryParamKeyOrgID = "fullName.orgID"
)

// New creates a new cluster group Flux CD helm feature resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster group Flux CD helm feature resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for VmwareTanzuManageV1alpha1ClustergroupHelmResourceService Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceCreate(request *helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmRequest) (*helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmResponse, error)

	VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceDelete(fn *helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmFullName) error

	VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceList(rp *helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupHelmListHelmRequestParameters) (*helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmListHelmsResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceCreate creates a Flux CD helm feature scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceCreate(request *helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmRequest) (*helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Helm.FullName.ClusterGroupName, apiKind).String()
	helmClusterGroupResponse := &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmResponse{}
	err := p.Create(requestURL, request, helmClusterGroupResponse)

	return helmClusterGroupResponse, err
}

/*
VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceDelete deletes a Flux CD helm feature scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceDelete(fn *helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmFullName) error {
	queryParams := url.Values{}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiKind).AppendQueryParams(queryParams).String()

	return p.Delete(requestURL)
}

/*
VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceList lists Flux CD continuous deliveries scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupHelmResourceServiceList(rp *helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupHelmListHelmRequestParameters) (*helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmListHelmsResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, rp.SearchScope.ClusterGroupName, apiKind).String()
	helmClusterGroupResponse := &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmListHelmsResponse{}
	err := p.Get(requestURL, helmClusterGroupResponse)

	return helmClusterGroupResponse, err
}
