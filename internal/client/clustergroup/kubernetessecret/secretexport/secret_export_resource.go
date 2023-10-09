/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package secretexportclustergroupclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	secret "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup/secretexport"
)

const (
	apiVersionAndGroup         = "v1alpha1/clustergroups"
	apiKind                    = "secretexports"
	apiSubGroup                = "namespace"
	queryParamKeyNamespaceName = "fullName.namespaceName"
	queryParamKeyOrgID         = "fullName.orgID"
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
	SecretExportResourceServiceCreate(request *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportRequest) (*secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportResponse, error)

	SecretExportResourceServiceDelete(fn *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportFullName) error

	SecretExportResourceServiceGet(fn *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportFullName) (*secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportGetSecretExportResponse, error)
}

/*
SecretExportResourceServiceCreate creates a secret export.
*/
func (c *Client) SecretExportResourceServiceCreate(request *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportRequest) (*secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.SecretExport.FullName.ClusterGroupName, apiSubGroup, apiKind).String()
	secretexportResponse := &secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportResponse{}
	err := c.Create(requestURL, request, secretexportResponse)

	return secretexportResponse, err
}

/*
SecretExportResourceServiceDelete deletes a secret export.
*/
func (c *Client) SecretExportResourceServiceDelete(fn *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportFullName) error {
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
SecretExportResourceServiceGet gets a secret export.
*/
func (c *Client) SecretExportResourceServiceGet(fn *secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportFullName) (*secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportGetSecretExportResponse, error) {
	queryParams := url.Values{}

	if fn.NamespaceName != "" {
		queryParams.Add(queryParamKeyNamespaceName, fn.NamespaceName)
	}

	if fn.OrgID != "" {
		queryParams.Add(queryParamKeyOrgID, fn.OrgID)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterGroupName, apiSubGroup, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	secretexportResponse := &secret.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportGetSecretExportResponse{}
	err := c.Get(requestURL, secretexportResponse)

	return secretexportResponse, err
}
