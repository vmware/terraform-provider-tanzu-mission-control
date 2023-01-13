/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgservicevspheremodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereDistribution Distribution of the tkg service vsphere cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgservicevsphere.Distribution
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereDistribution struct {

	// Arch of the OS used for the cluster.
	OsArch string `json:"osArch,omitempty"`

	// Name of the OS used for the cluster.
	OsName string `json:"osName,omitempty"`

	// Version of the OS used for the cluster.
	OsVersion string `json:"osVersion,omitempty"`

	// Version of the cluster.
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereDistribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereDistribution) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereDistribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
