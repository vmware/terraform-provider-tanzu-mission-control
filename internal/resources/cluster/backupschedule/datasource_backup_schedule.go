/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	backupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/backupschedule"
)

func DataSourceBackupSchedule() *schema.Resource {
	// Unpack resource map to datasource map.
	constructTFModelDataSourceResponseMap()

	return &schema.Resource{
		ReadContext: dataSourceTargetLocationRead,
		Schema:      backupScheduleDataSourceSchema,
	}
}

func dataSourceTargetLocationRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	var (
		resp *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse
		err  error
	)

	request := tfModelDataSourceRequestConverter.ConvertTFSchemaToAPIModel(data, []string{})
	resp, err = config.TMCConnection.BackupScheduleService.BackupScheduleResourceServiceList(request)

	switch {
	case err != nil:
		return diag.FromErr(errors.Wrap(err, "Couldn't list backup schedules"))
	case resp.Schedules == nil:
		data.SetId("NO_DATA")
	default:
		err = tfModelDataSourceResponseConverter.FillTFSchema(resp, data)

		if err != nil {
			diags = diag.FromErr(err)
		}

		fullNameList := []string{request.SearchScope.ManagementClusterName, request.SearchScope.ProvisionerName, request.SearchScope.ClusterName, request.SearchScope.Name}

		data.SetId(strings.Join(fullNameList, "/"))
	}

	return diags
}
