//go:build gitrepository
// +build gitrepository

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package gitrepository

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	gitrepositoryscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForGitRepositoryDataSource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	t.Log("start git repository data source acceptance tests!")

	// Test case for git repository data source.
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
				Config: testConfig.getTestGitRepositoryDataSourceBasicConfigValue(commonscope.ClusterScope, WithGitImplementation(fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT))),
				Check:  testConfig.checkGitRepositoryDataSourceAttributes(),
			},
			{
				Config: testConfig.getTestGitRepositoryDataSourceBasicConfigValue(commonscope.ClusterScope, WithGitImplementation(fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2)), WithInterval("10m")),
				Check:  testConfig.checkGitRepositoryDataSourceAttributes(),
			},
			{
				Config: testConfig.getTestGitRepositoryDataSourceBasicConfigValue(commonscope.ClusterGroupScope, WithGitImplementation(fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT))),
				Check:  testConfig.checkGitRepositoryDataSourceAttributes(),
			},
			{
				Config: testConfig.getTestGitRepositoryDataSourceBasicConfigValue(commonscope.ClusterGroupScope, WithGitImplementation(fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2)), WithInterval("10m")),
				Check:  testConfig.checkGitRepositoryDataSourceAttributes(),
			},
		},
	},
	)

	t.Log("git repository data source acceptance test completed")
}

func (testConfig *testAcceptanceConfig) getTestGitRepositoryDataSourceBasicConfigValue(scope commonscope.Scope, opts ...OperationOption) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestResourceHelperAndScope(scope, gitrepositoryscope.ScopesAllowed[:])
	gitRepoSpec := testConfig.getTestGitRepositoryResourceSpec(opts...)

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

	data "%s" "%s" {
		name = tanzu-mission-control_git_repository.test_git_repository.name

		namespace_name = "tanzu-continuousdelivery-resources"

		%s
	}
	`, helperBlock, testConfig.GitRepositoryResource, testConfig.GitRepositoryResourceVar, testConfig.GitRepositoryName, scopeBlock, gitRepoSpec, testConfig.GitRepositoryResource, testConfig.GitRepositoryDataSourceVar, scopeBlock)
}

// checkGitRepositoryDataSourceAttributes checks to get git repository creation.
func (testConfig *testAcceptanceConfig) checkGitRepositoryDataSourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyGitRepositoryDataSourceCreation(testConfig.GitRepositoryDataSourceName),
		resource.TestCheckResourceAttrPair(testConfig.GitRepositoryDataSourceName, "name", testConfig.GitRepositoryResourceName, "name"),
		resource.TestCheckResourceAttrSet(testConfig.GitRepositoryDataSourceName, "id"),
	}

	check = append(check, MetaDataSourceAttributeCheck(testConfig.GitRepositoryDataSourceName, testConfig.GitRepositoryResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyGitRepositoryDataSourceCreation(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have source secret resource %s", name)
		}

		return nil
	}
}
