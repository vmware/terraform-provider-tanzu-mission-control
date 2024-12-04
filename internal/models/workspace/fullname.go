// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package workspacemodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1WorkspaceFullName Full name of the workspace. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.FullName
type VmwareTanzuManageV1alpha1WorkspaceFullName struct {

	// Name of this Workspace.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspaceFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
