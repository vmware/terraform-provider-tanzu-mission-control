/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupscheduletests

import (
	"fmt"
	"strings"

	backupscheduleres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backupschedule"
	clusterres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	cgres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	dataprotectiontests "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection/tests"
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

	FullClusterCGBackupScheduleResourceName = "test_full_cluster_cg_backup"
	FullClusterCGBackupScheduleName         = "full-cluster-group-backup"

	NamespacesCGBackupScheduleResourceName = "test_namespaces_cg_backup"
	NamespacesCGBackupScheduleName         = "namespaces-cg-backup"

	LabelsCGBackupScheduleResourceName = "test_labels_cg_backup"
	LabelsCGBackupScheduleName         = "labels-cg-backup"
)

var (
	FullClusterBackupScheduleResourceFullName = fmt.Sprintf("%s.%s", backupscheduleres.ResourceName, FullClusterBackupScheduleResourceName)
	NamespacesBackupScheduleResourceFullName  = fmt.Sprintf("%s.%s", backupscheduleres.ResourceName, NamespacesBackupScheduleResourceName)
	LabelsBackupScheduleResourceFullName      = fmt.Sprintf("%s.%s", backupscheduleres.ResourceName, LabelsBackupScheduleResourceName)

	FullClusterCGBackupScheduleResourceFullName = fmt.Sprintf("%s.%s", backupscheduleres.ResourceName, FullClusterCGBackupScheduleResourceName)
	NamespacesCGBackupScheduleResourceFullName  = fmt.Sprintf("%s.%s", backupscheduleres.ResourceName, NamespacesCGBackupScheduleResourceName)
	LabelsCGBackupScheduleResourceFullName      = fmt.Sprintf("%s.%s", backupscheduleres.ResourceName, LabelsCGBackupScheduleResourceName)
)

type ResourceTFConfigBuilder struct {
	DataProtectionRequiredResource   string
	CGDataProtectionRequiredResource string
	TargetLocationRequiredResource   string
	CGTargetLocationRequiredResource string
	ClusterInfo                      string
	TargetLocationInfo               string
	ClusterGroupInfo                 string
}

func InitResourceTFConfigBuilder(scopeHelper *commonscope.ScopeHelperResources, bMode ResourceBuildMode, tmcManageCredentials string) *ResourceTFConfigBuilder {
	var (
		dataProtectionRequiredResource   string
		cgDataProtectionRequiredResource string
		targetLocationRequiredResource   string
		cgTargetLocationRequiredResource string
	)

	switch bMode {
	case RsFullBuild:
		dataProtectionConfigBuilder := dataprotectiontests.InitResourceTFConfigBuilder(scopeHelper, dataprotectiontests.RsFullBuild)
		targetLocationConfigBuilder := targetlocationtests.InitResourceTFConfigBuilder(scopeHelper, targetlocationtests.RsClusterOnlyNoParentRs)
		cgTargetLocationConfigBuilder := targetlocationtests.InitResourceTFConfigBuilder(scopeHelper, targetlocationtests.RsClusterGroupOnlyNoParentRs)

		dataProtectionRequiredResource = dataProtectionConfigBuilder.GetEnableDataProtectionConfig()
		cgDataProtectionRequiredResource = dataProtectionConfigBuilder.GetEnableClusterGroupDataProtectionConfig()
		targetLocationRequiredResource = targetLocationConfigBuilder.GetTMCManagedTargetLocationConfig(tmcManageCredentials)
		cgTargetLocationRequiredResource = cgTargetLocationConfigBuilder.GetTMCManagedTargetLocationConfig(tmcManageCredentials)

	case RsDataProtectionParentRsOnly:
		dataProtectionConfigBuilder := dataprotectiontests.InitResourceTFConfigBuilder(scopeHelper, dataprotectiontests.RsFullBuild)

		dataProtectionRequiredResource = dataProtectionConfigBuilder.GetEnableDataProtectionConfig()
		cgDataProtectionRequiredResource = dataProtectionConfigBuilder.GetEnableClusterGroupDataProtectionConfig()
	case RsTargetLocationParentRsOnly:
		targetLocationConfigBuilder := targetlocationtests.InitResourceTFConfigBuilder(scopeHelper, targetlocationtests.RsFullBuild)
		cgTargetLocationConfigBuilder := targetlocationtests.InitResourceTFConfigBuilder(scopeHelper, targetlocationtests.RsClusterGroupOnlyNoParentRs)
		targetLocationRequiredResource = targetLocationConfigBuilder.GetTMCManagedTargetLocationConfig(tmcManageCredentials)
		cgTargetLocationRequiredResource = cgTargetLocationConfigBuilder.GetTMCManagedTargetLocationConfig(tmcManageCredentials)
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

	cgName := fmt.Sprintf("%s.%s", scopeHelper.ClusterGroup.ResourceName, cgres.NameKey)
	cgInfo := fmt.Sprintf(`
		%s = %s
		`,
		backupscheduleres.ClusterGroupNameKey, cgName)

	targetLocationInfo := fmt.Sprintf("storage_location = %s.%s", targetlocationtests.TmcManagedResourceFullName, targetlocationres.NameKey)

	tfConfigBuilder := &ResourceTFConfigBuilder{
		DataProtectionRequiredResource:   strings.Trim(dataProtectionRequiredResource, " "),
		CGDataProtectionRequiredResource: strings.Trim(cgDataProtectionRequiredResource, " "),
		TargetLocationRequiredResource:   strings.Trim(targetLocationRequiredResource, " "),
		CGTargetLocationRequiredResource: strings.Trim(cgTargetLocationRequiredResource, " "),
		ClusterInfo:                      clusterInfo,
		TargetLocationInfo:               targetLocationInfo,
		ClusterGroupInfo:                 cgInfo,
	}

	return tfConfigBuilder
}

func (builder *ResourceTFConfigBuilder) GetFullClusterBackupScheduleConfig() string {
	return fmt.Sprintf(`
		%s

		%s

		resource "%s" "%s" {
			name = "%s"
			scope {
				cluster {
                      %s
				}
			}

          backup_scope = "%s"

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
		builder.ClusterInfo,
		backupscheduleres.FullClusterBackupScope,
		builder.TargetLocationInfo,
		dataprotectiontests.EnableDataProtectionResourceFullName,
		targetlocationtests.TmcManagedResourceFullName)
}

func (builder *ResourceTFConfigBuilder) GetNamespacesBackupScheduleConfig() string {
	return fmt.Sprintf(`
		%s
		
		%s
		
		resource "%s" "%s" {
			name = "%s"
			scope {
				cluster {
					%s
				}
			}

          backup_scope = "%s"
		
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
		builder.ClusterInfo,
		backupscheduleres.NamespacesBackupScope,
		builder.TargetLocationInfo,
		dataprotectiontests.EnableDataProtectionResourceFullName,
		targetlocationtests.TmcManagedResourceFullName)
}

func (builder *ResourceTFConfigBuilder) GetLabelsBackupScheduleConfig() string {
	return fmt.Sprintf(`
		%s

		%s

		resource "%s" "%s" {
		  name = "%s"
		  scope {
			cluster {
				%s
			  }
		  }

		  backup_scope = "%s"

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
		builder.ClusterInfo,
		backupscheduleres.LabelSelectorBackupScope,
		builder.TargetLocationInfo,
		dataprotectiontests.EnableDataProtectionResourceFullName,
		targetlocationtests.TmcManagedResourceFullName)
}

func (builder *ResourceTFConfigBuilder) GetFullClusterCGBackupScheduleConfig() string {
	return fmt.Sprintf(`
		%s

		%s

		resource "%s" "%s" {
			name = "%s"
			scope {
				cluster_group {
                      %s
				}
			}

          backup_scope = "%s"

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
		builder.CGDataProtectionRequiredResource,
		builder.CGTargetLocationRequiredResource,
		backupscheduleres.ResourceName,
		FullClusterCGBackupScheduleResourceName,
		FullClusterCGBackupScheduleName,
		builder.ClusterGroupInfo,
		backupscheduleres.FullClusterBackupScope,
		builder.TargetLocationInfo,
		dataprotectiontests.EnableClusterGroupLevelDataProtectionResourceFullName,
		targetlocationtests.TmcManagedResourceFullName)
}

func (builder *ResourceTFConfigBuilder) GetNamespacesCGBackupScheduleConfig() string {
	return fmt.Sprintf(`
		%s
		
		%s
		
		resource "%s" "%s" {
			name = "%s"
			scope {
				cluster_group {
					%s
				}
			}

          backup_scope = "%s"
		
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
		builder.CGDataProtectionRequiredResource,
		builder.CGTargetLocationRequiredResource,
		backupscheduleres.ResourceName,
		NamespacesCGBackupScheduleResourceName,
		NamespacesCGBackupScheduleName,
		builder.ClusterGroupInfo,
		backupscheduleres.NamespacesBackupScope,
		builder.TargetLocationInfo,
		dataprotectiontests.EnableClusterGroupLevelDataProtectionResourceFullName,
		targetlocationtests.TmcManagedResourceFullName)
}

func (builder *ResourceTFConfigBuilder) GetLabelsCGBackupScheduleConfig() string {
	return fmt.Sprintf(`
		%s

		%s

		resource "%s" "%s" {
		  name = "%s"
		  scope {
			cluster_group {
				%s
			  }
		  }

		  backup_scope = "%s"

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
		builder.CGDataProtectionRequiredResource,
		builder.CGTargetLocationRequiredResource,
		backupscheduleres.ResourceName,
		LabelsCGBackupScheduleResourceName,
		LabelsCGBackupScheduleName,
		builder.ClusterGroupInfo,
		backupscheduleres.LabelSelectorBackupScope,
		builder.TargetLocationInfo,
		dataprotectiontests.EnableClusterGroupLevelDataProtectionResourceFullName,
		targetlocationtests.TmcManagedResourceFullName)
}
