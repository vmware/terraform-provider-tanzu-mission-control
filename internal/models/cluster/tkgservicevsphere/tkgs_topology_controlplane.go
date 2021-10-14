/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgservicevspheremodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane VSphere specific control plane configuration for workload cluster object.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgservicevsphere.ControlPlane
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane struct {

	// Control plane instance type.
	Class string `json:"class,omitempty"`

	// High Availability or Non High Availability Cluster. HA cluster
	// creates three controlplane machines, and non HA creates just one.
	HighAvailability bool `json:"highAvailability,omitempty"`

	// Storage Class to be used for storage of the disks which store the root filesystems of the nodes.
	StorageClass string `json:"storageClass,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
