// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest Request to create an EksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.CreateEksClusterRequest
type VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest struct {

	// EksCluster to create.
	EksCluster *VmwareTanzuManageV1alpha1EksclusterEksCluster `json:"eksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse Response from creating an EksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.CreateEksClusterResponse
type VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse struct {

	// EksCluster created.
	EksCluster *VmwareTanzuManageV1alpha1EksclusterEksCluster `json:"eksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
