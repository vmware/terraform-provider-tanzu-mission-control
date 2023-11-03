//go:build inspections
// +build inspections

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionstests

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	inspectionsres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/inspections"
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

func TestAcceptanceInspectionsDataSources(t *testing.T) {
	err := context.Setup()

	if err != nil {
		t.Error(errors.Wrap(err, "unable to set the context"))
		t.FailNow()
	}

	// See cluster_env_vars.go for required environment variables.
	environmentVars, errs := ReadClusterEnvironmentVariables()

	if len(errs) > 0 {
		errMsg := ""

		for _, e := range errs {
			errMsg = fmt.Sprintf("%s\n%s", errMsg, e.Error())
		}

		t.Error(errors.Errorf("Required environment variables are missing: %s", errMsg))
		t.FailNow()
	}

	provider := initTestProvider(t)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: GetInspectionListConfig(environmentVars),
				Check: resource.ComposeTestCheckFunc(
					verifyInspectionListDataSource(provider, InspectionListDataSourceFullName),
				),
			},
			{
				Config: GetInspectionResultsConfig(environmentVars),
				Check: resource.ComposeTestCheckFunc(
					verifyInspectionResultsDataSource(provider, InspectionResultsDataSourceFullName),
				),
			},
		},
	},
	)

	t.Log("Inspections data sources acceptance test complete!")
}

func verifyInspectionListDataSource(provider *schema.Provider, dataSourceName string) resource.TestCheckFunc {
	verifyFunc := func(rs *terraform.ResourceState) (err error) {
		inspectionsCount, countExist := rs.Primary.Attributes[fmt.Sprintf("%s.#", inspectionsres.InspectionListKey)]

		if !countExist || (countExist && (inspectionsCount == "" || inspectionsCount == "0")) {
			err = errors.New("Inspection list is empty")
		}

		return err
	}

	return verifyBackupScheduleDataSource(provider, dataSourceName, verifyFunc)
}

func verifyInspectionResultsDataSource(provider *schema.Provider, dataSourceName string) resource.TestCheckFunc {
	verifyFunc := func(rs *terraform.ResourceState) (err error) {
		reportData, reportExists := rs.Primary.Attributes[fmt.Sprintf("%s.%s", inspectionsres.StatusKey, inspectionsres.ReportKey)]

		if !reportExists || (reportExists && reportData == "") {
			err = errors.New("Inspection results is empty")
		}

		return err
	}

	return verifyBackupScheduleDataSource(provider, dataSourceName, verifyFunc)
}

func verifyBackupScheduleDataSource(
	provider *schema.Provider,
	dataSourceName string,
	verificationFunc func(*terraform.ResourceState) error,
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

		return verificationFunc(rs)
	}
}
