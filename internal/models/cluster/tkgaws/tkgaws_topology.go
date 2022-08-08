/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgawsmodel

import (
	"github.com/go-openapi/swag"

	nodepoolmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology Topology is the topology definition for the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgaws.Topology
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology struct {

	// Control plane specific configuration.
	ControlPlane *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsControlPlane `json:"controlPlane,omitempty"`

	// Nodepool specific configuration.
	NodePools []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition `json:"nodePools"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
