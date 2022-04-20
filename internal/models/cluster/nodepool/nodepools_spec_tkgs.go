/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepool

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool TKGServiceVsphereNodepool is the nodepool spec for TKG service vsphere cluster.
// The values will flow via cluster:options api.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.TKGServiceVsphereNodepool
type VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool struct {

	// Nodepool instance type.
	// The potential values could be found using cluster:options api.
	Class string `json:"class,omitempty"`

	// Storage Class to be used for storage of the disks which store the root filesystem of the nodes.
	// The potential values could be found using cluster:options api.
	StorageClass string `json:"storageClass,omitempty"`

	// Configure volumes for node pool nodes.
	Volumes []*VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume `json:"volumes"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
