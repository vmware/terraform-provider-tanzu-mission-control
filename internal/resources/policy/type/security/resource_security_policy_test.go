/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package security

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
	clusterresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	clustergroupresource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	scoperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	securityPolicyResource    = ResourceName
	securityPolicyResourceVar = "test_security_policy"
	securityPolicyNamePrefix  = "tf-sp-test"

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

type testAcceptanceConfig struct {
	Provider                   *schema.Provider
	SecurityPolicyResource     string
	SecurityPolicyResourceVar  string
	SecurityPolicyResourceName string
	SecurityPolicyName         string
	Meta                       string
	Cluster                    *Cluster
	ClusterGroup               *ClusterGroup
	OrgID                      string
}

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                   initTestProvider(t),
		SecurityPolicyResource:     securityPolicyResource,
		SecurityPolicyResourceVar:  securityPolicyResourceVar,
		SecurityPolicyResourceName: fmt.Sprintf("%s.%s", securityPolicyResource, securityPolicyResourceVar),
		SecurityPolicyName:         acctest.RandomWithPrefix(securityPolicyNamePrefix),
		Meta:                       testhelper.MetaTemplate,
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
					if testConfig.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(policy.ClusterScope, baselineRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(policy.ClusterScope),
			},
			{
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(policy.ClusterGroupScope, baselineRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(policy.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(policy.OrganizationScope, baselineRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(policy.OrganizationScope),
			},
		},
	},
	)

	t.Log("security policy resource acceptance test complete for baseline recipe!")
	time.Sleep(2 * time.Minute)

	// Test case for security policy resource with custom recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(policy.ClusterScope, customRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(policy.ClusterScope),
			},
			{
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(policy.ClusterGroupScope, customRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(policy.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(policy.OrganizationScope, customRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(policy.OrganizationScope),
			},
		},
	},
	)

	t.Log("security policy resource acceptance test complete for custom recipe!")
	time.Sleep(2 * time.Minute)

	// Test case for security policy resource with strict recipe.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(policy.ClusterScope, strictRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(policy.ClusterScope),
			},
			{
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(policy.ClusterGroupScope, strictRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(policy.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.OrgID == "" {
						t.Skip("ORG_ID env var is not set for organization scoped security policy acceptance test")
					}
				},
				Config: testConfig.getTestSecurityPolicyResourceBasicConfigValue(policy.OrganizationScope, strictRecipe),
				Check:  testConfig.checkSecurityPolicyResourceAttributes(policy.OrganizationScope),
			},
		},
	},
	)

	t.Log("security policy resource acceptance test complete for strict recipe!")
	t.Log("all security policy resource acceptance tests complete!")
}

func (testConfig *testAcceptanceConfig) getTestSecurityPolicyResourceBasicConfigValue(scope policy.Scope, recipe recipe) string {
	helperBlock, scopeBlock, inputBlock := testConfig.getTestSecurityPolicyResourceHelperScopeAndInput(scope, recipe)

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

// getTestSecurityPolicyResourceHelperScopeAndInput builds the helper resource, scope and the input blocks for security policy resource based on a scope type and a recipe.
func (testConfig *testAcceptanceConfig) getTestSecurityPolicyResourceHelperScopeAndInput(scope policy.Scope, recipe recipe) (string, string, string) {
	var (
		helperBlock string
		scopeBlock  string
		inputBlock  string
	)

	switch scope {
	case policy.ClusterScope:
		helperBlock = testConfig.getTestResourceClusterConfigValue()
		scopeBlock = fmt.Sprintf(`
  scope {
    cluster {
      management_cluster_name = %[1]s.management_cluster_name
	  provisioner_name        = %[1]s.provisioner_name
	  name                    = %[1]s.name
	}
  }
`, testConfig.Cluster.ResourceName)
	case policy.ClusterGroupScope:
		helperBlock = testConfig.getTestResourceClusterGroupConfigValue()
		scopeBlock = fmt.Sprintf(`
  scope {
    cluster_group {
      cluster_group = %s.name
	}
  }
`, testConfig.ClusterGroup.ResourceName)
	case policy.OrganizationScope:
		helperBlock = ""
		scopeBlock = fmt.Sprintf(`
  scope {
    organization {
      organization = "%s"
	}
  }
`, testConfig.OrgID)
	case policy.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policy.ScopesAllowed[:], `, `))
	}

	switch recipe {
	case baselineRecipe:
		inputBlock = `
    input {
      baseline {
        audit              = true
        disable_native_psp = true
      }
    }
`
	case customRecipe:
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
	case strictRecipe:
		inputBlock = `
    input {
      strict {
        audit              = true
        disable_native_psp = false
      }
    }
`
	case unknownRecipe:
		log.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(recipesAllowed[:], `, `))
	}

	return helperBlock, scopeBlock, inputBlock
}

func (testConfig *testAcceptanceConfig) getTestResourceClusterConfigValue() string {
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
`, testConfig.Cluster.Resource, testConfig.Cluster.ResourceVar, testConfig.Cluster.ManagementClusterName, testConfig.Cluster.ProvisionerName, testConfig.Cluster.Name, testConfig.Meta, testConfig.Cluster.KubeConfigPath, testConfig.Cluster.ClusterGroupName)
}

func (testConfig *testAcceptanceConfig) getTestResourceClusterGroupConfigValue() string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"

  %s
}
`, testConfig.ClusterGroup.Resource, testConfig.ClusterGroup.ResourceVar, testConfig.ClusterGroup.Name, testConfig.Meta)
}

// checkSecurityPolicyResourceAttributes checks for security policy creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkSecurityPolicyResourceAttributes(scope policy.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifySecurityPolicyResourceCreation(scope),
		resource.TestCheckResourceAttr(testConfig.SecurityPolicyResourceName, "name", testConfig.SecurityPolicyName),
	}

	switch scope {
	case policy.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.SecurityPolicyResourceName, "scope.0.cluster.0.name", testConfig.Cluster.Name))
	case policy.ClusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.SecurityPolicyResourceName, "scope.0.cluster_group.0.cluster_group", testConfig.ClusterGroup.Name))
	case policy.OrganizationScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.SecurityPolicyResourceName, "scope.0.organization.0.organization", testConfig.OrgID))
	case policy.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policy.ScopesAllowed[:], `, `))
	}

	check = append(check, policy.MetaResourceAttributeCheck(testConfig.SecurityPolicyResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifySecurityPolicyResourceCreation(scope policy.Scope) resource.TestCheckFunc {
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

		switch scope {
		case policy.ClusterScope:
			fn := &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName{
				ClusterName:           testConfig.Cluster.Name,
				ManagementClusterName: scoperesource.AttachedValue,
				Name:                  testConfig.SecurityPolicyName,
				ProvisionerName:       scoperesource.AttachedValue,
			}

			resp, err := config.TMCConnection.ClusterPolicyResourceService.ManageV1alpha1ClusterPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped security policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped security policy resource is empty, resource: %s", testConfig.SecurityPolicyResourceName)
			}
		case policy.ClusterGroupScope:
			fn := &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName{
				ClusterGroupName: testConfig.ClusterGroup.Name,
				Name:             testConfig.SecurityPolicyName,
			}

			resp, err := config.TMCConnection.ClusterGroupPolicyResourceService.ManageV1alpha1ClustergroupPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster group scoped security policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster group scoped security policy resource is empty, resource: %s", testConfig.SecurityPolicyResourceName)
			}
		case policy.OrganizationScope:
			fn := &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
				OrgID: testConfig.OrgID,
				Name:  testConfig.SecurityPolicyName,
			}

			resp, err := config.TMCConnection.OrganizationPolicyResourceService.ManageV1alpha1OrganizationPolicyResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "organization scoped security policy resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "organization scoped security policy resource is empty, resource: %s", testConfig.SecurityPolicyResourceName)
			}
		case policy.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(policy.ScopesAllowed[:], `, `))
		}

		return nil
	}
}
