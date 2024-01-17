/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1PolicyTypeRecipeFullName Full name of the policy recipe. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.type.recipe.FullName
type VmwareTanzuManageV1alpha1PolicyTypeRecipeFullName struct {

	// Name of policy recipe.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of policy type.
	TypeName string `json:"typeName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTypeRecipeFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTypeRecipeFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1PolicyTypeRecipeFullName

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
