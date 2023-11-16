/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustergroupmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase Available phases for data protection object.
//
//   - PHASE_UNSPECIFIED: Phase_unspecified is the default phase.
//   - PENDING: Pending phase is set when the data protection object is being processed by the service (TMC).
//   - CREATING: Creating phase is set when data protection is being enabled on the cluster group.
//   - PENDING_DELETE: Pending delete is set when the data protection delete is being processed by the service.
//   - DELETING: Deleting the set when the data protection delete is in progress on the cluster group.
//   - READY: Ready phase is set when the data protection is successfully enabled.
//   - ERROR: Error phase is set when there was a failure while creating/deleting data protection.
//   - UPDATING: Updating is set when the data protection is being updated.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.Status.Phase.
type VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase string

func NewVmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase(value VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase) *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhase.
func (m VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase) Pointer() *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhasePENDING VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase = "PENDING"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseCREATING VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase = "CREATING"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhasePENDINGDELETE captures enum value "PENDING_DELETE".
	VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhasePENDINGDELETE VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase = "PENDING_DELETE"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseDELETING VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase = "DELETING"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseREADY captures enum value "READY".
	VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseREADY VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase = "READY"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseERROR VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase = "ERROR"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseUPDATING captures enum value "UPDATING".
	VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhaseUPDATING VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase = "UPDATING"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase

	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","CREATING","PENDING_DELETE","DELETING","READY","ERROR","UPDATING"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhaseEnum = append(vmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhaseEnum, v)
	}
}
