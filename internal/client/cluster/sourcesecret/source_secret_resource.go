/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package sourcesecretclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	sourcesecret "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/cluster"
)

const (
	apiVersionAndGroup                 = "v1alpha1/clusters"
	apiKind                            = "fluxcd/sourcesecrets"
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
	ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceCreate(request *sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourceSecretRequest) (*sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse, error)

	ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceDelete(fn *sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) error

	ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceGet(fn *sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) (*sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse, error)

	ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceUpdate(request *sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourceSecretRequest) (*sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse, error)
}

/*
ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceCreate creates a source secret.
*/
func (c *Client) ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceCreate(request *sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourceSecretRequest) (*sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.SourceSecret.FullName.ClusterName, apiKind).String()
	sourcesecretResponse := &sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse{}
	err := c.Create(requestURL, request, sourcesecretResponse)

	return sourcesecretResponse, err
}

/*
ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceDelete deletes a source secret.
*/
func (c *Client) ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceDelete(fn *sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) error {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams[queryParamKeyManagementClusterName] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams[queryParamKeyProvisionerName] = []string{fn.ProvisionerName}
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind, fn.Name).AppendQueryParams(queryParams).String()

	return c.Delete(requestURL)
}

/*
ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceGet gets a source secret.
*/
func (c *Client) ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceGet(fn *sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) (*sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse, error) {
	queryParams := url.Values{}

	if fn.ManagementClusterName != "" {
		queryParams[queryParamKeyManagementClusterName] = []string{fn.ManagementClusterName}
	}

	if fn.ProvisionerName != "" {
		queryParams[queryParamKeyProvisionerName] = []string{fn.ProvisionerName}
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.ClusterName, apiKind, fn.Name).AppendQueryParams(queryParams).String()
	sourcesecretResponse := &sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse{}
	err := c.Get(requestURL, sourcesecretResponse)

	return sourcesecretResponse, err
}

/*
ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceUpdate updates overwrite a source secret.
*/
func (c *Client) ManageV1alpha1ClusterFluxcdSourcesecretResourceServiceUpdate(request *sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourceSecretRequest) (*sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.SourceSecret.FullName.ClusterName, apiKind, request.SourceSecret.FullName.Name).String()
	secretResponse := &sourcesecret.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretResponse{}
	err := c.Update(requestURL, request, secretResponse)

	return secretResponse, err
}
