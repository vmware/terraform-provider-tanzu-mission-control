package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolPhase Phase of the nodepool resource.
//
//   - PHASE_UNSPECIFIED: Unspecified phase.
//   - PENDING: Resource is pending processing.
//   - CREATING: Resource is creating processing.
//   - READY: Resource is in ready state.
//   - ERROR: Error in processing.
//   - DELETING: Resource is being deleted.
//   - RESIZING: Resizing state.
//   - UPGRADING: An upgrade is in progress.
//   - UPGRADE_FAILED: An upgrade has failed.
//   - WAITING: Resource is not created yet. so wait till then.
//   - UPDATING: A generic phase for nodepool update.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.Phase
type VmwareTanzuManageV1alpha1AksclusterNodepoolPhase string

func NewVmwareTanzuManageV1alpha1AksclusterNodepoolPhase(value VmwareTanzuManageV1alpha1AksclusterNodepoolPhase) *VmwareTanzuManageV1alpha1AksclusterNodepoolPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterNodepoolPhase.
func (m VmwareTanzuManageV1alpha1AksclusterNodepoolPhase) Pointer() *VmwareTanzuManageV1alpha1AksclusterNodepoolPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterNodepoolPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	VmwareTanzuManageV1alpha1AksclusterNodepoolPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterNodepoolPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolPhasePENDING captures enum value "PENDING"
	VmwareTanzuManageV1alpha1AksclusterNodepoolPhasePENDING VmwareTanzuManageV1alpha1AksclusterNodepoolPhase = "PENDING"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseCREATING captures enum value "CREATING"
	VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseCREATING VmwareTanzuManageV1alpha1AksclusterNodepoolPhase = "CREATING"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseREADY captures enum value "READY"
	VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseREADY VmwareTanzuManageV1alpha1AksclusterNodepoolPhase = "READY"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseERROR captures enum value "ERROR"
	VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseERROR VmwareTanzuManageV1alpha1AksclusterNodepoolPhase = "ERROR"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseDELETING captures enum value "DELETING"
	VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseDELETING VmwareTanzuManageV1alpha1AksclusterNodepoolPhase = "DELETING"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseRESIZING captures enum value "RESIZING"
	VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseRESIZING VmwareTanzuManageV1alpha1AksclusterNodepoolPhase = "RESIZING"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseUPGRADING captures enum value "UPGRADING"
	VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseUPGRADING VmwareTanzuManageV1alpha1AksclusterNodepoolPhase = "UPGRADING"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseUPGRADEFAILED captures enum value "UPGRADE_FAILED"
	VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseUPGRADEFAILED VmwareTanzuManageV1alpha1AksclusterNodepoolPhase = "UPGRADE_FAILED"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseWAITING captures enum value "WAITING"
	VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseWAITING VmwareTanzuManageV1alpha1AksclusterNodepoolPhase = "WAITING"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseUPDATING captures enum value "UPDATING"
	VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseUPDATING VmwareTanzuManageV1alpha1AksclusterNodepoolPhase = "UPDATING"
)

// for schema
var vmwareTanzuManageV1alpha1AksclusterNodepoolPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterNodepoolPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","CREATING","READY","ERROR","DELETING","RESIZING","UPGRADING","UPGRADE_FAILED","WAITING","UPDATING"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterNodepoolPhaseEnum = append(vmwareTanzuManageV1alpha1AksclusterNodepoolPhaseEnum, v)
	}
}
