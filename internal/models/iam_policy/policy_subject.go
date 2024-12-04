// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package iammodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuCoreV1alpha1PolicySubject Representation of a subject in resource manager.
//
// swagger:model vmware.tanzu.core.v1alpha1.policy.Subject
type VmwareTanzuCoreV1alpha1PolicySubject struct {

	// Subject type.
	Kind *VmwareTanzuCoreV1alpha1PolicySubjectKind `json:"kind,omitempty"`

	// Subject name - allow max characters for email - 320.
	Name string `json:"name,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1PolicySubject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1PolicySubject) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1PolicySubject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
