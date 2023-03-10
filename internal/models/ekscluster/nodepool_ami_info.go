/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo AMI info for the nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.AmiInfo
type VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo struct {

	// AMI id if custom AMI is specified. It cannot be used if launch template id is specified.
	AmiID string `json:"amiId,omitempty"`

	// Override bootstrap command for custom AMI.
	OverrideBootstrapCmd string `json:"overrideBootstrapCmd,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
