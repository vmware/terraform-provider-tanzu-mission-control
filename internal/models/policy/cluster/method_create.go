/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest Request to create a Policy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.policy.CreatePolicyRequest
type VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest struct {

	// Policy to create.
	Policy *VmwareTanzuManageV1alpha1ClusterPolicyPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterPolicyPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse Response from creating a Policy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.policy.CreatePolicyResponse
type VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse struct {

	// Policy created.
	Policy *VmwareTanzuManageV1alpha1ClusterPolicyPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterPolicyPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
