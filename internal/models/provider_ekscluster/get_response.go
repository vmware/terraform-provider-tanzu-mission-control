/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterGetProviderEksClusterResponse Response from getting a ProviderEksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.manage.eks.providerekscluster.GetProviderEksClusterResponse
type VmwareTanzuManageV1alpha1ManageEksProvidereksclusterGetProviderEksClusterResponse struct {

	// ProviderEksCluster returned.
	ProviderEksCluster *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster `json:"providerEksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterGetProviderEksClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterGetProviderEksClusterResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManageEksProvidereksclusterGetProviderEksClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
