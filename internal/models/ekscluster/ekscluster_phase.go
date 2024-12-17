// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

// VmwareTanzuManageV1alpha1EksclusterPhase Phase of the cluster resource.
//
//   - PHASE_UNSPECIFIED: Unspecified phase.
//   - PENDING: Resource is pending processing.
//   - CREATING: Resource is being created.
//   - READY: Resource is ready state.
//   - DELETING: Resource is being deleted.
//   - ERROR: Error in processing.
//   - UPDATING: This phase is used to reflect the UPDATING state of EKS cluster.
//   - OVER_LIMIT: This phase indicates cluster has crossed resource limits set for the organization.
//
// For such cluster we no longer sync data back to TMC.
//   - UPGRADING: This phase indicates kubernetes version is being upgraded for the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.Phase
type VmwareTanzuManageV1alpha1EksclusterPhase string

func NewVmwareTanzuManageV1alpha1EksclusterPhase(value VmwareTanzuManageV1alpha1EksclusterPhase) *VmwareTanzuManageV1alpha1EksclusterPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1EksclusterPhase.
func (m VmwareTanzuManageV1alpha1EksclusterPhase) Pointer() *VmwareTanzuManageV1alpha1EksclusterPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1EksclusterPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1EksclusterPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1EksclusterPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1EksclusterPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1EksclusterPhasePENDING VmwareTanzuManageV1alpha1EksclusterPhase = "PENDING"

	// VmwareTanzuManageV1alpha1EksclusterPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1EksclusterPhaseCREATING VmwareTanzuManageV1alpha1EksclusterPhase = "CREATING"

	// VmwareTanzuManageV1alpha1EksclusterPhaseREADY captures enum value "READY".
	VmwareTanzuManageV1alpha1EksclusterPhaseREADY VmwareTanzuManageV1alpha1EksclusterPhase = "READY"

	// VmwareTanzuManageV1alpha1EksclusterPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1EksclusterPhaseDELETING VmwareTanzuManageV1alpha1EksclusterPhase = "DELETING"

	// VmwareTanzuManageV1alpha1EksclusterPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1EksclusterPhaseERROR VmwareTanzuManageV1alpha1EksclusterPhase = "ERROR"

	// VmwareTanzuManageV1alpha1EksclusterPhaseUPDATING captures enum value "UPDATING".
	VmwareTanzuManageV1alpha1EksclusterPhaseUPDATING VmwareTanzuManageV1alpha1EksclusterPhase = "UPDATING"

	// VmwareTanzuManageV1alpha1EksclusterPhaseOVERLIMIT captures enum value "OVER_LIMIT".
	VmwareTanzuManageV1alpha1EksclusterPhaseOVERLIMIT VmwareTanzuManageV1alpha1EksclusterPhase = "OVER_LIMIT"

	// VmwareTanzuManageV1alpha1EksclusterPhaseUPGRADING captures enum value "UPGRADING".
	VmwareTanzuManageV1alpha1EksclusterPhaseUPGRADING VmwareTanzuManageV1alpha1EksclusterPhase = "UPGRADING"
)
