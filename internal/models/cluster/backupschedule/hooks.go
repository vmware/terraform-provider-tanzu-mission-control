/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupHooks BackupHooks contains custom actions that should be executed at different phases of the backup.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.BackupHooks.
type VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupHooks struct {

	// Resources are hooks that should be executed when backing up individual instances of a resource.
	Resources []*VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupResourceHookSpec `json:"resources"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupHooks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation..
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupHooks) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupHooks

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
