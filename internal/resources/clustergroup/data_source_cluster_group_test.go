/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroup

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAcceptanceForClusterGroupDataSource(t *testing.T) {
	var provider = initTestProvider(t)

	clusterGroup := acctest.RandomWithPrefix("tf-cg-test")
	dataSourceName := "data.tmc_cluster_group.test_data_cluster_group"
	resourceName := "tmc_cluster_group.test_cluster_group"

	resource.Test(t, resource.TestCase{
		PreCheck:          testPreCheck(t),
		ProviderFactories: getTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTestDataSourceClusterGroupConfigValue(clusterGroup),
				Check: resource.ComposeTestCheckFunc(
					verifyClusterGroupDataSource(dataSourceName),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.description", resourceName, "meta.0.description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.labels.key1", resourceName, "meta.0.labels.key1"),
					resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.labels.key2", resourceName, "meta.0.labels.key2"),
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "meta.0.annotations.authoritativeRID"),
					resource.TestCheckResourceAttrSet(dataSourceName, "meta.0.uid"),
				),
			},
		},
	},
	)
	t.Log("cluster group data source acceptance test complete!")
}

func getTestDataSourceClusterGroupConfigValue(clusterGroupName string) string {
	return fmt.Sprintf(`
resource "tmc_cluster_group" "test_cluster_group" {
  name = "%s"
  meta {
    description = "cluster group with description"
    labels = {
      "key1" : "value1"
	  "key2" : "value2"
     }
   }
}

data "tmc_cluster_group" "test_data_cluster_group" {
  name = tmc_cluster_group.test_cluster_group.name
}
`, clusterGroupName)
}

func verifyClusterGroupDataSource(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has cluster group resource %s", name)
		}

		return nil
	}
}
