/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackagerepository

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	packagerepositoryscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForPackageRepositoryDataSource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute packagerepository tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_PKGREPO_ENV_TEST")
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

	t.Log("start packagerepository data source acceptance tests!")

	// Test case for packagerepository data source.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "registry.terraform.io/hashicorp/time",
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" && found {
						t.Skip("KUBECONFIG env var is not set for cluster scoped packagerepository acceptance test")
					}
				},
				Config: testConfig.getTestPackageRepositoryDataSourceBasicConfigValue(commonscope.ClusterScope, false, imageURL),
				Check:  testConfig.checkPkgRepositoryDataSourceAttributes(),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder...")
						testConfig.setupHTTPMocksUpdate(t)
					}
				},
				Config: testConfig.getTestPackageRepositoryDataSourceBasicConfigValue(commonscope.ClusterScope, true, updatedImageURL),
				Check:  testConfig.checkPkgRepositoryDataSourceAttributes(),
			},
		},
	},
	)

	t.Log("packagerepository data source acceptance test complete")
}

func (testConfig *testAcceptanceConfig) getTestPackageRepositoryDataSourceBasicConfigValue(scope commonscope.Scope, disabled bool, imageURL string) string {
	helperBlock, _ := testConfig.ScopeHelperResources.GetTestResourceHelperAndScope(scope, packagerepositoryscope.ScopesAllowed[:])

	if _, found := os.LookupEnv("ENABLE_PKGREPO_ENV_TEST"); !found {
		return fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"

		namespace_name = "%s"

		disabled = %t

		 scope {
		cluster {
			name = "%s"
			management_cluster_name = "attached"
			provisioner_name = "attached"
		}
	 }


		spec {
			imgpkg_bundle {
				image = "%s"
			}
		}
	}
	data "%s" "%s" {
		name = tanzu-mission-control_package_repository.test_pkg_repository.name

		scope {
			cluster {
				name = "%s"
				management_cluster_name = "attached"
				provisioner_name = "attached"
			}
		}
	}
	`, testConfig.PkgRepoResource, testConfig.PkgRepoResourceVar, testConfig.PkgRepoName, testConfig.Namespace, disabled, testConfig.ScopeHelperResources.Cluster.Name, imageURL, testConfig.PkgRepoResource, testConfig.PkgRepoDataSourceVar, testConfig.ScopeHelperResources.Cluster.Name)
	}

	return fmt.Sprintf(`
	%s

	resource "time_sleep" "wait_for_3m" {
		create_duration = "180s"

		depends_on = [tanzu-mission-control_cluster.test_cluster]
	}
	

	resource "%s" "%s" {
		name = "%s"

		namespace_name = "%s"

		disabled = %t

		scope {
			cluster {
				name = tanzu-mission-control_cluster.test_cluster.name
			}
		}

		spec {
			imgpkg_bundle {
				image = "%s"
			}
		}

		depends_on = [time_sleep.wait_for_3m]
	}

	data "%s" "%s" {
		name = tanzu-mission-control_package_repository.test_pkg_repository.name

		scope {
			cluster {
				name = tanzu-mission-control_cluster.test_cluster.name
			}
		}
	}
	`, helperBlock, testConfig.PkgRepoResource, testConfig.PkgRepoResourceVar, testConfig.PkgRepoName, testConfig.Namespace, disabled, imageURL, testConfig.PkgRepoResource, testConfig.PkgRepoDataSourceVar)
}

// checkPkgRepositoryDataSourceAttributes checks to get git repository creation.
func (testConfig *testAcceptanceConfig) checkPkgRepositoryDataSourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyPkgRepositoryDataSourceCreation(testConfig.PkgRepositoryDataSourceName),
		resource.TestCheckResourceAttrPair(testConfig.PkgRepositoryDataSourceName, "name", testConfig.PkgRepoResourceName, "name"),
		resource.TestCheckResourceAttrSet(testConfig.PkgRepositoryDataSourceName, "id"),
	}

	check = append(check, MetaDataSourceAttributeCheck(testConfig.PkgRepositoryDataSourceName, testConfig.PkgRepoResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyPkgRepositoryDataSourceCreation(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have source secret resource %s", name)
		}

		return nil
	}
}
