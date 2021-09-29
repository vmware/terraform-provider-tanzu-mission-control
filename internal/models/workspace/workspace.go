/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package workspacemodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1WorkspaceWorkspace A group of managed Kubenetes namespaces.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.Workspace
type VmwareTanzuManageV1alpha1WorkspaceWorkspace struct {

	// Full name for the Workspace.
	FullName *VmwareTanzuManageV1alpha1WorkspaceFullName `json:"fullName,omitempty"`

	// Metadata for the Workspace object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceWorkspace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceWorkspace) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspaceWorkspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
