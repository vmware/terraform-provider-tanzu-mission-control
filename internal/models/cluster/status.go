/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustermodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterStatus Status of the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.Status
type VmwareTanzuManageV1alpha1ClusterStatus struct {

	// CPU allocation of a cluster.
	AllocatedCPU *VmwareTanzuManageV1alpha1CommonClusterResourceAllocation `json:"allocatedCpu,omitempty"`

	// Memory allocation of a cluster.
	AllocatedMemory *VmwareTanzuManageV1alpha1CommonClusterResourceAllocation `json:"allocatedMemory,omitempty"`

	// Health of a resource.
	Health *VmwareTanzuManageV1alpha1CommonClusterHealth `json:"health,omitempty"`

	// Health details of the cluster.
	HealthDetails *VmwareTanzuManageV1alpha1CommonClusterHealthInfo `json:"healthDetails,omitempty"`

	// Cluster infrastructure provider.
	InfrastructureProvider *VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider `json:"infrastructureProvider,omitempty"`

	// Cluster infrastructure provider region.
	InfrastructureProviderRegion string `json:"infrastructureProviderRegion,omitempty"`

	// Installer link for TMC related K8s resource manifest.
	// Note: Applicable only for attached clusters.
	// If the cluster is attached with proxy, Get on this
	// URL would need user token with sufficient permission to read the
	// proxy set during the attach. In all other cases, this
	// URL can be fetched without user token.
	InstallerLink string `json:"installerLink,omitempty"`

	// Kubernetes Server Git Version.
	KubeServerVersion string `json:"kubeServerVersion,omitempty"`

	// Kubernetes Provider of the cluster.
	KubernetesProvider *VmwareTanzuManageV1alpha1CommonClusterKubernetesProvider `json:"kubernetesProvider,omitempty"`

	// Total number of nodes.
	NodeCount string `json:"nodeCount,omitempty"`

	// Phase of the cluster resource.
	Phase *VmwareTanzuManageV1alpha1ClusterPhase `json:"phase,omitempty"`

	// Type of the cluster.
	Type *VmwareTanzuManageV1alpha1ClusterType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
