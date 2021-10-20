/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgvspheremodel

import (
	"github.com/go-openapi/swag"

	nodepoolmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology Topology specific configuration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgvsphere.Topology
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology struct {

	// Control plane specific configuration.
	ControlPlane *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereControlPlane `json:"controlPlane,omitempty"`

	// Nodepool specific configuration.
	NodePools []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition `json:"nodePools"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
