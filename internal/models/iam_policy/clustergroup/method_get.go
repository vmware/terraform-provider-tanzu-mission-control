/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupiammodel

import (
	"github.com/go-openapi/swag"

	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

// VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupIAMPolicyResponse GetClusterGroupIAMPolicy response message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.GetClusterGroupIAMPolicyResponse
type VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupIAMPolicyResponse struct {

	// ClusterGroup policy.
	PolicyList []*iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policyList"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
