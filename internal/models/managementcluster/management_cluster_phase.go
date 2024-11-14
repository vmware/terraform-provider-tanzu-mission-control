// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

// Code generated by go-swagger; DO NOT EDIT.

package managementclustermodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// VmwareTanzuManageV1alpha1ManagementclusterPhase Phase of a resource.
//
//   - PHASE_UNSPECIFIED: Unspecified phase.
//   - PENDING: Resource is pending processing.
//   - PROCESSING: Processing the resource.
//   - CREATING: Resource is being created.
//   - READY: Resource is ready state.
//   - DELETING: Resource is being deleted.
//   - ERROR: Error in processing.
//   - DETACHING: Resource is being detached.
//   - READY_TO_ATTACH: Resource is ready to be attached.
//   - ATTACH_COMPLETE: Attach resource has been applied on the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.Phase
type VmwareTanzuManageV1alpha1ManagementclusterPhase string

func NewVmwareTanzuManageV1alpha1ManagementclusterPhase(value VmwareTanzuManageV1alpha1ManagementclusterPhase) *VmwareTanzuManageV1alpha1ManagementclusterPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ManagementclusterPhase.
func (m VmwareTanzuManageV1alpha1ManagementclusterPhase) Pointer() *VmwareTanzuManageV1alpha1ManagementclusterPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ManagementclusterPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	VmwareTanzuManageV1alpha1ManagementclusterPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ManagementclusterPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ManagementclusterPhasePENDING captures enum value "PENDING"
	VmwareTanzuManageV1alpha1ManagementclusterPhasePENDING VmwareTanzuManageV1alpha1ManagementclusterPhase = "PENDING"

	// VmwareTanzuManageV1alpha1ManagementclusterPhasePROCESSING captures enum value "PROCESSING"
	VmwareTanzuManageV1alpha1ManagementclusterPhasePROCESSING VmwareTanzuManageV1alpha1ManagementclusterPhase = "PROCESSING"

	// VmwareTanzuManageV1alpha1ManagementclusterPhaseCREATING captures enum value "CREATING"
	VmwareTanzuManageV1alpha1ManagementclusterPhaseCREATING VmwareTanzuManageV1alpha1ManagementclusterPhase = "CREATING"

	// VmwareTanzuManageV1alpha1ManagementclusterPhaseREADY captures enum value "READY"
	VmwareTanzuManageV1alpha1ManagementclusterPhaseREADY VmwareTanzuManageV1alpha1ManagementclusterPhase = "READY"

	// VmwareTanzuManageV1alpha1ManagementclusterPhaseDELETING captures enum value "DELETING"
	VmwareTanzuManageV1alpha1ManagementclusterPhaseDELETING VmwareTanzuManageV1alpha1ManagementclusterPhase = "DELETING"

	// VmwareTanzuManageV1alpha1ManagementclusterPhaseERROR captures enum value "ERROR"
	VmwareTanzuManageV1alpha1ManagementclusterPhaseERROR VmwareTanzuManageV1alpha1ManagementclusterPhase = "ERROR"

	// VmwareTanzuManageV1alpha1ManagementclusterPhaseDETACHING captures enum value "DETACHING"
	VmwareTanzuManageV1alpha1ManagementclusterPhaseDETACHING VmwareTanzuManageV1alpha1ManagementclusterPhase = "DETACHING"

	// VmwareTanzuManageV1alpha1ManagementclusterPhaseREADYTOATTACH captures enum value "READY_TO_ATTACH"
	VmwareTanzuManageV1alpha1ManagementclusterPhaseREADYTOATTACH VmwareTanzuManageV1alpha1ManagementclusterPhase = "READY_TO_ATTACH"

	// VmwareTanzuManageV1alpha1ManagementclusterPhaseATTACHCOMPLETE captures enum value "ATTACH_COMPLETE"
	VmwareTanzuManageV1alpha1ManagementclusterPhaseATTACHCOMPLETE VmwareTanzuManageV1alpha1ManagementclusterPhase = "ATTACH_COMPLETE"
)

// for schema
var vmwareTanzuManageV1alpha1ManagementclusterPhaseEnum []interface{}
