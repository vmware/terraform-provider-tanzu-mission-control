// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package nodepool

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig VM specific configuration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.TKGVsphereVMConfig
type VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig struct {

	// Number of CPUs per node.
	CPU string `json:"cpu,omitempty"`

	// Root disk size in gigabytes for the VM.
	DiskGib string `json:"diskGib,omitempty"`

	// Memory associated with the node in megabytes.
	MemoryMib string `json:"memoryMib,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
