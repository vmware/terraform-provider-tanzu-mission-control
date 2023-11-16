/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase The lifecycle phase of a schedule backup.
//
//   - PHASE_UNSPECIFIED: Phase_unspecified is the default phase.
//   - PENDING: Pending phase is set when the schedule object is being processed by the service (TMC).
//   - CREATING: Creating phase is set when schedule is being created on the cluster.
//   - NEW: The schedule has been created but not yet processed by velero.
//   - ENABLED: The schedule has been validated and will now be triggering backups according to the schedule spec.
//   - FAILEDVALIDATION: The schedule has failed the velero controller's validations and therefore will not trigger backups.
//   - PENDING_DELETE: Pending delete is set when the object deletion is being processed by the service.
//   - DELETING: The phase when schedule is being deleted.
//   - UPDATING: The phase when schedule is being updated.
//   - PAUSED: The phase when schedule is being paused.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.schedule.Status.Phase.
type VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase string

func NewVmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase(value VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase) *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase.
func (m VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase) Pointer() *VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhasePENDING VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase = "PENDING"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseCREATING VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase = "CREATING"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseNEW captures enum value "NEW".
	VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseNEW VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase = "NEW"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseENABLED captures enum value "ENABLED".
	VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseENABLED VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase = "ENABLED"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseFAILEDVALIDATION captures enum value "FAILEDVALIDATION".
	VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseFAILEDVALIDATION VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase = "FAILEDVALIDATION"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhasePENDINGDELETE captures enum value "PENDING_DELETE".
	VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhasePENDINGDELETE VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase = "PENDING_DELETE"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseDELETING VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase = "DELETING"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseUPDATING captures enum value "UPDATING".
	VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseUPDATING VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase = "UPDATING"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhasePAUSED captures enum value "PAUSED".
	VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhasePAUSED VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase = "PAUSED"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhase

	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","CREATING","NEW","ENABLED","FAILEDVALIDATION","PENDING_DELETE","DELETING","UPDATING","PAUSED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseEnum = append(vmwareTanzuManageV1alpha1ClusterDataprotectionScheduleStatusPhaseEnum, v)
	}
}
