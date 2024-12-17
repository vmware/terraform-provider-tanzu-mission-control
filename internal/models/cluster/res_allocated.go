// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustermodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1CommonClusterResourceAllocation ResourceAllocation is used for CPU and Memory metrics of a cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.ResourceAllocation
type VmwareTanzuManageV1alpha1CommonClusterResourceAllocation struct {

	// Allocatable is the quantity of compute resources that can be allocated by the user excluding reserved resources.
	Allocatable float32 `json:"allocatable,omitempty"`

	// Represents allocated percentage.
	AllocatedPercentage float32 `json:"allocatedPercentage,omitempty"`

	// Capacity is the total quantity of compute resources available including reserved resources.
	Capacity float32 `json:"capacity,omitempty"`

	// Requested is the requested quantity of compute resources.
	Requested float32 `json:"requested,omitempty"`

	// Units is the unit on which resource can be measured e.g. mb, millicores etc.
	Units string `json:"units,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterResourceAllocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterResourceAllocation) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonClusterResourceAllocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
