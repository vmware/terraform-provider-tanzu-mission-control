// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolMode The mode type options of the cluster nodepool
//
// - MODE_UNSPECIFIED: Unspecified mode.
//   - USER: User nodepools are primarily for hosting your application pods.
//   - SYSTEM: System agent pools are primarily for hosting critical system pods such as CoreDNS and metrics-server.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.Mode
type VmwareTanzuManageV1alpha1AksclusterNodepoolMode string

func NewVmwareTanzuManageV1alpha1AksclusterNodepoolMode(value VmwareTanzuManageV1alpha1AksclusterNodepoolMode) *VmwareTanzuManageV1alpha1AksclusterNodepoolMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterNodepoolMode.
func (m VmwareTanzuManageV1alpha1AksclusterNodepoolMode) Pointer() *VmwareTanzuManageV1alpha1AksclusterNodepoolMode {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterNodepoolModeMODEUNSPECIFIED captures enum value "MODE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1AksclusterNodepoolModeMODEUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterNodepoolMode = "MODE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolModeUSER captures enum value "USER".
	VmwareTanzuManageV1alpha1AksclusterNodepoolModeUSER VmwareTanzuManageV1alpha1AksclusterNodepoolMode = "USER"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolModeSYSTEM captures enum value "SYSTEM".
	VmwareTanzuManageV1alpha1AksclusterNodepoolModeSYSTEM VmwareTanzuManageV1alpha1AksclusterNodepoolMode = "SYSTEM"
)

// for schema.
var vmwareTanzuManageV1alpha1AksclusterNodepoolModeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterNodepoolMode
	if err := json.Unmarshal([]byte(`["MODE_UNSPECIFIED","USER","SYSTEM"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterNodepoolModeEnum = append(vmwareTanzuManageV1alpha1AksclusterNodepoolModeEnum, v)
	}
}
