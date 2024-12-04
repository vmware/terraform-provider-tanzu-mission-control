// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolUpgradeConfig Upgrade config for the nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.UpgradeConfig
type VmwareTanzuManageV1alpha1AksclusterNodepoolUpgradeConfig struct {

	// The maximum number or percentage of nodes that are surged during upgrade.
	// This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%').
	MaxSurge string `json:"maxSurge,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolUpgradeConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolUpgradeConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolUpgradeConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
