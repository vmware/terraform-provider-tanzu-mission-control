//go:build provisioner
// +build provisioner

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package provisioner

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForProvisionerDataSource(t *testing.T) {
	var provider = initTestProvider(t)

	provisionerResourceName := fmt.Sprintf("%s.%s", ResourceName, resourceVar)
	provisionerName := acctest.RandomWithPrefix("tf-prv-test")

	provisionerDataSource := fmt.Sprintf("data.%s.%s", ResourceName, dataSourceVar)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTestProvisionerWithDataSourceConfigValue(provisionerName),
				Check: resource.ComposeTestCheckFunc(
					checkDataSourceAttributes(provisionerDataSource, provisionerResourceName),
				),
			},
		},
	})
}

func getTestProvisionerWithDataSourceConfigValue(prvName string) string {
	prov_rsrc := fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"
		management_cluster = "%s"
		%s
	}

	data "%s" "%s" {
		name = tanzu-mission-control_provisioner.provisioner_resource.name
		management_cluster = tanzu-mission-control_provisioner.provisioner_resource.management_cluster
	}
	`, ResourceName, resourceVar, prvName, eksManagementCluster, testhelper.MetaTemplate, ResourceName, dataSourceVar)
	fmt.Println("resource defn is: ", prov_rsrc)
	return prov_rsrc
}

func checkDataSourceAttributes(dataSourceName, resourceName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyProvisionerDataSource(dataSourceName),
		resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
		resource.TestCheckResourceAttrSet(dataSourceName, "id"),
	}

	check = append(check, metaDataSourceAttributeCheck(dataSourceName, resourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func verifyProvisionerDataSource(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have provisioner resource %s", name)
		}

		return nil
	}
}

func metaDataSourceAttributeCheck(dataSourceName, resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.description", resourceName, "meta.0.description"),
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.labels.key1", resourceName, "meta.0.labels.key1"),
		resource.TestCheckResourceAttrPair(dataSourceName, "meta.0.labels.key2", resourceName, "meta.0.labels.key2"),
		resource.TestCheckResourceAttrSet(dataSourceName, "meta.0.uid"),
	}
}
