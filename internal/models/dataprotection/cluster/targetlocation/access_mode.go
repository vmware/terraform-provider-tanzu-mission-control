/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode The permissions for a BackupStorageLocation.
//
//   - READONLY: The read only access.
//   - READWRITE: Read and write access.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backuplocation.Status.BackupStorageLocationAccessMode.
type VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode string

func NewVmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode(value VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode) *VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode.
func (m VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode) Pointer() *VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeREADONLY captures enum value "READONLY".
	VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeREADONLY VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode = "READONLY"

	// VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeREADWRITE captures enum value "READWRITE".
	VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeREADWRITE VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode = "READWRITE"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessMode

	if err := json.Unmarshal([]byte(`["READONLY","READWRITE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeEnum = append(vmwareTanzuManageV1alpha1ClusterDataprotectionBackuplocationStatusBackupStorageLocationAccessModeEnum, v)
	}
}
