/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policy

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	clusterresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	clustergroupresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	scoperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	// Cluster.
	clusterResource            = clusterresource.ResourceName
	clusterResourceVar         = "test_cluster"
	managementClusterName      = scoperesource.AttachedValue
	provisionerName            = scoperesource.AttachedValue
	clusterName                = "tf-attach-test"
	clusterGroupNameForCluster = "default"

	// ClusterGroup.
	clusterGroupResource    = clustergroupresource.ResourceName
	clusterGroupResourceVar = "test_cluster_group"
	clusterGroupNamePrefix  = "tf-cg-test"
)

type Cluster struct {
	Resource              string
	ResourceVar           string
	ResourceName          string
	KubeConfigPath        string
	Name                  string
	ClusterGroupName      string
	ManagementClusterName string
	ProvisionerName       string
}

type ClusterGroup struct {
	ResourceName string
	Resource     string
	ResourceVar  string
	Name         string
}

type ScopeHelperResources struct {
	Meta         string
	Cluster      *Cluster
	ClusterGroup *ClusterGroup
	OrgID        string
}

func NewScopeHelperResources() *ScopeHelperResources {
	return &ScopeHelperResources{
		Meta: testhelper.MetaTemplate,
		Cluster: &Cluster{
			Resource:              clusterResource,
			ResourceVar:           clusterResourceVar,
			ResourceName:          fmt.Sprintf("%s.%s", clusterResource, clusterResourceVar),
			KubeConfigPath:        os.Getenv("KUBECONFIG"),
			Name:                  clusterName,
			ClusterGroupName:      clusterGroupNameForCluster,
			ManagementClusterName: managementClusterName,
			ProvisionerName:       provisionerName,
		},
		ClusterGroup: &ClusterGroup{
			ResourceName: fmt.Sprintf("%s.%s", clusterGroupResource, clusterGroupResourceVar),
			Resource:     clusterGroupResource,
			ResourceVar:  clusterGroupResourceVar,
			Name:         acctest.RandomWithPrefix(clusterGroupNamePrefix),
		},
		OrgID: os.Getenv("ORG_ID"),
	}
}

func (shr *ScopeHelperResources) getTestResourceClusterConfigValue() string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  management_cluster_name = "%s"
  provisioner_name        = "%s"
  name                    = "%s"

  %s

  attach_k8s_cluster {
    kubeconfig_file = "%s"
  }
 
  spec {
    cluster_group = "%s"
  }

  ready_wait_timeout      = "3m"
}
`, shr.Cluster.Resource, shr.Cluster.ResourceVar, shr.Cluster.ManagementClusterName, shr.Cluster.ProvisionerName, shr.Cluster.Name, shr.Meta, shr.Cluster.KubeConfigPath, shr.Cluster.ClusterGroupName)
}

func (shr *ScopeHelperResources) getTestResourceClusterGroupConfigValue() string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"

  %s
}
`, shr.ClusterGroup.Resource, shr.ClusterGroup.ResourceVar, shr.ClusterGroup.Name, shr.Meta)
}

// GetTestPolicyResourceHelperAndScope builds the helper resource and scope blocks for policy resource based on a scope type.
func (shr *ScopeHelperResources) GetTestPolicyResourceHelperAndScope(scope Scope) (string, string) {
	var (
		helperBlock string
		scopeBlock  string
	)

	switch scope {
	case ClusterScope:
		helperBlock = shr.getTestResourceClusterConfigValue()
		scopeBlock = fmt.Sprintf(`
	scope {
	  cluster {
	    management_cluster_name = %[1]s.management_cluster_name
		  provisioner_name        = %[1]s.provisioner_name
		  name                    = %[1]s.name
		}
	}
	`, shr.Cluster.ResourceName)
	case ClusterGroupScope:
		helperBlock = shr.getTestResourceClusterGroupConfigValue()
		scopeBlock = fmt.Sprintf(`
	scope {
	  cluster_group {
	    cluster_group = %s.name
		}
	}
	`, shr.ClusterGroup.ResourceName)
	case OrganizationScope:
		helperBlock = ""
		scopeBlock = fmt.Sprintf(`
	scope {
	  organization {
	    organization = "%s"
		}
	}
	`, shr.OrgID)
	case UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(ScopesAllowed[:], `, `))
	}

	return helperBlock, scopeBlock
}

func MetaResourceAttributeCheck(resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.uid"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.resource_version"),
	}
}
