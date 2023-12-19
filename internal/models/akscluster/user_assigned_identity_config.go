/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1AksclusterUserAssignedIdentityTypeConfig The managed identity config.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.UserAssignedIdentityTypeConfig

type VmwareTanzuManageV1alpha1AksclusterUserAssignedIdentityTypeConfig struct {
	ManagedResourceID string `json:"resourceId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterUserAssignedIdentityTypeConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterUserAssignedIdentityTypeConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterUserAssignedIdentityTypeConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
