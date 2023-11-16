/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionBackupExecHook ExecHook is a hook that uses the pod exec API to execute a command in a container in a pod.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.ExecHook.
type VmwareTanzuManageV1alpha1ClusterDataprotectionBackupExecHook struct {

	// Command is the command and arguments to execute.
	Command []string `json:"command"`

	// Container is the container in the pod where the command should be executed. If not specified,
	// the pod's first container is used.
	Container string `json:"container,omitempty"`

	// OnError specifies how Velero should behave if it encounters an error executing this hook.
	OnError *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorMode `json:"onError,omitempty"`

	// Timeout defines the maximum amount of time Velero should wait for the hook to complete before
	// considering the execution a failure.
	Timeout string `json:"timeout,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupExecHook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupExecHook) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionBackupExecHook

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
