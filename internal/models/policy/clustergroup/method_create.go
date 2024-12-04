// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policyclustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest Request to create a Policy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.policy.CreatePolicyRequest
type VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest struct {

	// Policy to create.
	Policy *VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse Response from creating a Policy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.policy.CreatePolicyResponse
type VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse struct {

	// Policy created.
	Policy *VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupPolicyPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
