/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package manifestclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	manifestmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/manifest"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters:manifest"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
)

// New creates a new cluster manifest helper API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster manifest helper API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ClusterManifestHelperGetManifest(params *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) (*manifestmodel.VmwareTanzuManageV1alpha1ClusterClusterManifestGetResponse, error)
}

/*
ClusterManifestHelperGetManifest gets attach manifest for a cluster.
*/
func (a *Client) ClusterManifestHelperGetManifest(params *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) (*manifestmodel.VmwareTanzuManageV1alpha1ClusterClusterManifestGetResponse, error) {
	queryParams := url.Values{}

	if params.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, params.ManagementClusterName)
	}

	if params.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, params.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, params.Name).AppendQueryParams(queryParams).String()
	response := &manifestmodel.VmwareTanzuManageV1alpha1ClusterClusterManifestGetResponse{}

	err := a.Get(requestURL, response)

	return response, err
}
