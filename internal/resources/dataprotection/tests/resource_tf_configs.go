/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectiontests

import (
	"fmt"
	dataprotectionres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection"
	"strings"

	clusterres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

const (
	EnableDataProtectionResourceName = "test_enable_data_protection"
)

var (
	EnableDataProtectionResourceFullName = fmt.Sprintf("%s.%s", dataprotectionres.ResourceName, EnableDataProtectionResourceName)
)

type ResourceBuildMode string

const (
	RsFullBuild        ResourceBuildMode = "FULL"
	RsNoParentResource ResourceBuildMode = "NO_PARENT_RESOURCE"
)

type ResourceTFConfigBuilder struct {
	ClusterRequiredResource string
	ClusterInfo             string
}

func InitResourceTFConfigBuilder(scopeHelper *commonscope.ScopeHelperResources, bMode ResourceBuildMode) *ResourceTFConfigBuilder {
	var clusterRequiredResource string

	if bMode != RsNoParentResource {
		clusterRequiredResource, _ = scopeHelper.GetTestResourceHelperAndScope(commonscope.ClusterScope, []string{commonscope.ClusterKey})
	}

	mgmtClusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ManagementClusterNameKey)
	clusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.NameKey)
	provisionerName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ProvisionerNameKey)
	clusterInfo := fmt.Sprintf(`
		%s = %s
		%s = %s        
		%s = %s  
		`,
		dataprotectionres.ClusterNameKey, clusterName,
		dataprotectionres.ManagementClusterNameKey, mgmtClusterName,
		dataprotectionres.ProvisionerNameKey, provisionerName)

	tfConfigBuilder := &ResourceTFConfigBuilder{
		ClusterRequiredResource: strings.Trim(clusterRequiredResource, " "),
		ClusterInfo:             clusterInfo,
	}

	return tfConfigBuilder
}

func (builder *ResourceTFConfigBuilder) GetEnableDataProtectionConfig() string {
	return fmt.Sprintf(`
		%s

		resource "%s" "%s" {
			%s
		}
		`,
		builder.ClusterRequiredResource,
		dataprotectionres.ResourceName,
		EnableDataProtectionResourceName,
		builder.ClusterInfo)
}
