/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// Root Directive Keys.
	SortByKey            = "sort_by"
	QueryKey             = "query"
	IncludeTotalCountKey = "include_total_count"
	SchedulesKey         = "schedules"
	TotalCountKey        = "total_count"
	ScopeKey             = "scope"
	ClusterScopeKey      = "cluster"
	ClusterGroupScopeKey = "cluster_group"
	ClusterGroupNameKey  = "cluster_group_name"
)

var (
	nameDSSchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "The name of the backup schedule",
		Required:    true,
		ForceNew:    true,
	}

	managementClusterNameDSSchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Management cluster name",
		Required:    true,
		ForceNew:    true,
	}

	provisionerNameDSSchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Cluster provisioner name",
		Required:    true,
		ForceNew:    true,
	}

	sortBySchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Sort backups by field.",
		Optional:    true,
	}

	querySchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Define a query for listing backups",
		Optional:    true,
	}

	includeTotalSchema = &schema.Schema{
		Type:        schema.TypeBool,
		Description: "Whether to include total count of backups.\n(Default: True)",
		Optional:    true,
		Default:     true,
	}

	schedulesSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "A list of schedules returned",
		Computed:    true,
		Elem: &schema.Resource{
			Schema: backupScheduleResourceSchema,
		},
	}

	totalCountSchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Total count of schedules returned",
		Computed:    true,
	}

	backupScheduleDataSourceSchema = map[string]*schema.Schema{
		NameKey:              nameDSSchema,
		ScopeKey:             searchScopeSchema,
		SortByKey:            sortBySchema,
		QueryKey:             querySchema,
		IncludeTotalCountKey: includeTotalSchema,
		SchedulesKey:         schedulesSchema,
		TotalCountKey:        totalCountSchema,
	}
)

var searchScopeSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Search scope block",
	MaxItems:    1,
	Required:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ClusterGroupScopeKey: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Cluster group scope block",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ClusterGroupNameKey: {
							Type:        schema.TypeString,
							Description: "Cluster group name",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			ClusterScopeKey: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Cluster scope block",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ClusterNameKey:           clusterNameSchema,
						ManagementClusterNameKey: managementClusterNameDSSchema,
						ProvisionerNameKey:       provisionerNameDSSchema,
					},
				},
			},
		},
	},
}
