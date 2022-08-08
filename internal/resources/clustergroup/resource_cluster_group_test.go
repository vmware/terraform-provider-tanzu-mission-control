/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroup

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
	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForClusterGroupResource(t *testing.T) {
	var provider = initTestProvider(t)

	resourceName := fmt.Sprintf("%s.%s", clusterGroupResource, clusterGroupResourceVar)
	clusterGroupName := acctest.RandomWithPrefix("tf-cg-test")

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTestResourceClusterGroupBasicConfigValue(clusterGroupName),
				Check: resource.ComposeTestCheckFunc(
					verifyClusterGroupResourceCreation(provider, resourceName, clusterGroupName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterGroupName),
				),
			},
			{
				Config: getTestResourceClusterGroupConfigValue(clusterGroupName),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, resourceName, clusterGroupName),
				),
			},
		},
	},
	)
	t.Log("cluster group resource acceptance test complete!")
}

func getTestResourceClusterGroupBasicConfigValue(clusterGroupName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"
}
`, clusterGroupResource, clusterGroupResourceVar, clusterGroupName)
}

func getTestResourceClusterGroupConfigValue(clusterGroupName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"
  %s
}
`, clusterGroupResource, clusterGroupResourceVar, clusterGroupName, testhelper.MetaTemplate)
}

func checkResourceAttributes(provider *schema.Provider, resourceName, clusterGroupName string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyClusterGroupResourceCreation(provider, resourceName, clusterGroupName),
		resource.TestCheckResourceAttr(resourceName, "name", clusterGroupName),
	}

	check = append(check, testhelper.MetaResourceAttributeCheck(resourceName)...)

	return resource.ComposeTestCheckFunc(check...)
}

func verifyClusterGroupResourceCreation(
	provider *schema.Provider,
	resourceName string,
	clusterGroupName string,
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
		}

		err := config.Setup()
		if err != nil {
			return errors.Wrap(err, "unable to set the context")
		}

		fn := &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
			Name: clusterGroupName,
		}

		resp, err := config.TMCConnection.ClusterGroupResourceService.ManageV1alpha1ClusterGroupResourceServiceGet(fn)
		if err != nil {
			return fmt.Errorf("cluster group resource not found: %s", err)
		}

		if resp == nil {
			return fmt.Errorf("cluster group resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
