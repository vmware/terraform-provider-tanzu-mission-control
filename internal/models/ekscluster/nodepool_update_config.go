/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig Update config for the nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.UpdateConfig
type VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig struct {

	// Maximum number of nodes unavailable at once during a version update.
	MaxUnavailableNodes string `json:"maxUnavailableNodes,omitempty"`

	// Maximum percentage of nodes unavailable during a version update.
	MaxUnavailablePercentage string `json:"maxUnavailablePercentage,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
