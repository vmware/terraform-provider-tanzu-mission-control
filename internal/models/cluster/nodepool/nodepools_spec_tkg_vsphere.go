// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package nodepool

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool TKGVsphereNodepool is the nodepool spec for TKG vsphere cluster.
// The values will flow via cluster:options api.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.TKGVsphereNodepool
type VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool struct {

	// VM specific configuration.
	VMConfig *VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig `json:"vmConfig,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
