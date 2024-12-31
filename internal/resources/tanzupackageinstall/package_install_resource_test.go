//go:build packageinstall
// +build packageinstall

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackageinstall

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	packageinstallclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackageinstall"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	packageinstallscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackageinstall/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:               initTestProvider(t),
		PkgInstallResource:     pkgInstallResource,
		PkgInstallResourceVar:  pkgInstallResourceVar,
		PkgInstallResourceName: fmt.Sprintf("%s.%s", pkgInstallResource, pkgInstallResourceVar),
		PkgInstallName:         acctest.RandomWithPrefix(pkgInstallNamePrefix),
		PkgRepoName:            acctest.RandomWithPrefix(pkgRepoNamePrefix),
		PkgName1:               pkgName1,
		PkgName2:               pkgName2,
		ScopeHelperResources:   commonscope.NewScopeHelperResources(),
		Namespace:              acctest.RandomWithPrefix(namespaceNamePrefix),
	}
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_PKGINS_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

func TestAcceptanceForPackageInstallResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute package install tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_PKGINS_ENV_TEST")
	if !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "dummy.tmc.mock.vmware.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.tanzu.broadcom.com")

		log.Println("Setting up the mock endpoints...")

		testConfig.setupHTTPMocks(t)
	} else {
		// Environment variables with non default values required for a successful call to package deployment service.
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

	t.Log("start package install resource acceptance tests!")

	// Test case for package install resource.
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
						t.Skip("KUBECONFIG env var is not set for cluster scoped package install acceptance test")
					}
				},
				Config: testConfig.getTestPackageInstallResourceBasicConfigValue(commonscope.ClusterScope, "2.0.0", false),
				Check:  testConfig.checkPackageInstallResourceAttributes(commonscope.ClusterScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder...")
						testConfig.setupHTTPMocksUpdate(t)
					}
				},
				Config: testConfig.getTestPackageInstallResourceBasicConfigValue(commonscope.ClusterScope, constraints, false),
				Check:  testConfig.checkPackageInstallResourceAttributes(commonscope.ClusterScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Reset GET mock responder...")
						testConfig.setupHTTPMocks(t)
					}

					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" && found {
						t.Skip("KUBECONFIG env var is not set for cluster scoped package install acceptance test")
					}
				},
				Config: testConfig.getTestPackageInstallResourceBasicConfigValue(commonscope.ClusterScope, "2.0.0", true),
				Check:  testConfig.checkPackageInstallResourceAttributes(commonscope.ClusterScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder...")
						testConfig.setupHTTPMocksUpdate(t)
					}
				},
				Config: testConfig.getTestPackageInstallResourceBasicConfigValue(commonscope.ClusterScope, constraints, true),
				Check:  testConfig.checkPackageInstallResourceAttributes(commonscope.ClusterScope),
			},
		},
	},
	)

	t.Log("package install resource acceptance test complete")
}

// nolint: unparam
func (testConfig *testAcceptanceConfig) getTestPackageInstallResourceBasicConfigValue(scope commonscope.Scope, constraints string, inlineValuesFromFile bool) string {
	helperBlock, _ := testConfig.ScopeHelperResources.GetTestResourceHelperAndScope(scope, packageinstallscope.ScopesAllowed[:])

	inlineValuesForPackageInstall := "inline_values = { \"bar\" : \"foo\" }"
	if inlineValuesFromFile {
		inlineValuesForPackageInstall = "path_to_inline_values = \"test.yaml\""
	}

	if _, found := os.LookupEnv("ENABLE_PKGINS_ENV_TEST"); !found {
		return fmt.Sprintf(`
		resource "%s" "%s" {
			name = "%s"

			namespace = "%s"

			scope {
				cluster {
					name = "%s"
				}
			}

			spec {
				package_ref {
					package_metadata_name = "pkg.test.carvel.dev"
					version_selection {
						constraints = "%s"
					}
				}

				%s
			}
		}
	`, testConfig.PkgInstallResource, testConfig.PkgInstallResourceVar, testConfig.PkgInstallName, testConfig.Namespace, testConfig.ScopeHelperResources.Cluster.Name, constraints, inlineValuesForPackageInstall)
	}

	return fmt.Sprintf(`
	%s

	resource "time_sleep" "wait_for_3m" {
		duration = "60s"

		depends_on = [tanzu-mission-control_cluster.test_cluster]
	}


	resource "%s" "%s" {
		name = "%s"

		scope {
			cluster {
				name = tanzu-mission-control_cluster.test_cluster.name
			}
		}

		spec {
			imgpkg_bundle {
				image = "projects.registry.vmware.com/tmc/build-integrations/package/repository/e2e-test-unauth-repo@sha256:87a5f7e0c44523fbc35a9432c657bebce246138bbd0f16d57f5615933ceef632"
			}
		}

		depends_on = [time_sleep.wait_for_3m]
	}

	resource "time_sleep" "wait_for_2m" {
		duration = "40s"

		depends_on = [tanzu-mission-control_package_repository.test_pkg_repository]
	}

	resource "%s" "%s" {
		name = "%s"

		namespace = "%s"

		scope {
			cluster {
				name = tanzu-mission-control_cluster.test_cluster.name
			}
		}

		spec {
			package_ref {
				package_metadata_name = "pkg.test.carvel.dev"
				version_selection {
					constraints = "%s"
				}
			}

			%s
		}

		depends_on = [time_sleep.wait_for_2m]
	}
	`, helperBlock, pkgRepoResource, pkgRepoResourceVar, testConfig.PkgRepoName,
		testConfig.PkgInstallResource, testConfig.PkgInstallResourceVar, testConfig.PkgInstallName,
		testConfig.Namespace, constraints, inlineValuesForPackageInstall)
}

// nolint: unparam
// checkPackageInstallResourceAttributes checks for package install creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkPackageInstallResourceAttributes(scopeType commonscope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyPackageInstallResourceCreation(scopeType),
		resource.TestCheckResourceAttr(testConfig.PkgInstallResourceName, "name", testConfig.PkgInstallName),
	}

	switch scopeType {
	case commonscope.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.PkgInstallResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	case commonscope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(packageinstallscope.ScopesAllowed[:], `, `))
	}

	check = append(check, commonscope.MetaResourceAttributeCheck(testConfig.PkgInstallResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyPackageInstallResourceCreation(scopeType commonscope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.PkgInstallResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.PkgInstallResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.PkgInstallResourceName)
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
		case commonscope.ClusterScope:
			fn := &packageinstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName{
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				ManagementClusterName: commonscope.AttachedValue,
				Name:                  testConfig.PkgInstallName,
				NamespaceName:         testConfig.Namespace,
				ProvisionerName:       commonscope.AttachedValue,
			}

			resp, err := config.TMCConnection.PackageInstallResourceService.InstallResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped package install resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped package install resource is empty, resource: %s", testConfig.PkgInstallResourceName)
			}
		case commonscope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(packageinstallscope.ScopesAllowed[:], `, `))
		}

		return nil
	}
}
