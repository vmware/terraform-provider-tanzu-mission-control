//go:build permissiontemplate
// +build permissiontemplate

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package permissiontemplatetests

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
	permissiontemplateres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/permissiontemplate"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

var (
	context = authctx.TanzuContext{
		ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
		Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
		VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
		TLSConfig:        &proxy.TLSConfig{},
	}

	dataProtectionPermissionTemplateDataSourceFullName = fmt.Sprintf("data.%s.%s", permissiontemplateres.ResourceName, dataProtectionPermissionTemplateDataSourceName)
	eksPermissionTemplateDataSourceFullName            = fmt.Sprintf("data.%s.%s", permissiontemplateres.ResourceName, eksPermissionTemplateDataSourceName)
)

func TestAcceptancePermissionTemplateDataSource(t *testing.T) {
	err := context.Setup()

	if err != nil {
		t.Error(errors.Wrap(err, "unable to set the context"))
		t.FailNow()
	}

	provider := initTestProvider(t)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: GetDataProtectionPermissionTemplateConfig(),
				Check: resource.ComposeTestCheckFunc(
					verifyPermissionTemplateDataSource(provider, dataProtectionPermissionTemplateDataSourceFullName),
				),
			},
			{
				Config: GetEKSPermissionTemplateConfig(),
				Check: resource.ComposeTestCheckFunc(
					verifyPermissionTemplateDataSource(provider, eksPermissionTemplateDataSourceFullName),
				),
			},
		},
	},
	)

	t.Log("Permission template data source acceptance test complete!")
}

func verifyPermissionTemplateDataSource(
	provider *schema.Provider,
	dataSourceName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[dataSourceName]

		if !ok {
			return fmt.Errorf("could not find data source %s", dataSourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, data source %s", dataSourceName)
		}

		if rs.Primary.Attributes[permissiontemplateres.TemplateKey] == "" {
			return errors.New("Permission template data source is empty!")
		}

		return nil
	}
}
