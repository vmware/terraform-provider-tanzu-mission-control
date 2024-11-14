// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package extension

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterExtensionPhase Extension lifecycle phase.
//
//   - PHASE_UNSPECIFIED: Unspecified  phase.
//   - ROLLING_BACK: Rolling back phase.
//   - ROLLED_BACK: Rolled back phase.
//   - PROCESSING: Processing phase.
//   - PROCESSED: Processed phase.
//   - FAILED: Failed phase.
//   - PENDING: Pending phase.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.extension.Phase
type VmwareTanzuManageV1alpha1ClusterExtensionPhase string

func NewVmwareTanzuManageV1alpha1ClusterExtensionPhase(value VmwareTanzuManageV1alpha1ClusterExtensionPhase) *VmwareTanzuManageV1alpha1ClusterExtensionPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterExtensionPhase.
func (m VmwareTanzuManageV1alpha1ClusterExtensionPhase) Pointer() *VmwareTanzuManageV1alpha1ClusterExtensionPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterExtensionPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterExtensionPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterExtensionPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterExtensionPhaseROLLINGBACK captures enum value "ROLLING_BACK".
	VmwareTanzuManageV1alpha1ClusterExtensionPhaseROLLINGBACK VmwareTanzuManageV1alpha1ClusterExtensionPhase = "ROLLING_BACK"

	// VmwareTanzuManageV1alpha1ClusterExtensionPhaseROLLEDBACK captures enum value "ROLLED_BACK".
	VmwareTanzuManageV1alpha1ClusterExtensionPhaseROLLEDBACK VmwareTanzuManageV1alpha1ClusterExtensionPhase = "ROLLED_BACK"

	// VmwareTanzuManageV1alpha1ClusterExtensionPhasePROCESSING captures enum value "PROCESSING".
	VmwareTanzuManageV1alpha1ClusterExtensionPhasePROCESSING VmwareTanzuManageV1alpha1ClusterExtensionPhase = "PROCESSING"

	// VmwareTanzuManageV1alpha1ClusterExtensionPhasePROCESSED captures enum value "PROCESSED".
	VmwareTanzuManageV1alpha1ClusterExtensionPhasePROCESSED VmwareTanzuManageV1alpha1ClusterExtensionPhase = "PROCESSED"

	// VmwareTanzuManageV1alpha1ClusterExtensionPhaseFAILED captures enum value "FAILED".
	VmwareTanzuManageV1alpha1ClusterExtensionPhaseFAILED VmwareTanzuManageV1alpha1ClusterExtensionPhase = "FAILED"

	// VmwareTanzuManageV1alpha1ClusterExtensionPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1ClusterExtensionPhasePENDING VmwareTanzuManageV1alpha1ClusterExtensionPhase = "PENDING"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterExtensionPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterExtensionPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","ROLLING_BACK","ROLLED_BACK","PROCESSING","PROCESSED","FAILED","PENDING"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterExtensionPhaseEnum = append(vmwareTanzuManageV1alpha1ClusterExtensionPhaseEnum, v)
	}
}
