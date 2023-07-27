/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package aksnodepool

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	aksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
)

const (
	apiVersionAndGroup                     = "/v1alpha1/aksclusters"
	apiNodepoolsPath                       = "nodepools"
	queryParamKeyFullNameCredentialName    = "fullName.credentialName" //nolint:gosec
	queryParamKeyFullNameSubscriptionID    = "fullName.subscriptionId"
	queryParamKeyFullNameResourceGroupName = "fullName.resourceGroupName"

	queryParamKeySearchCredentialName    = "searchScope.credentialName" //nolint:gosec
	queryParamKeySearchSubscriptionID    = "searchScope.subscriptionId"
	queryParamKeySearchResourceGroupName = "searchScope.resourceGroupName"
)

// New creates a new aks cluster node pool resource service API client.
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
	AksNodePoolResourceServiceCreate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse, error)

	AksNodePoolResourceServiceList(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse, error)

	AksNodePoolResourceServiceGet(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolGetNodepoolResponse, error)

	AksNodePoolResourceServiceUpdate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolResponse, error)

	AksNodePoolResourceServiceDelete(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName) error
}

// AksNodePoolResourceServiceCreate implements ClientService.
func (c *Client) AksNodePoolResourceServiceCreate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse, error) {
	response := &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Nodepool.FullName.AksClusterName, apiNodepoolsPath).String()

	err := c.Create(requestURL, request, response)

	return response, err
}

// AksNodePoolResourceServiceList implements ClientService.
func (c *Client) AksNodePoolResourceServiceList(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse, error) {
	queryParams := url.Values{}

	if fn.CredentialName != "" {
		queryParams.Add(queryParamKeySearchCredentialName, fn.CredentialName)
	}

	if fn.SubscriptionID != "" {
		queryParams.Add(queryParamKeySearchSubscriptionID, fn.SubscriptionID)
	}

	if fn.ResourceGroupName != "" {
		queryParams.Add(queryParamKeySearchResourceGroupName, fn.ResourceGroupName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name, apiNodepoolsPath).AppendQueryParams(queryParams).String()
	response := &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{}

	err := c.Get(requestURL, response)

	return response, err
}

func (c *Client) AksNodePoolResourceServiceGet(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolGetNodepoolResponse, error) {
	queryParams := url.Values{}

	if fn.CredentialName != "" {
		queryParams.Add(queryParamKeyFullNameCredentialName, fn.CredentialName)
	}

	if fn.SubscriptionID != "" {
		queryParams.Add(queryParamKeyFullNameSubscriptionID, fn.SubscriptionID)
	}

	if fn.ResourceGroupName != "" {
		queryParams.Add(queryParamKeyFullNameResourceGroupName, fn.ResourceGroupName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.AksClusterName, apiNodepoolsPath, fn.Name).AppendQueryParams(queryParams).String()
	response := &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolGetNodepoolResponse{}
	err := c.Get(requestURL, response)

	return response, err
}

// AksNodePoolResourceServiceUpdate implements ClientService.
func (c *Client) AksNodePoolResourceServiceUpdate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolResponse, error) {
	response := &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Nodepool.FullName.AksClusterName, apiNodepoolsPath, request.Nodepool.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}

// AksNodePoolResourceServiceDelete implements ClientService.
func (c *Client) AksNodePoolResourceServiceDelete(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName) error {
	queryParams := url.Values{}

	if fn.CredentialName != "" {
		queryParams.Add(queryParamKeyFullNameCredentialName, fn.CredentialName)
	}

	if fn.SubscriptionID != "" {
		queryParams.Add(queryParamKeyFullNameSubscriptionID, fn.SubscriptionID)
	}

	if fn.ResourceGroupName != "" {
		queryParams.Add(queryParamKeyFullNameResourceGroupName, fn.ResourceGroupName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.AksClusterName, apiNodepoolsPath, fn.Name).AppendQueryParams(queryParams).String()

	return c.Delete(requestURL)
}
