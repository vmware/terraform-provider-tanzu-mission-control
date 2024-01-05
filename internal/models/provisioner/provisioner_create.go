/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

// Code generated by go-swagger; DO NOT EDIT.

package provisionermodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerCreateProvisionerRequest Request to create a Provisioner.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.CreateProvisionerRequest
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerCreateProvisionerRequest struct {

	// Provisioner to create.
	Provisioner *VmwareTanzuManageV1alpha1ManagementclusterProvisionerProvisioner `json:"provisioner,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerCreateProvisionerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerCreateProvisionerRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerCreateProvisionerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}