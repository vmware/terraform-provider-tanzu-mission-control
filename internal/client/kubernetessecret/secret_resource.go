/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecretclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	secret "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiKind                            = "secrets"
	namespaces                         = "namespaces"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
)

// New creates a new secret resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for secret resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	SecretResourceServiceCreate(request *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest) (*secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse, error)

	SecretResourceServiceDelete(fn *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName) error

	SecretResourceServiceGet(fn *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName) (*secret.VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretResponse, error)

	SecretResourceServiceUpdate(request *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest) (*secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse, error)
}

/*
SecretResourceServiceCreate creates a secret.
*/
func (c *Client) SecretResourceServiceCreate(request *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest) (*secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Secret.FullName.ClusterName, namespaces, request.Secret.FullName.NamespaceName, apiKind).String()
	secretResponse := &secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse{}
	err := c.Create(requestURL, request, secretResponse)

	return secretResponse, err
}

/*
SecretResourceServiceDelete deletes a secret.
*/
func (c *Client) SecretResourceServiceDelete(fn *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams[queryParamKeyManagementClusterName] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams[queryParamKeyProvisionerName] = []string{fn.ProvisionerName}
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, namespaces, fn.NamespaceName, apiKind, fn.Name).AppendQueryParams(queryParams).String()

	return c.Delete(requestURL)
}

/*
SecretResourceServiceGet gets a secret.
*/
func (c *Client) SecretResourceServiceGet(fn *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName) (*secret.VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams[queryParamKeyManagementClusterName] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams[queryParamKeyProvisionerName] = []string{fn.ProvisionerName}
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, namespaces, fn.NamespaceName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	secretResponse := &secret.VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretResponse{}
	err := c.Get(requestURL, secretResponse)

	return secretResponse, err
}

/*
SecretResourceServiceUpdate updates overwrite a secret.
*/
func (c *Client) SecretResourceServiceUpdate(request *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest) (*secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Secret.FullName.ClusterName, namespaces, request.Secret.FullName.NamespaceName, apiKind, request.Secret.FullName.Name).String()
	secretResponse := &secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse{}
	err := c.Update(requestURL, request, secretResponse)

	return secretResponse, err
}
