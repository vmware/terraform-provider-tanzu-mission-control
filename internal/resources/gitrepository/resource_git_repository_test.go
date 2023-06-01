/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package gitrepository

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
	}
}

func TestAcceptanceForGitRepositoryResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	t.Log("start git repository resource acceptance tests!")

	// Test case for git repository resource for GO_GIT git implementation type.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped git repository acceptance test")
					}
				},
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterScope, fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT)),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterScope),
			},
			{
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterScope, fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT), WithInterval("10m")),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterScope),
			},
			{
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterGroupScope, fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT)),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterGroupScope),
			},
			{
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterGroupScope, fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT), WithInterval("10m")),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterGroupScope),
			},
		},
	},
	)

	t.Log("git repository resource acceptance test complete for GO_GIT git implementation type")

	// Test case for git repository resource for LIB_GIT2 git implementation type.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterGroupScope, fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2)),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterGroupScope),
			},
			{
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterGroupScope, fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2), WithInterval("10m")),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterGroupScope),
			},
			{
				PreConfig: func() {
					if testConfig.ScopeHelperResources.Cluster.KubeConfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for cluster scoped git repository acceptance test")
					}
				},
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterScope, fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2)),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterScope),
			},
			{
				Config: testConfig.getTestGitRepositoryResourceBasicConfigValue(commonscope.ClusterScope, fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2), WithInterval("10m")),
				Check:  testConfig.checkGitRepositoryResourceAttributes(commonscope.ClusterScope),
			},
		},
	},
	)

	t.Log("git repository resource acceptance test complete for LIB_GIT2 git implementation type")
}

func (testConfig *testAcceptanceConfig) getTestGitRepositoryResourceBasicConfigValue(scope commonscope.Scope, gitImplementation string, opts ...OperationOption) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestResourceHelperAndScope(scope, gitrepositoryscope.ScopesAllowed[:])
	gitRepoSpec := testConfig.getTestGitRepositoryResourceSpec(gitImplementation, opts...)

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
		interval string
	}

	OperationOption func(*OperationConfig)
)

func WithInterval(val string) OperationOption {
	return func(config *OperationConfig) {
		config.interval = val
	}
}

// getTestGitRepositoryResourceSpec builds the input block for git repository resource based a recipe.
func (testConfig *testAcceptanceConfig) getTestGitRepositoryResourceSpec(allowedCredential string, opts ...OperationOption) string {
	cfg := &OperationConfig{
		interval: "5m",
	}

	for _, o := range opts {
		o(cfg)
	}

	var gitRepoSpec string

	switch allowedCredential {
	case fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT):
		gitRepoSpec = fmt.Sprintf(`
		url = "https://github.com/tmc-build-integrations/sample-update-configmap"
		interval = "%s"
		ref {
			branch = "master"
		}
	`, cfg.interval)
	case fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2):
		gitRepoSpec = fmt.Sprintf(`
		url = "https://github.com/tmc-build-integrations/sample-update-configmap"
		interval = "%s"
		git_implementation = "LIB_GIT2"
		ref {
			branch = "master"
		}
	`, cfg.interval)
	}

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

		err := config.Setup()
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		switch scopeType {
		case commonscope.ClusterScope:
			fn := &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName{
				ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
				ManagementClusterName: commonscope.AttachedValue,
				Name:                  testConfig.GitRepositoryName,
				NamespaceName:         "tanzu-continuousdelivery-resources",
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
				NamespaceName:    "tanzu-continuousdelivery-resources",
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
