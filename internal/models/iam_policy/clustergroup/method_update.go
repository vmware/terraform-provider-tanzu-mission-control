/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupiammodel

import (
	"github.com/go-openapi/swag"

	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

// VmwareTanzuManageV1alpha1ClustergroupUpdateClusterGroupIAMPolicyResponse UpdateClusterGroupIAMPolicy response message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.UpdateClusterGroupIAMPolicyResponse
type VmwareTanzuManageV1alpha1ClustergroupUpdateClusterGroupIAMPolicyResponse struct {

	// ClusterGroup policy set.
	Policy *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupUpdateClusterGroupIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupUpdateClusterGroupIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupUpdateClusterGroupIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
