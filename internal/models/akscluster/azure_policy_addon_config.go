package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterAzurePolicyAddonConfig The azure-policy addon config.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.AzurePolicyAddonConfig
type VmwareTanzuManageV1alpha1AksclusterAzurePolicyAddonConfig struct {

	// Whether the azure-policy addon is enabled or not.
	Enabled bool `json:"enabled,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAzurePolicyAddonConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAzurePolicyAddonConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterAzurePolicyAddonConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
