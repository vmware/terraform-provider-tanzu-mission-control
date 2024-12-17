// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

// VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect The effect of the taint on pods
// that do not tolerate the taint.
// Valid effects are NoSchedule, NoExecute, PreferNoSchedule and EffectUnspecified.
//
//   - EFFECT_UNSPECIFIED: Unspecified effect.
//   - NO_SCHEDULE: Pods that do not tolerate this taint are not scheduled on the node.
//   - NO_EXECUTE: Pods are evicted from the node if are already running on the node.
//   - PREFER_NO_SCHEDULE: Avoids scheduling Pods that do not tolerate this taint onto the node.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.Taint.Effect
type VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect string

func NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(value VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect) *VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect.
func (m VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect) Pointer() *VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectEFFECTUNSPECIFIED captures enum value "EFFECT_UNSPECIFIED".
	VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectEFFECTUNSPECIFIED VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect = "EFFECT_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOSCHEDULE captures enum value "NO_SCHEDULE".
	VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOSCHEDULE VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect = "NO_SCHEDULE"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOEXECUTE captures enum value "NO_EXECUTE".
	VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOEXECUTE VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect = "NO_EXECUTE"

	// VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE captures enum value "PREFER_NO_SCHEDULE".
	VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect = "PREFER_NO_SCHEDULE"
)
