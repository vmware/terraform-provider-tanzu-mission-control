/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package permissiontemplateclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
	permissiontemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/permissiontemplate"
)

const (
	// API Paths.
	apiPath = "v1alpha1/account/credentials:permissiontemplate"

	// Query Params.
	capabilityQueryParam = "capability"
	providerQueryParam   = "provider"
)

// New creates a new permission template resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for permission template resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	PermissionTemplateResourceServiceGenerate(request *permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateRequest) (*permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse, error)

	PermissionTemplateResourceServiceGet(request *permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateRequest) (*permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse, error)
}

/*
PermissionTemplateResourceServiceGenerate generates a permission template.
*/
func (c *Client) PermissionTemplateResourceServiceGenerate(request *permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateRequest) (*permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse, error) {
	response := &permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse{}
	err := c.Create(apiPath, request, response)

	return response, err
}

/*
PermissionTemplateResourceServiceGet gets an existing permission template.
*/
func (c *Client) PermissionTemplateResourceServiceGet(request *permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateRequest) (*permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse, error) {
	response := &permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse{}
	requestURL := helper.ConstructRequestURL(apiPath, request.FullName.Name)

	queryParams := url.Values{}

	if request.Capability != "" {
		queryParams.Add(capabilityQueryParam, request.Capability)
	}

	if *request.Provider != credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProviderPROVIDERUNSPECIFIED {
		queryParams.Add(providerQueryParam, string(*request.Provider))
	}

	if len(queryParams) > 0 {
		requestURL = requestURL.AppendQueryParams(queryParams)
	}

	err := c.Get(requestURL.String(), response)

	return response, err
}
