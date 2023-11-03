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

func DataSourceInspections() *schema.Resource {
	// Unpack resource map to datasource map.
	constructTFListModelDataMap()

	return &schema.Resource{
		ReadContext: dataSourceInspectionsRead,
		Schema:      inspectionListDataSourceSchema,
	}
}

func dataSourceInspectionsRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfInspectionModelConverter.ConvertTFSchemaToAPIModel(data, []string{ClusterNameKey, ManagementClusterNameKey, ProvisionerNameKey})
	inspectionFullName := model.FullName

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Converting schema failed."))
	}

	resp, err := config.TMCConnection.InspectionsResourceService.InspectionsResourceServiceList(inspectionFullName)

	switch {
	case err != nil:
		return diag.FromErr(errors.Wrapf(err, "Couldn't list inspections."))
	case resp.Scans == nil:
		data.SetId("NO_DATA")
	default:
		err = tfListModelConverter.FillTFSchema(resp, data)

		if err != nil {
			return diag.FromErr(err)
		}

		inspectionFullName := resp.Scans[0].FullName

		var idKeys = []string{
			inspectionFullName.ManagementClusterName,
			inspectionFullName.ProvisionerName,
			inspectionFullName.ClusterName,
		}

		if inspectionFullName.Name != "" {
			idKeys = append(idKeys, inspectionFullName.Name)
		}

		data.SetId(strings.Join(idKeys, "/"))
	}

	return diags
}
