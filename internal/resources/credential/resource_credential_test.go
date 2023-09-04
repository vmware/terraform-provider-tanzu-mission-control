//go:build credential
// +build credential

/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package credential

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
	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForCredentialResource(t *testing.T) {
	var provider = initTestProvider(t)

	resourceName := fmt.Sprintf("%s.%s", credentialResource, credentialResourceVar)
	credentialName := acctest.RandomWithPrefix("tf-cred-test")

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTestResourceCredentialImageRegConfigValue(credentialName),
				Check: resource.ComposeTestCheckFunc(
					verifyCredentialResourceCreation(provider, resourceName, credentialName),
					resource.TestCheckResourceAttr(resourceName, "name", credentialName),
				),
			},
			{
				Config: getTestResourceCredentialEKSValue(credentialName),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, resourceName, credentialName),
				),
			},
		},
	},
	)
	t.Log("credential resource acceptance test complete!")
}

func getTestResourceCredentialImageRegConfigValue(credentialName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"
  
meta {
    description = "resource with description"
    labels = {
      "key1" : "value1",
      "key2" : "value2"
    }
    annotations = {
      "repository-path" : "something"
    }
  }

  spec {
    provider = "GENERIC_KEY_VALUE"
    capability = "IMAGE_REGISTRY"
    data {
      key_value{
        data  = {
          "registry-url" = "somethingnew"
        }
      }
    }
  }
}
`, credentialResource, credentialResourceVar, credentialName)
}

func getTestResourceCredentialEKSValue(credentialName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"

 spec {
    capability = "MANAGED_K8S_PROVIDER"
	provider = "AWS_EKS"
    data {
		aws_credential {
			account_id = ""
			generic_credential = ""
			iam_role{
				arn = "arn:aws:iam::4987398738934:role/clusterlifecycle-test.tmc.cloud.vmware.com"
				ext_id =""
			}
		}
	}
 }

}
`, credentialResource, credentialResourceVar, credentialName)
}

func checkResourceAttributes(provider *schema.Provider, resourceName, credentialName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyCredentialResourceCreation(provider, resourceName, credentialName),
		resource.TestCheckResourceAttr(resourceName, "name", credentialName),
	}

	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
		resource.TestCheckResourceAttrSet(resourceName, "meta.0.uid"),
	}

	check = append(check, checks...)

	return resource.ComposeTestCheckFunc(check...)
}

func verifyCredentialResourceCreation(
	provider *schema.Provider,
	resourceName string,
	credName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("not found resource %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource %s", resourceName)
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

		fn := &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialFullName{
			Name: credName,
		}

		resp, err := config.TMCConnection.CredentialResourceService.CredentialResourceServiceGet(fn)
		if err != nil {
			return fmt.Errorf("credential resource not found: %s", err)
		}

		if resp == nil {
			return fmt.Errorf("credential resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
