/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcmodels

import (
	"github.com/go-openapi/swag"

	tkccommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/nodepool"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology The cluster topology.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.Topology
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology struct {

	// The name of the cluster class for the cluster.
	ClusterClass string `json:"clusterClass,omitempty"`

	// Control plane specific configuration.
	ControlPlane *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterControlPlane `json:"controlPlane,omitempty"`

	// The core addons.
	CoreAddons []*VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCoreAddon `json:"coreAddons"`

	// Network specific configuration.
	Network *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings `json:"network,omitempty"`

	// Nodepool definition for the cluster.
	NodePools []*nodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition `json:"nodePools"`

	// Variables configuration for the cluster.
	Variables []*tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable `json:"variables"`

	// Kubernetes version of the cluster.
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
