//go:build iampolicy
// +build iampolicy

/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
	namespacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/namespace"
	organizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/organization"
	workspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/workspace"
	clusterresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	clustergroupresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	namespaceresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/namespace"
	scoperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
	workspaceresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/workspace"
)

const (
	IAMPolicyResource     = ResourceName
	IAMPolicyResourceVar1 = "test_iam_policy_1"
	IAMPolicyResourceVar2 = "test_iam_policy_2"
	managementClusterName = scoperesource.AttachedValue
	provisionerName       = scoperesource.AttachedValue
	subject1Name          = "test-1"
	subject1Kind          = "GROUP"
	// Cluster.
	clusterResource            = clusterresource.ResourceName
	clusterResourceVar         = "test_cluster"
	clusterName                = "tf-attach-test"
	clusterGroupNameForCluster = "default"
	clusterRole                = "cluster.view"
	// ClusterGroup.
	clusterGroupResource    = clustergroupresource.ResourceName
	clusterGroupResourceVar = "test_cluster_group"
	clusterGroupNamePrefix  = "tf-cg-test"
	clusterGroupRole        = "clustergroup.view"
	// Workspace.
	workspaceResource    = workspaceresource.ResourceName
	workspaceResourceVar = "test_workspace"
	workspaceNamePrefix  = "tf-ws-test"
	workspaceRole        = "workspace.view"
	// Namespace.
	namespaceResource             = namespaceresource.ResourceName
	namespaceResourceVar          = "test_namespace"
	namespaceName                 = "tf-namespace"
	namespaceRole                 = "namespace.view"
	clusterNamePrefixForNamespace = "tf-c-test"
	workspaceNameForNamespace     = "default"
	// Org.
	orgRole = "organization.view"
)

type Org struct {
	ID    string
	Role1 string
}

type ClusterGroup struct {
	ResourceName string
	Resource     string
	ResourceVar  string
	Name         string
	Role1        string
}

type Cluster struct {
	Resource         string
	ResourceVar      string
	ResourceName     string
	KubeConfigPath   string
	Name             string
	Role1            string
	ClusterGroupName string
}

type Workspace struct {
	ResourceName string
	Resource     string
	ResourceVar  string
	Name         string
	Role1        string
}

type Namespace struct {
	ResourceName  string
	Resource      string
	ResourceVar   string
	WorkspaceName string
	ClusterName   string
	Name          string
	Role1         string
}

type testAcceptanceConfig struct {
	Provider               *schema.Provider
	Meta                   string
	ManagementClusterName  string
	ProvisionerName        string
	IAMPolicyResource      string
	IAMPolicyResourceVar1  string
	IAMPolicyResourceVar2  string
	IAMPolicyResourceName1 string
	IAMPolicyResourceName2 string
	Subject1Name           string
	Subject1Kind           string
	Cluster                *Cluster
	ClusterGroup           *ClusterGroup
	Workspace              *Workspace
	Namespace              *Namespace
	Org                    *Org
}

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:               initTestProvider(t),
		IAMPolicyResource:      IAMPolicyResource,
		IAMPolicyResourceVar1:  IAMPolicyResourceVar1,
		IAMPolicyResourceVar2:  IAMPolicyResourceVar2,
		IAMPolicyResourceName1: fmt.Sprintf("%s.%s", IAMPolicyResource, IAMPolicyResourceVar1),
		IAMPolicyResourceName2: fmt.Sprintf("%s.%s", IAMPolicyResource, IAMPolicyResourceVar2),
		Subject1Name:           subject1Name,
		Subject1Kind:           subject1Kind,
		Meta:                   testhelper.MetaTemplate,
		ManagementClusterName:  managementClusterName,
		ProvisionerName:        provisionerName,
		Org: &Org{
			ID:    os.Getenv("ORG_ID"),
			Role1: orgRole,
		},
		ClusterGroup: &ClusterGroup{
			Resource:     clusterGroupResource,
			ResourceVar:  clusterGroupResourceVar,
			ResourceName: fmt.Sprintf("%s.%s", clusterGroupResource, clusterGroupResourceVar),
			Name:         acctest.RandomWithPrefix(clusterGroupNamePrefix),
			Role1:        clusterGroupRole,
		},
		Cluster: &Cluster{
			Resource:         clusterResource,
			ResourceVar:      clusterResourceVar,
			ResourceName:     fmt.Sprintf("%s.%s", clusterResource, clusterResourceVar),
			KubeConfigPath:   os.Getenv("KUBECONFIG"),
			Name:             clusterName,
			Role1:            clusterRole,
			ClusterGroupName: clusterGroupNameForCluster,
		},
		Workspace: &Workspace{
			Resource:     workspaceResource,
			ResourceVar:  workspaceResourceVar,
			ResourceName: fmt.Sprintf("%s.%s", workspaceResource, workspaceResourceVar),
			Name:         acctest.RandomWithPrefix(workspaceNamePrefix),
			Role1:        workspaceRole,
		},
		Namespace: &Namespace{
			Resource:      namespaceResource,
			ResourceVar:   namespaceResourceVar,
			ResourceName:  fmt.Sprintf("%s.%s", namespaceResource, namespaceResourceVar),
			ClusterName:   acctest.RandomWithPrefix(clusterNamePrefixForNamespace),
			WorkspaceName: workspaceNameForNamespace,
			Name:          namespaceName,
			Role1:         namespaceRole,
		},
	}
}

func TestAcceptanceForIAMPolicyResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	t.Log("start IAM policy resource acceptance tests!")

	// Test case for single IAM policy resource.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped IAM policy acceptance test")
					}
				},
				Config: testConfig.getTestIAMPolicyResourceBasicConfigValue(clusterScope, false),
				Check:  testConfig.checkIAMPolicyResourceAttributes(clusterScope),
			},
			{
				PreConfig: func() {
					if testConfig.Org.ID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped IAM policy acceptance test")
					}
				},
				Config: testConfig.getTestIAMPolicyResourceBasicConfigValue(organizationScope, false),
				Check:  testConfig.checkIAMPolicyResourceAttributes(organizationScope),
			},
			{
				Config: testConfig.getTestIAMPolicyResourceBasicConfigValue(clusterGroupScope, false),
				Check:  testConfig.checkIAMPolicyResourceAttributes(clusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for namespace scoped IAM policy acceptance test")
					}
				},
				Config: testConfig.getTestIAMPolicyResourceBasicConfigValue(namespaceScope, false),
				Check:  testConfig.checkIAMPolicyResourceAttributes(namespaceScope),
			},
			{
				Config: testConfig.getTestIAMPolicyResourceBasicConfigValue(workspaceScope, false),
				Check:  testConfig.checkIAMPolicyResourceAttributes(workspaceScope),
			},
		},
	},
	)

	t.Log("IAM policy resource acceptance test complete for single resource!")

	// Test case for multiple IAM policy resources.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.Org.ID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped IAM policy acceptance test")
					}
				},
				Config: testConfig.getTestIAMPolicyResourceBasicConfigValue(organizationScope, true),
				Check: resource.ComposeTestCheckFunc(
					testConfig.checkIAMPolicyResourceAttributes(organizationScope),
					testConfig.verifyIAMPolicyResourceCreation(organizationScope, testConfig.IAMPolicyResourceName2),
				),
			},
			{
				PreConfig: func() {
					if testConfig.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped IAM policy acceptance test")
					}
				},
				Config: testConfig.getTestIAMPolicyResourceBasicConfigValue(clusterScope, true),
				Check: resource.ComposeTestCheckFunc(
					testConfig.checkIAMPolicyResourceAttributes(clusterScope),
					testConfig.verifyIAMPolicyResourceCreation(clusterScope, testConfig.IAMPolicyResourceName2),
				),
			},
			{
				Config: testConfig.getTestIAMPolicyResourceBasicConfigValue(clusterGroupScope, true),
				Check: resource.ComposeTestCheckFunc(
					testConfig.checkIAMPolicyResourceAttributes(clusterGroupScope),
					testConfig.verifyIAMPolicyResourceCreation(clusterGroupScope, testConfig.IAMPolicyResourceName2),
				),
			},
			{
				Config: testConfig.getTestIAMPolicyResourceBasicConfigValue(workspaceScope, true),
				Check: resource.ComposeTestCheckFunc(
					testConfig.checkIAMPolicyResourceAttributes(workspaceScope),
					testConfig.verifyIAMPolicyResourceCreation(workspaceScope, testConfig.IAMPolicyResourceName2),
				),
			},
			{
				PreConfig: func() {
					if testConfig.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for namespace scoped IAM policy acceptance test")
					}
				},
				Config: testConfig.getTestIAMPolicyResourceBasicConfigValue(namespaceScope, true),
				Check: resource.ComposeTestCheckFunc(
					testConfig.checkIAMPolicyResourceAttributes(namespaceScope),
					testConfig.verifyIAMPolicyResourceCreation(namespaceScope, testConfig.IAMPolicyResourceName2),
				),
			},
		},
	},
	)

	t.Log("IAM policy resource acceptance test complete for multiple resources!")
	t.Log("all IAM policy resource acceptance tests complete!")
}

func (testConfig *testAcceptanceConfig) getTestIAMPolicyResourceBasicConfigValue(scope scope, multipleIAMPolicies bool) string {
	helperBlock, scopeBlock, roles := testConfig.getTestIAMPolicyResourceHelperScopeAndRole(scope)

	var configValue string

	if multipleIAMPolicies {
		configValue = fmt.Sprintf(`
%s

resource "%s" "%s" {

  %s
  role_bindings {
    role = "%s"
    subjects {
      name = "%s"
      kind = "%s"
    }
  }
}

resource "%s" "%s" {

  %s
  role_bindings {
    role = "%s"
    subjects {
      name = "test-2"
      kind = "USER"
    }
  }
}
`, helperBlock, testConfig.IAMPolicyResource, testConfig.IAMPolicyResourceVar1, scopeBlock, roles[0], testConfig.Subject1Name, testConfig.Subject1Kind, testConfig.IAMPolicyResource, testConfig.IAMPolicyResourceVar2, scopeBlock, roles[1])
	} else {
		configValue = fmt.Sprintf(`
%s

resource "%s" "%s" {

  %s
  role_bindings {
    role = "%s"
    subjects {
      name = "%s"
      kind = "%s"
    }
  }
}
`, helperBlock, testConfig.IAMPolicyResource, testConfig.IAMPolicyResourceVar1, scopeBlock, roles[0], testConfig.Subject1Name, testConfig.Subject1Kind)
	}

	return configValue
}

// getTestIAMPolicyResourceHelperScope builds the helper resource and scope block for IAM policy resource based on a scope type.
func (testConfig *testAcceptanceConfig) getTestIAMPolicyResourceHelperScopeAndRole(scope scope) (string, string, []string) {
	var (
		helperBlock string
		scopeBlock  string
		roles       []string
	)

	switch scope {
	case organizationScope:
		helperBlock = ""
		scopeBlock = fmt.Sprintf(`
  scope {
    organization {
      org_id = "%s"
	}
  }
`, testConfig.Org.ID)
		roles = []string{"organization.view", "organization.edit"}
	case clusterGroupScope:
		helperBlock = testConfig.getTestResourceClusterGroupConfigValue()
		scopeBlock = fmt.Sprintf(`
  scope {
    cluster_group {
      name = %s.name
	}
  }
`, testConfig.ClusterGroup.ResourceName)
		roles = []string{"clustergroup.view", "clustergroup.edit"}
	case clusterScope:
		helperBlock = testConfig.getTestResourceClusterConfigValue(testConfig.Cluster.Name)
		scopeBlock = fmt.Sprintf(`
  scope {
    cluster {
      management_cluster_name = %[1]s.management_cluster_name
	  provisioner_name        = %[1]s.provisioner_name
	  name                    = %[1]s.name
	}
  }
`, testConfig.Cluster.ResourceName)
		roles = []string{"cluster.view", "cluster.edit"}
	case workspaceScope:
		helperBlock = testConfig.getTestResourceWorkspaceConfigValue()
		scopeBlock = fmt.Sprintf(`
  scope {
    workspace {
      name = %s.name
	}
  }
`, testConfig.Workspace.ResourceName)
		roles = []string{"workspace.view", "workspace.edit"}
	case namespaceScope:
		helperBlock = fmt.Sprintf(`%s

%s`, testConfig.getTestResourceClusterConfigValue(testConfig.Namespace.ClusterName), testConfig.getTestResourceNamespaceConfigValue())
		scopeBlock = fmt.Sprintf(`
  scope {
    namespace {
      management_cluster_name = %[1]s.management_cluster_name
	  provisioner_name        = %[1]s.provisioner_name
      cluster_name            = %[1]s.cluster_name
	  name                    = %[1]s.name
	}
  }
`, testConfig.Namespace.ResourceName)
		roles = []string{"namespace.view", "namespace.edit"}
	case unknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
	}

	return helperBlock, scopeBlock, roles
}

func (testConfig *testAcceptanceConfig) getTestResourceClusterConfigValue(localClusterName string) string {
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
`, testConfig.Cluster.Resource, testConfig.Cluster.ResourceVar, testConfig.ManagementClusterName, testConfig.ProvisionerName, localClusterName, testConfig.Meta, testConfig.Cluster.KubeConfigPath, testConfig.Cluster.ClusterGroupName)
}

func (testConfig *testAcceptanceConfig) getTestResourceClusterGroupConfigValue() string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"

  %s
}
`, testConfig.ClusterGroup.Resource, testConfig.ClusterGroup.ResourceVar, testConfig.ClusterGroup.Name, testConfig.Meta)
}

func (testConfig *testAcceptanceConfig) getTestResourceWorkspaceConfigValue() string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"

  %s
}
`, testConfig.Workspace.Resource, testConfig.Workspace.ResourceVar, testConfig.Workspace.Name, testConfig.Meta)
}

func (testConfig *testAcceptanceConfig) getTestResourceNamespaceConfigValue() string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  management_cluster_name = "%s"
  provisioner_name        = "%s"
  cluster_name 			  = "%s"
  name                    = "%s"

  %s

  spec {
    workspace_name = "%s"
    attach         = "false"
  }

}
`, testConfig.Namespace.Resource, testConfig.Namespace.ResourceVar, testConfig.ManagementClusterName, testConfig.ProvisionerName, testConfig.Namespace.ClusterName, testConfig.Namespace.Name, testConfig.Meta, testConfig.Namespace.WorkspaceName)
}

// checkIAMPolicyResourceAttributes checks for IAM policy creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkIAMPolicyResourceAttributes(scope scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyIAMPolicyResourceCreation(scope, testConfig.IAMPolicyResourceName1),
		resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "role_bindings.0.subjects.0.name", testConfig.Subject1Name),
		resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "role_bindings.0.subjects.0.kind", testConfig.Subject1Kind),
	}

	switch scope {
	case organizationScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "scope.0.organization.0.org_id", testConfig.Org.ID))
		check = append(check, resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "role_bindings.0.role", testConfig.Org.Role1))
	case clusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "scope.0.cluster_group.0.name", testConfig.ClusterGroup.Name))
		check = append(check, resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "role_bindings.0.role", testConfig.ClusterGroup.Role1))
	case clusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "scope.0.cluster.0.name", testConfig.Cluster.Name))
		check = append(check, resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "role_bindings.0.role", testConfig.Cluster.Role1))
	case workspaceScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "scope.0.workspace.0.name", testConfig.Workspace.Name))
		check = append(check, resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "role_bindings.0.role", testConfig.Workspace.Role1))
	case namespaceScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "scope.0.namespace.0.name", testConfig.Namespace.Name))
		check = append(check, resource.TestCheckResourceAttr(testConfig.IAMPolicyResourceName1, "role_bindings.0.role", testConfig.Namespace.Role1))
	case unknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
	}

	check = append(check, MetaResourceAttributeCheck(testConfig.IAMPolicyResourceName1)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyIAMPolicyResourceCreation(scope scope, iamPolicyResourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		var (
			policyList []*iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy
			found      bool
		)

		rs, ok := s.RootModule().Resources[iamPolicyResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", iamPolicyResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", iamPolicyResourceName)
		}

		config := authctx.TanzuContext{
			ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
			Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
			VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
			TLSConfig:        &proxy.TLSConfig{},
		}

		err := config.Setup()
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		switch scope {
		case organizationScope:
			fn := &organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName{
				OrgID: testConfig.Org.ID,
			}

			resp, err := config.TMCConnection.OrganizationIAMResourceService.ManageV1alpha1OrganizationIAMPolicyGet(fn)
			if err != nil {
				return errors.Wrap(err, "organization scoped IAM policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "organization scoped IAM policy resource is empty, resource: %s", iamPolicyResourceName)
			}

			policyList = resp.PolicyList
		case clusterGroupScope:
			fn := &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
				Name: testConfig.ClusterGroup.Name,
			}

			resp, err := config.TMCConnection.ClusterGroupIAMResourceService.ManageV1alpha1ClusterGroupIAMPolicyGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster group scoped IAM policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster group scoped IAM policy resource is empty, resource: %s", iamPolicyResourceName)
			}

			policyList = resp.PolicyList
		case clusterScope:
			fn := &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{
				ManagementClusterName: scoperesource.AttachedValue,
				ProvisionerName:       scoperesource.AttachedValue,
				Name:                  testConfig.Cluster.Name,
			}

			resp, err := config.TMCConnection.ClusterIAMResourceService.ManageV1alpha1ClusterIAMPolicyGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped IAM policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped IAM policy resource is empty, resource: %s", iamPolicyResourceName)
			}

			policyList = resp.PolicyList
		case workspaceScope:
			fn := &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
				Name: testConfig.Workspace.Name,
			}

			resp, err := config.TMCConnection.WorkspaceIAMResourceService.ManageV1alpha1WorkspaceIAMPolicyGet(fn)
			if err != nil {
				return errors.Wrap(err, "workspace scoped IAM policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "workspace scoped IAM policy resource is empty, resource: %s", iamPolicyResourceName)
			}

			policyList = resp.PolicyList
		case namespaceScope:
			fn := &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName{
				ManagementClusterName: scoperesource.AttachedValue,
				ProvisionerName:       scoperesource.AttachedValue,
				ClusterName:           testConfig.Namespace.ClusterName,
				Name:                  testConfig.Namespace.Name,
			}

			resp, err := config.TMCConnection.NamespaceIAMResourceService.ManageV1alpha1ClusterNamespaceIAMPolicyGet(fn)
			if err != nil {
				return errors.Wrap(err, "namespace scoped IAM policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "namespace scoped IAM policy resource is empty, resource: %s", iamPolicyResourceName)
			}

			policyList = resp.PolicyList
		case unknownScope:
			return errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scopesAllowed[:], `, `))
		}

		for _, item := range policyList {
			if item.Meta.UID == rs.Primary.ID {
				found = true
			}
		}

		if !found {
			return errors.Errorf("IAM policy resource not found.")
		}

		return nil
	}
}
