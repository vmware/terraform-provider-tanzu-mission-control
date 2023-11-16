/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupscheduletests

import (
	"fmt"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection/tests"
	"strings"

	clusterres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	backupscheduleres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/backupschedule"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	targetlocationres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/targetlocation"
	targetlocationtests "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/targetlocation/tests"
)

type ResourceBuildMode string

const (
	RsFullBuild                  ResourceBuildMode = "FULL"
	RsDataProtectionParentRsOnly ResourceBuildMode = "DATA_PROTECTION_ONLY"
	RsTargetLocationParentRsOnly ResourceBuildMode = "TARGET_LOCATION_ONLY"
	RsNoParentRs                 ResourceBuildMode = "NO_PARENT_RS"
)

const (
	FullClusterBackupScheduleResourceName = "test_full_cluster_backup"
	FullClusterBackupScheduleName         = "full-cluster-backup"

	NamespacesBackupScheduleResourceName = "test_namespaces_backup"
	NamespacesBackupScheduleName         = "namespaces-backup"

	LabelsBackupScheduleResourceName = "test_labels_backup"
	LabelsBackupScheduleName         = "labels-backup"
)

var (
	FullClusterBackupScheduleResourceFullName = fmt.Sprintf("%s.%s", backupscheduleres.ResourceName, FullClusterBackupScheduleResourceName)
	NamespacesBackupScheduleResourceFullName  = fmt.Sprintf("%s.%s", backupscheduleres.ResourceName, NamespacesBackupScheduleResourceName)
	LabelsBackupScheduleResourceFullName      = fmt.Sprintf("%s.%s", backupscheduleres.ResourceName, LabelsBackupScheduleResourceName)
)

type ResourceTFConfigBuilder struct {
	DataProtectionRequiredResource string
	TargetLocationRequiredResource string
	ClusterInfo                    string
	TargetLocationInfo             string
}

func InitResourceTFConfigBuilder(scopeHelper *commonscope.ScopeHelperResources, bMode ResourceBuildMode, tmcManageCredentials string) *ResourceTFConfigBuilder {
	var (
		dataProtectionRequiredResource string
		targetLocationRequiredResource string
	)

	switch bMode {
	case RsFullBuild:
		dataProtectionConfigBuilder := dataprotectiontests.InitResourceTFConfigBuilder(scopeHelper, dataprotectiontests.RsFullBuild)
		targetLocationConfigBuilder := targetlocationtests.InitResourceTFConfigBuilder(scopeHelper, targetlocationtests.RsClusterOnlyNoParentRs)

		dataProtectionRequiredResource = dataProtectionConfigBuilder.GetEnableDataProtectionConfig()
		targetLocationRequiredResource = targetLocationConfigBuilder.GetTMCManagedTargetLocationConfig(tmcManageCredentials)
	case RsDataProtectionParentRsOnly:
		dataProtectionConfigBuilder := dataprotectiontests.InitResourceTFConfigBuilder(scopeHelper, dataprotectiontests.RsFullBuild)

		dataProtectionRequiredResource = dataProtectionConfigBuilder.GetEnableDataProtectionConfig()
	case RsTargetLocationParentRsOnly:
		targetLocationConfigBuilder := targetlocationtests.InitResourceTFConfigBuilder(scopeHelper, targetlocationtests.RsFullBuild)

		targetLocationRequiredResource = targetLocationConfigBuilder.GetTMCManagedTargetLocationConfig(tmcManageCredentials)
	}

	mgmtClusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ManagementClusterNameKey)
	clusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.NameKey)
	provisionerName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ProvisionerNameKey)
	clusterInfo := fmt.Sprintf(`
		%s = %s
		%s = %s        
		%s = %s  
		`,
		backupscheduleres.ClusterNameKey, clusterName,
		backupscheduleres.ManagementClusterNameKey, mgmtClusterName,
		backupscheduleres.ProvisionerNameKey, provisionerName)

	targetLocationInfo := fmt.Sprintf("storage_location = %s.%s", targetlocationtests.TmcManagedResourceFullName, targetlocationres.NameKey)

	tfConfigBuilder := &ResourceTFConfigBuilder{
		DataProtectionRequiredResource: strings.Trim(dataProtectionRequiredResource, " "),
		TargetLocationRequiredResource: strings.Trim(targetLocationRequiredResource, " "),
		ClusterInfo:                    clusterInfo,
		TargetLocationInfo:             targetLocationInfo,
	}

	return tfConfigBuilder
}

func (builder *ResourceTFConfigBuilder) GetFullClusterBackupScheduleConfig() string {
	return fmt.Sprintf(`
		%s

		%s

		resource "%s" "%s" {
		  name         = "%s"
		  scope = "%s"
          %s

		  spec {
			schedule {
			  rate = "0 12 * * 1"
			}
		
			template {
              %s
			  backup_ttl = "2592000s"
			  excluded_namespaces = [
				"app-01",
				"app-02",
				"app-03",
				"app-04"
			  ]
			  excluded_resources = [
				"secrets",
				"configmaps"
			  ]
			}
		  }

          depends_on = [%s, %s]
		}
		`,
		builder.DataProtectionRequiredResource,
		builder.TargetLocationRequiredResource,
		backupscheduleres.ResourceName,
		FullClusterBackupScheduleResourceName,
		FullClusterBackupScheduleName,
		backupscheduleres.FullClusterBackupScope,
		builder.ClusterInfo,
		builder.TargetLocationInfo,
		dataprotectiontests.EnableDataProtectionResourceFullName,
		targetlocationtests.TmcManagedResourceFullName)
}

func (builder *ResourceTFConfigBuilder) GetNamespacesBackupScheduleConfig() string {
	return fmt.Sprintf(`
		%s
		
		%s
		
		resource "%s" "%s" {
		  name         = "%s"
		  scope = "%s"
          %s
		
		  spec {
			schedule {
			  rate = "30 * * * *"
			}
		
			template {
			  included_namespaces = [
				"app-01",
				"app-02",
				"app-03",
				"app-04"
			  ]
			  backup_ttl = "86400s"
			  excluded_resources = [
				"secrets",
				"configmaps"
			  ]
			  include_cluster_resources    = true
			  %s
			  hooks {
				resource {
				  name = "sample-config"
				  pre_hook {
					exec {
					  command = ["echo 'hello'"]
					  container = "workload"
					  on_error  = "CONTINUE"
					  timeout   = "10s"
					}
				  }
				  pre_hook {
					exec {
					  command = ["echo 'hello'"]
					  container = "db"
					  on_error  = "CONTINUE"
					  timeout   = "30s"
					}
				  }
				  post_hook {
					exec {
					  command = ["echo 'goodbye'"]
					  container = "db"
					  on_error  = "CONTINUE"
					  timeout   = "60s"
					}
				  }
				  post_hook {
					exec {
					  command = ["echo 'goodbye'"]
					  container = "workload"
					  on_error  = "FAIL"
					  timeout   = "20s"
					}
				  }
				}
			  }
			}
		  }

          depends_on = [%s, %s]
		}
		`,
		builder.DataProtectionRequiredResource,
		builder.TargetLocationRequiredResource,
		backupscheduleres.ResourceName,
		NamespacesBackupScheduleResourceName,
		NamespacesBackupScheduleName,
		backupscheduleres.NamespacesBackupScope,
		builder.ClusterInfo,
		builder.TargetLocationInfo,
		dataprotectiontests.EnableDataProtectionResourceFullName,
		targetlocationtests.TmcManagedResourceFullName)
}

func (builder *ResourceTFConfigBuilder) GetLabelsBackupScheduleConfig() string {
	return fmt.Sprintf(`
		%s

		%s

		resource "%s" "%s" {
		  name         = "%s"
		  scope = "%s"
          %s

          spec {
			schedule {
			  rate = "0 12 * * *"
			}
		
			template {
			  default_volumes_to_fs_backup = false
			  include_cluster_resources = true
			  backup_ttl = "604800s"
			  %s
			  label_selector {
				match_expression {
				  key      = "apps.tanzu.vmware.com/tap-ns"
				  operator = "Exists"
				}
				match_expression {
				  key      = "apps.tanzu.vmware.com/exclude-from-backup"
				  operator = "DoesNotExist"
				}
			  }
			}
		  }

          depends_on = [%s, %s]
		}
		`,
		builder.DataProtectionRequiredResource,
		builder.TargetLocationRequiredResource,
		backupscheduleres.ResourceName,
		LabelsBackupScheduleResourceName,
		LabelsBackupScheduleName,
		backupscheduleres.LabelSelectorBackupScope,
		builder.ClusterInfo,
		builder.TargetLocationInfo,
		dataprotectiontests.EnableDataProtectionResourceFullName,
		targetlocationtests.TmcManagedResourceFullName)
}
