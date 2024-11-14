// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

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
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
	workspaceresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/workspace"
)

const (
	// Cluster.
	clusterResource            = clusterresource.ResourceName
	clusterResourceVar         = "test_cluster"
	managementClusterName      = scope.AttachedValue
	provisionerName            = scope.AttachedValue
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

type ScopeHelperResources struct {
	Meta         string
	Cluster      *Cluster
	ClusterGroup *ClusterGroup
	Workspace    *Workspace
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

func (shr *ScopeHelperResources) getTestResourceWorkspaceConfigValue() string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"

  %s
}
`, shr.Workspace.Resource, shr.Workspace.ResourceVar, shr.Workspace.Name, shr.Meta)
}

// GetTestPolicyResourceHelperAndScope builds the helper resource and scope blocks for policy resource based on a scope type.
func (shr *ScopeHelperResources) GetTestPolicyResourceHelperAndScope(scopeType scope.Scope, scopesAllowed []string, mock bool) (string, string) {
	var (
		helperBlock string
		scopeBlock  string
	)

	switch scopeType {
	case scope.ClusterScope:
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
	case scope.ClusterGroupScope:
		helperBlock = shr.getTestResourceClusterGroupConfigValue()
		scopeBlock = fmt.Sprintf(`
	scope {
	  cluster_group {
	    cluster_group = %s.name
		}
	}
	`, shr.ClusterGroup.ResourceName)
	case scope.WorkspaceScope:
		if mock {
			helperBlock = ""
			scopeBlock = `
	scope {
	  workspace {
	    workspace = "workspace1"
		}
	}
	`

			break
		}

		helperBlock = shr.getTestResourceWorkspaceConfigValue()
		scopeBlock = fmt.Sprintf(`
	scope {
	  workspace {
	    workspace = %s.name
		}
	}
	`, shr.Workspace.ResourceName)
	case scope.OrganizationScope:
		helperBlock = ""
		scopeBlock = fmt.Sprintf(`
	scope {
	  organization {
	    organization = "%s"
		}
	}
	`, shr.OrgID)
	case scope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed, `, `))
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
