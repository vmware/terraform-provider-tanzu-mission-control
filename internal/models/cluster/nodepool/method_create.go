/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepool

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest Request to create a Nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.CreateNodepoolRequest
type VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest struct {

	// Nodepool to create.
	Nodepool *VmwareTanzuManageV1alpha1ClusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse Response from creating a Nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.CreateNodepoolResponse
type VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse struct {

	// Nodepool created.
	Nodepool *VmwareTanzuManageV1alpha1ClusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
