// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmchartclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	helmchartsorgmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmcharts"
)

const (
	apiVersionAndGroup = "v1alpha1/organization/fluxcd/helm/repositories"
	apiSubGroup        = "chartmetadatas"
	apiKind            = "charts"
	queryParamKeyOrgID = "fullName.orgID"
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
	VmwareTanzuManageV1alpha1ClusterFluxcdChartResourceServiceGet(fn *helmchartsorgmodel.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartFullName) (*helmchartsorgmodel.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartGetResponse, error)

	VmwareTanzuManageV1alpha1ClusterFluxcdChartResourceServiceList(request *helmchartsorgmodel.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope) (*helmchartsorgmodel.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartListResponse, error)
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdChartResourceServiceGet gets a Flux CD helm charts scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdChartResourceServiceGet(fn *helmchartsorgmodel.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartFullName) (*helmchartsorgmodel.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartGetResponse, error) {
	queryParams := url.Values{}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.RepositoryName, apiSubGroup, fn.ChartMetadataName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	helmchartResponse := &helmchartsorgmodel.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartGetResponse{}
	err := p.Get(requestURL, helmchartResponse)

	return helmchartResponse, err
}

/*
VmwareTanzuManageV1alpha1ClusterFluxcdHelmChartsResourceServiceUpdate updates overwrite a Flux CD helm charts scoped to a cluster resource.
*/
func (p *Client) VmwareTanzuManageV1alpha1ClusterFluxcdChartResourceServiceList(request *helmchartsorgmodel.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope) (*helmchartsorgmodel.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartListResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.RepositoryName, apiSubGroup, request.ChartMetadataName, apiKind).String()
	helmchartResponse := &helmchartsorgmodel.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartListResponse{}
	err := p.Get(requestURL, helmchartResponse)

	return helmchartResponse, err
}
