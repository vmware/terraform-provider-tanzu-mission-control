/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForAttachClusterDataSource(t *testing.T) {
	var provider = initTestProvider(t)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testGetResourceClusterDefinition(t, withClusterName("tf-attach-test-ds"), withDataSourceScript()),
				Check: resource.ComposeTestCheckFunc(
					checkDataSourceAttributes(),
				),
			},
		},
	})
	t.Log("cluster data source acceptance test complete!")
}

func checkDataSourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyClusterDataSource(dataSourceName),
		resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
		resource.TestCheckResourceAttrSet(dataSourceName, "id"),
	}

	check = append(check, testhelper.MetaDataSourceAttributeCheck(dataSourceName, resourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func verifyClusterDataSource(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have cluster resource %s", name)
		}

		return nil
	}
}
