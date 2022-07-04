/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusteriammodel

import (
	"github.com/go-openapi/swag"

	clustermodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster"
	iammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

// VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyRequest PatchClusterIAMPolicy request message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.PatchClusterIAMPolicyRequest
type VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyRequest struct {

	// Binding delta to be applied.
	BindingDeltaList []*iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta `json:"bindingDeltaList"`

	// Cluster full_name.
	FullName *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName `json:"fullName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyResponse PatchClusterIAMPolicy response message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.PatchClusterIAMPolicyResponse
type VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyResponse struct {

	// New policy object.
	Policy *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterPatchClusterIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
