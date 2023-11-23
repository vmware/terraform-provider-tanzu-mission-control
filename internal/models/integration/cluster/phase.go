/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterintegrationmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterIntegrationPhase Integration Lifecycle Phase.
//
//   - PHASE_UNSPECIFIED: Unspecified  phase.
//   - CREATING: CREATING phase when process for adding integration to cluster is started.
//   - UPDATING: Updating phase when need to update configuration for the added integration.
//   - READY: READY phase when integration is added to cluster.
//   - ERROR: Error phase when there is any issue during addition/update/deletion of the integration.
//   - DELETING: DELETING phase when when process for removing integration to cluster is started.
//   - PENDING: PENDING phase when the process is waiting for changes in the cluster after addition/update of the integration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.integration.Phase
type VmwareTanzuManageV1alpha1ClusterIntegrationPhase string

func NewVmwareTanzuManageV1alpha1ClusterIntegrationPhase(value VmwareTanzuManageV1alpha1ClusterIntegrationPhase) *VmwareTanzuManageV1alpha1ClusterIntegrationPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterIntegrationPhase.
func (m VmwareTanzuManageV1alpha1ClusterIntegrationPhase) Pointer() *VmwareTanzuManageV1alpha1ClusterIntegrationPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterIntegrationPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterIntegrationPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterIntegrationPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterIntegrationPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1ClusterIntegrationPhaseCREATING VmwareTanzuManageV1alpha1ClusterIntegrationPhase = "CREATING"

	// VmwareTanzuManageV1alpha1ClusterIntegrationPhaseUPDATING captures enum value "UPDATING".
	VmwareTanzuManageV1alpha1ClusterIntegrationPhaseUPDATING VmwareTanzuManageV1alpha1ClusterIntegrationPhase = "UPDATING"

	// VmwareTanzuManageV1alpha1ClusterIntegrationPhaseREADY captures enum value "READY".
	VmwareTanzuManageV1alpha1ClusterIntegrationPhaseREADY VmwareTanzuManageV1alpha1ClusterIntegrationPhase = "READY"

	// VmwareTanzuManageV1alpha1ClusterIntegrationPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1ClusterIntegrationPhaseERROR VmwareTanzuManageV1alpha1ClusterIntegrationPhase = "ERROR"

	// VmwareTanzuManageV1alpha1ClusterIntegrationPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1ClusterIntegrationPhaseDELETING VmwareTanzuManageV1alpha1ClusterIntegrationPhase = "DELETING"

	// VmwareTanzuManageV1alpha1ClusterIntegrationPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1ClusterIntegrationPhasePENDING VmwareTanzuManageV1alpha1ClusterIntegrationPhase = "PENDING"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterIntegrationPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterIntegrationPhase

	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","CREATING","UPDATING","READY","ERROR","DELETING","PENDING"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterIntegrationPhaseEnum = append(vmwareTanzuManageV1alpha1ClusterIntegrationPhaseEnum, v)
	}
}
