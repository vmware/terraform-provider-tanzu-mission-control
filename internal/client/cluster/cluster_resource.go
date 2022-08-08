/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	queryParamKeyForce                 = "force"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
)

// New creates a new cluster resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ManageV1alpha1ClusterResourceServiceCreate(request *clustermodel.VmwareTanzuManageV1alpha1ClusterRequest) (*clustermodel.VmwareTanzuManageV1alpha1ClusterResponse, error)

	ManageV1alpha1ClusterResourceServiceDelete(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName, force string) error

	ManageV1alpha1ClusterResourceServiceGet(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) (*clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse, error)

	ManageV1alpha1ClusterResourceServiceUpdate(request *clustermodel.VmwareTanzuManageV1alpha1ClusterRequest) (*clustermodel.VmwareTanzuManageV1alpha1ClusterResponse, error)
}

/*
ManageV1alpha1ClusterResourceServiceCreate creates a cluster.
*/
func (c *Client) ManageV1alpha1ClusterResourceServiceCreate(
	request *clustermodel.VmwareTanzuManageV1alpha1ClusterRequest,
) (*clustermodel.VmwareTanzuManageV1alpha1ClusterResponse, error) {
	response := &clustermodel.VmwareTanzuManageV1alpha1ClusterResponse{}
	err := c.Create(apiVersionAndGroup, request, response)

	return response, err
}

/*
ManageV1alpha1ClusterResourceServiceUpdate updates a cluster.
*/
func (c *Client) ManageV1alpha1ClusterResourceServiceUpdate(
	request *clustermodel.VmwareTanzuManageV1alpha1ClusterRequest,
) (*clustermodel.VmwareTanzuManageV1alpha1ClusterResponse, error) {
	response := &clustermodel.VmwareTanzuManageV1alpha1ClusterResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Cluster.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}

/*
ManageV1alpha1ClusterResourceServiceDelete deletes a cluster.
*/
func (c *Client) ManageV1alpha1ClusterResourceServiceDelete(
	fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName, force string,
) error {
	queryParams := url.Values{
		queryParamKeyForce: []string{force},
	}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).AppendQueryParams(queryParams).String()

	return c.Delete(requestURL)
}

/*
ManageV1alpha1ClusterResourceServiceGet gets a cluster.
*/
func (c *Client) ManageV1alpha1ClusterResourceServiceGet(
	fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName,
) (*clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyProvisionerName, fn.ProvisionerName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).AppendQueryParams(queryParams).String()
	clusterResponse := &clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse{}
	err := c.Get(requestURL, clusterResponse)

	return clusterResponse, err
}
