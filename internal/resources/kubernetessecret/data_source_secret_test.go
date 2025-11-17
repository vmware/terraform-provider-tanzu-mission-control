//go:build clustersecret

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kubernetessecret

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	secretscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForSecretResourceDataSource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	t.Log("start cluster secret data source acceptance tests!", testConfig.ScopeHelperResources.Cluster.Name)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestDataSourceBasicConfigValue(commonscope.ClusterGroupScope),
				Check:  testConfig.checkDataSourceAttributes(),
			},
			{
				Config: testConfig.getTestDataSourceBasicConfigValue(commonscope.ClusterScope),
				Check:  testConfig.checkDataSourceAttributes(),
			},
		},
	})
	t.Log("secret data source acceptance test complete!")
}

func (testConfig *testAcceptanceConfig) getTestDataSourceBasicConfigValue(scope commonscope.Scope) string {
	helperBlock, scopeBlock := testConfig.ScopeHelperResources.GetTestResourceHelperAndScope(scope, secretscope.ScopesAllowed[:])

	return fmt.Sprintf(`
	%s

	resource "%s" "%s" {
	 name = "%s"

	 namespace_name = "default"

	 %s

	 spec {
		docker_config_json {
			username = "someusername"
			password = "somepassword"
			image_registry_url = "someregistryurl"
		}
	  }
	}

	data "%s" "%s" {
		name = tanzu-mission-control_kubernetes_secret.test_secret.name

		namespace_name = "default"

		%s
	}
	`, helperBlock, testConfig.SecretResource, testConfig.SecretResourceVar, testConfig.SecretName, scopeBlock, testConfig.SecretResource, testConfig.DataSourceVar, scopeBlock)
}

func (testConfig *testAcceptanceConfig) checkDataSourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyDataSourceCreation(testConfig.DataSourceName),
		resource.TestCheckResourceAttrPair(testConfig.DataSourceName, "name", testConfig.SecretResourceName, "name"),
		resource.TestCheckResourceAttrSet(testConfig.DataSourceName, "id"),
	}

	check = append(check, MetaDataSourceAttributeCheck(testConfig.DataSourceName, testConfig.SecretResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyDataSourceCreation(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have source secret resource %s", name)
		}

		return nil
	}
}
