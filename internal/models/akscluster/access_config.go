/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterAccessConfig Configs for the authentication and authorization.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.AccessConfig
type VmwareTanzuManageV1alpha1AksclusterAccessConfig struct {

	// Azure Active Directory config.
	AadConfig *VmwareTanzuManageV1alpha1AksclusterAADConfig `json:"aadConfig,omitempty"`

	// Disable the local accounts on the cluster when it's true.
	DisableLocalAccounts bool `json:"disableLocalAccounts,omitempty"`

	// Whether to enable kubernetes role-based access control.
	EnableRbac bool `json:"enableRbac,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAccessConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAccessConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterAccessConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
