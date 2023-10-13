/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ManagementClusterNameKey = "management_cluster_name"
	ProvisionerNameKey       = "provisioner_name"
	ClusterNameKey           = "name"
	AttachedValue            = "attached"
	ClusterGroupNameKey      = "cluster_group"
	WorkspaceNameKey         = "workspace"
	OrganizationIDKey        = "organization"
	ScopeKey                 = "scope"
	ClusterKey               = "cluster"
	ClusterGroupKey          = "cluster_group"
	WorkspaceKey             = "workspace"
	OrganizationKey          = "organization"
)

// Allowed scopes.
const (
	UnknownScope Scope = iota
	ClusterScope
	ClusterGroupScope
	WorkspaceScope
	OrganizationScope
)

func getSchemaForScope() func(string) *schema.Schema {
	// Emulate a map with a closure, innerMap is captured in the closure returned below.
	// Since the return value is always the same it gives the pseudo-constant output, which can be referred to in the same map-alike fashion.
	innerMap := map[string]*schema.Schema{
		ClusterKey:      ClusterPolicyFullname,
		ClusterGroupKey: ClusterGroupPolicyFullname,
		WorkspaceKey:    WorkspacePolicyFullname,
		OrganizationKey: OrganizationPolicyFullname,
	}

	return func(key string) *schema.Schema {
		return innerMap[key]
	}
}
