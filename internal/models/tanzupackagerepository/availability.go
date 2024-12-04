// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackagerepository

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityRequest The request type for enabling/disabling a Package Repository for a cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.repository.SetRepositoryAvailabilityRequest.
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityRequest struct {

	// If true, Package Repository is disabled for cluster.
	Disabled bool `json:"disabled"`

	// Package Repository full_name.
	FullName *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName `json:"fullName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse The response type for enabling/disabling a Package Repository for a cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.repository.SetRepositoryAvailabilityResponse.
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse struct {

	// Enabled Package Repository resource.
	Repository *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository `json:"repository,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
