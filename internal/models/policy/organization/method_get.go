/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyorganizationmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse Response from getting a Policy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.policy.GetPolicyResponse
type VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse struct {

	// Policy returned.
	Policy *VmwareTanzuManageV1alpha1OrganizationPolicyPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationPolicyGetPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
