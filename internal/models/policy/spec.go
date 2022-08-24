/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policymodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1CommonPolicySpec The policy spec.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.Spec
type VmwareTanzuManageV1alpha1CommonPolicySpec struct {

	// Inputs needed for the selected policy recipe.
	// To find the required inputs, check the input schema of the selected policy recipe.
	Input interface{} `json:"input,omitempty"`

	// Label based Namespace Selector for the policy.
	NamespaceSelector *VmwareTanzuManageV1alpha1CommonPolicyLabelSelector `json:"namespaceSelector,omitempty"`

	// Name of the policy recipe (helper) used for creating a policy.
	// Use PolicyRecipe API to find the list of available recipe in each type.
	Recipe string `json:"recipe,omitempty"`

	// The version of the recipe used for the policy.
	// The latest version will be selected by default (if left empty).
	// Use PolicyRecipeVersion API to find the list of versions in each recipe.
	RecipeVersion string `json:"recipeVersion,omitempty"`

	// Generated yaml based policy resources (read-only).
	// These will only be seen when viewing effective policies on a cluster or namespace.
	Resources []string `json:"resources"`

	// Type of the policy object.
	// Use PolicyTypes API to find the list of available policy types.
	Type string `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
