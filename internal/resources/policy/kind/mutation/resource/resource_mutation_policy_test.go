//go:build mutationpolicy
// +build mutationpolicy

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package mutationpolicyresource

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindmutation "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/mutation"
	policyoperations "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/operations"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	mutationPolicyResource    = policykindmutation.ResourceName
	mutationPolicyResourceVar = "test_mutation_policy"
	mutationPolicyNamePrefix  = "tf-mp-test"
)

type testAcceptanceConfig struct {
	Provider                   *schema.Provider
	MutationPolicyResource     string
	MutationPolicyResourceVar  string
	MutationPolicyResourceName string
	MutationPolicyName         string
	ScopeHelperResources       *policy.ScopeHelperResources
}

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                   initMutationPolicyTestProvider(t),
		MutationPolicyResource:     mutationPolicyResource,
		MutationPolicyResourceVar:  mutationPolicyResourceVar,
		MutationPolicyResourceName: fmt.Sprintf("%s.%s", mutationPolicyResource, mutationPolicyResourceVar),
		MutationPolicyName:         acctest.RandomWithPrefix(mutationPolicyNamePrefix),
		ScopeHelperResources:       policy.NewScopeHelperResources(),
	}
}

func TestAcceptanceForMutationPolicyResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	t.Log("Start mutation policy resource acceptance tests!")

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(scope.ClusterGroupScope, getPodSecurityResourceInput()),
				Check:  testConfig.checkClusterGroupScopeMutationPolicyResourceAttributes(),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster group scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(scope.ClusterScope, getPodSecurityResourceInput()),
				Check:  testConfig.checkClusterScopeMutationPolicyResourceAttributes(),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(scope.OrganizationScope, getPodSecurityResourceInput()),
				Check:  testConfig.checkOrganizationScopeMutationPolicyResourceAttributes(),
			},
			{
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(scope.ClusterGroupScope, getLabelResourceInput()),
				Check:  testConfig.checkClusterGroupScopeMutationPolicyResourceAttributes(),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(scope.ClusterScope, getLabelResourceInput()),
				Check:  testConfig.checkClusterScopeMutationPolicyResourceAttributes(),
			},
			{

				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(scope.OrganizationScope, getLabelResourceInput()),
				Check:  testConfig.checkOrganizationScopeMutationPolicyResourceAttributes(),
			},
			{
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(scope.ClusterGroupScope, getAnnotationResourceInput()),
				Check:  testConfig.checkClusterGroupScopeMutationPolicyResourceAttributes(),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(scope.ClusterScope, getAnnotationResourceInput()),
				Check:  testConfig.checkClusterScopeMutationPolicyResourceAttributes(),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(scope.OrganizationScope, getAnnotationResourceInput()),
				Check:  testConfig.checkOrganizationScopeMutationPolicyResourceAttributes(),
			},
		},
	},
	)

	t.Log("Mutation policy resource acceptance test complete!")
}

func (testConfig *testAcceptanceConfig) getTestMutationPolicyResourceBasicConfigValue(scope scope.Scope, inputBlock string) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestPolicyResourceHelperAndScope(scope, policyoperations.ScopeMap[testConfig.MutationPolicyResource], false)

	return fmt.Sprintf(`
	%s
	
	resource "%s" "%s" {
	 name = "%s"
	
	 %s
	
	 spec {
	   %s
	
	   namespace_selector {
	     match_expressions {
	       key      = "component"
	       operator = "NotIn"
	       values   = [
	         "api-server",
	         "agent-gateway"
	       ]
	     }
	     match_expressions {
	       key      = "not-a-component"
	       operator = "DoesNotExist"
	       values   = []
	     }
	   }
	 }
	}
	`, helperBlock, testConfig.MutationPolicyResource, testConfig.MutationPolicyResourceVar, testConfig.MutationPolicyName, scopeBlock, inputBlock)
}

func (testConfig *testAcceptanceConfig) checkClusterScopeMutationPolicyResourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyClusterScopeMutationPolicyResourceCreation(),
		resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "name", testConfig.MutationPolicyName),
	}

	check = append(check, resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	check = append(check, policy.MetaResourceAttributeCheck(testConfig.MutationPolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) checkClusterGroupScopeMutationPolicyResourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyClusterGroupScopeMutationPolicyResourceCreation(),
		resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "name", testConfig.MutationPolicyName),
	}

	check = append(check, resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "scope.0.cluster_group.0.cluster_group", testConfig.ScopeHelperResources.ClusterGroup.Name))
	check = append(check, policy.MetaResourceAttributeCheck(testConfig.MutationPolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) checkOrganizationScopeMutationPolicyResourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyOrganizationScopeMutationPolicyResourceCreation(),
		resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "name", testConfig.MutationPolicyName),
	}

	check = append(check, resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "scope.0.organization.0.organization", testConfig.ScopeHelperResources.OrgID))
	check = append(check, policy.MetaResourceAttributeCheck(testConfig.MutationPolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyClusterScopeMutationPolicyResourceCreation() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config, err := testConfig.getContext(s)
		if err != nil {
			return err
		}

		fn := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName{
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			ManagementClusterName: scope.AttachedValue,
			Name:                  testConfig.MutationPolicyName,
			ProvisionerName:       scope.AttachedValue,
		}

		resp, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceGet(fn)
		if err != nil {
			return errors.Wrap(err, "cluster scoped mutation policy resource not found")
		}

		if resp == nil {
			return errors.Wrapf(err, "cluster scoped mutation policy resource is empty, resource: %s", testConfig.MutationPolicyResourceName)
		}

		return nil
	}
}

func (testConfig *testAcceptanceConfig) verifyClusterGroupScopeMutationPolicyResourceCreation() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config, err := testConfig.getContext(s)
		if err != nil {
			return err
		}

		fn := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName{
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
			Name:             testConfig.MutationPolicyName,
		}

		resp, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceGet(fn)
		if err != nil {
			return errors.Wrap(err, "cluster group scoped mutation policy resource not found")
		}

		if resp == nil {
			return errors.Wrapf(err, "cluster group scoped mutation policy resource is empty, resource: %s", testConfig.MutationPolicyResourceName)
		}

		return nil
	}
}

func (testConfig *testAcceptanceConfig) verifyOrganizationScopeMutationPolicyResourceCreation() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config, err := testConfig.getContext(s)
		if err != nil {
			return err
		}

		fn := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
			OrgID: testConfig.ScopeHelperResources.OrgID,
			Name:  testConfig.MutationPolicyName,
		}

		resp, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceGet(fn)
		if err != nil {
			return errors.Wrap(err, "organization scoped mutation policy resource not found")
		}

		if resp == nil {
			return errors.Wrapf(err, "organization scoped mutation policy resource is empty, resource: %s", testConfig.MutationPolicyResourceName)
		}

		return nil
	}
}
func (testConfig *testAcceptanceConfig) getContext(s *terraform.State) (*authctx.TanzuContext, error) {
	rs, ok := s.RootModule().Resources[testConfig.MutationPolicyResourceName]
	if !ok {
		return nil, fmt.Errorf("not found resource: %s", testConfig.MutationPolicyResourceName)
	}

	if rs.Primary.ID == "" {
		return nil, fmt.Errorf("ID not set, resource: %s", testConfig.MutationPolicyResourceName)
	}

	config := &authctx.TanzuContext{
		ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
		Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
		VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
		TLSConfig:        &proxy.TLSConfig{},
	}

	if err := config.Setup(); err != nil {
		return nil, errors.Wrap(err, "unable to set the context")
	}

	return config, nil
}

func getAnnotationResourceInput() string {
	return `
    input {
      annotation {
        target_kubernetes_resources {
          api_groups = [
            "apps"
          ]
          kinds = [
            "Event"
          ]
        }
        scope = "Cluster"
        annotation {
          key   = "test"
          value = "optional"
        }
      }
    }
`
}

func getLabelResourceInput() string {
	return `
    input {
      annotation {
        target_kubernetes_resources {
          api_groups = [
            "apps"
          ]
          kinds = [
            "Event"
          ]
        }
        scope = "Cluster"
        annotation {
          key   = "test"
          value = "optional"
        }
      }
    }
`
}

func getPodSecurityResourceInput() string {
	return `
 	input {
      pod_security {
        allow_privilege_escalation {
          condition = "Always"
          value = true
        }
        capabilities_add {
          operation = "merge"
          values = ["AUDIT_CONTROL", "AUDIT_WRITE"]
        }
        capabilities_drop {
          operation = "merge"
          values = ["AUDIT_WRITE"]
        }
        fs_group {
          condition = "Always"
          value = 4
        }
        privileged {
          condition = "Always"
          value = true
        }
        read_only_root_filesystem {
          condition = "Always"
          value = true
        }
        run_as_group {
          condition = "Always"
          value = 5
        }
        run_as_non_root {
          condition = "Always"
          value = true
        }
        run_as_user {
          condition = "Always"
          value = 7
        }
        se_linux_options {
          condition = "IfFieldDoesNotExist"
          level = "level_test"
          user = "user_test"
          role = "role_test"
          type = "type_test"
        }
        supplemental_groups {
          condition = "Always"
          values = [0,1,2,3]
        }
      }
	}
`
}
