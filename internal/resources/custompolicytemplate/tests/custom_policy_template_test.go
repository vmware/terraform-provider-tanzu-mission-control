//go:build custompolicytemplate
// +build custompolicytemplate

// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package custompolicytemplate

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
	custompolicytemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/custompolicytemplate"
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

func TestAcceptanceCustomPolicyTemplateResource(t *testing.T) {
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
				Config: tfResourceConfigBuilder.GetSlimCustomPolicyTemplateConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(CustomPolicyTemplateResourceFullName, "name", CustomPolicyTemplateName),
					verifyTanzuKubernetesClusterResource(provider, CustomPolicyTemplateResourceFullName, CustomPolicyTemplateName),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetFullCustomPolicyTemplateConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(CustomPolicyTemplateResourceFullName, "name", CustomPolicyTemplateName),
					verifyTanzuKubernetesClusterResource(provider, CustomPolicyTemplateResourceFullName, CustomPolicyTemplateName),
				),
			},
		},
	},
	)

	t.Log("Custom policy template resource acceptance test complete!")
}

func verifyTanzuKubernetesClusterResource(
	provider *schema.Provider,
	resourceName string,
	customPolicyTemplateName string,
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

		fn := &custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName{
			Name: customPolicyTemplateName,
		}

		resp, err := context.TMCConnection.CustomPolicyTemplateResourceService.CustomPolicyTemplateResourceServiceGet(fn)

		if err != nil {
			return errors.Errorf("Custom IAM Role resource not found, resource: %s | err: %s", resourceName, err)
		}

		if resp == nil {
			return errors.Errorf("Custom IAM Role resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
