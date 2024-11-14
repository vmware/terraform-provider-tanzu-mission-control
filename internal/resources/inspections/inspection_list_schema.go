// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package inspections

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ResourceNameInspections = "tanzu-mission-control_inspections"

	// Root Keys.
	InspectionListKey = "inspections"
	TotalCountKey     = "total_count"
)

var inspectionListDataSourceSchema = map[string]*schema.Schema{
	ClusterNameKey:           clusterNameSchema,
	ManagementClusterNameKey: managementClusterNameSchema,
	ProvisionerNameKey:       provisionerNameSchema,
	NameKey:                  getNameSchema(false),
	InspectionListKey:        computedInspectionSchema,
	TotalCountKey:            totalCountSchema,
}

var totalCountSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Total count of inspections returned.",
	Computed:    true,
}
