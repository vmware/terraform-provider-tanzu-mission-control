//go:build provisioner
// +build provisioner

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package provisioner

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

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
	t.Log("provisioner datasource acceptance test complete!")
}

func getTestProvisionerWithDataSourceConfigValue(prvName string) string {
	return fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"
		management_cluster = "%s"
		%s
	}

	data "%s" "%s" {
		provisioners {
			name = tanzu-mission-control_provisioner.provisioner_resource.name
			management_cluster = tanzu-mission-control_provisioner.provisioner_resource.management_cluster
		}
	}
	`, ResourceName, resourceVar, prvName, eksManagementCluster, testhelper.MetaTemplate, ResourceName, dataSourceVar)
}

func checkDataSourceAttributes(dataSourceName, resourceName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyProvisionerDataSource(dataSourceName),
		resource.TestCheckResourceAttrPair(dataSourceName, "provisioners.0.name", resourceName, "name"),
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
		resource.TestCheckResourceAttrPair(dataSourceName, "provisioners.0.meta.0.description", resourceName, "meta.0.description"),
		resource.TestCheckResourceAttrPair(dataSourceName, "provisioners.0.meta.0.labels.key1", resourceName, "meta.0.labels.key1"),
		resource.TestCheckResourceAttrPair(dataSourceName, "provisioners.0.meta.0.labels.key2", resourceName, "meta.0.labels.key2"),
		resource.TestCheckResourceAttrSet(dataSourceName, "provisioners.0.meta.0.uid"),
	}
}
