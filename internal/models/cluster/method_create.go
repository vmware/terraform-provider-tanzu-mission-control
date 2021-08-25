/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustermodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterCreateClusterRequest Request to create a Cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.CreateClusterRequest
type VmwareTanzuManageV1alpha1ClusterCreateClusterRequest struct {
	// Cluster to create.
	Cluster *VmwareTanzuManageV1alpha1ClusterCluster `json:"cluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterCreateClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterCreateClusterRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterCreateClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterCreateClusterResponse Response from creating a Cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.CreateClusterResponse.
type VmwareTanzuManageV1alpha1ClusterCreateClusterResponse struct {

	// Cluster created.
	Cluster *VmwareTanzuManageV1alpha1ClusterCluster `json:"cluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterCreateClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterCreateClusterResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterCreateClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
