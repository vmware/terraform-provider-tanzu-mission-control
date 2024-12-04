// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policyworkspacemodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest Request to create a Policy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.policy.CreatePolicyRequest
type VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest struct {

	// Policy to create.
	Policy *VmwareTanzuManageV1alpha1WorkspacePolicyPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspacePolicyPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse Response from creating a Policy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.policy.CreatePolicyResponse
type VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse struct {

	// Policy created.
	Policy *VmwareTanzuManageV1alpha1WorkspacePolicyPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspacePolicyPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
