/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespace

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	namespaceResource      = "tanzu_mission_control_namespace"
	namespaceResourceVar   = "test_namespace"
	namespaceDataSourceVar = "test_data_namespace"

	clusterResource    = "tanzu_mission_control_cluster"
	clusterResourceVar = "tmc_cluster_test"
)

func TestAcceptanceForNamespaceDataSource(t *testing.T) {
	var provider = initTestProvider(t)

	namespaceName := acctest.RandomWithPrefix("tf-ns-test")
	clusterName := acctest.RandomWithPrefix("tf-cluster")
	dataSourceName := fmt.Sprintf("data.%s.%s", namespaceResource, namespaceDataSourceVar)
	resourceName := fmt.Sprintf("%s.%s", namespaceResource, namespaceResourceVar)

	resource.Test(t, resource.TestCase{
		PreCheck:          testPreCheck(t),
		ProviderFactories: getTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTestNamespaceDataSourceConfigValue(t, clusterName, namespaceName, testhelper.MetaTemplate),
				Check: resource.ComposeTestCheckFunc(
					checkDataSourceAttributes(dataSourceName, resourceName),
				),
			},
		},
	})
	t.Log("namespace data source acceptance test complete!")
}

func getTestNamespaceDataSourceConfigValue(t *testing.T, clusterName, namespaceName, meta string) string {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		t.Skipf("KUBECONFIG env var is not set: %s", kubeconfigPath)
	}

	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"

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
  cluster_name = tanzu_mission_control_cluster.tmc_cluster_test.name
  
  spec {
    workspace_name = "default"
    attach         = false
  }
}

data "%s" "%s" {
  name = tanzu_mission_control_namespace.test_namespace.name
  cluster_name = tanzu_mission_control_namespace.test_namespace.cluster_name
}
`, clusterResource, clusterResourceVar, clusterName, kubeconfigPath, meta, namespaceResource,
		namespaceResourceVar, namespaceName, namespaceResource, namespaceDataSourceVar)
}

func checkDataSourceAttributes(dataSourceName, resourceName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyNamespaceDataSource(dataSourceName),
		resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
		resource.TestCheckResourceAttrSet(dataSourceName, "id"),
	}

	check = append(check, testhelper.MetaDataSourceAttributeCheck(dataSourceName, resourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func verifyNamespaceDataSource(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have namespace resource %s", name)
		}

		return nil
	}
}
