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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/authctx"
	namespacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/namespace"
	testhelper "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForNamespaceResource(t *testing.T) {
	var provider = initTestProvider(t)

	resourceName := fmt.Sprintf("%s.%s", namespaceResource, namespaceResourceVar)
	namespaceName := acctest.RandomWithPrefix("tf-ns-test")
	clusterName := acctest.RandomWithPrefix("tf-cluster")

	resource.Test(t, resource.TestCase{
		PreCheck:          testPreCheck(t),
		ProviderFactories: getTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTestResourceClusterGroupConfigValue(t, clusterName, namespaceName),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, resourceName, clusterName, namespaceName),
				),
			},
		},
	},
	)
	t.Log("namespace resource acceptance test complete!")
}

func getTestResourceClusterGroupConfigValue(t *testing.T, clusterName, namespaceName string) string {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		t.Skipf("KUBECONFIG env var is not set: %s", kubeconfigPath)
	}

	return fmt.Sprintf(`
resource "%s" "%s" {
  name                    = "%s"

  attach_k8s_cluster {
    kubeconfig_file = "%s"
  }
 
  spec {
    cluster_group = "default"
  }

  ready_wait_timeout = "3m"
}

resource "%s" "%s" {
  name = "%s"
  cluster_name = tanzu-mission-control_cluster.tmc_cluster_test.name
  provisioner_name        = "attached"
  management_cluster_name = "attached"

  %s

  spec {
    workspace_name = "default"
    attach         = false
  }

}

`, clusterResource, clusterResourceVar, clusterName, kubeconfigPath, namespaceResource, namespaceResourceVar, namespaceName, testhelper.MetaTemplate)
}

func checkResourceAttributes(provider *schema.Provider, resourceName, clusterName, namespaceName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyNamespaceResourceCreation(provider, resourceName, clusterName, namespaceName),
		resource.TestCheckResourceAttr(resourceName, "name", namespaceName),
	}

	check = append(check, testhelper.MetaResourceAttributeCheck(resourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func verifyNamespaceResourceCreation(
	provider *schema.Provider,
	resourceName string,
	clusterName string,
	namespaceName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("not found resource %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource %s", resourceName)
		}

		config := authctx.TanzuContext{
			ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
			Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
			VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
		}

		err := config.Setup()
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		fn := &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName{
			Name:                  namespaceName,
			ClusterName:           clusterName,
			ManagementClusterName: "attached",
			ProvisionerName:       "attached",
		}

		resp, err := config.TMCConnection.NamespaceResourceService.ManageV1alpha1NamespaceResourceServiceGet(fn)
		if err != nil {
			return fmt.Errorf("namespace resource not found: %s", err)
		}

		if resp == nil {
			return fmt.Errorf("namespace resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
