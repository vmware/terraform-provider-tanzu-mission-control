/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmcharts

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/helmcharts/spec"
)

const (
	ResourceName = "tanzu-mission-control_helm_charts"

	ChartMetadataNameKey = "chart_metadata_name"
	NameKey              = "name"
	RepositoryNameKey    = "repository_name"
)

func DataSourceHelmCharts() *schema.Resource {
	return &schema.Resource{
		Schema: helmSchema,
	}
}

var helmSchema = map[string]*schema.Schema{
	ChartMetadataNameKey: {
		Type:        schema.TypeString,
		Description: "Name of the helm chart.",
		Optional:    true,
		Default:     "*",
	},
	NameKey: {
		Type:        schema.TypeString,
		Description: "Version of helm chart such as 0.5.1",
		Optional:    true,
		Default:     "*",
	},
	RepositoryNameKey: {
		Type:        schema.TypeString,
		Description: "Name of helm repository.",
		Optional:    true,
		Default:     "*",
	},
	spec.ChartsKey: spec.Charts,
}
