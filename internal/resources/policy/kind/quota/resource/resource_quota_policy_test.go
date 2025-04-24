//go:build quotapolicy
// +build quotapolicy

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package quotapolicyresource

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	policykindquota "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/quota"
	policyoperations "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/operations"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	quotaPolicyResource    = policykindquota.ResourceName
	quotaPolicyResourceVar = "test_namespace_quota_policy"
	quotaPolicyNamePrefix  = "tf-ns-qt-test"
)

type testAcceptanceConfig struct {
	Provider                *schema.Provider
	QuotaPolicyResource     string
	QuotaPolicyResourceVar  string
	QuotaPolicyResourceName string
	QuotaPolicyName         string
	ScopeHelperResources    *policy.ScopeHelperResources
}

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                initTestProvider(t),
		QuotaPolicyResource:     quotaPolicyResource,
		QuotaPolicyResourceVar:  quotaPolicyResourceVar,
		QuotaPolicyResourceName: fmt.Sprintf("%s.%s", quotaPolicyResource, quotaPolicyResourceVar),
		QuotaPolicyName:         acctest.RandomWithPrefix(quotaPolicyNamePrefix),
		ScopeHelperResources:    policy.NewScopeHelperResources(),
	}
}

func TestAcceptanceForQuotaPolicyResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	t.Log("start namespace quota policy resource acceptance tests!")

	// Test case for namespace quota policy resource with custom recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped namespace quota policy acceptance test")
					}
				},
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.ClusterScope, policykindquota.CustomRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindquota.CustomRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped namespace quota policy acceptance test")
					}
				},
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindquota.CustomRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("namespace quota policy resource acceptance test complete for custom recipe!")

	// Test case for namespace quota policy resource with small recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped namespace quota policy acceptance test")
					}
				},
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.ClusterScope, policykindquota.SmallRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindquota.SmallRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped namespace quota policy acceptance test")
					}
				},
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindquota.SmallRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("namespace quota policy resource acceptance test complete for small recipe!")

	// Test case for namespace quota policy resource with medium recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped namespace quota policy acceptance test")
					}
				},
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.ClusterScope, policykindquota.MediumRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindquota.MediumRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped namespace quota policy acceptance test")
					}
				},
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindquota.MediumRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("namespace quota policy resource acceptance test complete for medium recipe!")

	// Test case for namespace quota policy resource with large recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped namespace quota policy acceptance test")
					}
				},
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.ClusterScope, policykindquota.LargeRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindquota.LargeRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped namespace quota policy acceptance test")
					}
				},
				Config: testConfig.getTestQuotaPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindquota.LargeRecipe),
				Check:  testConfig.checkQuotaPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("namespace quota policy resource acceptance test complete for large recipe!")
	t.Log("all namespace quota policy resource acceptance tests complete!")
}

func (testConfig *testAcceptanceConfig) getTestQuotaPolicyResourceBasicConfigValue(scope scope.Scope, recipe policykindquota.Recipe) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestPolicyResourceHelperAndScope(scope, policyoperations.ScopeMap[testConfig.QuotaPolicyResource], false)
	inputBlock := testConfig.getTestQuotaPolicyResourceInput(recipe)

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
	`, helperBlock, testConfig.QuotaPolicyResource, testConfig.QuotaPolicyResourceVar, testConfig.QuotaPolicyName, scopeBlock, inputBlock)
}

// getTestQuotaPolicyResourceInput builds the input block for namespace quota policy resource based on recipe.
func (testConfig *testAcceptanceConfig) getTestQuotaPolicyResourceInput(recipe policykindquota.Recipe) string {
	var inputBlock string

	switch recipe {
	case policykindquota.SmallRecipe:
		inputBlock = `
    input {
      small {}
    }
`
	case policykindquota.CustomRecipe:
		inputBlock = `
    input {
      custom {
        limits_cpu               = "4"
        limits_memory            = "8Mi"
        persistent_volume_claims = 2
        persistent_volume_claims_per_class = {
          ab : 2
          cd : 4
        }
        requests_cpu     = "2"
        requests_memory  = "4Mi"
        requests_storage = "2G"
        requests_storage_per_class = {
          test : "2G"
          twt : "4G"
        }
        resource_counts = {
          pods : 2
        }
      }
    }
`
	case policykindquota.MediumRecipe:
		inputBlock = `
    input {
      medium {}
    }
`
	case policykindquota.LargeRecipe:
		inputBlock = `
    input {
      large {}
    }
`
	case policykindquota.UnknownRecipe:
		log.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(policykindquota.RecipesAllowed[:], `, `))
	}

	return inputBlock
}

// checkQuotaPolicyResourceAttributes checks for namespace quota policy creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkQuotaPolicyResourceAttributes(scopeType scope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyQuotaPolicyResourceCreation(scopeType),
		resource.TestCheckResourceAttr(testConfig.QuotaPolicyResourceName, "name", testConfig.QuotaPolicyName),
	}

	switch scopeType {
	case scope.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.QuotaPolicyResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	case scope.ClusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.QuotaPolicyResourceName, "scope.0.cluster_group.0.cluster_group", testConfig.ScopeHelperResources.ClusterGroup.Name))
	case scope.OrganizationScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.QuotaPolicyResourceName, "scope.0.organization.0.organization", testConfig.ScopeHelperResources.OrgID))
	case scope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policyoperations.ScopeMap[testConfig.QuotaPolicyResource], `, `))
	}

	check = append(check, policy.MetaResourceAttributeCheck(testConfig.QuotaPolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyQuotaPolicyResourceCreation(scopeType scope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.QuotaPolicyResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.QuotaPolicyResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.QuotaPolicyResourceName)
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
				Name:                  testConfig.QuotaPolicyName,
				ProvisionerName:       scope.AttachedValue,
			}

			resp, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped namespace quota policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped namespace quota policy resource is empty, resource: %s", testConfig.QuotaPolicyResourceName)
			}
		case scope.ClusterGroupScope:
			fn := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName{
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				Name:             testConfig.QuotaPolicyName,
			}

			resp, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster group scoped namespace quota policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster group scoped namespace quota policy resource is empty, resource: %s", testConfig.QuotaPolicyResourceName)
			}
		case scope.OrganizationScope:
			fn := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
				OrgID: testConfig.ScopeHelperResources.OrgID,
				Name:  testConfig.QuotaPolicyName,
			}

			resp, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "organization scoped namespace quota policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "organization scoped namespace quota policy resource is empty, resource: %s", testConfig.QuotaPolicyResourceName)
			}
		case scope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policyoperations.ScopeMap[testConfig.QuotaPolicyResource], `, `))
		}

		return nil
	}
}
