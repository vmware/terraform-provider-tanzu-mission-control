/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority ScaleSetPriority of the auto scaling config.
//
//   - SCALE_SET_PRIORITY_UNSPECIFIED: Unspecified scale set priority.
//   - REGULAR: Regular VMs will be used.
//   - SPOT: Spot priority VMs will be used. There is no SLA for spot nodes.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.ScaleSetPriority
type VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority string

func NewVmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority(value VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority) *VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority.
func (m VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority) Pointer() *VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPrioritySCALESETPRIORITYUNSPECIFIED captures enum value "SCALE_SET_PRIORITY_UNSPECIFIED".
	VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPrioritySCALESETPRIORITYUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority = "SCALE_SET_PRIORITY_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriorityREGULAR captures enum value "REGULAR".
	VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriorityREGULAR VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority = "REGULAR"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPrioritySPOT captures enum value "SPOT".
	VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPrioritySPOT VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority = "SPOT"
)

// for schema.
var vmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriorityEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority
	if err := json.Unmarshal([]byte(`["SCALE_SET_PRIORITY_UNSPECIFIED","REGULAR","SPOT"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriorityEnum = append(vmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriorityEnum, v)
	}
}
