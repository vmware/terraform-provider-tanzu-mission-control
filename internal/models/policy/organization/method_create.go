// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policyorganizationmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest Request to create a Policy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.policy.CreatePolicyRequest
type VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest struct {

	// Policy to create.
	Policy *VmwareTanzuManageV1alpha1OrganizationPolicyPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationPolicyPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse Response from creating a Policy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.policy.CreatePolicyResponse
type VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse struct {

	// Policy created.
	Policy *VmwareTanzuManageV1alpha1OrganizationPolicyPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationPolicyPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
