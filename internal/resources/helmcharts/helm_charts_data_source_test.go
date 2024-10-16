//go:build helmcharts
// +build helmcharts

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmcharts

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func testGetDefaultAcceptanceConfig(t *testing.T) *testAcceptanceConfig {
	return &testAcceptanceConfig{
		Provider:            initTestProvider(t),
		ChartResource:       chartResource,
		ChartDataSourceVar:  chartDataSourceVar,
		ChartMetadataName:   chartMetadataName,
		ChartDataSourceName: fmt.Sprintf("data.%s.%s", ResourceName, chartDataSourceVar),
	}
}

func TestAcceptanceForHelmChratsDataSource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute helm charts tests is not found, run this as a mock test by setting up an http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_HELMCHART_ENV_TEST")
	if !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "dummy.tmc.mock.vmware.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.tanzu.broadcom.com")

		log.Println("Setting up the mock endpoints...")

		testConfig.setupHTTPMocks(t)
	} else {
		// Environment variables with non default values required for a successful call to Cluster Config Service.
		requiredVars := []string{
			"VMW_CLOUD_ENDPOINT",
			"TMC_ENDPOINT",
			"VMW_CLOUD_API_TOKEN",
			"ORG_ID",
		}

		// Check if the required environment variables are set.
		for _, name := range requiredVars {
			if _, found := os.LookupEnv(name); !found {
				t.Errorf("required environment variable '%s' missing", name)
			}
		}
	}

	t.Log("start helm charts data source acceptance tests!")

	// Test case for helm charts data source.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		ExternalProviders: map[string]resource.ExternalProvider{},
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestBasicDataSourceConfigValue(),
				Check:  testConfig.checkDataSourceAttributes(),
			},
		},
	},
	)

	t.Log("helm charts data source acceptance test complete")
}

func (testConfig *testAcceptanceConfig) getTestBasicDataSourceConfigValue() string {
	return fmt.Sprintf(`
	data "%s" "%s" {
		chart_metadata_name = "%s"
	}
	`, testConfig.ChartResource, testConfig.ChartDataSourceVar, testConfig.ChartMetadataName)
}

// checkDataSourceAttributes checks to get helm charts.
func (testConfig *testAcceptanceConfig) checkDataSourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyDataSourceCreation(testConfig.ChartDataSourceName),
		resource.TestCheckResourceAttrSet(testConfig.ChartDataSourceName, "id"),
	}

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyDataSourceCreation(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have helm charts resource %s", name)
		}

		return nil
	}
}
