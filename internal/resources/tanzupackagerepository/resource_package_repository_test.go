/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackagerepository

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	packagerepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackagerepository"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	packagerepositoryscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzupackagerepository/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

// nolint: gosec
const (
	PkgRepoResource      = ResourceName
	PkgRepoResourceVar   = "test_pkg_repository"
	pkgRepoDataSourceVar = "test_data_source_pkg_repository"
	pkgRepoNamePrefix    = "tf-pkg-repository-test"

	imageURL        = "extensions.aws-usw2.tmc-dev.cloud.vmware.com/packages/standard/repo:v2.2.0_update.2"
	updatedImageURL = "extensions.aws-usw2.tmc-dev.cloud.vmware.com/packages/standard/repo:v2.2.0_update.1"
)

type testAcceptanceConfig struct {
	Provider                    *schema.Provider
	PkgRepoResource             string
	PkgRepoResourceVar          string
	PkgRepoResourceName         string
	PkgRepoName                 string
	PkgRepoDataSourceVar        string
	PkgRepositoryDataSourceName string
	ScopeHelperResources        *commonscope.ScopeHelperResources
	Namespace                   string
}

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                    initTestProvider(t),
		PkgRepoResource:             PkgRepoResource,
		PkgRepoResourceVar:          PkgRepoResourceVar,
		PkgRepoResourceName:         fmt.Sprintf("%s.%s", PkgRepoResource, PkgRepoResourceVar),
		PkgRepoName:                 acctest.RandomWithPrefix(pkgRepoNamePrefix),
		ScopeHelperResources:        commonscope.NewScopeHelperResources(),
		Namespace:                   globalRepoNamespace,
		PkgRepoDataSourceVar:        pkgRepoDataSourceVar,
		PkgRepositoryDataSourceName: fmt.Sprintf("data.%s.%s", ResourceName, pkgRepoDataSourceVar),
	}
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_PKGREPO_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_PKGREPO_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

func TestAcceptanceForPackageRepositoryResource(t *testing.T) {
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

	t.Log("start packagerepository resource acceptance tests!")

	// Test case for packagerepository resource.
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
				Config: testConfig.getTestPackageRepositoryResourceBasicConfigValue(commonscope.ClusterScope, false, imageURL),
				Check:  testConfig.checkPackageRepositoryResourceAttributes(commonscope.ClusterScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder...")
						testConfig.setupHTTPMocksUpdate(t)
					}
				},
				Config: testConfig.getTestPackageRepositoryResourceBasicConfigValue(commonscope.ClusterScope, true, updatedImageURL),
				Check:  testConfig.checkPackageRepositoryResourceAttributes(commonscope.ClusterScope),
			},
		},
	},
	)

	t.Log("packagerepository resource acceptance test complete")
}

func (testConfig *testAcceptanceConfig) getTestPackageRepositoryResourceBasicConfigValue(scope commonscope.Scope, disabled bool, imageURL string) string {
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
	`, testConfig.PkgRepoResource, testConfig.PkgRepoResourceVar, testConfig.PkgRepoName, testConfig.Namespace, disabled, testConfig.ScopeHelperResources.Cluster.Name, imageURL)
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
	`, helperBlock, testConfig.PkgRepoResource, testConfig.PkgRepoResourceVar, testConfig.PkgRepoName, testConfig.Namespace, disabled, imageURL)
}

// checkPackageRepositoryResourceAttributes checks for packagerepository creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkPackageRepositoryResourceAttributes(scopeType commonscope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyPackageRepositoryResourceCreation(scopeType),
		resource.TestCheckResourceAttr(testConfig.PkgRepoResourceName, "name", testConfig.PkgRepoName),
	}

	switch scopeType {
	case commonscope.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.PkgRepoResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	case commonscope.ClusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.PkgRepoResourceName, "scope.0.cluster_group.0.name", testConfig.ScopeHelperResources.ClusterGroup.Name))
	case commonscope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(packagerepositoryscope.ScopesAllowed[:], `, `))
	}

	check = append(check, commonscope.MetaResourceAttributeCheck(testConfig.PkgRepoResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyPackageRepositoryResourceCreation(scopeType commonscope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.PkgRepoResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.PkgRepoResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.PkgRepoResourceName)
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
			fn := &packagerepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName{
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				ManagementClusterName: commonscope.AttachedValue,
				Name:                  testConfig.PkgRepoName,
				NamespaceName:         testConfig.Namespace,
				ProvisionerName:       commonscope.AttachedValue,
			}

			resp, err := config.TMCConnection.ClusterPackageRepositoryService.RepositoryResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped packagerepository resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped packagerepository resource is empty, resource: %s", testConfig.PkgRepoResourceName)
			}
		case commonscope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(packagerepositoryscope.ScopesAllowed[:], `, `))
		}

		return nil
	}
}
