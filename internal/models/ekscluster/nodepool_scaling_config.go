/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig Nodepool scaling config.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.ScalingConfig
type VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig struct {

	// Desired size of nodepool.
	DesiredSize int32 `json:"desiredSize,omitempty"`

	// Maximum size of nodepool.
	MaxSize int32 `json:"maxSize,omitempty"`

	// Minimum size of nodepool.
	MinSize int32 `json:"minSize,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
