// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1CommonBatchPhase Phase of the source resource application on its atomic targets.
// Note: The phase can move back to Pending from Applied when there are additions to the list of available atomic targets.
// In such a case, the system will automatically try to apply the changes to the new targets to get back to the Applied state.
//
//   - PHASE_UNSPECIFIED: UNSPECIFIED phase.
//   - PENDING: PENDING phase is set when source resource is currently being applied on at least one atomic target.
//   - APPLIED: APPLIED phase is set when source resource is successfully applied or skipped due to an override on all atomic targets.
//   - ERROR: ERROR phase is set when source resource has failed to apply on at-least one atomic target (not considering overrides).
//   - DELETING: DELETING phase is set when source resource is being deleted (only applicable on some source resource types).
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.batch.Phase
type VmwareTanzuManageV1alpha1CommonBatchPhase string

func NewVmwareTanzuManageV1alpha1CommonBatchPhase(value VmwareTanzuManageV1alpha1CommonBatchPhase) *VmwareTanzuManageV1alpha1CommonBatchPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1CommonBatchPhase.
func (m VmwareTanzuManageV1alpha1CommonBatchPhase) Pointer() *VmwareTanzuManageV1alpha1CommonBatchPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1CommonBatchPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1CommonBatchPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1CommonBatchPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1CommonBatchPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1CommonBatchPhasePENDING VmwareTanzuManageV1alpha1CommonBatchPhase = "PENDING"

	// VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED captures enum value "APPLIED".
	VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED VmwareTanzuManageV1alpha1CommonBatchPhase = "APPLIED"

	// VmwareTanzuManageV1alpha1CommonBatchPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1CommonBatchPhaseERROR VmwareTanzuManageV1alpha1CommonBatchPhase = "ERROR"

	// VmwareTanzuManageV1alpha1CommonBatchPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1CommonBatchPhaseDELETING VmwareTanzuManageV1alpha1CommonBatchPhase = "DELETING"
)

// for schema.
var vmwareTanzuManageV1alpha1CommonBatchPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1CommonBatchPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","APPLIED","ERROR","DELETING"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1CommonBatchPhaseEnum = append(vmwareTanzuManageV1alpha1CommonBatchPhaseEnum, v)
	}
}
