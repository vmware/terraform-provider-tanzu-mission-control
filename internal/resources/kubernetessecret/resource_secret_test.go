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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	secretmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

const (
	secretResourceVar   = "test_secret"
	secretDataSourceVar = "test_data_secret"
	clusterResource     = "tanzu-mission-control_cluster"
	clusterResourceVar  = "tmc_cluster_test"
)

type testAcceptanceConfig struct {
	Provider           *schema.Provider
	SecretResource     string
	SecretResourceVar  string
	SecretResourceName string
	SecretName         string
}

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:           initTestProvider(t),
		SecretResource:     ResourceName,
		SecretResourceVar:  secretResourceVar,
		SecretResourceName: fmt.Sprintf("%s.%s", ResourceName, secretResourceVar),
		SecretName:         acctest.RandomWithPrefix("tf-sc-test"),
	}
}

func TestAcceptanceForSecretResource(t *testing.T) {
	clusterName := acctest.RandomWithPrefix("tf-cluster")

	testConfig := testGetDefaultAcceptanceConfig(t)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestResourceClusterGroupConfigValue(t, clusterName),
				Check: resource.ComposeTestCheckFunc(
					testConfig.checkResourceAttributes(testConfig.Provider, testConfig.SecretResource, clusterName, testConfig.SecretName),
				),
			},
		},
	},
	)
	t.Log("secret resource acceptance test complete!")
}

func (testConfig *testAcceptanceConfig) getTestResourceClusterGroupConfigValue(t *testing.T, clusterName string) string {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		t.Skipf("KUBECONFIG env var is not set: %s", kubeconfigPath)
	}

	return fmt.Sprintf(`
resource "%s" "%s" {
  management_cluster_name = "%s"
  provisioner_name        = "%s"
  name                    = "%s"

  %s

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

`, clusterResource, clusterResourceVar, scope.AttachedValue, scope.AttachedValue, clusterName, testhelper.MetaTemplate, kubeconfigPath,
		testConfig.SecretResource, testConfig.SecretResourceVar, testConfig.SecretName)
}

func (testConfig *testAcceptanceConfig) checkResourceAttributes(provider *schema.Provider, resourceName, clusterName, secretName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifySecretResourceCreation(provider, resourceName, clusterName, secretName),
		resource.TestCheckResourceAttr(testConfig.SecretResourceName, "name", testConfig.SecretName),
	}

	check = append(check, MetaResourceAttributeCheck(testConfig.SecretResourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifySecretResourceCreation(
	provider *schema.Provider,
	resourceName string,
	clusterName string,
	secretName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[testConfig.SecretResourceName]

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
			TLSConfig:        &proxy.TLSConfig{},
		}

		err := config.Setup()
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		fn := &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName{
			Name:                  secretName,
			ClusterName:           clusterName,
			ManagementClusterName: "attached",
			ProvisionerName:       "attached",
			NamespaceName:         "default",
		}

		resp, err := config.TMCConnection.SecretResourceService.SecretResourceServiceGet(fn)
		if err != nil {
			return fmt.Errorf("sceret resource not found: %s", err)
		}

		if resp == nil {
			return fmt.Errorf("sceret resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
