/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgservicevspheremodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology Topology is the topology for tkg service vsphere cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgservicevsphere.Topology
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology struct {

	// Control plane specific configuration.
	ControlPlane *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane `json:"controlPlane,omitempty"`

	// Nodepool specific configuration.
	NodePools []*VmwareTanzuManageV1alpha1ClusterNodepoolDefinition `json:"nodePools"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
