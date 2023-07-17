/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterDeleteAksClusterResponse Response from deleting an AksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.DeleteAksClusterResponse
type VmwareTanzuManageV1alpha1AksclusterDeleteAksClusterResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterDeleteAksClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterDeleteAksClusterResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterDeleteAksClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
