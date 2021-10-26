/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepool

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNodepoolStatus Status of node pool resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.Status
type VmwareTanzuManageV1alpha1ClusterNodepoolStatus struct {

	// Conditions for the nodepool resource.
	Conditions map[string]VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Phase of the nodepool resource.
	Phase *VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase `json:"phase,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
