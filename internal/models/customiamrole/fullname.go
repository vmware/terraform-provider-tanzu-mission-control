/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrolemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1IamRoleFullName Full name for role.
//
// swagger:model vmware.tanzu.manage.v1alpha1.iam.role.FullName
type VmwareTanzuManageV1alpha1IamRoleFullName struct {

	// Name of the role.
	Name string `json:"name,omitempty"`

	// Org Id.
	OrgID string `json:"orgId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRoleFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRoleFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1IamRoleFullName

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
