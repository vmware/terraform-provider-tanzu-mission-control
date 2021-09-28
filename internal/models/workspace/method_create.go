/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package workspacemodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1WorkspaceRequest Request to create a Workspace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.CreateWorkspaceRequest
type VmwareTanzuManageV1alpha1WorkspaceRequest struct {

	// Workspace to create.
	Workspace *VmwareTanzuManageV1alpha1WorkspaceWorkspace `json:"workspace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alphaWorkspaceResponse Response from creating a Workspace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.CreateWorkspaceResponse
type VmwareTanzuManageV1alphaWorkspaceResponse struct {

	// Workspace created.
	Workspace *VmwareTanzuManageV1alpha1WorkspaceWorkspace `json:"workspace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alphaWorkspaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alphaWorkspaceResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alphaWorkspaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
