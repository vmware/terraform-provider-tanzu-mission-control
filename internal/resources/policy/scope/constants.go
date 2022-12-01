/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

const (
	ManagementClusterNameKey = "management_cluster_name"
	ProvisionerNameKey       = "provisioner_name"
	ClusterNameKey           = "name"
	AttachedValue            = "attached"
	ClusterGroupNameKey      = "cluster_group"
	WorkspaceNameKey         = "workspace"
	OrganizationIDKey        = "organization"
	ScopeKey                 = "scope"
	clusterKey               = "cluster"
	clusterGroupKey          = "cluster_group"
	workspaceKey             = "workspace"
	organizationKey          = "organization"
)

// Allowed scopes.
const (
	UnknownScope Scope = iota
	ClusterScope
	ClusterGroupScope
	WorkspaceScope
	OrganizationScope
)
