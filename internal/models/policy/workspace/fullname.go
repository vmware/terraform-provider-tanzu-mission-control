/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyworkspacemodel

import (
	"fmt"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1WorkspacePolicyFullName Full name of the workspace policy. This includes the object
// name along with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.policy.FullName
type VmwareTanzuManageV1alpha1WorkspacePolicyFullName struct {

	// Name of the policy.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of the workspace.
	WorkspaceName string `json:"workspaceName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePolicyFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePolicyFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspacePolicyFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

func (m *VmwareTanzuManageV1alpha1WorkspacePolicyFullName) ToString() string {
	if m == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s:%s", m.OrgID, m.WorkspaceName, m.Name)
}
