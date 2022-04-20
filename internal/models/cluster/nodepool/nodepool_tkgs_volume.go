/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepool

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume TKGServiceVsphereVolume defines a Persistent Volume Claim attached for the workload cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.TKGServiceVsphereVolume
type VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume struct {

	// Volume capacity is in gib.
	Capacity float32 `json:"capacity,omitempty"`

	// MountPath is the directory where the volume device is to be mounted.
	MountPath string `json:"mountPath,omitempty"`

	// Volume name.
	Name string `json:"name,omitempty"`

	// Storage class for PVC
	// If omitted, default storage class will be used for the disks.
	StorageClass string `json:"storageClass,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
