// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterAddonsConfig The addons config.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.AddonsConfig
type VmwareTanzuManageV1alpha1AksclusterAddonsConfig struct {

	// The azure-keyvault-secrets-provider addon config.
	AzureKeyvaultSecretsProviderConfig *VmwareTanzuManageV1alpha1AksclusterAzureKeyvaultSecretsProviderAddonConfig `json:"azureKeyvaultSecretsProviderConfig,omitempty"`

	// The azure-policy addon config.
	AzurePolicyConfig *VmwareTanzuManageV1alpha1AksclusterAzurePolicyAddonConfig `json:"azurePolicyConfig,omitempty"`

	// The monitoring config.
	MonitoringConfig *VmwareTanzuManageV1alpha1AksclusterMonitoringAddonConfig `json:"monitoringConfig,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAddonsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAddonsConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterAddonsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
