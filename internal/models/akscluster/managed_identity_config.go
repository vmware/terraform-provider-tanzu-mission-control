/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1AksclusterManagedIdentityConfig The managed identity config.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.ManagedIdentityConfig
type VmwareTanzuManageV1alpha1AksclusterManagedIdentityConfig struct {
	Type *VmwareTanzuManageV1alpha1AksclusterManagedIdentityType `json:"type,omitempty"`

	UserAssignedIdentityType *VmwareTanzuManageV1alpha1AksclusterUserAssignedIdentityTypeConfig `json:"userAssigned,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterManagedIdentityConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterManagedIdentityConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterManagedIdentityConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
