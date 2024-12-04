// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package inspections

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ResourceNameInspectionResults = "tanzu-mission-control_inspection_results"
)

var inspectionResultsDataSourceSchema = map[string]*schema.Schema{
	ClusterNameKey:           clusterNameSchema,
	ManagementClusterNameKey: managementClusterNameSchema,
	ProvisionerNameKey:       provisionerNameSchema,
	NameKey:                  getNameSchema(true),
	StatusKey:                computedInspectionSchema.Elem.(*schema.Resource).Schema[StatusKey],
}
