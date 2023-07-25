/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterPhase Phase of the cluster resource.
//
//   - PHASE_UNSPECIFIED: Unspecified phase.
//   - PENDING: Resource is pending processing.
//   - CREATING: Resource is being created.
//   - READY: Resource is ready state.
//   - DELETING: Resource is being deleted.
//   - ERROR: Error in processing.
//   - UPDATING: This phase is used to reflect the UPDATING state of AKS cluster.
//   - OVER_LIMIT: This phase indicates cluster has crossed resource limits set for the organization.
//
// For such cluster we no longer sync data back to TMC.
//   - UPGRADING: This phase indicates kubernetes version is being upgraded for the cluster.
//   - STARTING: The AKS cluster is being started.
//   - STOPPING: The AKS cluster is being stopped.
//   - STOPPED: The AKS cluster is stopped.
//   - PENDING_MANAGE: This phase indicates the cluster is in the process of being managed by TMC.
//   - PENDING_UNMANAGE: This phase indicates the cluster is in the process of being unmanaged by TMC.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.Phase
type VmwareTanzuManageV1alpha1AksclusterPhase string

func NewVmwareTanzuManageV1alpha1AksclusterPhase(value VmwareTanzuManageV1alpha1AksclusterPhase) *VmwareTanzuManageV1alpha1AksclusterPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterPhase.
func (m VmwareTanzuManageV1alpha1AksclusterPhase) Pointer() *VmwareTanzuManageV1alpha1AksclusterPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1AksclusterPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1AksclusterPhasePENDING VmwareTanzuManageV1alpha1AksclusterPhase = "PENDING"

	// VmwareTanzuManageV1alpha1AksclusterPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1AksclusterPhaseCREATING VmwareTanzuManageV1alpha1AksclusterPhase = "CREATING"

	// VmwareTanzuManageV1alpha1AksclusterPhaseREADY captures enum value "READY".
	VmwareTanzuManageV1alpha1AksclusterPhaseREADY VmwareTanzuManageV1alpha1AksclusterPhase = "READY"

	// VmwareTanzuManageV1alpha1AksclusterPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1AksclusterPhaseDELETING VmwareTanzuManageV1alpha1AksclusterPhase = "DELETING"

	// VmwareTanzuManageV1alpha1AksclusterPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1AksclusterPhaseERROR VmwareTanzuManageV1alpha1AksclusterPhase = "ERROR"

	// VmwareTanzuManageV1alpha1AksclusterPhaseUPDATING captures enum value "UPDATING".
	VmwareTanzuManageV1alpha1AksclusterPhaseUPDATING VmwareTanzuManageV1alpha1AksclusterPhase = "UPDATING"

	// VmwareTanzuManageV1alpha1AksclusterPhaseOVERLIMIT captures enum value "OVER_LIMIT".
	VmwareTanzuManageV1alpha1AksclusterPhaseOVERLIMIT VmwareTanzuManageV1alpha1AksclusterPhase = "OVER_LIMIT"

	// VmwareTanzuManageV1alpha1AksclusterPhaseUPGRADING captures enum value "UPGRADING".
	VmwareTanzuManageV1alpha1AksclusterPhaseUPGRADING VmwareTanzuManageV1alpha1AksclusterPhase = "UPGRADING"

	// VmwareTanzuManageV1alpha1AksclusterPhaseSTARTING captures enum value "STARTING".
	VmwareTanzuManageV1alpha1AksclusterPhaseSTARTING VmwareTanzuManageV1alpha1AksclusterPhase = "STARTING"

	// VmwareTanzuManageV1alpha1AksclusterPhaseSTOPPING captures enum value "STOPPING".
	VmwareTanzuManageV1alpha1AksclusterPhaseSTOPPING VmwareTanzuManageV1alpha1AksclusterPhase = "STOPPING"

	// VmwareTanzuManageV1alpha1AksclusterPhaseSTOPPED captures enum value "STOPPED".
	VmwareTanzuManageV1alpha1AksclusterPhaseSTOPPED VmwareTanzuManageV1alpha1AksclusterPhase = "STOPPED"

	// VmwareTanzuManageV1alpha1AksclusterPhasePENDINGMANAGE captures enum value "PENDING_MANAGE".
	VmwareTanzuManageV1alpha1AksclusterPhasePENDINGMANAGE VmwareTanzuManageV1alpha1AksclusterPhase = "PENDING_MANAGE"

	// VmwareTanzuManageV1alpha1AksclusterPhasePENDINGUNMANAGE captures enum value "PENDING_UNMANAGE".
	VmwareTanzuManageV1alpha1AksclusterPhasePENDINGUNMANAGE VmwareTanzuManageV1alpha1AksclusterPhase = "PENDING_UNMANAGE"
)

// for schema.
var vmwareTanzuManageV1alpha1AksclusterPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","CREATING","READY","DELETING","ERROR","UPDATING","OVER_LIMIT","UPGRADING","STARTING","STOPPING","STOPPED","PENDING_MANAGE","PENDING_UNMANAGE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterPhaseEnum = append(vmwareTanzuManageV1alpha1AksclusterPhaseEnum, v)
	}
}
