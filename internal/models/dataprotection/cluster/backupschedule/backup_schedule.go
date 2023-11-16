/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSchedule A pre-scheduled or periodic Backup that should be run.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.schedule.Schedule.
type VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSchedule struct {

	// Full name for the Schedule.
	FullName *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleFullName `json:"fullName,omitempty"`

	// Metadata for the schedule object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the schedule.
	Spec *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSpec `json:"spec,omitempty"`

	// Status of the schedule.
	Status *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSchedule) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSchedule

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
