// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustercommon

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig The advanced configuration for TKGm cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.AdvancedConfig
type VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig struct {

	// The key of the advanced configuration parameters.
	Key string `json:"key,omitempty"`

	// The value of the advanced configuration parameters.
	Value string `json:"value,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
