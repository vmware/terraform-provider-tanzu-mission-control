/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSpec The schedule spec.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.schedule.Spec.
type VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSpec struct {

	// Paused specifies whether the schedule is paused or not.
	Paused bool `json:"paused"`

	// Rate at which the backup is to be run.
	Schedule *VmwareTanzuManageV1alpha1CommonScheduleSchedule `json:"schedule,omitempty"`

	// The definition of the Backup to be run on the provided schedule.
	Template *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupSpec `json:"template,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
