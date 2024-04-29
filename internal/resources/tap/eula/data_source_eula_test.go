//go:build tapeula
// +build tapeula

/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tapeula

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
		Provider:                    initTestProvider(t),
		TAPEULAResource:             tapEULAResource,
		TAPEULAResourceVar:          tapEULAResourceVar,
		TAPEULAResourceTAPVersion:   fmt.Sprintf("%s.%s", tapEULAResource, tapEULAResourceVar),
		TAPEULATAPVersion:           tapEULATAPVersion,
		TAPEULADataSourceVar:        tapEULADataSourceVar,
		TAPEULADataSourceTAPVersion: fmt.Sprintf("data.%s.%s", ResourceName, tapEULADataSourceVar),
	}
}

func TestAcceptanceForTAPEULADataSource(t *testing.T) {
	testConfig := testGetDefaultAcceptanceConfig(t)

	// If the flag to execute tap eula tests is not found, run this as a mock test by setting up a http intercept for each endpoint.
	_, found := os.LookupEnv("ENABLE_TAPEULA_ENV_TEST")
	if !found {
		os.Setenv("TF_ACC", "true")
		os.Setenv("TMC_ENDPOINT", "dummy.tmc.mock.vmware.com")
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.cloud.vmware.com")

		log.Println("Setting up the mock endpoints...")
		testConfig.setupHTTPMocks(t)
	} else {
		// Environment variables with non default values required for a successful call to Package Deployment Service.
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

	t.Log("start tap eula data source acceptance tests!")

	// Test case for tap eula data source.
	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(testConfig.Provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testConfig.getTestTAPEULADataSourceBasicConfigValue(),
				Check:  testConfig.checkTAPEULADataSourceAttributes(),
			},
		},
	},
	)

	t.Log("tap eula data source acceptance test completed")
}

func (testConfig *testAcceptanceConfig) getTestTAPEULADataSourceBasicConfigValue() string {
	return fmt.Sprintf(`
	resource "%s" "%s" {
	 tap_version = "%s"
	}

	data "%s" "%s" {
		tap_version = tanzu-mission-control_tap_eula.test_tap_eula.tap_version
	}
	`, testConfig.TAPEULAResource, testConfig.TAPEULAResourceVar, testConfig.TAPEULATAPVersion, testConfig.TAPEULAResource, testConfig.TAPEULADataSourceVar)
}

// checkTAPEULADataSourceAttributes checks for tap eula creation attributes.
func (testConfig *testAcceptanceConfig) checkTAPEULADataSourceAttributes() resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		testConfig.verifyTAPEULADataSourceCreation(testConfig.TAPEULADataSourceTAPVersion),
		resource.TestCheckResourceAttrPair(testConfig.TAPEULADataSourceTAPVersion, "tap_version", testConfig.TAPEULAResourceTAPVersion, "tap_version"),
	}

	return resource.ComposeTestCheckFunc(check...)
}

func (testConfig *testAcceptanceConfig) verifyTAPEULADataSourceCreation(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module does not have tap eula resource %s", name)
		}

		return nil
	}
}
