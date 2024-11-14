// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig Auto scaling config for the nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.AutoScalingConfig
type VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig struct {

	// Whether to enable auto-scaler.
	Enabled bool `json:"enabled,omitempty"`

	// The maximum number of nodes for auto-scaling.
	MaxCount int32 `json:"maxCount,omitempty"`

	// The minimum number of nodes for auto-scaling.
	MinCount int32 `json:"minCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
