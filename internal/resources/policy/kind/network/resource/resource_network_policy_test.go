//go:build networkpolicy
// +build networkpolicy

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package networkpolicyresource

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	policyworkspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/workspace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindNetwork "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network"
	policyoperations "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/operations"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                  initTestProvider(t),
		NetworkPolicyResource:     networkPolicyResource,
		NetworkPolicyResourceVar:  networkPolicyResourceVar,
		NetworkPolicyResourceName: fmt.Sprintf("%s.%s", networkPolicyResource, networkPolicyResourceVar),
		NetworkPolicyName:         acctest.RandomWithPrefix(networkPolicyNamePrefix),
		ScopeHelperResources:      policy.NewScopeHelperResources(),
	}
}

// Function to set up config based on the type of testing- Mock/Actual environment.
func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_POLICY_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

// nolint: dupl
func TestAcceptanceForNetworkPolicyResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	var found bool
	// If the flag to execute policy tests is not found, run this as a unit test by setting up an http intercept for each endpoint
	if _, found = os.LookupEnv("ENABLE_POLICY_ENV_TEST"); !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "play.abc.def.ghi.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.tanzu.broadcom.com")

		testConfig.ScopeHelperResources.OrgID = "dummy_org"
		testConfig.setupHTTPMocks(t)
	} else {
		// Environment variables with non default values required for a successful call to MKP
		requiredVars := []string{
			"VMW_CLOUD_ENDPOINT",
			"TMC_ENDPOINT",
			"VMW_CLOUD_API_TOKEN",
			"ORG_ID",
		}

		// Check if the required environment variables are set
		for _, name := range requiredVars {
			if _, found := os.LookupEnv(name); !found {
				t.Errorf("required environment variable '%s' missing", name)
			}
		}
	}

	t.Log("start network policy resource acceptance tests!")

	// Test case for network policy resource with allow-all recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: func() []resource.TestStep {
			steps := []resource.TestStep{
				{
					Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.AllowAllRecipe),
					Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.AllowAllRecipe),
				},
				{
					PreConfig: func() {
						if testConfig.ScopeHelperResources.OrgID == "" {
							t.Skip("ORG_ID env var is not set for organization scoped network policy acceptance test")
						}
					},
					Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindNetwork.AllowAllRecipe),
					Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.OrganizationScope, policykindNetwork.AllowAllRecipe),
				},
			}

			if found {
				steps = append(steps, resource.TestStep{
					Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.AllowAllRecipe, WithFromOwnNamespace("true")),
					Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.AllowAllRecipe),
				})
			}
			return steps
		}(),
	})

	t.Log("Network policy resource acceptance test complete for allow-all recipe!")

	// Test case for network policy resource with allow-all-to-pods recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: func() []resource.TestStep {
			steps := []resource.TestStep{
				{
					Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.AllowAllToPodsRecipe),
					Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.AllowAllToPodsRecipe),
				},
				{
					PreConfig: func() {
						if testConfig.ScopeHelperResources.OrgID == "" {
							t.Skip("ORG_ID env var is not set for organization scoped network policy acceptance test")
						}
					},
					Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindNetwork.AllowAllToPodsRecipe),
					Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.OrganizationScope, policykindNetwork.AllowAllToPodsRecipe),
				},
			}

			if found {
				steps = append(steps, resource.TestStep{
					Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.AllowAllToPodsRecipe, WithFromOwnNamespace("true"), WithPodLabel("key3", "value3")),
					Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.AllowAllToPodsRecipe),
				})
			}
			return steps
		}(),
	})

	t.Log("Network policy resource acceptance test complete for allow-all-to-pods recipe!")

	// Test case for network policy resource with allow-all-egress recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.AllowAllEgressRecipe),
				Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.AllowAllEgressRecipe),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped network policy acceptance test")
					}
				},
				Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindNetwork.AllowAllEgressRecipe),
				Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.OrganizationScope, policykindNetwork.AllowAllEgressRecipe),
			},
		},
	})

	t.Log("Network policy resource acceptance test complete for allow-all-egress recipe!")

	// Test case for network policy resource with deny-all recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.DenyAllRecipe),
				Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.DenyAllRecipe),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped network policy acceptance test")
					}
				},
				Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindNetwork.DenyAllRecipe),
				Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.OrganizationScope, policykindNetwork.DenyAllRecipe),
			},
		},
	})

	t.Log("Network policy resource acceptance test complete for deny-all recipe!")

	// Test case for network policy resource with deny-all-to-pods recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: func() []resource.TestStep {
			steps := []resource.TestStep{
				{
					Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.DenyAllToPodsRecipe),
					Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.DenyAllToPodsRecipe),
				},
				{
					PreConfig: func() {
						if testConfig.ScopeHelperResources.OrgID == "" {
							t.Skip("ORG_ID env var is not set for organization scoped network policy acceptance test")
						}
					},
					Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindNetwork.DenyAllToPodsRecipe),
					Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.OrganizationScope, policykindNetwork.DenyAllToPodsRecipe),
				},
			}

			if found {
				steps = append(steps, resource.TestStep{
					Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.DenyAllToPodsRecipe, WithPodLabel("key3", "value3")),
					Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.DenyAllToPodsRecipe),
				})
			}
			return steps
		}(),
	})

	t.Log("Network policy resource acceptance test complete for deny-all-to-pods recipe!")

	// Test case for network policy resource with deny-all-egress recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.DenyAllEgressRecipe),
				Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.DenyAllEgressRecipe),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped network policy acceptance test")
					}
				},
				Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindNetwork.DenyAllEgressRecipe),
				Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.OrganizationScope, policykindNetwork.DenyAllEgressRecipe),
			},
		},
	})

	t.Log("Network policy resource acceptance test complete for deny-all-egress recipe!")

	// TODO: Skip custom-egress and custom-ingress tests currently for mock until the test issue is fixed.
	// Test case for network policy resource with custom-egress recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: func() []resource.TestStep {
			if found {
				steps := []resource.TestStep{
					{
						Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.CustomEgressRecipe),
						Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.CustomEgressRecipe),
					},
					{
						PreConfig: func() {
							if testConfig.ScopeHelperResources.OrgID == "" {
								t.Skip("ORG_ID env var is not set for organization scoped network policy acceptance test")
							}
						},
						Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindNetwork.CustomEgressRecipe),
						Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.OrganizationScope, policykindNetwork.CustomEgressRecipe),
					},
					{
						Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.CustomEgressRecipe, WithPodLabel("key3", "value3")),
						Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.CustomEgressRecipe),
					},
				}
				return steps
			}
			return []resource.TestStep{}
		}(),
	})

	t.Log("Network policy resource acceptance test complete for custom-egress recipe!")

	// Test case for network policy resource with custom-ingress recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: func() []resource.TestStep {
			if found {
				steps := []resource.TestStep{
					{
						Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.CustomIngressRecipe),
						Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.CustomIngressRecipe),
					},
					{
						PreConfig: func() {
							if testConfig.ScopeHelperResources.OrgID == "" {
								t.Skip("ORG_ID env var is not set for organization scoped network policy acceptance test")
							}
						},
						Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindNetwork.CustomIngressRecipe),
						Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.OrganizationScope, policykindNetwork.CustomIngressRecipe),
					},
					{
						Config: testConfig.getTestNetworkPolicyResourceBasicConfigValue(scope.WorkspaceScope, policykindNetwork.CustomIngressRecipe, WithPodLabel("key3", "value3")),
						Check:  testConfig.checkNetworkPolicyResourceAttributes(scope.WorkspaceScope, policykindNetwork.CustomIngressRecipe),
					},
				}
				return steps
			}
			return []resource.TestStep{}
		}(),
	})
	t.Log("Network policy resource acceptance test complete for custom-ingress recipe!")

	t.Log("all network policy resource acceptance tests complete!")
}

type (
	OperationConfig struct {
		fromOwnNamespace string
		podLabelKey      string
		podLabelValue    string
	}

	OperationOption func(*OperationConfig)
)

func WithFromOwnNamespace(fn string) OperationOption {
	return func(config *OperationConfig) {
		config.fromOwnNamespace = fn
	}
}

func WithPodLabel(k, v string) OperationOption {
	return func(config *OperationConfig) {
		config.podLabelKey = k
		config.podLabelValue = v
	}
}

func (testConfig *testAcceptanceConfig) getTestNetworkPolicyResourceBasicConfigValue(scope scope.Scope, recipe policykindNetwork.Recipe, opts ...OperationOption) string {
	var (
		helperBlock string
		scopeBlock  string
	)

	inputBlock := testConfig.getTestNetworkPolicyResourceInput(recipe, opts...)

	_, found := os.LookupEnv("ENABLE_POLICY_ENV_TEST")
	helperBlock, scopeBlock = testConfig.ScopeHelperResources.GetTestPolicyResourceHelperAndScope(scope, policyoperations.ScopeMap[testConfig.NetworkPolicyResource], !found)

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
	`, helperBlock, testConfig.NetworkPolicyResource, testConfig.NetworkPolicyResourceVar, testConfig.NetworkPolicyName+strings.ReplaceAll(string(recipe), "_", "-"), scopeBlock, inputBlock)
}

// getTestNetworkPolicyResourceInput builds the input block for network policy resource based a recipe.
func (testConfig *testAcceptanceConfig) getTestNetworkPolicyResourceInput(recipe policykindNetwork.Recipe, opts ...OperationOption) string {
	cfg := &OperationConfig{
		fromOwnNamespace: "false",
		podLabelKey:      "key1",
		podLabelValue:    "value1",
	}

	for _, o := range opts {
		o(cfg)
	}

	var inputBlock string

	switch recipe {
	case policykindNetwork.AllowAllRecipe:
		inputBlock = `
    input {
      allow_all {
        from_own_namespace = %s
      }
    }
`
		inputBlock = fmt.Sprintf(inputBlock, cfg.fromOwnNamespace)
	case policykindNetwork.AllowAllToPodsRecipe:
		inputBlock = `
    input {
      allow_all_to_pods {
        from_own_namespace = %s
        to_pod_labels = {
          "%s" = "%s"
          "key2" = "value2"
        }
      }
    }
`
		inputBlock = fmt.Sprintf(inputBlock, cfg.fromOwnNamespace, cfg.podLabelKey, cfg.podLabelValue)

	case policykindNetwork.AllowAllEgressRecipe:
		inputBlock = `
    input {
      allow_all_egress {}
    }
`
	case policykindNetwork.DenyAllRecipe:
		inputBlock = `
    input {
      deny_all {}
    }
`
	case policykindNetwork.DenyAllToPodsRecipe:
		inputBlock = `
    input {
      deny_all_to_pods {
        to_pod_labels = {
          "%s" = "%s"
          "key2" = "value2"
        }
      }
    }
`
		inputBlock = fmt.Sprintf(inputBlock, cfg.podLabelKey, cfg.podLabelValue)
	case policykindNetwork.DenyAllEgressRecipe:
		inputBlock = `
    input {
      deny_all_egress {}
    }
`
	case policykindNetwork.CustomEgressRecipe:
		inputBlock = `
    input {
      custom_egress {
        rules {
          ports {
            port = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_ip {
              ip_block {
                cidr = "192.168.1.1/24"
                except = [
                  "2001:db9::/64",
                ]
              }
            }
          }
          rule_spec {
            custom_selector {
              namespace_selector = {
                "key-1" = "value-1"
                "key-2" = "value-2"
              }
              pod_selector = {
                "key-1" = "value-1"
                "key-2" = "value-2"
              }
            }
          }
        }
        to_pod_labels = {
          "%s" = "%s"
          "key2" = "value2"
        }
      }
    }
`
		inputBlock = fmt.Sprintf(inputBlock, cfg.podLabelKey, cfg.podLabelValue)
	case policykindNetwork.CustomIngressRecipe:
		inputBlock = `
    input {
      custom_ingress {
        rules {
          ports {
            port = "8443"
            protocol = "TCP"
          }
          rule_spec {
            custom_ip {
              ip_block {
                cidr = "192.168.1.1/24"
                except = [
                  "2001:db9::/64",
                ]
              }
            }
          }
          rule_spec {
            custom_selector {
              namespace_selector = {
                "key-1" = "value-1"
                "key-2" = "value-2"
              }
              pod_selector = {
                "key-1" = "value-1"
                "key-2" = "value-2"
              }
            }
          }
        }
        to_pod_labels = {
          "%s" = "%s"
          "key2" = "value2"
        }
      }
    }
`
		inputBlock = fmt.Sprintf(inputBlock, cfg.podLabelKey, cfg.podLabelValue)
	case policykindNetwork.UnknownRecipe:
		log.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(policykindNetwork.RecipesAllowed[:], `, `))
	}

	return inputBlock
}

// checkNetworkPolicyResourceAttributes checks for network policy creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkNetworkPolicyResourceAttributes(scopeType scope.Scope, recipe policykindNetwork.Recipe) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyNetworkPolicyResourceCreation(scopeType, recipe),
		resource.TestCheckResourceAttr(testConfig.NetworkPolicyResourceName, "name", testConfig.NetworkPolicyName+strings.ReplaceAll(string(recipe), "_", "-")),
	}

	switch scopeType {
	case scope.WorkspaceScope:
		workspaceName := testConfig.ScopeHelperResources.Workspace.Name

		if _, found := os.LookupEnv("ENABLE_POLICY_ENV_TEST"); !found {
			workspaceName = "workspace1"
		}

		check = append(check, resource.TestCheckResourceAttr(testConfig.NetworkPolicyResourceName, "scope.0.workspace.0.workspace", workspaceName))
	case scope.OrganizationScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.NetworkPolicyResourceName, "scope.0.organization.0.organization", testConfig.ScopeHelperResources.OrgID))
	case scope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policyoperations.ScopeMap[testConfig.NetworkPolicyResource], `, `))
	}

	check = append(check, policy.MetaResourceAttributeCheck(testConfig.NetworkPolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyNetworkPolicyResourceCreation(scopeType scope.Scope, recipe policykindNetwork.Recipe) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.NetworkPolicyResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.NetworkPolicyResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.NetworkPolicyResourceName)
		}

		config := authctx.TanzuContext{
			ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
			Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
			VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
			TLSConfig:        &proxy.TLSConfig{},
		}

		err := getSetupConfig(&config)
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		switch scopeType {
		case scope.WorkspaceScope:
			workspaceName := testConfig.ScopeHelperResources.Workspace.Name

			if _, found := os.LookupEnv("ENABLE_POLICY_ENV_TEST"); !found {
				workspaceName = "workspace1"
			}

			fn := &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName{
				WorkspaceName: workspaceName,
				Name:          testConfig.NetworkPolicyName + strings.ReplaceAll(string(recipe), "_", "-"),
			}

			resp, err := config.TMCConnection.WorkspacePolicyResourceService.ManageV1alpha1WorkspacePolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "workspace scoped network policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "workspace scoped network policy resource is empty, resource: %s", testConfig.NetworkPolicyResourceName)
			}
		case scope.OrganizationScope:
			fn := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
				OrgID: testConfig.ScopeHelperResources.OrgID,
				Name:  testConfig.NetworkPolicyName + strings.ReplaceAll(string(recipe), "_", "-"),
			}

			resp, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "organization scoped network policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "organization scoped network policy resource is empty, resource: %s", testConfig.NetworkPolicyResourceName)
			}
		case scope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policyoperations.ScopeMap[testConfig.NetworkPolicyResource], `, `))
		}

		return nil
	}
}
