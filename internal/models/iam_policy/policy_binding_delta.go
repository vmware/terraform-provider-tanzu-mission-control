/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iammodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuCoreV1alpha1PolicyBindingDelta Used for modify policy apis.
//
// swagger:model vmware.tanzu.core.v1alpha1.policy.BindingDelta
type VmwareTanzuCoreV1alpha1PolicyBindingDelta struct {

	// Type of operation.
	Op *VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType `json:"op,omitempty"`

	// Role for binding.
	Role string `json:"role,omitempty"`

	// Subject of rolebinding.
	Subject *VmwareTanzuCoreV1alpha1PolicySubject `json:"subject,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1PolicyBindingDelta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1PolicyBindingDelta) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1PolicyBindingDelta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
