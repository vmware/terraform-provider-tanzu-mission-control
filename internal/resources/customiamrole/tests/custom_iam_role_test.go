//go:build customiamrole
// +build customiamrole

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamroletests

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
	customiamrolemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/customiamrole"
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

func TestAcceptanceCustomIAMRoleResource(t *testing.T) {
	err := context.Setup()

	if err != nil {
		t.Error(errors.Wrap(err, "unable to set the context"))
		t.FailNow()
	}

	var (
		provider                = initTestProvider(t)
		tfResourceConfigBuilder = InitResourceTFConfigBuilder()
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: tfResourceConfigBuilder.GetCustomSlimIAMRoleConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(CustomIAMRoleResourceFullName, "name", CustomIAMRoleName),
					verifyTanzuKubernetesClusterResource(provider, CustomIAMRoleResourceFullName, CustomIAMRoleName),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetCustomFullIAMRoleConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(CustomIAMRoleResourceFullName, "name", CustomIAMRoleName),
					verifyTanzuKubernetesClusterResource(provider, CustomIAMRoleResourceFullName, CustomIAMRoleName),
				),
			},
		},
	},
	)

	t.Log("Custom IAM role resource acceptance test complete!")
}

func verifyTanzuKubernetesClusterResource(
	provider *schema.Provider,
	resourceName string,
	customRoleName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("could not find resource %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource %s", resourceName)
		}

		fn := &customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleFullName{
			Name: customRoleName,
		}

		resp, err := context.TMCConnection.CustomIAMRoleResourceService.CustomIARoleResourceServiceGet(fn)

		if err != nil {
			return errors.Errorf("Custom IAM Role resource not found, resource: %s | err: %s", resourceName, err)
		}

		if resp == nil {
			return errors.Errorf("Custom IAM Role resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
