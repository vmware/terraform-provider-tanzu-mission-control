/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterRequest Request to update (overwrite) a ProviderEksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.manage.eks.providerekscluster.UpdateProviderEksClusterRequest
type VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterRequest struct {

	// Update ProviderEksCluster.
	ProviderEksCluster *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster `json:"providerEksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterResponse Response from updating a ProviderEksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.manage.eks.providerekscluster.UpdateProviderEksClusterResponse
type VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterResponse struct {

	// ProviderEksCluster updated.
	ProviderEksCluster *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster `json:"providerEksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
