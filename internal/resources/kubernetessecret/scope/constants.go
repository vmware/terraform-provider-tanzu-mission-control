/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

const (
	AttachedValue            = "attached"
	ScopeKey                 = "scope"
	ClusterKey               = "cluster"
	ClusterGroupKey          = "cluster_group"
	ManagementClusterNameKey = "management_cluster_name"
	ProvisionerNameKey       = "provisioner_name"
	ClusterNameKey           = "cluster_name"
)

const (
	UnknownScope Scope = iota
	ClusterScope
	ClusterGroupScope
)
