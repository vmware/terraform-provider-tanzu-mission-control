/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"fmt"
	"strings"

	clusterres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	clustergroupres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	integrationschema "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/integration/schema"
)

const (
	ClusterIntegrationResourceName      = "test_cluster_integration"
	ClusterGroupIntegrationResourceName = "test_cluster_group_integration"
)

var (
	ClusterIntegrationResourceFullName      = fmt.Sprintf("%s.%s", integrationschema.ResourceName, ClusterIntegrationResourceName)
	ClusterGroupIntegrationResourceFullName = fmt.Sprintf("%s.%s", integrationschema.ResourceName, ClusterGroupIntegrationResourceName)
)

type ResourceTFConfigBuilder struct {
	ClusterResource      string
	ClusterScope         string
	ClusterGroupResource string
	clusterGroupScope    string
}

func InitResourceTFConfigBuilder(scopeHelper *commonscope.ScopeHelperResources) *ResourceTFConfigBuilder {
	clusterRes, _ := scopeHelper.GetTestResourceHelperAndScope(commonscope.ClusterScope, []string{commonscope.ClusterKey})
	clusterGroupRes, _ := scopeHelper.GetTestResourceHelperAndScope(commonscope.ClusterGroupScope, []string{commonscope.ClusterGroupKey})

	mgmtClusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ManagementClusterNameKey)
	clusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.NameKey)
	provisionerName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ProvisionerNameKey)
	clusterScope := fmt.Sprintf(`
		%s = %s
		%s = %s        
		%s = %s  
		`,
		integrationschema.NameKey, clusterName,
		integrationschema.ManagementClusterNameKey, mgmtClusterName,
		integrationschema.ProvisionerNameKey, provisionerName)

	clusterGroupName := fmt.Sprintf("%s.%s", scopeHelper.ClusterGroup.ResourceName, clustergroupres.NameKey)
	clusterGroupScope := fmt.Sprintf(`
		%s = %s 
		`,
		integrationschema.NameKey, clusterGroupName)

	tfConfigBuilder := &ResourceTFConfigBuilder{
		ClusterResource:      strings.Trim(clusterRes, " "),
		ClusterScope:         clusterScope,
		ClusterGroupResource: strings.Trim(clusterGroupRes, " "),
		clusterGroupScope:    clusterGroupScope,
	}

	return tfConfigBuilder
}

func (builder *ResourceTFConfigBuilder) GetClusterTOIntegrationConfig(credentialsName string) string {
	return fmt.Sprintf(`
		%s

		resource "%s" "%s" {
          scope {
             cluster {
                %s
             }
          }
		  
		
		  spec {
			credentials_name = "%s"
		  }
		}
		`,
		builder.ClusterResource,
		integrationschema.ResourceName,
		ClusterIntegrationResourceName,
		builder.ClusterScope,
		credentialsName,
	)
}

func (builder *ResourceTFConfigBuilder) GetClusterGroupTOIntegrationConfig(credentialsName string) string {
	return fmt.Sprintf(`
		%s

		resource "%s" "%s" {
          scope {
             cluster_group {
                %s
             }
          }
		  
		
		  spec {
			credentials_name = "%s"
		  }
		}
		`,
		builder.ClusterGroupResource,
		integrationschema.ResourceName,
		ClusterIntegrationResourceName,
		builder.clusterGroupScope,
		credentialsName,
	)
}
