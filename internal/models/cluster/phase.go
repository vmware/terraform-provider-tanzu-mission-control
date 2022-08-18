/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustermodel

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterPhase Phase of the cluster resource.
/*
  - PHASE_UNSPECIFIED: Unspecified phase.
  - PENDING: Resource is pending processing.
  - PROCESSING: Processing the resource.
  - CREATING: Resource is being created.
  - READY: Resource is ready state.
  - DELETING: Resource is being deleted.
  - ERROR: Error in processing.
  - DETACHING: Resource is being detached.
  - UPGRADING: An upgrade is in progress.
  - UPGRADE_FAILED: An upgrade has failed.

 swagger:model vmware.tanzu.manage.v1alpha1.cluster.Phase
*/
type VmwareTanzuManageV1alpha1ClusterPhase string

func NewVmwareTanzuManageV1alpha1ClusterPhase(value VmwareTanzuManageV1alpha1ClusterPhase) *VmwareTanzuManageV1alpha1ClusterPhase {
	v := value
	return &v
}

const (

	// VmwareTanzuManageV1alpha1ClusterPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1ClusterPhasePENDING VmwareTanzuManageV1alpha1ClusterPhase = "PENDING"

	// VmwareTanzuManageV1alpha1ClusterPhasePROCESSING captures enum value "PROCESSING".
	VmwareTanzuManageV1alpha1ClusterPhasePROCESSING VmwareTanzuManageV1alpha1ClusterPhase = "PROCESSING"

	// VmwareTanzuManageV1alpha1ClusterPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1ClusterPhaseCREATING VmwareTanzuManageV1alpha1ClusterPhase = "CREATING"

	// VmwareTanzuManageV1alpha1ClusterPhaseREADY captures enum value "READY".
	VmwareTanzuManageV1alpha1ClusterPhaseREADY VmwareTanzuManageV1alpha1ClusterPhase = "READY"

	// VmwareTanzuManageV1alpha1ClusterPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1ClusterPhaseDELETING VmwareTanzuManageV1alpha1ClusterPhase = "DELETING"

	// VmwareTanzuManageV1alpha1ClusterPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1ClusterPhaseERROR VmwareTanzuManageV1alpha1ClusterPhase = "ERROR"

	// VmwareTanzuManageV1alpha1ClusterPhaseDETACHING captures enum value "DETACHING".
	VmwareTanzuManageV1alpha1ClusterPhaseDETACHING VmwareTanzuManageV1alpha1ClusterPhase = "DETACHING"

	// VmwareTanzuManageV1alpha1ClusterPhaseUPGRADING captures enum value "UPGRADING".
	VmwareTanzuManageV1alpha1ClusterPhaseUPGRADING VmwareTanzuManageV1alpha1ClusterPhase = "UPGRADING"

	// VmwareTanzuManageV1alpha1ClusterPhaseUPGRADEFAILED captures enum value "UPGRADE_FAILED".
	VmwareTanzuManageV1alpha1ClusterPhaseUPGRADEFAILED VmwareTanzuManageV1alpha1ClusterPhase = "UPGRADE_FAILED"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","PROCESSING","CREATING","READY","DELETING","ERROR","DETACHING","UPGRADING","UPGRADE_FAILED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterPhaseEnum = append(vmwareTanzuManageV1alpha1ClusterPhaseEnum, v)
	}
}
