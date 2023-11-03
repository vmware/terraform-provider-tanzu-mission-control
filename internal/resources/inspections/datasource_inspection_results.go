/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspections

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
)

func DataSourceInspectionResults() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceInspectionResultsRead,
		Schema:      inspectionResultsDataSourceSchema,
	}
}

func dataSourceInspectionResultsRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfInspectionModelConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey, ClusterNameKey, ManagementClusterNameKey, ProvisionerNameKey})
	inspectionFullName := model.FullName

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Converting schema failed."))
	}

	resp, err := config.TMCConnection.InspectionsResourceService.InspectionsResourceServiceGet(inspectionFullName)

	switch {
	case err != nil:
		return diag.FromErr(errors.Wrapf(err, "Couldn't read inspection results."))
	case resp.Scan == nil:
		data.SetId("NO_DATA")
	default:
		err = tfInspectionModelConverter.FillTFSchema(resp.Scan, data)

		if err != nil {
			return diag.FromErr(err)
		}

		inspectionFullName := resp.Scan.FullName

		var idKeys = []string{
			inspectionFullName.ManagementClusterName,
			inspectionFullName.ProvisionerName,
			inspectionFullName.ClusterName,
			inspectionFullName.Name,
		}

		data.SetId(strings.Join(idKeys, "/"))
	}

	return diags
}
