// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

// Package models
// nolint: dupl
package models

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest Request to update (overwrite) an AksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.UpdateAksClusterRequest
type VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest struct {

	// Update AksCluster.
	AksCluster *VmwareTanzuManageV1alpha1AksCluster `json:"aksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse Response from updating an AksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.UpdateAksClusterResponse
type VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse struct {

	// AksCluster updated.
	AksCluster *VmwareTanzuManageV1alpha1AksCluster `json:"aksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
