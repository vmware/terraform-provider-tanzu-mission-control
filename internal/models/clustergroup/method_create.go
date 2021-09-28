/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterGroupRequest Request to create a ClusterGroup.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.CreateClusterGroupRequest
type VmwareTanzuManageV1alpha1ClusterGroupRequest struct {

	// ClusterGroup to create.
	ClusterGroup *VmwareTanzuManageV1alpha1ClustergroupClusterGroup `json:"clusterGroup,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterGroupResponse Response from creating a ClusterGroup.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.CreateClusterGroupResponse
type VmwareTanzuManageV1alpha1ClusterGroupResponse struct {

	// ClusterGroup created.
	ClusterGroup *VmwareTanzuManageV1alpha1ClustergroupClusterGroup `json:"clusterGroup,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
