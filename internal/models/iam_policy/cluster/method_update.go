// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clusteriammodel

import (
	"github.com/go-openapi/swag"

	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

// VmwareTanzuManageV1alpha1ClusterUpdateClusterIAMPolicyResponse UpdateClusterIAMPolicy response message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.UpdateClusterIAMPolicyResponse
type VmwareTanzuManageV1alpha1ClusterUpdateClusterIAMPolicyResponse struct {

	// Cluster policy set.
	Policy *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterUpdateClusterIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterUpdateClusterIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterUpdateClusterIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
