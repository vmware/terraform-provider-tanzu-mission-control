// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package iamworkspaceclient

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
	workspaceiammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy/workspace"
	workspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/workspace"
)

const apiVersionAndGroup = "v1alpha1/workspaces:iam"

// New creates a new workspace IAM Policy resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for workspace IAM Policy resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for ManageV1alpha1WorkspaceIAMPolicy Client methods.
type ClientService interface {
	ManageV1alpha1WorkspaceIAMPolicyGet(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName) (*workspaceiammodel.VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceIAMPolicyResponse, error)

	ManageV1alpha1WorkspaceIAMPolicyPatch(request *workspaceiammodel.VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyRequest) (*workspaceiammodel.VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyResponse, error)

	ManageV1alpha1WorkspaceIAMPolicyUpdate(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName, request *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) (*workspaceiammodel.VmwareTanzuManageV1alpha1WorkspaceUpdateWorkspaceIAMPolicyResponse, error)
}

/*
  ManageV1alpha1WorkspaceIAMPolicyGet gets all iam policies scoped to a workspace.
*/

func (c *Client) ManageV1alpha1WorkspaceIAMPolicyGet(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName) (*workspaceiammodel.VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceIAMPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).String()
	response := &workspaceiammodel.VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceIAMPolicyResponse{}
	err := c.Get(requestURL, response)

	return response, err
}

/*
  ManageV1alpha1WorkspaceIAMPolicyPatch patches all iam policies scoped to a workspace.
*/

func (c *Client) ManageV1alpha1WorkspaceIAMPolicyPatch(request *workspaceiammodel.VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyRequest) (*workspaceiammodel.VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyResponse, error) {
	response := &workspaceiammodel.VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.FullName.Name).String()
	err := c.Patch(requestURL, request, response)

	return response, err
}

/*
  ManageV1alpha1WorkspaceIAMPolicyUpdate updates all iam policies scoped to a workspace, deletes if body is empty.
*/

func (c *Client) ManageV1alpha1WorkspaceIAMPolicyUpdate(fn *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName, request *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy) (*workspaceiammodel.VmwareTanzuManageV1alpha1WorkspaceUpdateWorkspaceIAMPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).String()
	response := &workspaceiammodel.VmwareTanzuManageV1alpha1WorkspaceUpdateWorkspaceIAMPolicyResponse{}
	err := c.Update(requestURL, request, response)

	return response, err
}
