//go:build mutationpolicy

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package mutationpolicyresource

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policyoperations "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/operations"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

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

	var found bool
	if _, found = os.LookupEnv("ENABLE_POLICY_ENV_TEST"); !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "play.abc.def.ghi.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.tanzu.broadcom.com")

		testConfig.ScopeHelperResources.OrgID = "dummy_org"
		testConfig.setupHTTPMocks(t)
	} else {
		requiredVars := []string{
			"VMW_CLOUD_ENDPOINT",
			"TMC_ENDPOINT",
			"VMW_CLOUD_API_TOKEN",
			"ORG_ID",
		}

		for _, name := range requiredVars {
			if _, found := os.LookupEnv(name); !found {
				t.Errorf("required environment variable '%s' missing", name)
			}
		}
	}

	t.Log("Start mutation policy resource acceptance tests!")

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(podSecurity, scope.ClusterGroupScope, getPodSecurityResourceInput()),
				Check:  testConfig.checkClusterGroupScopeMutationPolicyResourceAttributes(podSecurity),
			},
			{
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(label, scope.ClusterGroupScope, getLabelResourceInput()),
				Check:  testConfig.checkClusterGroupScopeMutationPolicyResourceAttributes(label),
			},
			{
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(annotation, scope.ClusterGroupScope, getAnnotationResourceInput()),
				Check:  testConfig.checkClusterGroupScopeMutationPolicyResourceAttributes(annotation),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(podSecurity, scope.OrganizationScope, getPodSecurityResourceInput()),
				Check:  testConfig.checkOrganizationScopeMutationPolicyResourceAttributes(podSecurity),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(label, scope.OrganizationScope, getLabelResourceInput()),
				Check:  testConfig.checkOrganizationScopeMutationPolicyResourceAttributes(label),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(annotation, scope.OrganizationScope, getAnnotationResourceInput()),
				Check:  testConfig.checkOrganizationScopeMutationPolicyResourceAttributes(annotation),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster group scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(podSecurity, scope.ClusterScope, getPodSecurityResourceInput()),
				Check:  testConfig.checkClusterScopeMutationPolicyResourceAttributes(podSecurity),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(label, scope.ClusterScope, getLabelResourceInput()),
				Check:  testConfig.checkClusterScopeMutationPolicyResourceAttributes(label),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped mutation policy acceptance test")
					}
				},
				Config: testConfig.getTestMutationPolicyResourceBasicConfigValue(annotation, scope.ClusterScope, getAnnotationResourceInput()),
				Check:  testConfig.checkClusterScopeMutationPolicyResourceAttributes(annotation),
			},
		},
	},
	)

	t.Log("Mutation policy resource acceptance test complete!")
}

func (testConfig *testAcceptanceConfig) getTestMutationPolicyResourceBasicConfigValue(recipe string, scope scope.Scope, inputBlock string) string {
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
	`, helperBlock, testConfig.MutationPolicyResource, testConfig.MutationPolicyResourceVar, testConfig.MutationPolicyName+recipe, scopeBlock, inputBlock)
}

func (testConfig *testAcceptanceConfig) checkClusterScopeMutationPolicyResourceAttributes(recipe string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyClusterScopeMutationPolicyResourceCreation(recipe),
		resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "name", testConfig.MutationPolicyName+recipe),
	}

	check = append(check, resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	check = append(check, policy.MetaResourceAttributeCheck(testConfig.MutationPolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) checkClusterGroupScopeMutationPolicyResourceAttributes(recipe string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyClusterGroupScopeMutationPolicyResourceCreation(recipe),
		resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "name", testConfig.MutationPolicyName+recipe),
	}

	check = append(check, resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "scope.0.cluster_group.0.cluster_group", testConfig.ScopeHelperResources.ClusterGroup.Name))
	check = append(check, policy.MetaResourceAttributeCheck(testConfig.MutationPolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) checkOrganizationScopeMutationPolicyResourceAttributes(recipe string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyOrganizationScopeMutationPolicyResourceCreation(recipe),
		resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "name", testConfig.MutationPolicyName+recipe),
	}

	check = append(check, resource.TestCheckResourceAttr(testConfig.MutationPolicyResourceName, "scope.0.organization.0.organization", testConfig.ScopeHelperResources.OrgID))
	check = append(check, policy.MetaResourceAttributeCheck(testConfig.MutationPolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyClusterScopeMutationPolicyResourceCreation(recipe string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config, err := testConfig.getContext(s)
		if err != nil {
			return err
		}

		fn := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName{
			ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
			ManagementClusterName: scope.AttachedValue,
			Name:                  testConfig.MutationPolicyName + recipe,
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

func (testConfig *testAcceptanceConfig) verifyClusterGroupScopeMutationPolicyResourceCreation(recipe string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config, err := testConfig.getContext(s)
		if err != nil {
			return err
		}

		fn := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName{
			ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
			Name:             testConfig.MutationPolicyName + recipe,
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

func (testConfig *testAcceptanceConfig) verifyOrganizationScopeMutationPolicyResourceCreation(recipe string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config, err := testConfig.getContext(s)
		if err != nil {
			return err
		}

		fn := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
			OrgID: testConfig.ScopeHelperResources.OrgID,
			Name:  testConfig.MutationPolicyName + recipe,
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

	err := getSetupConfig(config)
	if err != nil {
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
      label {
        target_kubernetes_resources {
          api_groups = [
            "apps"
          ]
          kinds = [
            "Event"
          ]
        }
        scope = "Cluster"
        label {
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
