/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionBackupSpec The backup spec.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.Spec.
type VmwareTanzuManageV1alpha1ClusterDataprotectionBackupSpec struct {

	// Specifies the time used to wait for CSI VolumeSnapshot status turns to ReadyToUse
	// during creation, before returning error as timeout. The default value is 10 minute.
	CsiSnapshotTimeout string `json:"csiSnapshotTimeout,omitempty"`

	// Specifies whether all pod volumes should be backed up via file system backup by default.
	DefaultVolumesToFsBackup bool `json:"defaultVolumesToFsBackup"`

	// Specifies whether restic should be used to take a backup of all pod volumes by default.
	// Deprecated - use default_volumes_to_fs_backup instead.
	DefaultVolumesToRestic bool `json:"defaultVolumesToRestic"`

	// The namespaces to be excluded in the backup.
	ExcludedNamespaces []string `json:"excludedNamespaces"`

	// The name list for the resources to be excluded in backup.
	ExcludedResources []string `json:"excludedResources"`

	// Hooks represent custom actions that should be executed at different phases of the backup.
	Hooks *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackupHooks `json:"hooks,omitempty"`

	// A flag which specifies whether cluster-scoped resources should be included for consideration in the backup.
	// If set to true, all cluster-scoped resources will be backed up. If set to false, all cluster-scoped resources.
	// will be excluded from the backup. If unset, all cluster-scoped resources are included if and only if all.
	// namespaces are included and there are no excluded namespaces.
	// Otherwise, only cluster-scoped resources associated with namespace-scoped resources.
	// included in the backup spec are backed up. For example, if a PersistentVolumeClaim is included in the backup,
	// its associated PersistentVolume (which is cluster-scoped) would also be backed up.
	IncludeClusterResources bool `json:"includeClusterResources"`

	// The namespace to be included for backup from. If empty, all namespaces are included.
	IncludedNamespaces []string `json:"includedNamespaces"`

	// The name list for the resources to be included into backup. If empty, all resources are included.
	IncludedResources []string `json:"includedResources"`

	// The label selector to selectively adding individual objects to the backup. If empty.
	// or nil, all objects are included. Optional.
	LabelSelector *K8sIoApimachineryPkgApisMetaV1LabelSelector `json:"labelSelector,omitempty"`

	// A list of metav1.LabelSelector to filter with when adding individual objects to the backup.
	// If multiple provided they will be joined by the OR operator. LabelSelector as well as
	// OrLabelSelectors cannot co-exist in backup request, only one of them can be used.
	OrLabelSelectors []*K8sIoApimachineryPkgApisMetaV1LabelSelector `json:"orLabelSelectors"`

	// Specifies the backup order of resources of specific Kind. The map key is the Kind name and.
	// value is a list of resource names separated by commas. Each resource name has format "namespace/resourcename".
	// For cluster resources, simply use "resourcename".
	OrderedResources map[string]string `json:"orderedResources,omitempty"`

	// A flag which specifies whether to take cloud snapshots of any PV's referenced in the set of objects.
	// included in the Backup. If set to true, snapshots will be taken. If set to false, snapshots will be skipped.
	// If left unset, snapshots will be attempted if volume snapshots are configured for the cluster.
	SnapshotVolumes bool `json:"snapshotVolumes"`

	// The name of a BackupStorageLocation where the backup should be stored.
	StorageLocation string `json:"storageLocation,omitempty"`

	// The backup retention period.
	TTL string `json:"ttl,omitempty"`

	// A list containing names of VolumeSnapshotLocations associated with this backup.
	VolumeSnapshotLocations []string `json:"volumeSnapshotLocations"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionBackupSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionBackupSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
