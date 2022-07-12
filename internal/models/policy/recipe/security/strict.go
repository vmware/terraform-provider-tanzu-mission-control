/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyrecipesecuritymodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict Input schema for security policy strict recipe version v1.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.spec.security.v1.Strict
type VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict struct {

	// Audit (dry-run).
	Audit bool `json:"audit,omitempty"`

	// Disable native pod security policy.
	DisableNativePsp bool `json:"disableNativePsp,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
