/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package workspacemodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceResponse Response from getting a Workspace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.GetWorkspaceResponse
type VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceResponse struct {

	// Workspace returned.
	Workspace *VmwareTanzuManageV1alpha1WorkspaceWorkspace `json:"workspace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
