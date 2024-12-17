// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterAzureKeyvaultSecretsProviderAddonConfig The azure-keyvault-secrets-provider addon config.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.AzureKeyvaultSecretsProviderAddonConfig
type VmwareTanzuManageV1alpha1AksclusterAzureKeyvaultSecretsProviderAddonConfig struct {

	// Whether to enable secret rotation for the azure-keyvault-secrets-provider addon.
	EnableSecretRotation bool `json:"enableSecretRotation,omitempty"`

	// Whether the azure-keyvault-secrets-provider addon is enabled or not.
	Enabled bool `json:"enabled,omitempty"`

	// The interval at which to rotate the secrets in the azure-keyvault-secrets-provider addon.
	RotationPoolInterval string `json:"rotationPoolInterval,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAzureKeyvaultSecretsProviderAddonConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAzureKeyvaultSecretsProviderAddonConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterAzureKeyvaultSecretsProviderAddonConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
