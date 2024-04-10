/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupscheduletests

import (
	"fmt"

	backupscheduleres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backupschedule"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

type DataSourceBuildMode string

const (
	DsFullBuild  DataSourceBuildMode = "FULL"
	DsNoParentRs DataSourceBuildMode = "NO_PARENT_RESOURCE"
)

const (
	DataSourceName   = "test_cluster_scope"
	DataSourceCGName = "test_cluster_group_scope"
)

var (
	DataSourceFullName   = fmt.Sprintf("data.%s.%s", backupscheduleres.ResourceName, DataSourceName)
	DataSourceCGFullName = fmt.Sprintf("data.%s.%s", backupscheduleres.ResourceName, DataSourceCGName)
)

type DataSourceTFConfigBuilder struct {
	BackupScheduleRequiredResource   string
	CGBackupScheduleRequiredResource string
	ClusterInfo                      string
	ClusterGroupInfo                 string
}

func InitDataSourceTFConfigBuilder(scopeHelper *commonscope.ScopeHelperResources, resourceConfigBuilder *ResourceTFConfigBuilder, bMode DataSourceBuildMode) *DataSourceTFConfigBuilder {
	var backupScheduleRequiredResource, cgBackupScheduleRequiredResource string

	if bMode != DsNoParentRs {
		backupScheduleRequiredResource = resourceConfigBuilder.GetLabelsBackupScheduleConfig()
		cgBackupScheduleRequiredResource = resourceConfigBuilder.GetLabelsCGBackupScheduleConfig()
	}

	tfConfigBuilder := &DataSourceTFConfigBuilder{
		BackupScheduleRequiredResource:   backupScheduleRequiredResource,
		CGBackupScheduleRequiredResource: cgBackupScheduleRequiredResource,
		ClusterInfo:                      fmt.Sprintf("%s = \"%s\"", backupscheduleres.ClusterNameKey, scopeHelper.Cluster.Name),
		ClusterGroupInfo:                 fmt.Sprintf("%s = \"%s\"", backupscheduleres.ClusterGroupNameKey, scopeHelper.ClusterGroup.Name),
	}

	return tfConfigBuilder
}

func (builder *DataSourceTFConfigBuilder) GetDataSourceConfig() string {
	return fmt.Sprintf(`
		%s

		data "%s" "%s" {
          name = "%s"
		  scope {
			cluster {
				%s
			}
		  }

          depends_on = [%s]
		}
		`,
		builder.BackupScheduleRequiredResource,
		backupscheduleres.ResourceName,
		DataSourceName,
		LabelsBackupScheduleName,
		builder.ClusterInfo,
		LabelsBackupScheduleResourceFullName)
}

func (builder *DataSourceTFConfigBuilder) GetCGDataSourceConfig() string {
	return fmt.Sprintf(`
		%s

		data "%s" "%s" {
          name = "%s"
		  scope {
			cluster_group {
				%s
			}
		  }

          depends_on = [%s]
		}
		`,
		builder.CGBackupScheduleRequiredResource,
		backupscheduleres.ResourceName,
		DataSourceCGName,
		LabelsCGBackupScheduleName,
		builder.ClusterGroupInfo,
		LabelsCGBackupScheduleResourceFullName)
}
