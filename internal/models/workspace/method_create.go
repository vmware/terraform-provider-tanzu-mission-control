/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package workspacemodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceRequest Request to create a Workspace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.CreateWorkspaceRequest
type VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceRequest struct {

	// Workspace to create.
	Workspace *VmwareTanzuManageV1alpha1WorkspaceWorkspace `json:"workspace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceResponse Response from creating a Workspace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.CreateWorkspaceResponse
type VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceResponse struct {

	// Workspace created.
	Workspace *VmwareTanzuManageV1alpha1WorkspaceWorkspace `json:"workspace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspaceCreateWorkspaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
