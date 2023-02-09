/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package credential

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForClusterGroupDataSource(t *testing.T) {
	var provider = initTestProvider(t)

	cred := acctest.RandomWithPrefix("tf-cred-test")
	dataSourceName := fmt.Sprintf("data.%s.%s", credentialResource, credentialDataSourceVar)
	resourceName := fmt.Sprintf("%s.%s", credentialResource, credentialResourceVar)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTestDataSourceCredConfigValue(cred, testhelper.MetaTemplate),
				Check: resource.ComposeTestCheckFunc(
					checkDataSourceAttributes(dataSourceName, resourceName),
				),
			},
		},
	},
	)
	t.Log("credential data source acceptance test complete!")
}

func getTestDataSourceCredConfigValue(credName, meta string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"
  %s

spec {
    capability = "MANAGED_K8S_PROVIDER"
	provider = "AWS_EKS"
    data {
		aws_credential {
			account_id = ""
			generic_credential = ""
			iam_role{
				arn = "arn:aws:iam::4987398738934:role/clusterlifecycle-test.tmc.cloud.vmware.com"
				ext_id =""
			}
		}
	}
 }

}

data "%s" "%s" {
  name = tanzu_mission_control_credential.test_credential.name
}
`, credentialResource, credentialResourceVar, credName, meta, credentialResource, credentialDataSourceVar)
}

func checkDataSourceAttributes(dataSourceName, resourceName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyClusterGroupDataSource(dataSourceName),
		resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
		resource.TestCheckResourceAttrSet(dataSourceName, "id"),
	}

	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.uid"),
	}

	check = append(check, checks...)

	return resource.ComposeTestCheckFunc(check...)
}

func verifyClusterGroupDataSource(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have credential resource %s", name)
		}

		return nil
	}
}
