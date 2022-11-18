/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest Request to create a Nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.CreateNodepoolRequest
type VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest struct {

	// Nodepool to create/update/get.
	Nodepool *VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse Response from creating a Nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.CreateNodepoolResponse
type VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse struct {

	// Nodepool created/updated/fetched.
	Nodepool *VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterNodepoolAPIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
