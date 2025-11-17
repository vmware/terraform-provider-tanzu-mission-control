//go:build cluster

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

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
				Config: testGetResourceClusterDefinition(t, testhelper.WithClusterName("tf-attach-test-ds"), testhelper.WithDataSourceScript()),
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
		verifyClusterDataSource(testhelper.ClusterDataSourceName),
		resource.TestCheckResourceAttrPair(testhelper.ClusterDataSourceName, "name", testhelper.ClusterResourceName, "name"),
		resource.TestCheckResourceAttrSet(testhelper.ClusterDataSourceName, "id"),
	}

	check = append(check, testhelper.MetaDataSourceAttributeCheck(testhelper.ClusterDataSourceName, testhelper.ClusterResourceName)...)

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
