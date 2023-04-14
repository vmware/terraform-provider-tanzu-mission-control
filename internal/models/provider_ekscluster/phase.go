/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase Phase of the cluster resource.
//
//   - PHASE_UNSPECIFIED: Unspecified phase.
//   - PENDING_UNMANAGE: This phase indicates the cluster is in the process of being unmanaged by TMC.
//   - PENDING_MANAGE: This phase indicates the cluster is in the process of being managed by TMC.
//   - UNMANAGED: This phase indicates the cluster is not managed by TMC.
//   - MANAGED: This phase indicates the cluster is managed by TMC.
//   - ERROR: Error in processing.
//
// swagger:model vmware.tanzu.manage.v1alpha1.manage.eks.providerekscluster.Phase
type VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase string

func NewVmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase(value VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase) *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase.
func (m VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase) Pointer() *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhasePENDINGUNMANAGE captures enum value "PENDING_UNMANAGE".
	VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhasePENDINGUNMANAGE VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase = "PENDING_UNMANAGE"

	// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhasePENDINGMANAGE captures enum value "PENDING_MANAGE".
	VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhasePENDINGMANAGE VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase = "PENDING_MANAGE"

	// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhaseUNMANAGED captures enum value "UNMANAGED".
	VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhaseUNMANAGED VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase = "UNMANAGED"

	// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhaseMANAGED captures enum value "MANAGED".
	VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhaseMANAGED VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase = "MANAGED"

	// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhaseERROR VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase = "ERROR"
)

// for schema.
var vmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING_UNMANAGE","PENDING_MANAGE","UNMANAGED","MANAGED","ERROR"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhaseEnum = append(vmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhaseEnum, v)
	}
}
