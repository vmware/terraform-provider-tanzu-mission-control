/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package continuousdeliveryclusterclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	continuousdeliveryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/continuousdelivery/cluster"
)

const (
	apiVersionAndGroup                            = "v1alpha1/clusters"
	apiKind                                       = "fluxcd/continuousdelivery"
	queryParamKeyFullNameManagementClusterName    = "fullName.managementClusterName"
	queryParamKeyFullNameProvisionerName          = "fullName.provisionerName"
	queryParamKeySearchScopeManagementClusterName = "searchScope.managementClusterName"
	queryParamKeySearchScopeProvisionerName       = "searchScope.provisionerName"
)

// New creates a new cluster Flux CD continuous delivery resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster Flux CD continuous delivery resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryResourceService Client methods.
type ClientService interface {
	VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryResourceServiceCreate(request *continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest) (*continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse, error)

	VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryResourceServiceDelete(fn *continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryFullName) error

	VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryResourceServiceList(rp *continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesRequestParameters) (*continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryResourceServiceCreate creates a Flux CD continuous delivery scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryResourceServiceCreate(request *continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest) (*continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.ContinuousDelivery.FullName.ClusterName, apiKind).String()
	fluxCDContinuousDeliveryClusterResponse := &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse{}
	err := p.Create(requestURL, request, fluxCDContinuousDeliveryClusterResponse)

	return fluxCDContinuousDeliveryClusterResponse, err
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryResourceServiceDelete deletes a Flux CD continuous delivery scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryResourceServiceDelete(fn *continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyFullNameManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyFullNameProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind).AppendQueryParams(queryParams).String()

	return p.Delete(requestURL)
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryResourceServiceList lists Flux CD continuous deliveries scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryResourceServiceList(rp *continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesRequestParameters) (*continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse, error) {
	queryParams := url.Values{}

	if rp.SearchScope.ManagementClusterName != "" {
		queryParams.Add(queryParamKeySearchScopeManagementClusterName, rp.SearchScope.ManagementClusterName)
	}

	if rp.SearchScope.ProvisionerName != "" {
		queryParams.Add(queryParamKeySearchScopeProvisionerName, rp.SearchScope.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, rp.SearchScope.ClusterName, apiKind).AppendQueryParams(queryParams).String()
	fluxCDContinuousDeliveryClusterResponse := &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse{}
	err := p.Get(requestURL, fluxCDContinuousDeliveryClusterResponse)

	return fluxCDContinuousDeliveryClusterResponse, err
}
