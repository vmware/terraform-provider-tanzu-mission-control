/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyworkspaceclient

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyworkspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/workspace"
)

const (
	apiVersionAndGroup = "v1alpha1/workspaces"
	apiKind            = "policies"
)

// New creates a new workspace policy resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for workspace policy resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for ManageV1alpha1WorkspacePolicyResourceService Client methods.
type ClientService interface {
	ManageV1alpha1WorkspacePolicyResourceServiceCreate(request *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse, error)

	ManageV1alpha1WorkspacePolicyResourceServiceDelete(fn *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName) error

	ManageV1alpha1WorkspacePolicyResourceServiceGet(fn *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse, error)

	ManageV1alpha1WorkspacePolicyResourceServiceUpdate(request *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse, error)
}

/*
ManageV1alpha1WorkspacePolicyResourceServiceCreate creates a policy scoped to a workspace resource.
*/
func (p *Client) ManageV1alpha1WorkspacePolicyResourceServiceCreate(request *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Policy.FullName.WorkspaceName, apiKind).String()
	policyWorkspaceResponse := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse{}
	err := p.Create(requestURL, request, policyWorkspaceResponse)

	return policyWorkspaceResponse, err
}

/*
ManageV1alpha1WorkspacePolicyResourceServiceDelete deletes a policy scoped to a workspace resource.
*/
func (p *Client) ManageV1alpha1WorkspacePolicyResourceServiceDelete(fn *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName) error {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.WorkspaceName, apiKind, fn.Name).String()

	return p.Delete(requestURL)
}

/*
ManageV1alpha1WorkspacePolicyResourceServiceGet gets a policy scoped to a workspace resource.
*/
func (p *Client) ManageV1alpha1WorkspacePolicyResourceServiceGet(fn *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.WorkspaceName, apiKind, fn.Name).String()
	policyWorkspaceResponse := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse{}
	err := p.Get(requestURL, policyWorkspaceResponse)

	return policyWorkspaceResponse, err
}

/*
ManageV1alpha1WorkspacePolicyResourceServiceUpdate updates overwrite a policy scoped to a workspace resource.
*/
func (p *Client) ManageV1alpha1WorkspacePolicyResourceServiceUpdate(request *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest) (*policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse, error) {
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.Policy.FullName.WorkspaceName, apiKind, request.Policy.FullName.Name).String()
	policyWorkspaceResponse := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse{}
	err := p.Update(requestURL, request, policyWorkspaceResponse)

	return policyWorkspaceResponse, err
}
