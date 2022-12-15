/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterNodepoolStatus Status of node pool resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.Status
type VmwareTanzuManageV1alpha1EksclusterNodepoolStatus struct {

	// Conditions for the nodepool resource.
	Conditions map[string]VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Phase of the nodepool resource.
	Phase *VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase `json:"phase,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterNodepoolStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
