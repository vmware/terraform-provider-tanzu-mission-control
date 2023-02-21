/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package secretexportclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	secret "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster/secretexport"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiKind                            = "secretexports"
	namespaces                         = "namespaces"
	queryParamKeyManagementClusterName = "fullName.managementClusterName"
	queryParamKeyProvisionerName       = "fullName.provisionerName"
)

// New creates a new secret export resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for secret export resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	SecretExportResourceServiceCreate(request *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportRequest) (*secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportResponse, error)

	SecretExportResourceServiceDelete(fn *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretexportFullName) error

	SecretExportResourceServiceGet(fn *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretexportFullName) (*secret.VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretExportResponse, error)
}

/*
SecretExportResourceServiceCreate creates a secret export.
*/
func (c *Client) SecretExportResourceServiceCreate(request *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportRequest) (*secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.SecretExport.FullName.ClusterName, namespaces, request.SecretExport.FullName.NamespaceName, apiKind).String()
	secretexportResponse := &secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportResponse{}
	err := c.Create(requestURL, request, secretexportResponse)

	return secretexportResponse, err
}

/*
SecretExportResourceServiceDelete deletes a secret export.
*/
func (c *Client) SecretExportResourceServiceDelete(fn *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretexportFullName) error {
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
SecretExportResourceServiceGet gets a secret export.
*/
func (c *Client) SecretExportResourceServiceGet(fn *secret.VmwareTanzuManageV1alpha1ClusterNamespaceSecretexportFullName) (*secret.VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretExportResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams[queryParamKeyManagementClusterName] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams[queryParamKeyProvisionerName] = []string{fn.ProvisionerName}
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, namespaces, fn.NamespaceName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	secretexportResponse := &secret.VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretExportResponse{}
	err := c.Get(requestURL, secretexportResponse)

	return secretexportResponse, err
}
