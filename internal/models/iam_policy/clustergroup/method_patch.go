// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustergroupiammodel

import (
	"github.com/go-openapi/swag"

	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

// VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyRequest PatchClusterGroupIAMPolicy request message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.PatchClusterGroupIAMPolicyRequest
type VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyRequest struct {

	// Binding delta to be applied.
	BindingDeltaList []*iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta `json:"bindingDeltaList"`

	// ClusterGroup full_name.
	FullName *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName `json:"fullName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyResponse PatchClusterGroupIAMPolicy response message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.PatchClusterGroupIAMPolicyResponse
type VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyResponse struct {

	// New policy object.
	Policy *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupPatchClusterGroupIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
