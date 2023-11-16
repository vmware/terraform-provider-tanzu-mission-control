/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1CommonScheduleSchedule Holds the schedule options for scheduling a task.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.schedule.Schedule.
type VmwareTanzuManageV1alpha1CommonScheduleSchedule struct {

	// A Cron expression defining when to run a task.
	Rate string `json:"rate,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonScheduleSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonScheduleSchedule) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonScheduleSchedule

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
