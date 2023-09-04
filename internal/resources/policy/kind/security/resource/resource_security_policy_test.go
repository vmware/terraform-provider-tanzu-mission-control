//go:build securitypolicy
// +build securitypolicy

/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package securitypolicyresource

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
	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	policykindsecurity "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security"
	policyoperations "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/operations"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	securityPolicyResource    = policykindsecurity.ResourceName
	securityPolicyResourceVar = "test_security_policy"
	securityPolicyNamePrefix  = "tf-sp-test"
)

type testAcceptanceConfig struct {
	Provider                   *schema.Provider
	SecurityPolicyResource     string
	SecurityPolicyResourceVar  string
	SecurityPolicyResourceName string
	SecurityPolicyName         string
	ScopeHelperResources       *policy.ScopeHelperResources
}

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                   initTestProvider(t),
		SecurityPolicyResource:     securityPolicyResource,
		SecurityPolicyResourceVar:  securityPolicyResourceVar,
		SecurityPolicyResourceName: fmt.Sprintf("%s.%s", securityPolicyResource, securityPolicyResourceVar),
		SecurityPolicyName:         acctest.RandomWithPrefix(securityPolicyNamePrefix),
		ScopeHelperResources:       policy.NewScopeHelperResources(),
	}
}

func TestAcceptanceForSecurityPolicyResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	t.Log("start security policy resource acceptance tests!")

	// Test case for security policy resource with baseline recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(scope.ClusterScope, policykindsecurity.BaselineRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindsecurity.BaselineRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindsecurity.BaselineRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("security policy resource acceptance test complete for baseline recipe!")

	// Test case for security policy resource with custom recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(scope.ClusterScope, policykindsecurity.CustomRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindsecurity.CustomRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindsecurity.CustomRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("security policy resource acceptance test complete for custom recipe!")

	// Test case for security policy resource with strict recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(scope.ClusterScope, policykindsecurity.StrictRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(scope.ClusterScope),
			},
			{
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(scope.ClusterGroupScope, policykindsecurity.StrictRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(scope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(scope.OrganizationScope, policykindsecurity.StrictRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(scope.OrganizationScope),
			},
		},
	},
	)

	t.Log("security policy resource acceptance test complete for strict recipe!")
	t.Log("all security policy resource acceptance tests complete!")
}

func (testConfig *testAcceptanceConfig) getTestSecurityPolicyResourceBasicConfigValue(scope scope.Scope, recipe policykindsecurity.Recipe) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestPolicyResourceHelperAndScope(scope, policyoperations.ScopeMap[testConfig.SecurityPolicyResource], false)
	inputBlock := testConfig.getTestSecurityPolicyResourceInput(recipe)

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
	`, helperBlock, testConfig.SecurityPolicyResource, testConfig.SecurityPolicyResourceVar, testConfig.SecurityPolicyName, scopeBlock, inputBlock)
}

// getTestSecurityPolicyResourceInput builds the input block for security policy resource based a recipe.
func (testConfig *testAcceptanceConfig) getTestSecurityPolicyResourceInput(recipe policykindsecurity.Recipe) string {
	var inputBlock string

	switch recipe {
	case policykindsecurity.BaselineRecipe:
		inputBlock = `
    input {
      baseline {
        audit              = true
        disable_native_psp = true
      }
    }
`
	case policykindsecurity.CustomRecipe:
		inputBlock = `
    input {
      custom {
        audit                        = true
        disable_native_psp           = false
        allow_privileged_containers  = true
        allow_privilege_escalation   = true
        allow_host_namespace_sharing = true
        allow_host_network           = true
        read_only_root_file_system   = true

        allowed_host_port_range {
          min = 3000
          max = 5000
        }

        allowed_volumes              = [
          "configMap",
          "nfs",
          "vsphereVolume"
        ]

        run_as_user {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        run_as_group {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        supplemental_groups {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        fs_group {
          rule = "RunAsAny"

          ranges {
            min = 3
            max = 5
          }
          ranges {
            min = 7
            max = 12
          }
        }

        linux_capabilities {
          allowed_capabilities       = [
            "CHOWN",
            "IPC_LOCK"
          ]
          required_drop_capabilities = [
            "SYS_TIME"
          ]
        }

        allowed_host_paths {
          path_prefix = "p1"
          read_only  = true
        }
        allowed_host_paths {
          path_prefix = "p2"
          read_only  = false
        }
        allowed_host_paths {
          path_prefix = "p3"
          read_only  = true
        }

        allowed_se_linux_options {
          level = "s0"
          role = "sysadm_r"
          type = "httpd_sys_content_t"
          user = "root"
        }

        sysctls {
          forbidden_sysctls = [
            "kernel.msgmax",
            "kernel.sem"
          ]
        }

        seccomp {
          allowed_profiles        = [
            "Localhost"
          ]
          allowed_localhost_files = [
            "profiles/audit.json",
            "profiles/violation.json"
          ]
        }
      }
    }
`
	case policykindsecurity.StrictRecipe:
		inputBlock = `
    input {
      strict {
        audit              = true
        disable_native_psp = false
      }
    }
`
	case policykindsecurity.UnknownRecipe:
		log.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(policykindsecurity.RecipesAllowed[:], `, `))
	}

	return inputBlock
}

// checkSecurityPolicyResourceAttributes checks for security policy creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkSecurityPolicyResourceAttributes(scopeType scope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifySecurityPolicyResourceCreation(scopeType),
		resource.TestCheckResourceAttr(testConfig.SecurityPolicyResourceName, "name", testConfig.SecurityPolicyName),
	}

	switch scopeType {
	case scope.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.SecurityPolicyResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	case scope.ClusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.SecurityPolicyResourceName, "scope.0.cluster_group.0.cluster_group", testConfig.ScopeHelperResources.ClusterGroup.Name))
	case scope.OrganizationScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.SecurityPolicyResourceName, "scope.0.organization.0.organization", testConfig.ScopeHelperResources.OrgID))
	case scope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policyoperations.ScopeMap[testConfig.SecurityPolicyResource], `, `))
	}

	check = append(check, policy.MetaResourceAttributeCheck(testConfig.SecurityPolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifySecurityPolicyResourceCreation(scopeType scope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.SecurityPolicyResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.SecurityPolicyResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.SecurityPolicyResourceName)
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
				Name:                  testConfig.SecurityPolicyName,
				ProvisionerName:       scope.AttachedValue,
			}

			resp, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped security policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped security policy resource is empty, resource: %s", testConfig.SecurityPolicyResourceName)
			}
		case scope.ClusterGroupScope:
			fn := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName{
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				Name:             testConfig.SecurityPolicyName,
			}

			resp, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster group scoped security policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster group scoped security policy resource is empty, resource: %s", testConfig.SecurityPolicyResourceName)
			}
		case scope.OrganizationScope:
			fn := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
				OrgID: testConfig.ScopeHelperResources.OrgID,
				Name:  testConfig.SecurityPolicyName,
			}

			resp, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "organization scoped security policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "organization scoped security policy resource is empty, resource: %s", testConfig.SecurityPolicyResourceName)
			}
		case scope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policyoperations.ScopeMap[testConfig.SecurityPolicyResource], `, `))
		}

		return nil
	}
}
