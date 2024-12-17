// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgvspheremodel

import (
	"github.com/go-openapi/swag"

	clustercommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/common"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec The tkg vsphere cluster spec.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgvsphere.Spec
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec struct {

	// Advanced configurations for vSphere cluster.
	AdvancedConfigs []*clustercommon.VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig `json:"advancedConfigs"`

	// VSphere specific distribution.
	Distribution *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereDistribution `json:"distribution,omitempty"`

	// VSphere related settings for workload cluster.
	Settings *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSettings `json:"settings,omitempty"`

	// Topology specific configuration.
	Topology *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology `json:"topology,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
