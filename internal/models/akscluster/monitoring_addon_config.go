/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterMonitoringAddonConfig The monitoring addon config.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.MonitoringAddonConfig
type VmwareTanzuManageV1alpha1AksclusterMonitoringAddonConfig struct {

	// Whether the monitoring addon is enabled or not.
	Enabled bool `json:"enabled,omitempty"`

	// The log analytics workspace id for the monitoring addon.
	LogAnalyticsWorkspaceID string `json:"logAnalyticsWorkspaceId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterMonitoringAddonConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterMonitoringAddonConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterMonitoringAddonConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
