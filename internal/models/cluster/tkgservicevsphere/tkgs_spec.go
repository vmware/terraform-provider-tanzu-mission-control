// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgservicevspheremodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec The tkg service vsphere cluster spec.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgservicevsphere.Spec
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec struct {

	// VSphere specific distribution.
	Distribution *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereDistribution `json:"distribution,omitempty"`

	// VSphere related settings for workload cluster.
	Settings *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSettings `json:"settings,omitempty"`

	// Topology specific configuration.
	Topology *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology `json:"topology,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
