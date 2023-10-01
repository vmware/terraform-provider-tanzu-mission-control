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
	sortByKey            = "sort_by"
	queryKey             = "query"
	includeTotalCountKey = "include_total_count"
	schedulesKey         = "schedules"
	totalCountKey        = "total_count"
)

var backupScheduleDataSourceSchema = map[string]*schema.Schema{
	scopeKey:             searchScopeSchema,
	sortByKey:            sortBySchema,
	queryKey:             querySchema,
	includeTotalCountKey: includeTotalSchema,
	schedulesKey:         schedulesSchema,
	totalCountKey:        totalCountSchema,
}

var searchScopeSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	Required: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			clusterNameKey:           clusterNameSchema,
			managementClusterNameKey: managementClusterNameSchema,
			provisionerNameKey:       provisionerNameSchema,
			nameKey:                  nameSchema,
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

var schedulesSchema = &schema.Schema{
	Type:     schema.TypeList,
	Computed: true,
	Elem: &schema.Resource{
		Schema: backupScheduleResourceSchema,
	},
}

var totalCountSchema = &schema.Schema{
	Type:     schema.TypeString,
	Computed: true,
}
