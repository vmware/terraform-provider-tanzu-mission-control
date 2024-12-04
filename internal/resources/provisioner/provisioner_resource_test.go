//go:build provisioner
// +build provisioner

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package provisioner

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
	provisionermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/provisioner"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForProvisionerResource(t *testing.T) {
	var provider = initTestProvider(t)

	provisionerResourceName := fmt.Sprintf("%s.%s", ResourceName, resourceVar)
	provisionerName := acctest.RandomWithPrefix("tf-prv-test")

	_, enableProvisionerEnvTest := os.LookupEnv("ENABLE_PROVISIONER_ENV_TEST")

	if !enableProvisionerEnvTest {
		endpoint := "play.abc.def.ghi.com"

		os.Setenv("TF_ACC", "true")
		os.Setenv(authctx.ServerEndpointEnvVar, endpoint)
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.tanzu.broadcom.com")

		setupHTTPMocks(t)
		setUpProvisionerEndPointMocks(t, endpoint, provisionerName)
	} else {
		requiredVars := []string{
			"VMW_CLOUD_ENDPOINT",
			"TMC_ENDPOINT",
			"VMW_CLOUD_API_TOKEN",
		}

		for _, name := range requiredVars {
			if _, found := os.LookupEnv(name); !found {
				t.Errorf("required environment variable '%s' missing", name)
			}
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTestProvisionerWithResourceConfigValue(provisionerName),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, provisionerResourceName, provisionerName),
				),
			},
			{
				PreConfig: func() {
					if !enableProvisionerEnvTest {
						t.Log("Setting up the updated GET mock responder...")
						setupHTTPMocks(t)
						setUpHTTPMockUpdate(t, provisionerName)
					}
				},
				Config: updateTestProvisionerWithResourceConfigValue(provisionerName),
				Check: resource.ComposeTestCheckFunc(
					checkUpdateResourceAttributes(provider, provisionerResourceName, provisionerName),
				),
			},
		},
	})
	t.Log("provisioner resource acceptance test complete!")
}

func checkResourceAttributes(provider *schema.Provider, resourceName, prvName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyProvisionerResourceCreation(provider, resourceName, prvName),
	}

	check = append(check, metaResourceAttributeCheck(resourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func checkUpdateResourceAttributes(provider *schema.Provider, resourceName, prvName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyProvisionerResourceCreation(provider, resourceName, prvName),
	}

	check = append(check, metaUpdateResourceAttributeCheck(resourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func getTestProvisionerWithResourceConfigValue(prvName string) string {
	return fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"
		management_cluster = "%s"
		%s
	}
	`, ResourceName, resourceVar, prvName, eksManagementCluster, testhelper.MetaTemplate)
}

func updateTestProvisionerWithResourceConfigValue(prvName string) string {
	return fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"
		management_cluster = "%s"
		meta {
		description = "resource with updated description"
		labels = {
			"key1" : "value1"
			"key2" : "value2"
			"key3" : "value3"
		}
	  }
	}
	`, ResourceName, resourceVar, prvName, eksManagementCluster)
}

func verifyProvisionerResourceCreation(
	provider *schema.Provider,
	resourceName string,
	provisionerName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found resource: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource: %s", resourceName)
		}

		config := &authctx.TanzuContext{
			ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
			Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
			VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
			TLSConfig:        &proxy.TLSConfig{},
		}

		err := getSetupConfig(config)
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		fn := &provisionermodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerFullName{
			ManagementClusterName: eksManagementCluster,
			Name:                  provisionerName,
		}

		resp, err := config.TMCConnection.ProvisionerResourceService.ProvisionerResourceServiceGet(fn)
		if err != nil {
			return errors.Wrap(err, "provisioner resource not found")
		}

		if resp == nil {
			return errors.Wrapf(err, "provisioner resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_PROVISIONER_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}

func metaResourceAttributeCheck(resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "meta.0.description", "resource with description"),
		resource.TestCheckResourceAttr(resourceName, "meta.0.labels.key1", "value1"),
		resource.TestCheckResourceAttr(resourceName, "meta.0.labels.key2", "value2"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.uid"),
	}
}

func metaUpdateResourceAttributeCheck(resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "meta.0.description", "resource with updated description"),
		resource.TestCheckResourceAttr(resourceName, "meta.0.labels.key1", "value1"),
		resource.TestCheckResourceAttr(resourceName, "meta.0.labels.key2", "value2"),
		resource.TestCheckResourceAttr(resourceName, "meta.0.labels.key3", "value3"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.uid"),
	}
}
