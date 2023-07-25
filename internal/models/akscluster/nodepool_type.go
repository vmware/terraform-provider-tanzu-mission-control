/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolType The type options of nodepool.
//
//   - TYPE_UNSPECIFIED: Unspecified type.
//   - VIRTUAL_MACHINE_SCALE_SETS: Create a nodepool backed by a Virtual Machine Scale Set.
//   - AVAILABILITY_SET: Use of this is strongly discouraged.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.Type
type VmwareTanzuManageV1alpha1AksclusterNodepoolType string

func NewVmwareTanzuManageV1alpha1AksclusterNodepoolType(value VmwareTanzuManageV1alpha1AksclusterNodepoolType) *VmwareTanzuManageV1alpha1AksclusterNodepoolType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterNodepoolType.
func (m VmwareTanzuManageV1alpha1AksclusterNodepoolType) Pointer() *VmwareTanzuManageV1alpha1AksclusterNodepoolType {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterNodepoolTypeTYPEUNSPECIFIED captures enum value "TYPE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1AksclusterNodepoolTypeTYPEUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterNodepoolType = "TYPE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolTypeVIRTUALMACHINESCALESETS captures enum value "VIRTUAL_MACHINE_SCALE_SETS".
	VmwareTanzuManageV1alpha1AksclusterNodepoolTypeVIRTUALMACHINESCALESETS VmwareTanzuManageV1alpha1AksclusterNodepoolType = "VIRTUAL_MACHINE_SCALE_SETS"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolTypeAVAILABILITYSET captures enum value "AVAILABILITY_SET".
	VmwareTanzuManageV1alpha1AksclusterNodepoolTypeAVAILABILITYSET VmwareTanzuManageV1alpha1AksclusterNodepoolType = "AVAILABILITY_SET"
)

// for schema.
var vmwareTanzuManageV1alpha1AksclusterNodepoolTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterNodepoolType
	if err := json.Unmarshal([]byte(`["TYPE_UNSPECIFIED","VIRTUAL_MACHINE_SCALE_SETS","AVAILABILITY_SET"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterNodepoolTypeEnum = append(vmwareTanzuManageV1alpha1AksclusterNodepoolTypeEnum, v)
	}
}
