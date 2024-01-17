/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipemodels

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1PolicyTypeRecipeSpec Spec of policy recipe.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.type.recipe.Spec
type VmwareTanzuManageV1alpha1PolicyTypeRecipeSpec struct {

	// Deprecated specifies whether this version (latest version) of the recipe is deprecated.
	// Deprecated recipes will not be assignable to new policy instances nor visible in the UI.
	Deprecated bool `json:"deprecated,omitempty"`

	// InputSchema defines the set of variable inputs needed to create a policy using this recipe, in JsonSchema format.
	// This input schema is for the latest version of the recipe. For previous versions, check Versions API.
	InputSchema string `json:"inputSchema,omitempty"`

	// Policy templates are references to kubernetes resources (policy pre-requisites) associated with this recipe.
	// These templates will be applied on clusters where policy instances using this recipe are effective.
	// A recipe can have 0 or more templates associated with it.
	// These references are for the latest version of the recipe. For previous versions, check Versions API.
	PolicyTemplates []*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference `json:"policyTemplates"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTypeRecipeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTypeRecipeSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1PolicyTypeRecipeSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
