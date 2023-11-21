//go:build clusterclass
// +build clusterclass

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclasstests

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
	clusterclassres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clusterclass"
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

func TestAcceptanceClusterClassDataSource(t *testing.T) {
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
				Config: GetClusterClassConfig(environmentVars),
				Check: resource.ComposeTestCheckFunc(
					verifyBackupScheduleDataSource(provider, ClusterClassDataSourceFullName),
				),
			},
		},
	},
	)

	t.Log("Cluster class data source acceptance test complete!")
}

func verifyBackupScheduleDataSource(
	provider *schema.Provider,
	dataSourceName string,
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

		if rs.Primary.Attributes[clusterclassres.VariablesSchemaKey] == "" {
			return errors.New("Cluster class variables not found!")
		}

		return nil
	}
}
