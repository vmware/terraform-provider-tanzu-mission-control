//go:build targetlocation

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package targetlocationtests

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	targetlocationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/targetlocation"
	targetlocationres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/targetlocation"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

var (
	context = authctx.TanzuContext{
		ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
		Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
		VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
		TLSConfig:        &proxy.TLSConfig{},
	}
)

func TestAcceptanceTargetLocationResource(t *testing.T) {
	err := context.Setup()

	if err != nil {
		t.Error(errors.Wrap(err, "unable to set the context"))
		t.FailNow()
	}

	tmcManagedCredentialsName, tmcManagedCredentialsExist := os.LookupEnv(tmcManagedCredentialsEnv)
	azureCredentialsName, azureCredentialsExist := os.LookupEnv(azureCredentialsNameEnv)

	if !tmcManagedCredentialsExist {
		t.Error("TMC Managed credentials name is missing!")
		t.FailNow()
	}

	if !azureCredentialsExist {
		t.Error("Azure credentials name is missing!")
		t.FailNow()
	}

	var (
		tfResourceConfigBuilder   = InitResourceTFConfigBuilder(testScopeHelper, RsFullBuild)
		tfDataSourceConfigBuilder = InitDataSourceTFConfigBuilder(testScopeHelper, tfResourceConfigBuilder, DsFullBuild, tmcManagedCredentialsName)
		provider                  = initTestProvider(t)
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: tfResourceConfigBuilder.GetAWSSelfManagedTargetLocationConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(AwsSelfManagedResourceFullName, "name", TargetLocationAWSSelfManagedName),
					verifyTargetLocationResourceCreation(provider, AwsSelfManagedResourceFullName, TargetLocationAWSSelfManagedName),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetAzureSelfManagedTargetLocationConfig(azureCredentialsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(AzureSelfManagedResourceFullName, "name", TargetLocationAzureSelfManagedName),
					verifyTargetLocationResourceCreation(provider, AzureSelfManagedResourceFullName, TargetLocationAzureSelfManagedName),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetTMCManagedTargetLocationConfig(tmcManagedCredentialsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(TmcManagedResourceFullName, "name", TargetLocationTMCManagedName),
					verifyTargetLocationResourceCreation(provider, TmcManagedResourceFullName, TargetLocationTMCManagedName),
				),
			},
			{
				Config: tfDataSourceConfigBuilder.GetProviderTargetLocationDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					verifyTargetLocationDataSource(provider, ProviderDataSourceFullName, TargetLocationTMCManagedName),
				),
			},
			{
				Config: tfDataSourceConfigBuilder.GetClusterTargetLocationDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					verifyTargetLocationDataSource(provider, ClusterDataSourceFullName, TargetLocationTMCManagedName),
				),
			},
		},
	},
	)

	t.Log("target location resource acceptance test complete!")
}

func verifyTargetLocationResourceCreation(
	provider *schema.Provider,
	resourceName string,
	targetLocationName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("could not found resource %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource %s", resourceName)
		}

		fn := &targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName{
			Name:         targetLocationName,
			ProviderName: "tmc",
		}

		resp, err := context.TMCConnection.TargetLocationService.TargetLocationResourceServiceGet(fn)

		if err != nil {
			return fmt.Errorf("target location resource not found, resource: %s | err: %s", resourceName, err)
		}

		if resp == nil {
			return fmt.Errorf("target location resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}

func verifyTargetLocationDataSource(
	provider *schema.Provider,
	dataSourceName string,
	targetLocationName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[dataSourceName]

		if !ok {
			return fmt.Errorf("could not found data source %s", dataSourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, data source %s", dataSourceName)
		}

		firstTargetLocation := fmt.Sprintf("%s.0.%s", targetlocationres.TargetLocationsKey, targetlocationres.NameKey)

		if rs.Primary.Attributes[firstTargetLocation] != targetLocationName {
			return fmt.Errorf("target location wasn't found at index 0 (%s)", targetLocationName)
		}

		return nil
	}
}
