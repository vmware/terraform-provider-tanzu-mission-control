/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package repositorycredentialclustermodel

// Full name of the Source Secret.
type VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName struct {
	// ID of Organization.
	OrgId string `json:"orgId,omitempty"`
	// Name of management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`
	// Name of Provisioner.
	ProvisionerName string `json:"provisionerName,omitempty"`
	// Name of Cluster.
	ClusterName string `json:"clusterName,omitempty"`
	// Name of Source Secret.
	Name string `json:"name,omitempty"`
}
