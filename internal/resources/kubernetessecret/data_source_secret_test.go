//go:build clustersecret
// +build clustersecret

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForSecretDataSource(t *testing.T) {
	var provider = initTestProvider(t)

	secretName := acctest.RandomWithPrefix("tf-sct-test")
	clusterName := acctest.RandomWithPrefix("tf-cluster-data")
	dataSourceName := fmt.Sprintf("data.%s.%s", ResourceName, secretDataSourceVar)
	resourceName := fmt.Sprintf("%s.%s", ResourceName, secretResourceVar)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTestSecretDataSourceConfigValue(t, clusterName, secretName, testhelper.MetaTemplate),
				Check: resource.ComposeTestCheckFunc(
					checkDataSourceAttributes(dataSourceName, resourceName),
				),
			},
		},
	})
	t.Log("secret data source acceptance test complete!")
}

func getTestSecretDataSourceConfigValue(t *testing.T, clusterName, secretName, meta string) string {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		t.Skipf("KUBECONFIG env var is not set: %s", kubeconfigPath)
	}

	return fmt.Sprintf(`
resource "%s" "%s" {
	management_cluster_name = "attached"
	provisioner_name        = "attached"
	name                    = "%s"

  attach_k8s_cluster {
    kubeconfig_file = "%s"
  }

  %s

  spec {
    cluster_group = "default"
  }

  ready_wait_timeout = "3m"
}

resource "%s" "%s" {
  name = "%s"

  namespace_name = "default"

  scope {
	cluster {
		management_cluster_name = "attached"
		provisioner_name = "attached"
		cluster_name = tanzu-mission-control_cluster.tmc_cluster_test.name
	}
  }

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

  scope {
	cluster {
		management_cluster_name = "attached"
		provisioner_name = "attached"
		cluster_name = tanzu-mission-control_kubernetes_secret.test_secret.scope[0].cluster[0].cluster_name
	}
  }
}
`, clusterResource, clusterResourceVar, clusterName, kubeconfigPath, meta, ResourceName,
		secretResourceVar, secretName, ResourceName, secretDataSourceVar)
}

func checkDataSourceAttributes(dataSourceName, resourceName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifySecretDataSource(dataSourceName),
		resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
		resource.TestCheckResourceAttrSet(dataSourceName, "id"),
	}

	check = append(check, MetaDataSourceAttributeCheck(dataSourceName, resourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func verifySecretDataSource(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have secret resource %s", name)
		}

		return nil
	}
}
