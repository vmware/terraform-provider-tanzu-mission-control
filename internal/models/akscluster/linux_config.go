// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterLinuxConfig The linux VMs config.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.LinuxConfig
type VmwareTanzuManageV1alpha1AksclusterLinuxConfig struct {

	// The administrator username to use for Linux VMs.
	AdminUsername string `json:"adminUsername,omitempty"`

	// Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.
	SSHKeys []string `json:"sshKeys"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterLinuxConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterLinuxConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterLinuxConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
