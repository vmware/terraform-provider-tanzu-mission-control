//go:build clustersecret
// +build clustersecret

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

import (
	"fmt"
	"log"
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
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:           initTestProvider(t),
		SecretResource:     ResourceName,
		SecretResourceVar:  secretResourceVar,
		SecretResourceName: fmt.Sprintf("%s.%s", ResourceName, secretResourceVar),
		SecretName:         acctest.RandomWithPrefix("tf-sc-test"),
		NamespaceName:      "default",
		ClusterName:        acctest.RandomWithPrefix("tf-cluster"),
	}
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_SECRET_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

func TestAcceptanceForSecretResource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute kubernetes secret tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	if _, found := os.LookupEnv("ENABLE_SECRET_ENV_TEST"); !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "dummy.tmc.mock.vmware.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.cloud.vmware.com")

		log.Println("Setting up the mock endpoints...")

		testConfig.setupHTTPMocks(t)
	} else {
		// Environment variables with non default values required for a successful call to MKP
		requiredVars := []string{
			"VMW_CLOUD_ENDPOINT",
			"TMC_ENDPOINT",
			"VMW_CLOUD_API_TOKEN",
			"ORG_ID",
		}

		// Check if the required environment variables are set
		for _, name := range requiredVars {
			if _, found := os.LookupEnv(name); !found {
				t.Errorf("required environment variable '%s' missing", name)
			}
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestResourceClusterGroupConfigValue(t),
				Check: resource.ComposeTestCheckFunc(
					testConfig.checkResourceAttributes(testConfig.Provider, testConfig.SecretResource, testConfig.SecretName),
				),
			},
		},
	},
	)
	t.Log("secret resource acceptance test complete!")
}

func (testConfig *testAcceptanceConfig) getTestResourceClusterGroupConfigValue(t *testing.T, opts ...OperationOption) string {
	secretSpec := testConfig.getTestSecretResourceSpec(opts...)

	if _, found := os.LookupEnv("ENABLE_SECRET_ENV_TEST"); !found {
		return fmt.Sprintf(`

resource "%s" "%s" {
  name = "%s"

  namespace_name = "default"

  scope {
	cluster {
		management_cluster_name = "attached"
		provisioner_name = "attached"
		cluster_name = "%s"
	}
  }

  %s

}

`, testConfig.SecretResource, testConfig.SecretResourceVar, testConfig.SecretName, testConfig.ClusterName, secretSpec)
	}

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

  %s

}

`, clusterResource, clusterResourceVar, commonscope.AttachedValue, commonscope.AttachedValue, testConfig.ClusterName, testhelper.MetaTemplate, kubeconfigPath,
		testConfig.SecretResource, testConfig.SecretResourceVar, testConfig.SecretName, secretSpec)
}

func (testConfig *testAcceptanceConfig) checkResourceAttributes(provider *schema.Provider, resourceName, secretName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifySecretResourceCreation(provider, resourceName, testConfig.ClusterName, secretName),
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

		err := getSetupConfig(&config)
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		fn := &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName{
			Name:                  secretName,
			ClusterName:           clusterName,
			ManagementClusterName: "attached",
			ProvisionerName:       "attached",
			NamespaceName:         testConfig.NamespaceName,
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

type (
	OperationConfig struct {
		username         string
		password         string
		imageRegistryURL string
	}

	OperationOption func(*OperationConfig)
)

func WithUsername(val string) OperationOption {
	return func(config *OperationConfig) {
		config.username = val
	}
}

func WithPassword(val string) OperationOption {
	return func(config *OperationConfig) {
		config.password = val
	}
}

func WithURL(val string) OperationOption {
	return func(config *OperationConfig) {
		config.imageRegistryURL = val
	}
}

// getTestSecretResourceSpec builds the input block for cluster secret resource based a recipe.
func (testConfig *testAcceptanceConfig) getTestSecretResourceSpec(opts ...OperationOption) string {
	cfg := &OperationConfig{
		username:         "someusername",
		password:         "somepassword",
		imageRegistryURL: "someregistryurl",
	}

	for _, o := range opts {
		o(cfg)
	}

	secretSpec := fmt.Sprintf(`  spec {
	docker_config_json {
		username = "%s"
		password = "%s"
		image_registry_url = "%s"
	}
  }`, cfg.username, cfg.password, cfg.imageRegistryURL)

	return secretSpec
}
