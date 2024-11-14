// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

// Package models
// nolint: dupl
package models

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest Request to create an AksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.CreateAksClusterRequest
type VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest struct {

	// AksCluster to create.
	AksCluster *VmwareTanzuManageV1alpha1AksCluster `json:"aksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse Response from creating an AksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.CreateAksClusterResponse
type VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse struct {

	// AksCluster created.
	AksCluster *VmwareTanzuManageV1alpha1AksCluster `json:"aksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
