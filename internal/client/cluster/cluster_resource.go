/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclient

import (
	"fmt"
	"net/url"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/client/transport"
	clustermodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/cluster"
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
func (c *Client) ManageV1alpha1ClusterResourceServiceCreate(request *clustermodel.VmwareTanzuManageV1alpha1ClusterRequest) (*clustermodel.VmwareTanzuManageV1alpha1ClusterResponse, error) {
	response := &clustermodel.VmwareTanzuManageV1alpha1ClusterResponse{}
	err := c.Create("v1alpha1/clusters", request, response)

	return response, err
}

/*
ManageV1alpha1ClusterResourceServiceUpdate updates a cluster.
*/
func (c *Client) ManageV1alpha1ClusterResourceServiceUpdate(request *clustermodel.VmwareTanzuManageV1alpha1ClusterRequest) (*clustermodel.VmwareTanzuManageV1alpha1ClusterResponse, error) {
	response := &clustermodel.VmwareTanzuManageV1alpha1ClusterResponse{}
	requestURL := fmt.Sprintf("%s/%s", "v1alpha1/clusters", request.Cluster.FullName.Name)
	err := c.Update(requestURL, request, response)

	return response, err
}

/*
ManageV1alpha1ClusterResourceServiceDelete deletes a cluster.
*/
func (c *Client) ManageV1alpha1ClusterResourceServiceDelete(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName, force string) error {
	queryParams := url.Values{
		"force": []string{force},
	}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s/%s?%s", "v1alpha1/clusters", fn.Name, queryParams.Encode())

	return c.Delete(requestURL)
}

/*
ManageV1alpha1ClusterResourceServiceGet gets a cluster.
*/
func (c *Client) ManageV1alpha1ClusterResourceServiceGet(fn *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) (*clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s/%s?%s", "v1alpha1/clusters", fn.Name, queryParams.Encode())
	clusterResponse := &clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse{}
	err := c.Get(requestURL, clusterResponse)

	return clusterResponse, err
}
