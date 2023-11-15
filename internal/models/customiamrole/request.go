/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrolemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1IamRoleCreateRoleRequest Request to create a Role.
//
// swagger:model vmware.tanzu.manage.v1alpha1.iam.role.CreateRoleRequest
type VmwareTanzuManageV1alpha1IamRoleData struct {

	// Role to create.
	Role *VmwareTanzuManageV1alpha1IamRole `json:"role,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRoleData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRoleData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1IamRoleData

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
