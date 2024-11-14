// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackagerepository

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRequest Request to create a Repository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.repository.CreateRepositoryRequest.
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRequest struct {

	// Repository to create.
	Repository *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository `json:"repository,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryResponse Response from creating a Repository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.repository.CreateRepositoryResponse.
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryResponse struct {

	// Repository created.
	Repository *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository `json:"repository,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
