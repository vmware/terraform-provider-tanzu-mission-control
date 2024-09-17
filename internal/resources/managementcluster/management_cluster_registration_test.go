//go:build managementcluster
// +build managementcluster

/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package managementcluster

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	registrationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/managementcluster"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func TestAcceptanceForManagementClusterRegistrationResource(t *testing.T) {
	var provider = initTestProvider(t)

	tKGsResourceName := fmt.Sprintf("%s.%s", "tanzu-mission-control_management_cluster", "test_tkgs")
	tKGmResourceName := fmt.Sprintf("%s.%s", "tanzu-mission-control_management_cluster", "test_tkgm")

	tkgsSimpleName := acctest.RandomWithPrefix("a-tf-tkgs-test")
	tkgmSimpleName := acctest.RandomWithPrefix("a-tf-tkgm-test")

	tkgmKubeconfigFilePathName := acctest.RandomWithPrefix("a-tf-tkgm-kubeconfig-filepath-test")

	kubeconfigPath := os.Getenv("KUBECONFIG")
	_, enablePolicyEnvTest := os.LookupEnv("ENABLE_MANAGEMENT_CLUSTER_ENV_TEST")

	if !enablePolicyEnvTest {
		os.Setenv("TF_ACC", "true")

		endpoint := "play.abc.def.ghi.com"

		os.Setenv(authctx.ServerEndpointEnvVar, endpoint)
		os.Setenv("VMW_CLOUD_API_TOKEN", "dummy")
		os.Setenv("VMW_CLOUD_ENDPOINT", "console.tanzu.broadcom.com")

		setupHTTPMocks(t)
		setUpOrgPolicyEndPointMocks(t, endpoint, tkgsSimpleName, clustermodel.NewVmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType("VMWARE_TANZU_KUBERNETES_GRID_SERVICE"))
		setUpOrgPolicyEndPointMocks(t, endpoint, tkgmSimpleName, clustermodel.NewVmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType("VMWARE_TANZU_KUBERNETES_GRID"))
	} else {
		requiredVars := []string{
			"VMW_CLOUD_ENDPOINT",
			"TMC_ENDPOINT",
			"VMW_CLOUD_API_TOKEN",
			"ORG_ID",
		}

		for _, name := range requiredVars {
			if _, found := os.LookupEnv(name); !found {
				t.Errorf("required environment variable '%s' missing", name)
			}
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          testhelper.TestPreCheck(t),
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: getTKGsResourceWithoutKubeconfigWithDataSource(tkgsSimpleName),
				Check:  resource.ComposeTestCheckFunc(checkResourceAttributes(tKGsResourceName, tkgsSimpleName, false)),
			},
			{
				Config: getTKGmResourceWithoutKubeconfigWithDataSource(tkgmSimpleName),
				Check:  resource.ComposeTestCheckFunc(checkResourceAttributes(tKGmResourceName, tkgmSimpleName, false)),
			},
			{
				PreConfig: func() {
					if kubeconfigPath == "" {
						t.Skip("KUBECONFIG env var is not set for management cluster registration acceptance test")
					}
					if !enablePolicyEnvTest {
						t.Skip("Acceptance tests against outside systems are not enabled")
					}
				},
				Config: getTKGmResourceWithDataSourceWithKubeConfigFilePath(tkgmKubeconfigFilePathName, kubeconfigPath),
				Check:  resource.ComposeTestCheckFunc(checkResourceAttributes(tKGmResourceName, tkgmKubeconfigFilePathName, true)),
			},
		},
	},
	)
	t.Log("management cluster registration resource acceptance test complete!")
}

func getTKGsResourceWithoutKubeconfigWithDataSource(name string) string {
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

func getTKGmResourceWithoutKubeconfigWithDataSource(name string) string {
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

func getTKGmResourceWithDataSourceWithKubeConfigFilePath(name string, kubeconfigPath string) string {
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

func checkResourceAttributes(resourceName, name string, checkReadyState bool) resource.TestCheckFunc {
	var check = []resource.TestCheckFunc{
		verifyManagementClusterRegistrationResourceCreation(resourceName, name, checkReadyState),
		resource.TestCheckResourceAttr(resourceName, "name", name),
	}

	return resource.ComposeTestCheckFunc(check...)
}

func verifyManagementClusterRegistrationResourceCreation(
	resourceName string,
	name string,
	checkReadyState bool,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config, err := getContext(s, resourceName)
		if err != nil {
			return err
		}

		request := &registrationmodel.VmwareTanzuManageV1alpha1ManagementclusterFullName{
			Name: name,
		}

		resp, err := config.TMCConnection.ManagementClusterRegistrationResourceService.ManagementClusterResourceServiceGet(request)
		if err != nil || resp == nil {
			return fmt.Errorf("management cluster registration resource not found: %s", err)
		}

		if checkReadyState && !strings.EqualFold(string(registrationmodel.VmwareTanzuManageV1alpha1ManagementclusterPhaseREADY), string(*resp.ManagementCluster.Status.Phase)) {
			return fmt.Errorf("registration has not finilalised, received non READY phase: %s", *resp.ManagementCluster.Status.Phase)
		}

		if resp == nil {
			return fmt.Errorf("management cluster registration resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}

func getContext(s *terraform.State, resourceName string) (*authctx.TanzuContext, error) {
	rs, ok := s.RootModule().Resources[resourceName]
	if !ok {
		return nil, fmt.Errorf("not found resource: %s", resourceName)
	}

	if rs.Primary.ID == "" {
		return nil, fmt.Errorf("ID not set, resource: %s", resourceName)
	}

	config := &authctx.TanzuContext{
		ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
		Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
		VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
		TLSConfig:        &proxy.TLSConfig{},
	}

	err := getSetupConfig(config)

	if err != nil {
		return nil, errors.Wrap(err, "unable to set the context")
	}

	return config, nil
}

func getSetupConfig(config *authctx.TanzuContext) error {
	if _, found := os.LookupEnv("ENABLE_MANAGEMENT_CLUSTER_ENV_TEST"); !found {
		return config.SetupWithDefaultTransportForTesting()
	}

	return config.Setup()
}
