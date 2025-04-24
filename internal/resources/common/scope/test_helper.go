// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package commonscope

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	clusterresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	clustergroupresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
	workspaceresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/workspace"
)

const (
	// Cluster.
	clusterResource            = clusterresource.ResourceName
	clusterResourceVar         = "test_cluster"
	managementClusterName      = AttachedValue
	provisionerName            = AttachedValue
	clusterNamePrefix          = "tf-attach-test"
	clusterGroupNameForCluster = "default"

	// ClusterGroup.
	clusterGroupResource    = clustergroupresource.ResourceName
	clusterGroupResourceVar = "test_cluster_group"
	clusterGroupNamePrefix  = "tf-cg-test"

	// Workspace.
	workspaceResource    = workspaceresource.ResourceName
	workspaceResourceVar = "test_workspace"
	workspaceNamePrefix  = "tf-workspace-test"
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

type Workspace struct {
	ResourceName string
	Resource     string
	ResourceVar  string
	Name         string
}

type (
	ScopeHelperResources struct {
		Meta         string
		Cluster      *Cluster
		ClusterGroup *ClusterGroup
		Workspace    *Workspace
		OrgID        string
	}

	ScopeHelperResourcesOption func(*ScopeHelperResources)
)

func WithRandomClusterGroupNameForCluster() ScopeHelperResourcesOption {
	return func(shr *ScopeHelperResources) {
		randomClusterGroupName := acctest.RandomWithPrefix(clusterGroupNamePrefix)
		shr.Cluster.ClusterGroupName = randomClusterGroupName
		shr.ClusterGroup.Name = randomClusterGroupName
	}
}

func NewScopeHelperResources(opts ...ScopeHelperResourcesOption) *ScopeHelperResources {
	shr := &ScopeHelperResources{
		Meta: testhelper.MetaTemplate,
		Cluster: &Cluster{
			Resource:              clusterResource,
			ResourceVar:           clusterResourceVar,
			ResourceName:          fmt.Sprintf("%s.%s", clusterResource, clusterResourceVar),
			KubeConfigPath:        os.Getenv("KUBECONFIG"),
			Name:                  acctest.RandomWithPrefix(clusterNamePrefix),
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
		Workspace: &Workspace{
			ResourceName: fmt.Sprintf("%s.%s", workspaceResource, workspaceResourceVar),
			Resource:     workspaceResource,
			ResourceVar:  workspaceResourceVar,
			Name:         acctest.RandomWithPrefix(workspaceNamePrefix),
		},
		OrgID: os.Getenv("ORG_ID"),
	}

	for _, o := range opts {
		o(shr)
	}

	return shr
}

func (shr *ScopeHelperResources) getTestResourceWorkspaceConfigValue() string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"

  %s
}
`, shr.Workspace.Resource, shr.Workspace.ResourceVar, shr.Workspace.Name, shr.Meta)
}

// GetTestResourceHelperAndScope builds the helper resource and scope blocks for git repository resource based on a scope type.
func (shr *ScopeHelperResources) GetTestResourceHelperAndScope(scopeType Scope, scopesAllowed []string) (string, string) {
	var (
		helperBlock string
		scopeBlock  string
	)

	switch scopeType {
	case ClusterScope:
		helperBlock = shr.getTestResourceClusterConfigValue()

		// For cases in which WithRandomClusterGroupNameForCluster option is used.
		if shr.Cluster.ClusterGroupName == shr.ClusterGroup.Name && shr.Cluster.ClusterGroupName != clusterGroupNameForCluster {
			preRequisiteForHelperBlock := shr.getTestResourceClusterGroupConfigValue()
			helperBlock = preRequisiteForHelperBlock + helperBlock
		}

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
	    name = %s.name
		}
	}
	`, shr.ClusterGroup.ResourceName)
	case WorkspaceScope:
		helperBlock = shr.getTestResourceWorkspaceConfigValue()
		scopeBlock = fmt.Sprintf(`
	scope {
	  workspace {
	    workspace = %s.name
		}
	}
	`, shr.Workspace.ResourceName)
	case UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed, `, `))
	}

	return helperBlock, scopeBlock
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

func MetaResourceAttributeCheck(resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.uid"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.resource_version"),
	}
}
