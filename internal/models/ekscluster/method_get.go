// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterGetEksClusterResponse Response from getting an EksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.GetEksClusterResponse
type VmwareTanzuManageV1alpha1EksclusterGetEksClusterResponse struct {

	// EksCluster returned.
	EksCluster *VmwareTanzuManageV1alpha1EksclusterEksCluster `json:"eksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterGetEksClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterGetEksClusterResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterGetEksClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
