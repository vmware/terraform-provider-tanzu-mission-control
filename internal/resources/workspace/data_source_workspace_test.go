//go:build workspace
// +build workspace

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package workspace

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	workspaceResource      = "tanzu-mission-control_workspace"
	workspaceResourceVar   = "test_workspace"
	workspaceDataSourceVar = "test_data_workspace"
)

func TestAcceptanceForWorkspaceDataSource(t *testing.T) {
	var provider = initTestProvider(t)

	workspace := acctest.RandomWithPrefix("tf-ws-test")
	dataSourceName := fmt.Sprintf("data.%s.%s", workspaceResource, workspaceDataSourceVar)
	resourceName := fmt.Sprintf("%s.%s", workspaceResource, workspaceResourceVar)

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
resource "%s" "%s" {
  name = "%s"
  %s
}

data "%s" "%s" {
  name = tanzu-mission-control_workspace.test_workspace.name
}
`, workspaceResource, workspaceResourceVar, workspaceName, testhelper.MetaTemplate, workspaceResource, workspaceDataSourceVar)
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
