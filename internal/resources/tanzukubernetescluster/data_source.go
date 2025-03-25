// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzukubernetescluster

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tanzukubernetesclustermodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
)

func DataSourceTanzuKubernetesCluster() *schema.Resource {
	dsSchema := helper.DatasourceSchemaFromResourceSchema(tanzuKubernetesClusterSchema)

	// Set 'Required' schema elements
	helper.AddRequiredFieldsToSchema(dsSchema, "name", "management_cluster_name", "provisioner_name")

	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceTanzuKubernetesClusterRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: dsSchema,
	}
}

func dataSourceTanzuKubernetesClusterRead(_ context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var (
		resp *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData
	)

	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey, ProvisionerNameKey, ManagementClusterNameKey, TimeoutPolicyKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read TKG Cluster."))
	}

	clusterFn := model.FullName

	resp, err = readFullClusterResource(&config, clusterFn)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read TKG cluster.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name))
	} else if resp != nil {
		kubernetesClusterModel := resp.TanzuKubernetesCluster

		err = tfModelResourceConverter.FillTFSchema(kubernetesClusterModel, data)

		if err != nil {
			diags = diag.FromErr(err)
		}

		fullNameList := []string{kubernetesClusterModel.FullName.ManagementClusterName, kubernetesClusterModel.FullName.ProvisionerName, kubernetesClusterModel.FullName.Name}

		data.SetId(strings.Join(fullNameList, "/"))
	}

	return diags
}
