/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicyresource

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

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
	policykindCustom "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	customPolicyResource    = policykindCustom.ResourceName
	customPolicyResourceVar = "test_custom_policy"
	customPolicyNamePrefix  = "tf-cp-test"
)

type testAcceptanceConfig struct {
	Provider                 *schema.Provider
	CustomPolicyResource     string
	CustomPolicyResourceVar  string
	CustomPolicyResourceName string
	CustomPolicyName         string
	ScopeHelperResources     *policy.ScopeHelperResources
}

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                 initTestProvider(t),
		CustomPolicyResource:     customPolicyResource,
		CustomPolicyResourceVar:  customPolicyResourceVar,
		CustomPolicyResourceName: fmt.Sprintf("%s.%s", customPolicyResource, customPolicyResourceVar),
		CustomPolicyName:         acctest.RandomWithPrefix(customPolicyNamePrefix),
		ScopeHelperResources:     policy.NewScopeHelperResources(),
	}
}

func TestAcceptanceForCustomPolicyResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	t.Log("start custom policy resource acceptance tests!")

	// Test case for custom policy resource with tmc-https-ingress recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterScope, policykindCustom.TMCHTTPSIngressRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindCustom.TMCHTTPSIngressRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindCustom.TMCHTTPSIngressRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("Custom policy resource acceptance test complete for tmc-https-ingress recipe!")
	time.Sleep(2 * time.Minute)

	// Test case for custom policy resource with tmc-block-nodeport-service recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterScope, policykindCustom.TMCBlockNodeportServiceRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindCustom.TMCBlockNodeportServiceRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindCustom.TMCBlockNodeportServiceRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("Custom policy resource acceptance test complete for tmc-block-nodeport-service recipe!")
	time.Sleep(2 * time.Minute)

	// Test case for custom policy resource with tmc-block-rolebinding-subjects recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterScope, policykindCustom.TMCBlockRolebindingSubjectsRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindCustom.TMCBlockRolebindingSubjectsRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindCustom.TMCBlockRolebindingSubjectsRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("Custom policy resource acceptance test complete for tmc-block-rolebinding-subjects recipe!")
	time.Sleep(2 * time.Minute)

	// Test case for custom policy resource with tmc-block-resources recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterScope, policykindCustom.TMCBlockResourcesRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindCustom.TMCBlockResourcesRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindCustom.TMCBlockResourcesRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("Custom policy resource acceptance test complete for tmc-block-resources recipe!")
	time.Sleep(2 * time.Minute)

	// Test case for custom policy resource with tmc-external-ips recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterScope, policykindCustom.TMCExternalIPSRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindCustom.TMCExternalIPSRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindCustom.TMCExternalIPSRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("Custom policy resource acceptance test complete for tmc-external-ips recipe!")
	time.Sleep(2 * time.Minute)

	// Test case for custom policy resource with tmc-require-labels recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterScope, policykindCustom.TMCRequireLabelsRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindCustom.TMCRequireLabelsRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped custom policy acceptance test")
					}
				},
				Config: testConfig.getTestCustomPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindCustom.TMCRequireLabelsRecipe),
				Check:  testConfig.checkCustomPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("Custom policy resource acceptance test complete for tmc-require-labels recipe!")
	t.Log("all custom policy resource acceptance tests complete!")
}

func (testConfig *testAcceptanceConfig) getTestCustomPolicyResourceBasicConfigValue(scope scope.Scope, recipe policykindCustom.Recipe) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestPolicyResourceHelperAndScope(scope)
	inputBlock := testConfig.getTestCustomPolicyResourceInput(recipe)

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
	`, helperBlock, testConfig.CustomPolicyResource, testConfig.CustomPolicyResourceVar, testConfig.CustomPolicyName, scopeBlock, inputBlock)
}

// getTestCustomPolicyResourceInput builds the input block for custom policy resource based a recipe.
func (testConfig *testAcceptanceConfig) getTestCustomPolicyResourceInput(recipe policykindCustom.Recipe) string {
	var inputBlock string

	switch recipe {
	case policykindCustom.TMCHTTPSIngressRecipe:
		inputBlock = `
    input {
      tmc_https_ingress {
        audit              = false
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
      }
    }
`
	case policykindCustom.TMCBlockNodeportServiceRecipe:
		inputBlock = `
    input {
      tmc_block_nodeport_service {
        audit              = false
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
        target_kubernetes_resources {
          api_groups = [
            "Batch",
          ]
          kinds = [
            "Ingress",
          ]
        }
      }
    }
`
	case policykindCustom.TMCBlockResourcesRecipe:
		inputBlock = `
    input {
      tmc_block_resources {
        audit              = false
        target_kubernetes_resources {
          api_groups = [
            "batch",
          ]
          kinds = [
            "Event",
          ]
        }
      }
    }
`
	case policykindCustom.TMCBlockRolebindingSubjectsRecipe:
		inputBlock = `
    input {
      tmc_block_rolebinding_subjects {
        audit              = false
        parameters {
          disallowed_subjects {
            kind = "User"
            name = "subject-1"
          }
        }
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
      }
    }
`
	case policykindCustom.TMCExternalIPSRecipe:
		inputBlock = `
    input {
      tmc_external_ips {
        audit              = false
        parameters {
          allowed_ips = [
            "127.0.0.1",
          ]
        }
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
      }
    }
`
	case policykindCustom.TMCRequireLabelsRecipe:
		inputBlock = `
    input {
      tmc_require_labels {
        audit              = false
        parameters {
          labels {
            key = "test"
            value = "optional"
          }
        }
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
        }
      }
    }
`
	case policykindCustom.UnknownRecipe:
		log.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(policykindCustom.RecipesAllowed[:], `, `))
	}

	return inputBlock
}

// checkCustomPolicyResourceAttributes checks for custom policy creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkCustomPolicyResourceAttributes(scopeType scope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyCustomPolicyResourceCreation(scopeType),
		resource.TestCheckResourceAttr(testConfig.CustomPolicyResourceName, "name", testConfig.CustomPolicyName),
	}

	switch scopeType {
	case scope.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.CustomPolicyResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	case scope.ClusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.CustomPolicyResourceName, "scope.0.cluster_group.0.cluster_group", testConfig.ScopeHelperResources.ClusterGroup.Name))
	case scope.OrganizationScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.CustomPolicyResourceName, "scope.0.organization.0.organization", testConfig.ScopeHelperResources.OrgID))
	case scope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	check = append(check, policy.MetaResourceAttributeCheck(testConfig.CustomPolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyCustomPolicyResourceCreation(scopeType scope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.CustomPolicyResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.CustomPolicyResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.CustomPolicyResourceName)
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

		switch scopeType {
		case scope.ClusterScope:
			fn := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName{
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				ManagementClusterName: scope.AttachedValue,
				Name:                  testConfig.CustomPolicyName,
				ProvisionerName:       scope.AttachedValue,
			}

			resp, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped custom policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped custom policy resource is empty, resource: %s", testConfig.CustomPolicyResourceName)
			}
		case scope.ClusterGroupScope:
			fn := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName{
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				Name:             testConfig.CustomPolicyName,
			}

			resp, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster group scoped custom policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster group scoped custom policy resource is empty, resource: %s", testConfig.CustomPolicyResourceName)
			}
		case scope.OrganizationScope:
			fn := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
				OrgID: testConfig.ScopeHelperResources.OrgID,
				Name:  testConfig.CustomPolicyName,
			}

			resp, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "organization scoped custom policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "organization scoped custom policy resource is empty, resource: %s", testConfig.CustomPolicyResourceName)
			}
		case scope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
		}

		return nil
	}
}
