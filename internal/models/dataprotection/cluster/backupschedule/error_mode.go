/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode Specifies how Velero should behave if it encounters an error executing this hook.
//
//   - MODE_UNSPECIFIED: The default mode.
//   - CONTINUE: Means that an error from a hook is acceptable, and the operation can proceed.
//   - FAIL: Means that an error from a hook is problematic, and the operation should be in error.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.HookErrorMode.
type VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode string

func NewVmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode(value VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode) *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode.
func (m VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode) Pointer() *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorModeMODEUNSPECIFIED captures enum value "MODE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorModeMODEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode = "MODE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorModeCONTINUE captures enum value "CONTINUE".
	VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorModeCONTINUE VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode = "CONTINUE"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorModeFAIL captures enum value "FAIL".
	VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorModeFAIL VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode = "FAIL"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorModeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode

	if err := json.Unmarshal([]byte(`["MODE_UNSPECIFIED","CONTINUE","FAIL"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorModeEnum = append(vmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorModeEnum, v)
	}
}
