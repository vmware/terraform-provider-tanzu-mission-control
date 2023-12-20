/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatus Status of the schedule resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.schedule.Status.
type VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatus struct {

	// A list of available phases for schedule object.
	AvailablePhases []*VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase `json:"availablePhases"`

	// The conditions attached to this backup object.
	// The description of the conditions is as follows:
	// - "Scheduled" with status 'Unknown' indicates the schedule request has not been applied to the cluster yet.
	// - "Scheduled" with status 'False' indicates the request could not be forwarded to the cluster (e.g. intent generation failure).
	// - "Scheduled" with status 'True' and "Ready" with status 'Unknown' indicates the schedule create / delete intent has been applied / deleted but not yet acted upon.
	// - "Ready" with status 'True' indicates the the creation of schedule is complete.
	// - "Ready" with status 'False' indicates the the creation of schedule is in error state.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// The last time a Backup was run for this schedule.
	// Format: date-time
	LastBackup strfmt.DateTime `json:"lastBackup,omitempty"`

	// The resource generation the current status applies to.
	ObservedGeneration string `json:"observedGeneration,omitempty"`

	// The current phase of the Schedule.
	Phase *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase `json:"phase,omitempty"`

	// Additional info about the phase.
	PhaseInfo string `json:"phaseInfo,omitempty"`

	// The list of all validation errors (if applicable).
	ValidationErrors []string `json:"validationErrors"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatus

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
