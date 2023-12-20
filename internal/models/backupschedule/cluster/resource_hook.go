/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupResourceHookSpec BackupResourceHookSpec defines one or more BackupResourceHooks that should be executed based on
// the rules defined for namespaces and labels.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.BackupResourceHookSpec.
type VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupResourceHookSpec struct {

	// ExcludedNamespaces specifies the namespaces to which this hook spec does not apply.
	ExcludedNamespaces []string `json:"excludedNamespaces"`

	// IncludedNamespaces specifies the namespaces to which this hook spec applies. If empty, it applies
	// to all namespaces.
	IncludedNamespaces []string `json:"includedNamespaces"`

	// LabelSelector, if specified, filters the resources to which this hook spec applies.
	LabelSelector *K8sIoApimachineryPkgApisMetaV1LabelSelector `json:"labelSelector,omitempty"`

	// Name is the name of this hook.
	Name string `json:"name,omitempty"`

	// PostHooks is a list of BackupResourceHooks to execute after storing the item in the backup.
	// These are executed after all "additional items" from item actions are processed.
	PostHooks []*VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupResourceHook `json:"postHooks"`

	// PreHooks is a list of BackupResourceHooks to execute prior to storing the item in the backup.
	// These are executed before any "additional items" from item actions are processed.
	PreHooks []*VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupResourceHook `json:"preHooks"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupResourceHookSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupResourceHookSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupResourceHookSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
