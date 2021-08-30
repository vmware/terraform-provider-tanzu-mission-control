/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupRequest Request to create a ClusterGroup.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.CreateClusterGroupRequest
type VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupRequest struct {

	// ClusterGroup to create.
	ClusterGroup *VmwareTanzuManageV1alpha1ClustergroupClusterGroup `json:"clusterGroup,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupResponse Response from creating a ClusterGroup.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.CreateClusterGroupResponse
type VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupResponse struct {

	// ClusterGroup created.
	ClusterGroup *VmwareTanzuManageV1alpha1ClustergroupClusterGroup `json:"clusterGroup,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupCreateClusterGroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
