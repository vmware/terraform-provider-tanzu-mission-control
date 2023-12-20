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
	backupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backupschedule/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backupschedule/scope"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func DataSourceBackupSchedule() *schema.Resource {
	// Unpack resource map to datasource map.
	constructTFModelDataSourceResponseMap()

	return &schema.Resource{
		ReadContext: dataSourceBackupScheduleRead,
		Schema:      backupScheduleDataSourceSchema,
	}
}

func dataSourceBackupScheduleRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	backupScheduleName, ok := data.Get(NameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read backup schedule name")
	}

	scopedFullnameData := scope.ConstructScope(data, backupScheduleName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control git repository entry; Scope full name is empty")
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			request, err := tfModelDataSourceRequestConverter.ConvertTFSchemaToAPIModel(data, []string{})

			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Couldn't read Tanzu Mission Control backup schedule."))
			}

			var resp *backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse

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
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			request, err := tfModelCGDataSourceRequestConverter.ConvertTFSchemaToAPIModel(data, []string{})

			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Couldn't read Tanzu Mission Control backup schedule."))
			}

			resp, err := config.TMCConnection.ClusterGroupBackupScheduleService.VmwareTanzuManageV1alpha1ClustergroupBackupScheduleResourceServiceList(request)

			switch {
			case err != nil:
				return diag.FromErr(errors.Wrap(err, "Couldn't list backup schedules"))
			case resp.Schedules == nil:
				data.SetId("NO_DATA")
			default:
				err = tfModelCGDataSourceResponseConverter.FillTFSchema(resp, data)

				if err != nil {
					diags = diag.FromErr(err)
				}

				fullNameList := []string{request.SearchScope.ClusterGroupName, request.SearchScope.Name}

				data.SetId(strings.Join(fullNameList, "/"))
			}
		}
	case commonscope.UnknownScope:
		return diag.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	return diags
}
