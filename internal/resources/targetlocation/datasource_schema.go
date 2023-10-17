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
	Type:        schema.TypeList,
	Description: "Search scope block can contain either cluster scope or provider scope but not both.",
	MaxItems:    1,
	Required:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ProviderScopeKey: {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ProviderScopeNameKey: {
							Type:        schema.TypeString,
							Description: "The name of the target location.",
							Optional:    true,
						},
						ProviderScopeCredentialNameKey: {
							Type:        schema.TypeString,
							Description: "The name of the credentials used for the target location.",
							Optional:    true,
						},
						ProviderScopeAssignedGroupNameKey: {
							Type:        schema.TypeString,
							Description: "A cluster or cluster group assigned for the target location.",
							Optional:    true,
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
							Type:        schema.TypeString,
							Description: "Cluster name",
							Required:    true,
						},
						ClusterScopeManagementClusterNameKey: {
							Type:        schema.TypeString,
							Description: "Management cluster name",
							Optional:    true,
						},
						ClusterScopeProvisionerNameKey: {
							Type:        schema.TypeString,
							Description: "Cluster provisioner name",
							Optional:    true,
						},
						ClusterScopeNameKey: {
							Type:        schema.TypeString,
							Description: "The name of the target location",
							Optional:    true,
						},
					},
				},
			},
		},
	},
}

var sortBySchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Sort target locations by field.",
	Optional:    true,
}

var querySchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Define a query for listing target locations",
	Optional:    true,
}

var includeTotalSchema = &schema.Schema{
	Type:        schema.TypeBool,
	Description: "Whether to include total count of target locations.\n(Default: True)",
	Optional:    true,
	Default:     true,
}

var targetLocationsSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "A list of target locations returned",
	Computed:    true,
	Elem: &schema.Resource{
		Schema: backupTargetLocationResourceSchema,
	},
}

var totalCountSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Total count of target locations returned",
	Computed:    true,
}
