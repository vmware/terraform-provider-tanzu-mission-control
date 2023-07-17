/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolResponse Response from updating a Nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.UpdateNodepoolResponse
type VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolResponse struct {

	// Nodepool updated.
	Nodepool *VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
