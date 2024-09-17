//go:build tanzupackage
// +build tanzupackage

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackage

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	packagescope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/package/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:             initTestProvider(t),
		PkgResource:          PkgResource,
		PkgName:              pkgName,
		ScopeHelperResources: commonscope.NewScopeHelperResources(),
		PkgDataSourceVar:     pkgDataSourceVar,
		PkgDataSourceName:    fmt.Sprintf("data.%s.%s", ResourceName, pkgDataSourceVar),
	}
}

func TestAcceptanceForPackageDataSource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute package tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_PKG_ENV_TEST")
	if !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "dummy.tmc.mock.vmware.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.tanzu.broadcom.com")

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

	t.Log("start package data source acceptance tests!")

	// Test case for package data source.
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
						t.Skip("KUBECONFIG env var is not set for cluster scoped package acceptance test")
					}
				},
				Config: testConfig.getTestPackageDataSourceBasicConfigValue(commonscope.ClusterScope, imageURL),
				Check:  testConfig.checkPkgDataSourceAttributes(),
			},
		},
	},
	)

	t.Log("package data source acceptance test complete")
}

func (testConfig *testAcceptanceConfig) getTestPackageDataSourceBasicConfigValue(scope commonscope.Scope, imageURL string) string {
	helperBlock, _ := testConfig.ScopeHelperResources.GetTestResourceHelperAndScope(scope, packagescope.ScopeAllowed[:])

	if _, found := os.LookupEnv("ENABLE_PKG_ENV_TEST"); !found {
		return fmt.Sprintf(`
	data "%s" "%s" {
		name = "%s"

		metadata_name = "%s"

		scope {
			cluster {
				name = "%s"
				management_cluster_name = "attached"
				provisioner_name = "attached"
			}
		}
	}
	`, testConfig.PkgResource, testConfig.PkgDataSourceVar, testConfig.PkgName, pkgMetadataName, testConfig.ScopeHelperResources.Cluster.Name)
	}

	return fmt.Sprintf(`
	%s

	resource "time_sleep" "wait_for_3m" {
		create_duration = "180s"

		depends_on = [tanzu-mission-control_cluster.test_cluster]
	}
	

	resource "tanzu-mission-control_package_repository" "test_pkg_repository" {
		name = "test-repo"

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

	resource "time_sleep" "wait_for_1m" {
		create_duration = "60s"

		depends_on = [tanzu-mission-control_package_repository.test_pkg_repository]
	}

	data "%s" "%s" {
		name = "%s"

		metadata_name = "%s"

		scope {
			cluster {
				name = tanzu-mission-control_cluster.test_cluster.name
			}
		}

		depends_on = [time_sleep.wait_for_1m]
	}
	`, helperBlock, imageURL, testConfig.PkgResource, testConfig.PkgDataSourceVar, testConfig.PkgName, pkgMetadataName)
}

// checkPkgDataSourceAttributes checks to get package.
func (testConfig *testAcceptanceConfig) checkPkgDataSourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyPkgDataSourceCreation(testConfig.PkgDataSourceName),
		resource.TestCheckResourceAttrSet(testConfig.PkgDataSourceName, "id"),
	}

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyPkgDataSourceCreation(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have package resource %s", name)
		}

		return nil
	}
}
