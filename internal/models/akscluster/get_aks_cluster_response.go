// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse Response from getting an AksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.GetAksClusterResponse
type VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse struct {

	// AksCluster returned.
	AksCluster *VmwareTanzuManageV1alpha1AksCluster `json:"aksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
