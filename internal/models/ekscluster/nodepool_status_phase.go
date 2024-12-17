// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

// VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase Phase of the nodepool resource.
//
//   - PHASE_UNSPECIFIED: Unspecified phase.
//   - CREATING: Resource is pending processing.
//   - READY: Resource is in ready state.
//   - ERROR: Error in processing.
//   - DELETING: Resource is being deleted.
//   - RESIZING: Resizing state.
//   - UPGRADING: An upgrade is in progress.
//   - UPGRADE_FAILED: An upgrade has failed.
//   - WAITING: The cluster is not created yet. so wait till then.
//   - UPDATING: A generic phase for nodepool update.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.Status.Phase
type VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase string

func NewVmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase(value VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase) *VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase.
func (m VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase) Pointer() *VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseCREATING VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase = "CREATING"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseREADY captures enum value "READY".
	VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseREADY VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase = "READY"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseERROR VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase = "ERROR"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseDELETING VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase = "DELETING"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseRESIZING captures enum value "RESIZING".
	VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseRESIZING VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase = "RESIZING"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseUPGRADING captures enum value "UPGRADING".
	VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseUPGRADING VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase = "UPGRADING"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseUPGRADEFAILED captures enum value "UPGRADE_FAILED".
	VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseUPGRADEFAILED VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase = "UPGRADE_FAILED"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseWAITING captures enum value "WAITING".
	VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseWAITING VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase = "WAITING"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseUPDATING captures enum value "UPDATING".
	VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseUPDATING VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhase = "UPDATING"
)
