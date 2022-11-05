/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package credentialclient

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
)

const (
	apiVersionAndGroup = "v1alpha1/account/credentials"
)

// New creates a new cluster resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for credentials resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	CredentialResourceServiceCreate(request *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialCreateCredentialRequest) (*credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialCreateCredentialResponse, error)

	CredentialResourceServiceDelete(fn *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialFullName) error

	CredentialResourceServiceGet(fn *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialFullName) (*credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialGetCredentialResponse, error)
}

/*
CredentialResourceServiceCreate creates a credential.
*/
func (c *Client) CredentialResourceServiceCreate(request *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialCreateCredentialRequest,
) (*credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialCreateCredentialResponse, error) {
	response := &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialCreateCredentialResponse{}
	err := c.Create(apiVersionAndGroup, request, response)

	return response, err
}

/*
CredentialResourceServiceDelete deletes a credential.
*/
func (c *Client) CredentialResourceServiceDelete(
	fn *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialFullName,
) error {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).String()

	return c.Delete(requestURL)
}

/*
CredentialResourceServiceGet gets a credential.
*/
func (c *Client) CredentialResourceServiceGet(
	fn *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialFullName,
) (*credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialGetCredentialResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).String()
	resp := &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialGetCredentialResponse{}
	err := c.Get(requestURL, resp)

	return resp, err
}
