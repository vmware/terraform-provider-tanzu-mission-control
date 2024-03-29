//go:build helmrepository
// +build helmrepository

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmrepository

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:             initTestProvider(t),
		RepoResource:         repoResource,
		ScopeHelperResources: commonscope.NewScopeHelperResources(),
		RepoDataSourceVar:    repoDataSourceVar,
		Namespace:            "tanzu-helm-resources",
		RepoDataSourceName:   fmt.Sprintf("data.%s.%s", ResourceName, repoDataSourceVar),
	}
}

func TestAcceptanceForHelmRepositoryDataSource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute helm repository tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_HELMREPO_ENV_TEST")
	if !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "dummy.tmc.mock.vmware.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.cloud.vmware.com")

		log.Println("Setting up the mock endpoints...")

		testConfig.setupHTTPMocks(t)
	} else {
		// Environment variables with non default values required for a successful call to Cluster Config Service.
		requiredVars := []string{
			"VMW_CLOUD_ENDPOINT",
			"TMC_ENDPOINT",
			"VMW_CLOUD_API_TOKEN",
			"ORG_ID",
		}

		// Check if the required environment variables are set.
		for _, name := range requiredVars {
			if _, found := os.LookupEnv(name); !found {
				t.Errorf("required environment variable '%s' missing", name)
			}
		}
	}

	t.Log("start helm repository data source acceptance tests!")

	// Test case for helm repository data source.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "registry.terraform.io/hashicorp/time",
			},
			"null": {
				Source: "registry.terraform.io/hashicorp/null",
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" && found {
						t.Skip("KUBECONFIG env var is not set for cluster scoped helm repository acceptance test")
					}
				},
				Config: testConfig.getTestBasicDataSourceConfigValue(),
				Check:  testConfig.checkDataSourceAttributes(),
			},
		},
	},
	)

	t.Log("helm repository data source acceptance test complete")
}

func (testConfig *testAcceptanceConfig) getTestBasicDataSourceConfigValue() string {
	if _, found := os.LookupEnv("ENABLE_HELMREPO_ENV_TEST"); !found {
		return fmt.Sprintf(`
	data "%s" "%s" {
		namespace_name = "tanzu-helm-resources"

		scope {
			cluster {
				name = "%s"
				management_cluster_name = "attached"
				provisioner_name = "attached"
			}
		}
	}
	`, testConfig.RepoResource, testConfig.RepoDataSourceVar, testConfig.ScopeHelperResources.Cluster.Name)
	}

	return fmt.Sprintf(`
	resource "tanzu-mission-control_cluster_group" "test_cluster_group" {
		name = "%s"
	  
		%s
	}

	resource "tanzu-mission-control_cluster" "test_cluster" {
		management_cluster_name = "%s"
		provisioner_name        = "%s"
		name                    = "%s"
	  
		%s
	  
		attach_k8s_cluster {
		  kubeconfig_file = "%s"
		}
	   
		spec {
		  cluster_group = tanzu-mission-control_cluster_group.test_cluster_group.name
		}
	  
		ready_wait_timeout      = "3m"
	}

	resource "time_sleep" "wait_for_3m" {
		create_duration = "10s"

		depends_on = [tanzu-mission-control_cluster.test_cluster]
	}

	resource "tanzu-mission-control_helm_feature" "test_helm_feature" {

		scope {
			cluster {
				name = tanzu-mission-control_cluster.test_cluster.name
			}
		}

		depends_on = [time_sleep.wait_for_3m]
	}

	resource "time_sleep" "wait_for_2m" {
		destroy_duration = "240s"

		depends_on = [tanzu-mission-control_helm_feature.test_helm_feature]
	}

	resource "null_resource" "next" {
		depends_on = [time_sleep.wait_for_2m]
	}

	resource "time_sleep" "wait_for_1m" {
		create_duration = "180s"

		depends_on = [tanzu-mission-control_helm_feature.test_helm_feature]
	}

	data "%s" "%s" {
		namespace_name = "%s"

		scope {
			cluster {
				name = tanzu-mission-control_cluster.test_cluster.name
			}
		}

		depends_on = [time_sleep.wait_for_1m]
	}
	`, testConfig.ScopeHelperResources.ClusterGroup.Name, testConfig.ScopeHelperResources.Meta,
		testConfig.ScopeHelperResources.Cluster.ManagementClusterName, testConfig.ScopeHelperResources.Cluster.ProvisionerName, testConfig.ScopeHelperResources.Cluster.Name,
		testConfig.ScopeHelperResources.Meta, testConfig.ScopeHelperResources.Cluster.KubeConfigPath,
		testConfig.RepoResource, testConfig.RepoDataSourceVar, testConfig.Namespace,
	)
}

// checkDataSourceAttributes checks to get helm repo.
func (testConfig *testAcceptanceConfig) checkDataSourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyDataSourceCreation(testConfig.RepoDataSourceName),
		resource.TestCheckResourceAttrSet(testConfig.RepoDataSourceName, "id"),
	}

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyDataSourceCreation(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have helm repository resource %s", name)
		}

		return nil
	}
}
