/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustermodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionSpec The spec collects all the options for installing backup and restore solution into a Kubernetes cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.Spec.
type VmwareTanzuManageV1alpha1ClusterDataprotectionSpec struct {
	// A flag to indicate whether to skip installation of restic server (https://github.com/restic/restic).
	// Otherwise, restic would be enabled by default as part of Data Protection installation.
	DisableRestic bool `json:"disableRestic"`

	// A flag to indicate whether to backup all the supported API Group versions of a resource on the cluster.
	EnableAllAPIGroupVersionsBackup bool `json:"enableAllApiGroupVersionsBackup"`

	// A flag to indicate whether to install CSI snapshotting related capabilities.
	EnableCsiSnapshots bool `json:"enableCsiSnapshots"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
