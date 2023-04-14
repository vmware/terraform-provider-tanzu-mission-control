/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package providerekscluster

import (
	"net/url"

	"github.com/go-openapi/runtime"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/provider_ekscluster"
)

const (
	apiVersionAndGroup          = "/v1alpha1/manage/providereksclusters"
	queryParamKeyCredentialName = "fullName.credentialName" //nolint:gosec
	queryParamKeyRegion         = "fullName.region"
)

// New creates a new provider eks cluster resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for provider eks cluster resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientOption is the option for Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods.
type ClientService interface {
	ProviderEksClusterResourceServiceGet(request *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName) (*models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterGetProviderEksClusterResponse, error)

	ProviderEksClusterResourceServiceUpdate(params *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterRequest) (*models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterResponse, error)
}

/*
ProviderEksClusterResourceServiceGet gets a provider eks cluster.
*/
func (c *Client) ProviderEksClusterResourceServiceGet(fn *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName) (*models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterGetProviderEksClusterResponse, error) {
	queryParams := url.Values{}
	if fn.CredentialName != "" {
		queryParams.Add(queryParamKeyCredentialName, fn.CredentialName)
	}

	if fn.Region != "" {
		queryParams.Add(queryParamKeyRegion, fn.Region)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).AppendQueryParams(queryParams).String()
	clusterResponse := &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterGetProviderEksClusterResponse{}

	err := c.Get(requestURL, clusterResponse)

	return clusterResponse, err
}

/*
ProviderEksClusterResourceServiceUpdate updates overwrite a provider eks cluster.
*/
func (c *Client) ProviderEksClusterResourceServiceUpdate(request *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterRequest) (*models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterResponse, error) {
	response := &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.ProviderEksCluster.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}
