/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/
// Code generated by go-swagger; DO NOT EDIT.

package managementclusterregistration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterFullName FullName of the managementcluster. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.FullName
type VmwareTanzuManageV1alpha1ManagementclusterFullName struct {

	// Unique identifier of the ManagementCluster.
	Name string `json:"name,omitempty"`

	// ID of Organization. Generally a GUID
	OrgID string `json:"orgId,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ManagementclusterFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ManagementclusterFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func (m *VmwareTanzuManageV1alpha1ManagementclusterFullName) ToString() string {
	if m == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s", m.OrgID, m.Name)
}
