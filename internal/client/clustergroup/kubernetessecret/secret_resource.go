// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kubernetessecretclustergroupclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	secret "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup"
)

const (
	apiVersionAndGroup         = "v1alpha1/clustergroups"
	apiKind                    = "secrets"
	apiSubGroup                = "namespace"
	queryParamKeyNamespaceName = "fullName.namespaceName"
	queryParamKeyOrgID         = "fullName.orgID"
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
	SecretResourceServiceCreate(request *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest) (*secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse, error)

	SecretResourceServiceDelete(fn *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName) error

	SecretResourceServiceGet(fn *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName) (*secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretGetSecretResponse, error)

	SecretResourceServiceUpdate(request *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest) (*secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse, error)
}

/*
SecretResourceServiceCreate creates a secret.
*/
func (c *Client) SecretResourceServiceCreate(request *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest) (*secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Secret.FullName.ClusterGroupName, apiSubGroup, apiKind).String()
	secretResponse := &secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse{}
	err := c.Create(requestURL, request, secretResponse)

	return secretResponse, err
}

/*
SecretResourceServiceDelete deletes a secret.
*/
func (c *Client) SecretResourceServiceDelete(fn *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName) error {
	queryParams := url.Values{}

	if fn.NamespaceName != "" {
		queryParams.Add(queryParamKeyNamespaceName, fn.NamespaceName)
	}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiSubGroup, apiKind, fn.Name).AppendQueryParams(queryParams).String()

	return c.Delete(requestURL)
}

/*
SecretResourceServiceGet gets a secret.
*/
func (c *Client) SecretResourceServiceGet(fn *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName) (*secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretGetSecretResponse, error) {
	queryParams := url.Values{}

	if fn.NamespaceName != "" {
		queryParams.Add(queryParamKeyNamespaceName, fn.NamespaceName)
	}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiSubGroup, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	secretResponse := &secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretGetSecretResponse{}
	err := c.Get(requestURL, secretResponse)

	return secretResponse, err
}

/*
SecretResourceServiceUpdate updates overwrite a secret.
*/
func (c *Client) SecretResourceServiceUpdate(request *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest) (*secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Secret.FullName.ClusterGroupName, apiSubGroup, apiKind, request.Secret.FullName.Name).String()
	secretResponse := &secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse{}
	err := c.Update(requestURL, request, secretResponse)

	return secretResponse, err
}
