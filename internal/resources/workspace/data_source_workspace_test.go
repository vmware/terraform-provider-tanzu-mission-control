/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package workspace

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	testhelper "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/testing"
)

func TestAcceptanceForWorkspaceDataSource(t *testing.T) {
	var provider = initTestProvider(t)

	workspace := acctest.RandomWithPrefix("tf-ws-test")
	dataSourceName := "data.tmc_workspace.test_data_workspace"
	resourceName := "tmc_workspace.test_workspace"

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTestWorkspaceDataSourceConfigValue(workspace),
				Check: resource.ComposeTestCheckFunc(
					checkDataSourceAttributes(dataSourceName, resourceName),
				),
			},
		},
	})
	t.Log("workspace data source acceptance test complete!")
}

func getTestWorkspaceDataSourceConfigValue(workspaceName string) string {
	return fmt.Sprintf(`
resource "tmc_workspace" "test_workspace" {
  name = "%s"
  %s
}

data "tmc_workspace" "test_data_workspace" {
  name = tmc_workspace.test_workspace.name
}
`, workspaceName, testhelper.MetaTemplate)
}

func checkDataSourceAttributes(dataSourceName, resourceName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyWorkspaceDataSource(dataSourceName),
		resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
		resource.TestCheckResourceAttrSet(dataSourceName, "id"),
	}

	check = append(check, testhelper.MetaDataSourceAttributeCheck(dataSourceName, resourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func verifyWorkspaceDataSource(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have workspace resource %s", name)
		}

		return nil
	}
}
