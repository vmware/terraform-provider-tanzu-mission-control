/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupResponse Response from getting a ClusterGroup.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.GetClusterGroupResponse
type VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupResponse struct {

	// ClusterGroup returned.
	ClusterGroup *VmwareTanzuManageV1alpha1ClustergroupClusterGroup `json:"clusterGroup,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupGetClusterGroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
