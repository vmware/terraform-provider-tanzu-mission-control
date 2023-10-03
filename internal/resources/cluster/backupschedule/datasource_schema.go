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
)

var backupScheduleDataSourceSchema = map[string]*schema.Schema{
	ScopeKey:             searchScopeSchema,
	SortByKey:            sortBySchema,
	QueryKey:             querySchema,
	IncludeTotalCountKey: includeTotalSchema,
	SchedulesKey:         schedulesSchema,
	TotalCountKey:        totalCountSchema,
}

var searchScopeSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Search scope block",
	MaxItems:    1,
	Required:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ClusterNameKey:           clusterNameSchema,
			ManagementClusterNameKey: managementClusterNameDSSchema,
			ProvisionerNameKey:       provisionerNameDSSchema,
			NameKey:                  nameDSSchema,
		},
	},
}

var nameDSSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "The name of the backup schedule",
	Optional:    true,
}

var managementClusterNameDSSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Management cluster name",
	Optional:    true,
}

var provisionerNameDSSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Cluster provisioner name",
	Optional:    true,
}

var sortBySchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Sort backups by field.",
	Optional:    true,
}

var querySchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Define a query for listing backups",
	Optional:    true,
}

var includeTotalSchema = &schema.Schema{
	Type:        schema.TypeBool,
	Description: "Whether to include total count of backups.\n(Default: True)",
	Optional:    true,
	Default:     true,
}

var schedulesSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "A list of schedules returned",
	Computed:    true,
	Elem: &schema.Resource{
		Schema: backupScheduleResourceSchema,
	},
}

var totalCountSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Total count of schedules returned",
	Computed:    true,
}
