/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipemodels

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1PolicyTypeRecipe A Recipe is an internal template for policy type.
//
// Recipe is a convenience decorator. It gives a friendly way to produce policy instances using simple parameters.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.type.recipe.Recipe
type VmwareTanzuManageV1alpha1PolicyTypeRecipe struct {

	// Full name for the policy recipe.
	FullName *VmwareTanzuManageV1alpha1PolicyTypeRecipeFullName `json:"fullName,omitempty"`

	// Metadata for the policy recipe object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the policy recipe.
	Spec *VmwareTanzuManageV1alpha1PolicyTypeRecipeSpec `json:"spec,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTypeRecipe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTypeRecipe) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1PolicyTypeRecipe

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
