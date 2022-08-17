/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepool

import "encoding/json"

// VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase Phase of the nodepool resource.
/*
  - PHASE_UNSPECIFIED: Unspecified phase.
  - CREATING: Resource is pending processing.
  - READY: Resource is in ready state.
  - ERROR: Error in processing.
  - DELETING: Resource is being deleted.
  - RESIZING: Resizing state.
  - UPGRADING: An upgrade is in progress.
  - UPGRADE_FAILED: An upgrade has failed.
  - WAITING: The cluster is not created yet. so wait till then.

 swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.Status.Phase
*/
type VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase string

func NewVmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase(value VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase) *VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase {
	v := value
	return &v
}

const (

	// VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseCREATING VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase = "CREATING"

	// VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseREADY captures enum value "READY".
	VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseREADY VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase = "READY"

	// VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseERROR VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase = "ERROR"

	// VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseDELETING VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase = "DELETING"

	// VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseRESIZING captures enum value "RESIZING".
	VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseRESIZING VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase = "RESIZING"

	// VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseUPGRADING captures enum value "UPGRADING".
	VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseUPGRADING VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase = "UPGRADING"

	// VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseUPGRADEFAILED captures enum value "UPGRADE_FAILED".
	VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseUPGRADEFAILED VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase = "UPGRADE_FAILED"

	// VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseWAITING captures enum value "WAITING".
	VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseWAITING VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase = "WAITING"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterNodepoolStatusPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","CREATING","READY","ERROR","DELETING","RESIZING","UPGRADING","UPGRADE_FAILED","WAITING"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseEnum = append(vmwareTanzuManageV1alpha1ClusterNodepoolStatusPhaseEnum, v)
	}
}
