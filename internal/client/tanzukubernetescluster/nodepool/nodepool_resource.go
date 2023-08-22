/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetesnodepoolclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tkcmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
	tkcnodepool "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/nodepool"
)

const (
	apiVersionAndGroup      = "/v1alpha1/managementclusters"
	apiNodepoolsPath        = "nodepools"
	provisioners            = "provisioners"
	queryParamKeySortBy     = "sortBy"
	tanzukubernetesclusters = "tanzukubernetesclusters"
)

// New creates a new tanzu kubernetes cluster node pool resource service API client.
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
	TanzuKubernetesNodePoolResourceServiceGet(fn *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName) (*tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIResponse, error)

	TanzuKubernetesNodePoolResourceServiceCreate(request *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIRequest) (*tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIResponse, error)

	TanzuKubernetesNodePoolResourceServiceList(cluster *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName) (*tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolListNodepoolsResponse, error)

	TanzuKubernetesNodePoolResourceServiceDelete(fn *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName) error

	TanzuKubernetesNodePoolResourceServiceUpdate(request *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIRequest) (*tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIResponse, error)
}

// TanzuKubernetesNodePoolResourceServiceGet implements ClientService.
func (c *Client) TanzuKubernetesNodePoolResourceServiceGet(fn *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName) (*tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIResponse, error) {
	queryParams := url.Values{}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ManagementClusterName, provisioners, fn.ProvisionerName, tanzukubernetesclusters, fn.TanzuKubernetesClusterName, apiNodepoolsPath, fn.Name).AppendQueryParams(queryParams).String()
	response := &tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIResponse{}

	err := c.Get(requestURL, response)

	return response, err
}

// TanzuKubernetesNodePoolResourceServiceCreate implements ClientService.
func (c *Client) TanzuKubernetesNodePoolResourceServiceCreate(request *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIRequest) (*tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIResponse, error) {
	response := &tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIResponse{}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Nodepool.FullName.ManagementClusterName, provisioners, request.Nodepool.FullName.ProvisionerName, tanzukubernetesclusters, request.Nodepool.FullName.TanzuKubernetesClusterName, apiNodepoolsPath).String()

	err := c.Create(requestURL, request, response)

	return response, err
}

// TanzuKubernetesNodePoolResourceServiceDelete implements ClientService.
func (c *Client) TanzuKubernetesNodePoolResourceServiceDelete(fn *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName) error {
	queryParams := url.Values{}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ManagementClusterName, provisioners, fn.ProvisionerName, tanzukubernetesclusters, fn.TanzuKubernetesClusterName, apiNodepoolsPath, fn.Name).AppendQueryParams(queryParams).String()

	return c.Delete(requestURL)
}

// TanzuKubernetesNodePoolResourceServiceList implements ClientService.
func (c *Client) TanzuKubernetesNodePoolResourceServiceList(cluster *tkcmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName) (*tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolListNodepoolsResponse, error) {
	queryParams := url.Values{}

	// for stability of the results
	queryParams.Add(queryParamKeySortBy, "createTime")

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, cluster.ManagementClusterName, provisioners, cluster.ProvisionerName, tanzukubernetesclusters, cluster.Name, apiNodepoolsPath).AppendQueryParams(queryParams).String()
	clusterNodePoolsResponse := &tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolListNodepoolsResponse{}
	err := c.Get(requestURL, clusterNodePoolsResponse)

	return clusterNodePoolsResponse, err
}

// TanzuKubernetesNodePoolResourceServiceUpdate implements ClientService.
func (c *Client) TanzuKubernetesNodePoolResourceServiceUpdate(request *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIRequest) (*tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIResponse, error) {
	response := &tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolAPIResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Nodepool.FullName.ManagementClusterName, provisioners, request.Nodepool.FullName.ProvisionerName, tanzukubernetesclusters, request.Nodepool.FullName.TanzuKubernetesClusterName, apiNodepoolsPath, request.Nodepool.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}
