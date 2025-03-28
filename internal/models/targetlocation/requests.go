// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package targetlocationmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest Request to create a BackupLocation.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.CreateBackupLocationRequest.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest struct {

	// BackupLocation to create.
	BackupLocation *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationBackupLocation `json:"backupLocation,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationRequest

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationCreateBackupLocationResponse Response from creating a BackupLocation.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.CreateBackupLocationResponse.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse struct {

	// BackupLocation created.
	BackupLocation *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationBackupLocation `json:"backupLocation,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationListBackupLocationsResponse Response from listing BackupLocations.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.ListBackupLocationsResponse.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationListBackupLocationsResponse struct {

	// List of backuplocations.
	BackupLocations []*VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationBackupLocation `json:"backupLocations"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationListBackupLocationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationListBackupLocationsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationListBackupLocationsResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// !!! NOT GENERATED BY SWAGGER !!!.

type ListBackupLocationsRequest struct {
	// Scope can be provider or cluster.
	SearchScope *ListBackupLocationsSearchScope `json:"searchScope"`

	// Sort results by.
	SortBy string `json:"sortBy,omitempty"`

	// Query to run against the API.
	Query string `json:"query,omitempty"`

	// Include Total.
	IncludeTotalCount bool `json:"includeTotal"`
}

// MarshalBinary interface implementation.
func (m *ListBackupLocationsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *ListBackupLocationsRequest) UnmarshalBinary(b []byte) error {
	var res ListBackupLocationsRequest

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
