/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iammodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuCoreV1alpha1PolicyIAMPolicy Representation of an iam policy.
//
// swagger:model vmware.tanzu.core.v1alpha1.policy.IAMPolicy
type VmwareTanzuCoreV1alpha1PolicyIAMPolicy struct {

	// Metadata for this policy.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// List of role bindings associated with the policy.
	RoleBindings []*VmwareTanzuCoreV1alpha1PolicyRoleBinding `json:"roleBindings"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1PolicyIAMPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1PolicyIAMPolicy) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1PolicyIAMPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
