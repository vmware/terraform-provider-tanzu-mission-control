/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionsmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets Targets is an enumeration of targets that CIS can run against.
//
//   - TARGETS_UNSPECIFIED: Unspecified Target refers to an unspecified target.
//   - LEADER_NODE: Target is the leader Node.
//   - NODE: Target is all nodes apart from the control plane nodes.
//   - ETCD: Target is the ETCD.
//   - CONTROL_PLANE: Target is the control plane components.
//   - POLICIES: Target is the policies on the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.CISSpec.Targets
type VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets string

func NewVmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets(value VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets) *VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets.
func (m VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets) Pointer() *VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsTARGETSUNSPECIFIED captures enum value "TARGETS_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsTARGETSUNSPECIFIED VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets = "TARGETS_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsLEADERNODE captures enum value "LEADER_NODE".
	VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsLEADERNODE VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets = "LEADER_NODE"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsNODE captures enum value "NODE".
	VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsNODE VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets = "NODE"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsETCD captures enum value "ETCD".
	VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsETCD VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets = "ETCD"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsCONTROLPLANE captures enum value "CONTROL_PLANE".
	VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsCONTROLPLANE VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets = "CONTROL_PLANE"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsPOLICIES captures enum value "POLICIES".
	VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsPOLICIES VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets = "POLICIES"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets

	if err := json.Unmarshal([]byte(`["TARGETS_UNSPECIFIED","LEADER_NODE","NODE","ETCD","CONTROL_PLANE","POLICIES"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsEnum = append(vmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargetsEnum, v)
	}
}
