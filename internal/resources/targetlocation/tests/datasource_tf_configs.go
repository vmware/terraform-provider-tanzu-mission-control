/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationtests

import (
	"fmt"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection/tests"

	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	targetlocationres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/targetlocation"
)

type DataSourceBuildMode string

const (
	DsFullBuild  DataSourceBuildMode = "FULL"
	DsNoParentRs DataSourceBuildMode = "NO_PARENT_RESOURCE"
)

const (
	ClusterDataSourceName  = "test_cluster_scope"
	ProviderDataSourceName = "test_provider_scope"
)

var (
	ClusterDataSourceFullName  = fmt.Sprintf("data.%s.%s", targetlocationres.ResourceName, ClusterDataSourceName)
	ProviderDataSourceFullName = fmt.Sprintf("data.%s.%s", targetlocationres.ResourceName, ProviderDataSourceName)
)

type DataSourceTFConfigBuilder struct {
	DataProtectionRequiredResource string
	TargetLocationRequiredResource string
	ClusterInfo                    string
}

func InitDataSourceTFConfigBuilder(scopeHelper *commonscope.ScopeHelperResources, resourceConfigBuilder *ResourceTFConfigBuilder, bMode DataSourceBuildMode, tmcManageCredentials string) *DataSourceTFConfigBuilder {
	var (
		dataProtectionRequiredRs string
		targetLocationRequiredRs string
	)

	if bMode != DsNoParentRs {
		dataProtectionRequiredRs = dataprotectiontests.InitResourceTFConfigBuilder(scopeHelper, dataprotectiontests.RsNoParentResource).GetEnableDataProtectionConfig()
		targetLocationRequiredRs = resourceConfigBuilder.GetTMCManagedTargetLocationConfig(tmcManageCredentials)
	}

	tfConfigBuilder := &DataSourceTFConfigBuilder{
		DataProtectionRequiredResource: dataProtectionRequiredRs,
		TargetLocationRequiredResource: targetLocationRequiredRs,
		ClusterInfo:                    fmt.Sprintf("%s = \"%s\"", targetlocationres.ClusterScopeClusterNameKey, scopeHelper.Cluster.Name),
	}

	return tfConfigBuilder
}

func (builder *DataSourceTFConfigBuilder) GetProviderTargetLocationDataSourceConfig() string {
	return fmt.Sprintf(`
		%s

		data "%s" "%s" {
		  scope {

			provider {
    		  name = "%s"
			}
		  }
			
          depends_on = [%s]
		}
		`,
		builder.TargetLocationRequiredResource,
		targetlocationres.ResourceName,
		ProviderDataSourceName,
		TargetLocationTMCManagedName,
		TmcManagedResourceFullName)
}

func (builder *DataSourceTFConfigBuilder) GetClusterTargetLocationDataSourceConfig() string {
	return fmt.Sprintf(`
		%s

		%s

		data "%s" "%s" {
		  scope {

			cluster {
			  %s
			  name = "%s"
			}
		  }

          depends_on = [%s, %s]
		}
		`,
		builder.TargetLocationRequiredResource,
		builder.DataProtectionRequiredResource,
		targetlocationres.ResourceName,
		ClusterDataSourceName,
		builder.ClusterInfo,
		TargetLocationTMCManagedName,
		TmcManagedResourceFullName,
		dataprotectiontests.EnableDataProtectionResourceFullName)
}
