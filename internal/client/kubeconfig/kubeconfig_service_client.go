/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubeconfig

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubeconfig"
	"net/url"
)

const (
	apiVersionAndGroup                         = "/v1alpha1/clusters"
	apiKubeconfigPath                          = "kubeconfig"
	queryParamKeyFullNameManagementClusterName = "full_name.managementClusterName"
	queryParamKeyFullNameProvisionerName       = "full_name.provisionerName"
)

// New creates a new kubeconfig service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for kubeconfig service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	KubeconfigServiceGet(fn *models.VmwareTanzuManageV1alpha1ClusterFullName) (*models.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponse, error)
}

/*
KubeconfigServiceGet gets cluster kubeconfig.
*/
func (c *Client) KubeconfigServiceGet(fn *models.VmwareTanzuManageV1alpha1ClusterFullName) (*models.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponse, error) {
	queryParams := url.Values{}
	if fn.ManagementClusterName != "" {
		queryParams.Add(queryParamKeyFullNameManagementClusterName, fn.ManagementClusterName)
	}

	if fn.ProvisionerName != "" {
		queryParams.Add(queryParamKeyFullNameProvisionerName, fn.ProvisionerName)
	}

	queryParams.Add("cli", string(models.VmwareTanzuManageV1alpha1ClusterKubeconfigCliTypeTANZUCLI))

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name, apiKubeconfigPath).AppendQueryParams(queryParams).String()
	clusterResponse := &models.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponse{}

	err := c.Get(requestURL, clusterResponse)

	return clusterResponse, err
}
