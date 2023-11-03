//go:build managementclusterregistration
// +build managementclusterregistration

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package managementclusterregistration

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
	managementclusterregistrationclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/managementclusterregistration"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForManagementClusterRegistrationResource(t *testing.T) {
	var provider = initTestProvider(t)

	tKGsresourceName := fmt.Sprintf("%s.%s", "tanzu-mission-control_management_cluster", "test_tkgs")
	tkgsName := acctest.RandomWithPrefix("a-tf-tkgs-test")

	tKGmResourceName := fmt.Sprintf("%s.%s", "tanzu-mission-control_management_cluster", "test_tkgm")
	tkgmName := acctest.RandomWithPrefix("a-tf-tkgm-test")

	kubeconfigPath := os.Getenv("KUBECONFIG")

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getSimpleTKGsResourceWithDataSource(tkgsName),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, tKGsresourceName, tkgsName),
				),
			},
			{
				Config: getSimpleTKGmResourceWithDataSource(tkgmName),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, tKGmResourceName, tkgmName),
				),
			},
			{
				PreConfig: func() {
					if kubeconfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for management cluster registration acceptance test")
					}
				},
				Config: getSimpleTKGmResourceWithDataSourceWithKubeConfig(tkgmName, kubeconfigPath),
				Check: resource.ComposeTestCheckFunc(
					checkResourceAttributes(provider, tKGmResourceName, tkgmName),
				),
			},
		},
	},
	)
	t.Log("management cluster registration resource acceptance test complete!")
}

func getSimpleTKGsResourceWithDataSource(name string) string {
	return fmt.Sprintf(`
		resource "tanzu-mission-control_management_cluster" "test_tkgs" {
		  name = "%s"
		  spec {
			cluster_group = "default" 
			kubernetes_provider_type = "VMWARE_TANZU_KUBERNETES_GRID_SERVICE" 
		  }
		}
		
		data "tanzu-mission-control_management_cluster" "read_tkgs_management_cluster_registration" {
			name = tanzu-mission-control_management_cluster.test_tkgs.name
		}
		`, name)
}

func getSimpleTKGmResourceWithDataSource(name string) string {
	return fmt.Sprintf(`
		resource "tanzu-mission-control_management_cluster" "test_tkgm" {
		  name = "%s"
		  spec {
			cluster_group = "default" 
			kubernetes_provider_type = "VMWARE_TANZU_KUBERNETES_GRID" 
		  }
		}
		
		data "tanzu-mission-control_management_cluster" "read_tkgm_management_cluster_registration" {
			name = tanzu-mission-control_management_cluster.test_tkgm.name
		}
		`, name)
}

func getSimpleTKGmResourceWithDataSourceWithKubeConfig(name string, kubeconfigPath string) string {
	return fmt.Sprintf(`
		resource "tanzu-mission-control_management_cluster" "test_tkgm" {
		  name = "%s"
		  spec {
			cluster_group = "default" 
			kubernetes_provider_type = "VMWARE_TANZU_KUBERNETES_GRID"
		  }
          register_management_cluster {
			tkgm_kubeconfig_file = "%s"
		  }
		}
		
		data "tanzu-mission-control_management_cluster" "read_tkgm_management_cluster_registration" {
			name = tanzu-mission-control_management_cluster.test_tkgm.name
		}
		`, name, kubeconfigPath)
}

func checkResourceAttributes(provider *schema.Provider, resourceName, name string) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyManagementClusterRegistrationResourceCreation(provider, resourceName, name),
		resource.TestCheckResourceAttr(resourceName, "name", name),
	}

	return resource.ComposeTestCheckFunc(check...)
}

func verifyManagementClusterRegistrationResourceCreation(
	provider *schema.Provider,
	resourceName string,
	name string,
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

		request := &managementclusterregistrationclient.VmwareTanzuManageV1alpha1ManagementclusterFullName{
			Name: name,
		}

		resp, err := config.TMCConnection.ManagementClusterRegistrationResourceService.ManagementClusterResourceServiceGet(request)
		if err != nil || resp == nil {
			return fmt.Errorf("management cluster registration resource not found: %s", err)
		}

		if resp == nil {
			return fmt.Errorf("management cluster registration resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
