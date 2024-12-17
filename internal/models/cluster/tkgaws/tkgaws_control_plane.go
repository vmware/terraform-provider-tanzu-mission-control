// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgawsmodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsControlPlane Control plane configuration for the AWS cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgaws.ControlPlane
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsControlPlane struct {

	// List of availability zones for the control plane nodes.
	AvailabilityZones []string `json:"availabilityZones"`

	// Flag which controls if the cluster needs to be highly available. A highly available cluster has three
	// controlplane machines, and a non highly available cluster has one.
	HighAvailability bool `json:"highAvailability,omitempty"`

	// Control plane instance type.
	InstanceType string `json:"instanceType,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsControlPlane) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsControlPlane) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsControlPlane
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
