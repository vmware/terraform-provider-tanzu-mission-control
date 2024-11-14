// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package continuousdeliveryclustergroupclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	continuousdeliveryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/continuousdelivery/clustergroup"
)

const (
	apiVersionAndGroup = "v1alpha1/clustergroups"
	apiKind            = "fluxcd/continuousdelivery"
	queryParamKeyOrgID = "fullName.orgID"
)

// New creates a new cluster group Flux CD continuous delivery resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster group Flux CD continuous delivery resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryResourceService Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryResourceServiceCreate(request *continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest) (*continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse, error)

	VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryResourceServiceDelete(fn *continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryFullName) error

	VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryResourceServiceList(rp *continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesRequestParameters) (*continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryResourceServiceCreate creates a Flux CD continuous delivery scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryResourceServiceCreate(request *continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest) (*continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.ContinuousDelivery.FullName.ClusterGroupName, apiKind).String()
	fluxCDContinuousDeliveryClusterGroupResponse := &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse{}
	err := p.Create(requestURL, request, fluxCDContinuousDeliveryClusterGroupResponse)

	return fluxCDContinuousDeliveryClusterGroupResponse, err
}

/*
VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryResourceServiceDelete deletes a Flux CD continuous delivery scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryResourceServiceDelete(fn *continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryFullName) error {
	queryParams := url.Values{}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiKind).AppendQueryParams(queryParams).String()

	return p.Delete(requestURL)
}

/*
VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryResourceServiceList lists Flux CD continuous deliveries scoped to a cluster group resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryResourceServiceList(rp *continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesRequestParameters) (*continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, rp.SearchScope.ClusterGroupName, apiKind).String()
	fluxCDContinuousDeliveryClusterGroupResponse := &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse{}
	err := p.Get(requestURL, fluxCDContinuousDeliveryClusterGroupResponse)

	return fluxCDContinuousDeliveryClusterGroupResponse, err
}
