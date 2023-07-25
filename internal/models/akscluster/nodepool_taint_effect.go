/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect The effect of the taint on pods
// that do not tolerate the taint.
// Valid effects are NoSchedule, NoExecute, PreferNoSchedule and EffectUnspecified.
//
//   - EFFECT_UNSPECIFIED: Unspecified effect.
//   - NO_SCHEDULE: Pods that do not tolerate this taint are not scheduled on the node.
//   - NO_EXECUTE: Pods are evicted from the node if are already running on the node.
//   - PREFER_NO_SCHEDULE: Avoids scheduling Pods that do not tolerate this taint onto the node.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.Taint.Effect
type VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect string

func NewVmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect(value VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect) *VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect.
func (m VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect) Pointer() *VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectEFFECTUNSPECIFIED captures enum value "EFFECT_UNSPECIFIED".
	VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectEFFECTUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect = "EFFECT_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectNOSCHEDULE captures enum value "NO_SCHEDULE".
	VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectNOSCHEDULE VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect = "NO_SCHEDULE"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectNOEXECUTE captures enum value "NO_EXECUTE".
	VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectNOEXECUTE VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect = "NO_EXECUTE"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectPREFERNOSCHEDULE captures enum value "PREFER_NO_SCHEDULE".
	VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectPREFERNOSCHEDULE VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect = "PREFER_NO_SCHEDULE"
)

// for schema.
var vmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect
	if err := json.Unmarshal([]byte(`["EFFECT_UNSPECIFIED","NO_SCHEDULE","NO_EXECUTE","PREFER_NO_SCHEDULE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectEnum = append(vmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectEnum, v)
	}
}
