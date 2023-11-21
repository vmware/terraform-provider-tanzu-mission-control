/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1PolicyTypeRecipeData Response from getting a Recipe.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.type.recipe.GetRecipeResponse
type VmwareTanzuManageV1alpha1PolicyTypeRecipeData struct {

	// Recipe returned.
	Recipe *VmwareTanzuManageV1alpha1PolicyTypeRecipe `json:"recipe,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTypeRecipeData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTypeRecipeData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1PolicyTypeRecipeData

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
