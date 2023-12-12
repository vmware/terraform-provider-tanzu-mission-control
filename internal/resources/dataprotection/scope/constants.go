/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

const (
	ManagementClusterNameKey = "management_cluster_name"
	ProvisionerNameKey       = "provisioner_name"
	ClusterNameKey           = "cluster_name"
	ClusterGroupNameKey      = "cluster_group_name"
	ScopeKey                 = "scope"
	ClusterKey               = "cluster"
	ClusterGroupKey          = "cluster_group"
)

// Allowed scopes.
const (
	UnknownScope Scope = iota
	ClusterScope
	ClusterGroupScope
)
