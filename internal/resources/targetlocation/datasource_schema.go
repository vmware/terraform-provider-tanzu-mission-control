/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocation

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// Root Keys.
	ScopeKey             = "scope"
	SortByKey            = "sort_by"
	QueryKey             = "query"
	IncludeTotalCountKey = "include_total_count"
	TargetLocationsKey   = "target_locations"
	TotalCountKey        = "total_count"

	// Scope Directive Keys.
	ProviderScopeKey = "provider"
	ClusterScopeKey  = "cluster"

	// Provider Scope Directive Keys.
	ProviderScopeProviderNameKey      = "provider_name"
	ProviderScopeNameKey              = "name"
	ProviderScopeCredentialNameKey    = "credential_name" // #nosec G101
	ProviderScopeAssignedGroupNameKey = "assigned_group_name"

	// Cluster Scope Directive Keys.
	ClusterScopeClusterNameKey           = "cluster_name"
	ClusterScopeManagementClusterNameKey = "management_cluster_name"
	ClusterScopeProvisionerNameKey       = "provisioner_name"
	ClusterScopeNameKey                  = "name"
)

var backupTargetLocationDataSourceSchema = map[string]*schema.Schema{
	ScopeKey:             scopeSchema,
	SortByKey:            sortBySchema,
	QueryKey:             querySchema,
	IncludeTotalCountKey: includeTotalSchema,
	TargetLocationsKey:   targetLocationsSchema,
	TotalCountKey:        totalCountSchema,
}

var scopeSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	Required: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ProviderScopeKey: {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ProviderScopeProviderNameKey: {
							Type:     schema.TypeString,
							Required: true,
						},
						ProviderScopeNameKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						ProviderScopeCredentialNameKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						ProviderScopeAssignedGroupNameKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			ClusterScopeKey: {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ClusterScopeClusterNameKey: {
							Type:     schema.TypeString,
							Required: true,
						},
						ClusterScopeManagementClusterNameKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						ClusterScopeProvisionerNameKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						ClusterScopeNameKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	},
}

var sortBySchema = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var querySchema = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var includeTotalSchema = &schema.Schema{
	Type:     schema.TypeBool,
	Optional: true,
	Default:  true,
}

var targetLocationsSchema = &schema.Schema{
	Type:     schema.TypeList,
	Computed: true,
	Elem: &schema.Resource{
		Schema: backupTargetLocationResourceSchema,
	},
}

var totalCountSchema = &schema.Schema{
	Type:     schema.TypeString,
	Computed: true,
}
