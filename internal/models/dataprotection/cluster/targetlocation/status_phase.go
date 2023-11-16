/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase Available phases for backup location object.
//
//   - PHASE_UNSPECIFIED: Phase_unspecified is the default phase.
//   - PENDING: Pending phase is set when the backup location object is being processed by the service (TMC).
//   - CREATING: Creating phase is set when backup location is being created by the service.
//   - PENDING_DELETE: Pending delete is set when the backup location delete is being processed by the service.
//   - DELETING: Deleting the set when the backup location delete is in progress.
//   - READY: Ready phase is set when the backup location is successfully created.
//   - ERROR: Error phase is set when there was a failure while creating/deleting backup location.
//   - UPDATING: Updating the set when the backup location is being updated.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.Status.Phase.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase string

func NewVmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase(value VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase) *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase.
func (m VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase) Pointer() *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhasePENDING VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase = "PENDING"

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseCREATING VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase = "CREATING"

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhasePENDINGDELETE captures enum value "PENDING_DELETE".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhasePENDINGDELETE VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase = "PENDING_DELETE"

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseDELETING VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase = "DELETING"

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseREADY captures enum value "READY".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseREADY VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase = "READY"

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseERROR VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase = "ERROR"

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseUPDATING captures enum value "UPDATING".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseUPDATING VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase = "UPDATING"
)

// for schema.
var vmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase

	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","CREATING","PENDING_DELETE","DELETING","READY","ERROR","UPDATING"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseEnum = append(vmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhaseEnum, v)
	}
}
