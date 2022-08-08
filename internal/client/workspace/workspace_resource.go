/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package workspaceclient

import (
	"fmt"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	workspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/workspace"
)

// New creates a new workspace resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for workspace resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	ManageV1alpha1WorkspaceResourceServiceCreate(request *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest) (*workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse, error)

	ManageV1alpha1WorkspaceResourceServiceDelete(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName) error

	ManageV1alpha1WorkspaceResourceServiceGet(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName) (*workspacemodel.VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceResponse, error)

	ManageV1alpha1WorkspaceResourceServiceUpdate(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest) (*workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse, error)
}

/*
  ManageV1alpha1WorkspaceResourceServiceUpdate updates a workspace.
*/
func (c *Client) ManageV1alpha1WorkspaceResourceServiceUpdate(
	request *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest,
) (*workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse, error) {
	requestURL := fmt.Sprintf("%s/%s", "v1alpha1/workspaces", request.Workspace.FullName.Name)
	namespaceResponse := &workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse{}
	err := c.Update(requestURL, request, namespaceResponse)

	return namespaceResponse, err
}

/*
  ManageV1alpha1WorkspaceResourceServiceCreate creates a workspace.
*/
func (c *Client) ManageV1alpha1WorkspaceResourceServiceCreate(
	request *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest,
) (*workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse, error) {
	namespaceResponse := &workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse{}
	err := c.Create("v1alpha1/workspaces", request, namespaceResponse)

	return namespaceResponse, err
}

/*
  ManageV1alpha1WorkspaceResourceServiceGet gets a workspace.
*/
func (c *Client) ManageV1alpha1WorkspaceResourceServiceGet(
	fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName,
) (*workspacemodel.VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceResponse, error) {
	requestURL := fmt.Sprintf("%s/%s", "v1alpha1/workspaces", fn.Name)
	workspaceResponse := &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceResponse{}

	err := c.Get(requestURL, workspaceResponse)

	return workspaceResponse, err
}

/*
  ManageV1alpha1WorkspaceResourceServiceDelete deletes a workspace.
*/
func (c *Client) ManageV1alpha1WorkspaceResourceServiceDelete(
	fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName,
) error {
	requestURL := fmt.Sprintf("%s/%s", "v1alpha1/workspaces", fn.Name)

	return c.Delete(requestURL)
}
