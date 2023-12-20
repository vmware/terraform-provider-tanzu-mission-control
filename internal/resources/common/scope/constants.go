/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package commonscope

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
)

const (
	ManagementClusterNameKey = "management_cluster_name"
	ProvisionerNameKey       = "provisioner_name"
	NameKey                  = "name"
	ScopeKey                 = "scope"
	AttachedValue            = "attached"
	ClusterKey               = "cluster"
	ClusterGroupKey          = "cluster_group"
	ClusterGroupNameKey      = "cluster_group_name"
)

// Scopes.
const (
	UnknownScope Scope = iota
	ClusterScope
	ClusterGroupScope
	WorkspaceScope
)

func getSchemaForScope() func(string) *schema.Schema {
	// Emulate a map with a closure, innerMap is captured in the closure returned below.
	// Since the return value is always the same it gives the pseudo-constant output, which can be referred to in the same map-alike fashion.
	innerMap := map[string]*schema.Schema{
		ClusterKey:      cluster.ClusterFullname,
		ClusterGroupKey: clustergroup.ClusterGroupFullname,
	}

	return func(key string) *schema.Schema {
		return innerMap[key]
	}
}
