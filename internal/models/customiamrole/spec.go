/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrolemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1IamRoleSpec Spec for role.
//
// swagger:model vmware.tanzu.manage.v1alpha1.iam.role.Spec
type VmwareTanzuManageV1alpha1IamRoleSpec struct {

	// AggregationRule.
	AggregationRule *VmwareTanzuManageV1alpha1IamRoleAggregationRule `json:"aggregationRule,omitempty"`

	// Flag representing whether role is deprecated.
	IsDeprecated bool `json:"isDeprecated"`

	// This flag will help the client identify if this is an inbuilt role.
	IsInbuilt bool `json:"isInbuilt"`

	// Valid resources for this role.
	Resources []*VmwareTanzuManageV1alpha1IamPermissionResource `json:"resources"`

	// KubernetesRule.
	Rules []*VmwareTanzuManageV1alpha1IamRoleKubernetesRule `json:"rules"`

	// Tanzu-specific permissions for the role.
	TanzuPermissions []string `json:"tanzuPermissions"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRoleSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRoleSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1IamRoleSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
