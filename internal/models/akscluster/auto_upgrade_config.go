/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterAutoUpgradeConfig The auto upgrade config.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.AutoUpgradeConfig
type VmwareTanzuManageV1alpha1AksclusterAutoUpgradeConfig struct {

	// The channel for the auto upgrade.
	Channel *VmwareTanzuManageV1alpha1AksclusterChannel `json:"channel,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAutoUpgradeConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAutoUpgradeConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterAutoUpgradeConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
