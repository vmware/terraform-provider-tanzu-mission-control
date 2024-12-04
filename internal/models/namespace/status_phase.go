// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package namespacemodel

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase The overall phase of the namespace.
/*
  - PHASE_UNSPECIFIED: Phase_unspecified is the default phase
  - CREATING: Creating phase is set when the namespace is being created.
  - ATTACHING: Attaching phase is set when the namespace is being attached.
  - UPDATING: Updating phase is set when the namespace is being updated.
  - READY: Ready phase is set when the namespace is successfully created/attached/updated.
  - ERROR: Error phase is set when there was a failure while creating/attaching/updating the namespace.

 swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.Status.Phase
*/
type VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase string

func NewVmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase(value VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase) *VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase {
	v := value
	return &v
}

const (

	// VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseCREATING VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase = "CREATING"

	// VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseATTACHING captures enum value "ATTACHING".
	VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseATTACHING VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase = "ATTACHING"

	// VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseUPDATING captures enum value "UPDATING".
	VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseUPDATING VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase = "UPDATING"

	// VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseREADY captures enum value "READY".
	VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseREADY VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase = "READY"

	// VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseERROR VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase = "ERROR"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","CREATING","ATTACHING","UPDATING","READY","ERROR"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseEnum = append(vmwareTanzuManageV1alpha1ClusterNamespaceStatusPhaseEnum, v)
	}
}
