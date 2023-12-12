/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectiontests

import (
	"fmt"
	"strings"

	clusterres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	cgres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	dataprotectionres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection"
	dataprotectionscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection/scope"
)

const (
	EnableDataProtectionResourceName                  = "test_enable_data_protection"
	EnableClusterGroupLevelDataProtectionResourceName = "test_enable_data_protection_cg"
)

var (
	EnableDataProtectionResourceFullName                  = fmt.Sprintf("%s.%s", dataprotectionres.ResourceName, EnableDataProtectionResourceName)
	EnableClusterGroupLevelDataProtectionResourceFullName = fmt.Sprintf("%s.%s", dataprotectionres.ResourceName, EnableClusterGroupLevelDataProtectionResourceName)
)

type ResourceBuildMode string

const (
	RsFullBuild        ResourceBuildMode = "FULL"
	RsNoParentResource ResourceBuildMode = "NO_PARENT_RESOURCE"
)

type ResourceTFConfigBuilder struct {
	ClusterRequiredResource      string
	ClusterInfo                  string
	ClusterGroupRequiredResource string
	ClusterGroupInfo             string
}

func InitResourceTFConfigBuilder(scopeHelper *commonscope.ScopeHelperResources, bMode ResourceBuildMode) *ResourceTFConfigBuilder {
	var clusterRequiredResource, cgRequiredResource string

	if bMode != RsNoParentResource {
		clusterRequiredResource, _ = scopeHelper.GetTestResourceHelperAndScope(commonscope.ClusterScope, []string{commonscope.ClusterKey})
		cgRequiredResource, _ = scopeHelper.GetTestResourceHelperAndScope(commonscope.ClusterGroupScope, []string{commonscope.ClusterGroupKey})
	}

	mgmtClusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ManagementClusterNameKey)
	clusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.NameKey)
	provisionerName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ProvisionerNameKey)
	clusterInfo := fmt.Sprintf(`
		%s = %s
		%s = %s        
		%s = %s  
		`,
		dataprotectionscope.ClusterNameKey, clusterName,
		dataprotectionscope.ManagementClusterNameKey, mgmtClusterName,
		dataprotectionscope.ProvisionerNameKey, provisionerName)

	cgName := fmt.Sprintf("%s.%s", scopeHelper.ClusterGroup.ResourceName, cgres.NameKey)
	cgInfo := fmt.Sprintf(`
		%s = %s
		`,
		dataprotectionscope.ClusterGroupNameKey, cgName)

	tfConfigBuilder := &ResourceTFConfigBuilder{
		ClusterRequiredResource:      strings.Trim(clusterRequiredResource, " "),
		ClusterInfo:                  clusterInfo,
		ClusterGroupRequiredResource: strings.Trim(cgRequiredResource, " "),
		ClusterGroupInfo:             cgInfo,
	}

	return tfConfigBuilder
}

func (builder *ResourceTFConfigBuilder) GetEnableDataProtectionConfig() string {
	return fmt.Sprintf(`
		%s

		resource "%s" "%s" {
			scope {
				cluster {
					%s
				}
			}
		    spec {
				disable_restic                       = false
				enable_csi_snapshots                 = true
				enable_all_api_group_versions_backup = false
			}
		}
		`,
		builder.ClusterRequiredResource,
		dataprotectionres.ResourceName,
		EnableDataProtectionResourceName,
		builder.ClusterInfo)
}

func (builder *ResourceTFConfigBuilder) GetEnableClusterGroupDataProtectionConfig() string {
	return fmt.Sprintf(`
		%s

		resource "%s" "%s" {
			scope {
				cluster_group {
					%s
				}
			}
			spec {
				disable_restic                       = false
				enable_csi_snapshots                 = true
				enable_all_api_group_versions_backup = false
			}
		}
		`,
		builder.ClusterGroupRequiredResource,
		dataprotectionres.ResourceName,
		EnableClusterGroupLevelDataProtectionResourceName,
		builder.ClusterGroupInfo)
}
