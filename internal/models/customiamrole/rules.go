/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrolemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1IamRoleKubernetesRule KubernetesRule for a role.
//
// swagger:model vmware.tanzu.manage.v1alpha1.iam.role.KubernetesRule
type VmwareTanzuManageV1alpha1IamRoleKubernetesRule struct {

	// API group.
	APIGroups []string `json:"apiGroups"`

	// Non-resource urls for the role.
	NonResourceUrls []string `json:"nonResourceUrls"`

	// ResourceNames to restrict the rule to resources by name
	ResourceNames []string `json:"resourceNames"`

	// Resources - added a validation to input.
	Resources []string `json:"resources"`

	// Verbs.
	Verbs []string `json:"verbs"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRoleKubernetesRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRoleKubernetesRule) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1IamRoleKubernetesRule

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
