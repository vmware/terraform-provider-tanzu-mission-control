/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolRequest Request to create a Nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.CreateNodepoolRequest
type VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolRequest struct {

	// Nodepool to create.
	Nodepool *VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
