/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package gitrepository

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
	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
	gitrepositoryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	gitrepositoryscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	gitRepositoryResource      = ResourceName
	gitRepositoryResourceVar   = "test_git_repository"
	gitRepositoryDataSourceVar = "test_data_source_git_repository"
	gitRepositoryNamePrefix    = "tf-gr-test"
)

type testAcceptanceConfig struct {
	Provider                    *schema.Provider
	GitRepositoryResource       string
	GitRepositoryResourceVar    string
	GitRepositoryResourceName   string
	GitRepositoryName           string
	ScopeHelperResources        *commonscope.ScopeHelperResources
	GitRepositoryDataSourceVar  string
	GitRepositoryDataSourceName string
	Namespace                   string
}

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:                    initTestProvider(t),
		GitRepositoryResource:       gitRepositoryResource,
		GitRepositoryResourceVar:    gitRepositoryResourceVar,
		GitRepositoryResourceName:   fmt.Sprintf("%s.%s", gitRepositoryResource, gitRepositoryResourceVar),
		GitRepositoryName:           acctest.RandomWithPrefix(gitRepositoryNamePrefix),
		ScopeHelperResources:        commonscope.NewScopeHelperResources(),
		GitRepositoryDataSourceVar:  gitRepositoryDataSourceVar,
		GitRepositoryDataSourceName: fmt.Sprintf("data.%s.%s", ResourceName, gitRepositoryDataSourceVar),
		Namespace:                   "tanzu-continuousdelivery-resources",
	}
}

func getConfigureContextFunc() func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if _, found := os.LookupEnv("ENABLE_GITREPO_ENV_TEST"); !found {
		return authctx.ProviderConfigureContextWithDefaultTransportForTesting
	}

	return authctx.ProviderConfigureContext
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_GITREPO_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

func TestAcceptanceForGitRepositoryResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute git repository tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_GITREPO_ENV_TEST")
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

	t.Log("start git repository resource acceptance tests!")

	// Test case for git repository resource.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" && found {
						t.Skip("KUBECONFIG env var is not set for cluster scoped git repository acceptance test")
					}
				},
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterScope, WithGitImplementation(fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT))),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder for cluster scope...")
						testConfig.setupHTTPMocksUpdate(t, commonscope.ClusterScope)
					}
				},
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterScope, WithGitImplementation(fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2)), WithInterval("10m")),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterScope),
			},
			{
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterGroupScope, WithGitImplementation(fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT))),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if !found {
						t.Log("Setting up the updated GET mock responder for cluster group scope...")
						testConfig.setupHTTPMocksUpdate(t, commonscope.ClusterGroupScope)
					}
				},
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterGroupScope, WithGitImplementation(fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2)), WithInterval("10m")),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterGroupScope),
			},
		},
	},
	)

	t.Log("git repository resource acceptance test completed")
}

func (testConfig *testAcceptanceConfig) getTestGitRepositoryResourceBasicConfigValue(scope commonscope.Scope, opts ...OperationOption) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestResourceHelperAndScope(scope, gitrepositoryscope.ScopesAllowed[:])
	gitRepoSpec := testConfig.getTestGitRepositoryResourceSpec(opts...)

	if _, found := os.LookupEnv("ENABLE_GITREPO_ENV_TEST"); !found {
		clStr := fmt.Sprintf(`
	resource "%s" "%s" {
	 name = "%s"

	 namespace_name = "tanzu-continuousdelivery-resources"
	
	 scope {
		cluster {
			name = "%s"
			management_cluster_name = "attached"
			provisioner_name = "attached"
		}
	 }
	
	 spec {
	   %s
	 }
	}
	`, testConfig.GitRepositoryResource, testConfig.GitRepositoryResourceVar, testConfig.GitRepositoryName, testConfig.ScopeHelperResources.Cluster.Name, gitRepoSpec)

		cgStr := fmt.Sprintf(`
	resource "%s" "%s" {
	 name = "%s"

	 namespace_name = "tanzu-continuousdelivery-resources"
	
	 scope {
		cluster_group {
			name = "%s"
		}
	 }
	
	 spec {
	   %s
	 }
	}
	`, testConfig.GitRepositoryResource, testConfig.GitRepositoryResourceVar, testConfig.GitRepositoryName, testConfig.ScopeHelperResources.ClusterGroup.Name, gitRepoSpec)

		switch scope {
		case commonscope.ClusterScope:
			return clStr
		case commonscope.ClusterGroupScope:
			return cgStr
		default:
			return ""
		}
	}

	return fmt.Sprintf(`
	%s

	resource "%s" "%s" {
	 name = "%s"

	 namespace_name = "tanzu-continuousdelivery-resources"

	 %s

	 spec {
	   %s
	 }
	}
	`, helperBlock, testConfig.GitRepositoryResource, testConfig.GitRepositoryResourceVar, testConfig.GitRepositoryName, scopeBlock, gitRepoSpec)
}

type (
	OperationConfig struct {
		interval          string
		gitImplementation string
	}

	OperationOption func(*OperationConfig)
)

func WithInterval(val string) OperationOption {
	return func(config *OperationConfig) {
		config.interval = val
	}
}

func WithGitImplementation(val string) OperationOption {
	return func(config *OperationConfig) {
		config.gitImplementation = val
	}
}

// getTestGitRepositoryResourceSpec builds the input block for git repository resource.
func (testConfig *testAcceptanceConfig) getTestGitRepositoryResourceSpec(opts ...OperationOption) string {
	cfg := &OperationConfig{
		interval:          "5m",
		gitImplementation: fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT),
	}

	for _, o := range opts {
		o(cfg)
	}

	gitRepoSpec := fmt.Sprintf(`
		url = "https://github.com/tmc-build-integrations/sample-update-configmap"
		git_implementation = "%s"
		interval = "%s"
		ref {
			branch = "master"
		}
	`, cfg.gitImplementation, cfg.interval)

	return gitRepoSpec
}

// checkGitRepositoryResourceAttributes checks for git repository creation along with meta attributes.
func (testConfig *testAcceptanceConfig) checkGitRepositoryResourceAttributes(scopeType commonscope.Scope) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyGitRepositoryResourceCreation(scopeType),
		resource.TestCheckResourceAttr(testConfig.GitRepositoryResourceName, "name", testConfig.GitRepositoryName),
	}

	switch scopeType {
	case commonscope.ClusterScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.GitRepositoryResourceName, "scope.0.cluster.0.name", testConfig.ScopeHelperResources.Cluster.Name))
	case commonscope.ClusterGroupScope:
		check = append(check, resource.TestCheckResourceAttr(testConfig.GitRepositoryResourceName, "scope.0.cluster_group.0.name", testConfig.ScopeHelperResources.ClusterGroup.Name))
	case commonscope.UnknownScope:
		log.Printf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(gitrepositoryscope.ScopesAllowed[:], `, `))
	}

	check = append(check, commonscope.MetaResourceAttributeCheck(testConfig.GitRepositoryResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyGitRepositoryResourceCreation(scopeType commonscope.Scope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testConfig.Provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.GitRepositoryResourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", testConfig.GitRepositoryResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", testConfig.GitRepositoryResourceName)
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
			fn := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName{
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				ManagementClusterName: commonscope.AttachedValue,
				Name:                  testConfig.GitRepositoryName,
				NamespaceName:         testConfig.Namespace,
				ProvisionerName:       commonscope.AttachedValue,
			}

			resp, err := config.TMCConnection.ClusterGitRepositoryResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdGitrepositoryResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster scoped git repository resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster scoped git repository resource is empty, resource: %s", testConfig.GitRepositoryResourceName)
			}
		case commonscope.ClusterGroupScope:
			fn := &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName{
				ClusterGroupName: testConfig.ScopeHelperResources.ClusterGroup.Name,
				Name:             testConfig.GitRepositoryName,
				NamespaceName:    testConfig.Namespace,
			}

			resp, err := config.TMCConnection.ClusterGroupGitRepositoryResourceService.VmwareTanzuManageV1alpha1ClustergroupFluxcdGitrepositoryResourceServiceGet(fn)
			if err != nil {
				return errors.Wrap(err, "cluster group scoped git repository resource not found")
			}

			if resp == nil {
				return errors.Wrapf(err, "cluster group scoped git repository resource is empty, resource: %s", testConfig.GitRepositoryResourceName)
			}
		case commonscope.UnknownScope:
			return errors.Errorf("[ERROR]: No valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(gitrepositoryscope.ScopesAllowed[:], `, `))
		}

		return nil
	}
}
