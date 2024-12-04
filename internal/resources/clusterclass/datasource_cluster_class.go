// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clusterclass

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clusterclassmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clusterclass"
)

func DataSourceClusterClass() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceClusterClassRead,
		Schema:      clusterClassSchema,
	}
}

func dataSourceClusterClassRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var resp *clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassListData

	config := m.(authctx.TanzuContext)
	request, err := tfModelDataSourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey, ManagementClusterNameKey, ProvisionerNameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read cluster class"))
	}

	clusterClassFn := request.FullName
	resp, err = config.TMCConnection.ClusterClassResourceService.ClusterClassResourceServiceGet(clusterClassFn)

	switch {
	case err != nil:
		return diag.FromErr(errors.Wrap(err, "Couldn't read cluster class"))
	case resp.ClusterClasses == nil || len(resp.ClusterClasses) == 0:
		data.SetId("NO_DATA")
	default:
		clusterClass := resp.ClusterClasses[0]
		err = tfModelDataSourceConverter.FillTFSchema(clusterClass, data)

		if err != nil {
			diags = diag.FromErr(err)
		}

		variablesMap := BuildClusterClassMap(clusterClass.Spec)
		variablesTemplateMap := generateClusterVariablesTemplate(variablesMap)
		variablesSchemaJSONBytes, _ := json.Marshal(variablesMap)
		variablesTemplateJSONBytes, _ := json.Marshal(variablesTemplateMap)

		_ = data.Set(VariablesSchemaKey, helper.ConvertToString(variablesSchemaJSONBytes, ""))
		_ = data.Set(VariablesTemplateKey, helper.ConvertToString(variablesTemplateJSONBytes, ""))

		fullNameList := []string{clusterClassFn.ManagementClusterName, clusterClassFn.ProvisionerName, clusterClassFn.Name}

		data.SetId(strings.Join(fullNameList, "/"))
	}

	return diags
}
