/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionsmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase Phase describes the phase] of the inspection scan.
//
//   - PHASE_UNSPECIFIED: Unknown - to be used if status of the current inspection scan is unknown.
//   - RUNNING: Running - to indicate the inspection scan is currently running.
//   - PENDING: Pending - to indicate that the inspectionscan  is waiting to be started.
//   - COMPLETE: Complete - to indicate that the sonobuoy open source has completed the inspection scan.
//   - UPLOAD: Upload - to indicate that the inspection scan results are being uploaded to S3.
//   - FINISH: Finish - to indicate that the inspection has completed inspection + uploaded results to S3 successfully.
//   - STOP: Stop - to stop the sonobuoy inspection.
//   - ERROR: Error - to indicate that an error had occurred during the inspection.
//   - QUEUED: Queued - to indicate that the inspection is queued and waiting to be applied.
//   - CANCEL: CANCEL - to indicate that the inspection is canceled.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.Status.Phase
type VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase string

func NewVmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase(value VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase) *VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase.
func (m VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase) Pointer() *VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseRUNNING captures enum value "RUNNING".
	VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseRUNNING VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase = "RUNNING"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhasePENDING VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase = "PENDING"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseCOMPLETE captures enum value "COMPLETE".
	VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseCOMPLETE VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase = "COMPLETE"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseUPLOAD captures enum value "UPLOAD".
	VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseUPLOAD VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase = "UPLOAD"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseFINISH captures enum value "FINISH".
	VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseFINISH VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase = "FINISH"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseSTOP captures enum value "STOP".
	VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseSTOP VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase = "STOP"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseERROR VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase = "ERROR"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseQUEUED captures enum value "QUEUED".
	VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseQUEUED VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase = "QUEUED"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseCANCEL captures enum value "CANCEL".
	VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseCANCEL VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase = "CANCEL"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase

	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","RUNNING","PENDING","COMPLETE","UPLOAD","FINISH","STOP","ERROR","QUEUED","CANCEL"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseEnum = append(vmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhaseEnum, v)
	}
}
