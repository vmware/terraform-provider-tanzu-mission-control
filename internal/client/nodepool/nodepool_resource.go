/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepools

import (
	"fmt"
	"net/url"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/client/transport"
	nodepoolsmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
)

// New creates a new cluster node pool resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for cluster node pool resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ManageV1alpha1ClusterNodePoolResourceServiceCreate(request *nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest) (*nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse, error)

	ManageV1alpha1ClusterNodePoolResourceServiceGet(fn *nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolFullName) (*nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse, error)

	ManageV1alpha1ClusterNodePoolResourceServiceDelete(fn *nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolFullName) error

	ManageV1alpha1ClusterNodePoolResourceServiceUpdate(request *nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest) (*nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse, error)
}

/*
ManageV1alpha1ClusterNodePoolsResourceServiceCreate creates a node pool resource.
*/
func (c *Client) ManageV1alpha1ClusterNodePoolResourceServiceCreate(request *nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest) (*nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse, error) {
	response := &nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse{}

	requestURL := fmt.Sprintf("%s/%s/%s", "v1alpha1/clusters", request.Nodepool.FullName.ClusterName, "nodepools")
	err := c.Create(requestURL, request, response)

	return response, err
}

/*
ManageV1alpha1ClusterNodePoolResourceServiceGet gets a node cluster.
*/
func (c *Client) ManageV1alpha1ClusterNodePoolResourceServiceGet(fn *nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolFullName) (*nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s?%s", "v1alpha1/clusters", fn.ClusterName, "nodepools", fn.Name, queryParams.Encode())
	clusterNodePoolResponse := &nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse{}
	err := c.Get(requestURL, clusterNodePoolResponse)

	return clusterNodePoolResponse, err
}

func (c *Client) ManageV1alpha1ClusterNodePoolResourceServiceDelete(fn *nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams["fullName.managementClusterName"] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams["fullName.provisionerName"] = []string{fn.ProvisionerName}
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s?%s", "v1alpha1/clusters", fn.ClusterName, "nodepools", fn.Name, queryParams.Encode())

	return c.Delete(requestURL)
}

func (c *Client) ManageV1alpha1ClusterNodePoolResourceServiceUpdate(request *nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest) (*nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse, error) {
	response := &nodepoolsmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse{}
	requestURL := fmt.Sprintf("%s/%s/%s/%s", "v1alpha1/clusters", request.Nodepool.FullName.ClusterName, "nodepools", request.Nodepool.FullName.Name)
	err := c.Update(requestURL, request, response)

	return response, err
}
