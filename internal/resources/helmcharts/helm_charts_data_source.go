// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmcharts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	chartsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmcharts"
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
		Schema:      helmSchema,
		ReadContext: dataSourceHelmChartsRead,
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

func dataSourceHelmChartsRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	metadataName, ok := d.Get(ChartMetadataNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read chart metadata name")
	}

	name, ok := d.Get(NameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read chart name")
	}

	repositoryName, ok := d.Get(RepositoryNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read repository name")
	}

	resp, err := config.TMCConnection.OrganizationHelmChartsResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdChartResourceServiceList(&chartsmodel.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope{
		Name:              name,
		ChartMetadataName: metadataName,
		RepositoryName:    repositoryName,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if len(resp.Charts) == 0 {
		return diag.Errorf("No entry found for Tanzu Mission Control Helm Charts with Metadata name : %s", metadataName)
	}

	d.SetId(resp.Charts[0].Meta.UID)

	flattenedSpec := spec.FlattenCharts(resp)

	if err := d.Set(ChartMetadataNameKey, metadataName); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(NameKey, name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(RepositoryNameKey, repositoryName); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(spec.ChartsKey, flattenedSpec); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
