// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policyclustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupPolicyGetPolicyResponse Response from getting a Policy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.policy.GetPolicyResponse
type VmwareTanzuManageV1alpha1ClustergroupPolicyGetPolicyResponse struct {

	// Policy returned.
	Policy *VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPolicyGetPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPolicyGetPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupPolicyGetPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
