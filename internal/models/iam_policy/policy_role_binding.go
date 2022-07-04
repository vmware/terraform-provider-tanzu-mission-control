/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iammodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuCoreV1alpha1PolicyRoleBinding Representation of an iam role-binding in resource manager.
//
// swagger:model vmware.tanzu.core.v1alpha1.policy.RoleBinding
type VmwareTanzuCoreV1alpha1PolicyRoleBinding struct {

	// Role for this rolebinding -max length for role is 126.
	Role string `json:"role,omitempty"`

	// Subject of rolebinding.
	Subjects []*VmwareTanzuCoreV1alpha1PolicySubject `json:"subjects"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1PolicyRoleBinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1PolicyRoleBinding) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1PolicyRoleBinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
