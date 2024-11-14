// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policyworkspacemodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse Response from getting a Policy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.policy.GetPolicyResponse
type VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse struct {

	// Policy returned.
	Policy *VmwareTanzuManageV1alpha1WorkspacePolicyPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspacePolicyGetPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
