/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustergroupmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpec The spec collects all the options for installing backup and restore solution into a cluster group.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.Spec.
type VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpec struct {
	// A flag to indicate whether to skip installation of restic server (https://github.com/restic/restic).
	// Otherwise, restic would be enabled by default as part of Data Protection installation.
	DisableRestic bool `json:"disableRestic"`

	// A flag to indicate whether to backup all the supported API Group versions of a resource on the cluster.
	EnableAllAPIGroupVersionsBackup bool `json:"enableAllApiGroupVersionsBackup"`

	// List of Backup Locations to install in the cluster.
	BackupLocationNames []string `json:"backup_location_names,omitempty"`

	// A flag to indicate whether to backup all the supported API Group versions of a resource on the cluster.
	EnableAllApiGroupVersionsBackup bool `json:"enable_all_api_group_versions_backup,omitempty"`

	// A flag to indicate whether to install CSI snapshotting related capabilities.
	EnableCsiSnapshots bool `json:"enableCsiSnapshots"`

	// A flag to indicate whether to install the node agent daemonset which is responsible for data transfer
	// to the target location.
	UseNodeAgent bool `json:"use_node_agent,omitempty"`

	// Selector to include/exclude specific clusters (optional).
	Selector *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelector `json:"selector,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
