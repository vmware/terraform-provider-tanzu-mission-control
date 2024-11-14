//go:build workspace
// +build workspace

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package workspace

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
	workspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/workspace"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForWorkspaceResource(t *testing.T) {
	var provider = initTestProvider(t)

	resourceName := fmt.Sprintf("%s.%s", workspaceResource, workspaceResourceVar)
	workspaceName := acctest.RandomWithPrefix("tf-ws-test")

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTestWorkspaceResourceBasicConfigValue(workspaceName),
				Check: resource.ComposeTestCheckFunc(
					verifyWorkspaceResourceCreation(provider, resourceName, workspaceName),
					resource.TestCheckResourceAttr(resourceName, "name", workspaceName),
				),
			},
			{
				Config: getTestWorkspaceResourceConfigValue(workspaceName),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, resourceName, workspaceName),
				),
			},
		},
	},
	)
	t.Log("workspace resource acceptance test complete!")
}

func checkResourceAttributes(provider *schema.Provider, resourceName, workspaceName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyWorkspaceResourceCreation(provider, resourceName, workspaceName),
		resource.TestCheckResourceAttr(resourceName, "name", workspaceName),
	}

	check = append(check, testhelper.MetaResourceAttributeCheck(resourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func getTestWorkspaceResourceBasicConfigValue(workspaceName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"
}
`, workspaceResource, workspaceResourceVar, workspaceName)
}

func getTestWorkspaceResourceConfigValue(workspaceName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"
  %s
}
`, workspaceResource, workspaceResourceVar, workspaceName, testhelper.MetaTemplate)
}

func verifyWorkspaceResourceCreation(
	provider *schema.Provider,
	resourceName string,
	workspaceName string,
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

		fn := &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
			Name: workspaceName,
		}

		resp, err := config.TMCConnection.WorkspaceResourceService.ManageV1alpha1WorkspaceResourceServiceGet(fn)
		if err != nil {
			return errors.Wrap(err, "workspace resource not found")
		}

		if resp == nil {
			return errors.Wrapf(err, "workspace resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
