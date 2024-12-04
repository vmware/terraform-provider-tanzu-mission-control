// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgawsmodel

import (
	"github.com/go-openapi/swag"

	clustercommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/common"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec TKG AWS cluster spec.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgaws.Spec
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec struct {

	// Advanced configurations for AWS cluster.
	AdvancedConfigs []*clustercommon.VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig `json:"advancedConfigs"`

	// Kubernetes version distribution for the cluster.
	Distribution *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsDistribution `json:"distribution,omitempty"`

	// Cluster settings for the AWS cluster.
	Settings *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings `json:"settings,omitempty"`

	// Topology configuration of the cluster.
	Topology *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology `json:"topology,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
