// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package backupschedulemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleUpdateScheduleResponse Response from updating a Schedule.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.schedule.UpdateScheduleResponse.
type VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest struct {

	// Schedule updated.
	Schedule *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSchedule `json:"schedule,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleRequest

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleUpdateScheduleResponse Response from updating a Schedule.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.schedule.UpdateScheduleResponse.
type VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse struct {

	// Schedule updated.
	Schedule *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSchedule `json:"schedule,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse Response from listing Schedules.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.schedule.ListSchedulesResponse.
type VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse struct {

	// List of schedules.
	Schedules []*VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSchedule `json:"schedules"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleDeleteScheduleResponse Response from deleting a Schedule.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.schedule.DeleteScheduleResponse.
type VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleDeleteScheduleResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleDeleteScheduleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleDeleteScheduleResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleDeleteScheduleResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// !!! NOT GENERATED BY SWAGGER !!!.

type ListBackupSchedulesRequest struct {
	// Scope can be provider or cluster.
	SearchScope *ListBackupSchedulesSearchScope `json:"searchScope"`

	// Sort results by.
	SortBy string `json:"sortBy,omitempty"`

	// Query to run against the API.
	Query string `json:"query,omitempty"`

	// Include Total.
	IncludeTotalCount bool `json:"includeTotal"`
}

// MarshalBinary interface implementation.
func (m *ListBackupSchedulesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *ListBackupSchedulesRequest) UnmarshalBinary(b []byte) error {
	var res ListBackupSchedulesRequest

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
