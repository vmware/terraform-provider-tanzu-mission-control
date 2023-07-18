/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	aksclients "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/akscluster"
	aksnodepool "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/akscluster/nodepool"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	aksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/akscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

func validateSetup(t *testing.T) {
	// Check if the required environment variables are set
	for _, name := range []string{"AKS_CREDENTIAL_NAME", "AKS_SUBSCRIPTION_ID"} {
		if _, found := os.LookupEnv(name); !found {
			t.Errorf("required environment variable '%s' missing", name)
		}
	}
}

func initTestProvider(t *testing.T, clusterClient *mockClusterClient, nodepoolClient *mockNodepoolClient) *schema.Provider {
	testAksClusterProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			akscluster.ResourceName:   akscluster.ResourceTMCAKSCluster(),
			clustergroup.ResourceName: clustergroup.ResourceClusterGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			akscluster.ResourceName:   akscluster.DataSourceTMCAKSCluster(),
			clustergroup.ResourceName: clustergroup.DataSourceClusterGroup(),
		},
		ConfigureContextFunc: getConfigContext(clusterClient, nodepoolClient),
	}
	if err := testAksClusterProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAksClusterProvider
}

func getConfigContext(clusterClient aksclients.ClientService, nodepoolClient aksnodepool.ClientService) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		if _, found := os.LookupEnv("ENABLE_EKS_ENV_TEST"); found {
			return &authctx.TanzuContext{
				TMCConnection: &client.TanzuMissionControl{
					AKSClusterResourceService:  clusterClient,
					AKSNodePoolResourceService: nodepoolClient,
				},
			}, nil
		}

		return authctx.ProviderConfigureContext(ctx, d)
	}
}

func TestAccAksCluster_basics(t *testing.T) {
	clusterClient := &mockClusterClient{}
	nodepoolClient := &mockNodepoolClient{}
	fn := getFullName()
	rname := fmt.Sprintf("tanzu-mission-control_akscluster.%v", fn.Name)

	var aksCluster aksmodel.VmwareTanzuManageV1alpha1AksclusterAksCluster

	var aksnodepool aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			validateSetup(t)
			testhelper.TestPreCheck(t)
		},
		ProviderFactories: testhelper.GetTestProviderFactories(initTestProvider(t, clusterClient, nodepoolClient)),
		CheckDestroy: func(state *terraform.State) error {
			_, err := clusterClient.AksClusterResourceServiceGet(fn)
			if !clienterrors.IsNotFoundError(err) {
				return err
			}

			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAKSCluster(fn),
				Check: resource.ComposeTestCheckFunc(
					testAksClusterExists(rname, fn, clusterClient, &aksCluster),
				),
			},
			{
				Config: testAKSClusterDisableCSI(fn),
				Check: resource.ComposeTestCheckFunc(
					testAksClusterExists(rname, fn, clusterClient, &aksCluster),
					resource.TestCheckResourceAttr(rname, "spec.0.config.0.storage_config.0.enable_disk_csi_driver", "true"),
					resource.TestCheckResourceAttr(rname, "spec.0.config.0.storage_config.0.enable_file_csi_driver", "true"),
				),
			},
			{
				Config: testAKSClusterAddUserNodepool(fn),
				Check: resource.ComposeTestCheckFunc(
					testAksNodepoolExists(getNodepoolFullName(fn, "userpool"), nodepoolClient, &aksnodepool),
					resource.TestCheckResourceAttr(rname, "spec.0.nodepool.#", "2"),
				),
			},
			{
				Config: testAKSClusterRemoveUserNodepool(fn),
				Check: resource.ComposeTestCheckFunc(
					testAksNodepoolDoesNotExists(getNodepoolFullName(fn, "userpool"), nodepoolClient),
					resource.TestCheckResourceAttr(rname, "spec.0.nodepool.#", "1"),
				),
			},
		},
	})
}

func getFullName() *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName {
	return &aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName{
		CredentialName:    os.Getenv("AKS_CREDENTIAL_NAME"),
		Name:              fmt.Sprintf("tf-acc-test-%s", acctest.RandString(5)),
		ResourceGroupName: os.Getenv("test-group"),
		SubscriptionID:    os.Getenv("AKS_SUBSCRIPTION_ID"),
	}
}

func getNodepoolFullName(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName, name string) *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName {
	return &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName{
		AksClusterName:    fn.Name,
		CredentialName:    fn.CredentialName,
		Name:              name,
		ResourceGroupName: fn.ResourceGroupName,
		SubscriptionID:    fn.SubscriptionID,
	}
}

func testAksNodepoolDoesNotExists(npfn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName, client aksnodepool.ClientService) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		np, err := client.AksNodePoolResourceServiceGet(npfn)
		if clienterrors.IsNotFoundError(err) {
			return nil
		}

		if err != nil {
			return fmt.Errorf("expected nodepool to not exist: %v", np.Nodepool.Status.Conditions)
		}

		return err
	}
}

func testAksNodepoolExists(npfn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName, client aksnodepool.ClientService, a *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		resp, err := client.AksNodePoolResourceServiceGet(npfn)
		if clienterrors.IsNotFoundError(err) {
			return err
		}

		a = resp.Nodepool

		return err
	}
}

func testAksClusterExists(name string, fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName, clusterClient aksclients.ClientService, cluster *aksmodel.VmwareTanzuManageV1alpha1AksclusterAksCluster) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", fn.Name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Cluster ID not set: %s", fn.Name)
		}

		resp, err := clusterClient.AksClusterResourceServiceGet(fn)
		if err != nil {
			return err
		}

		cluster = resp.AksCluster

		return nil
	}
}

func testAKSCluster(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName) string {
	return fmt.Sprintf(`resource "tanzu-mission-control_akscluster" "%s" {
  credential_name = "%s"
  subscription_id = "%s"
  resource_group  = "test-group"
  name            = "%s"
  spec {
    config {
      location = "eastus"
      kubernetes_version = "1.24.10"
      network_config {
        dns_prefix = "dns-tf-test"
      }
      storage_config {
        enable_disk_csi_driver = false
        enable_file_csi_driver = false
      }
    }
    nodepool {
      name = "systemnp"
      spec {
        count = 1
        mode = "SYSTEM"
        vm_size = "Standard_DS2_v2"
      }
    }
  }
}`, fn.Name, fn.CredentialName, fn.SubscriptionID, fn.Name)
}

func testAKSClusterDisableCSI(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName) string {
	return fmt.Sprintf(`resource "tanzu-mission-control_akscluster" "%s" {
  credential_name = "%s"
  subscription_id = "%s"
  resource_group  = "test-group"
  name            = "%s"
  spec {
    config {
      location = "eastus"
      kubernetes_version = "1.24.10"
      network_config {
        dns_prefix = "dns-tf-test"
      }
      storage_config {
        enable_disk_csi_driver = true
        enable_file_csi_driver = true
      } 
    }
    nodepool {
      name = "systemnp"
      spec {
        count = 1
        mode = "SYSTEM"
        vm_size = "Standard_DS2_v2"
      }
    }
  }
}`, fn.Name, fn.CredentialName, fn.SubscriptionID, fn.Name)
}

func testAKSClusterAddUserNodepool(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName) string {
	return fmt.Sprintf(`resource "tanzu-mission-control_akscluster" "%s" {
  credential_name = "%s"
  subscription_id = "%s"
  resource_group  = "test-group"
  name            = "%s"
  spec {
    config {
      location = "eastus"
      kubernetes_version = "1.24.10"
      network_config {
        dns_prefix = "dns-tf-test"
      }
      storage_config {
        enable_disk_csi_driver = true
        enable_file_csi_driver = true
      } 
    }
    nodepool {
      name = "systemnp"
      spec {
        count = 1
        mode = "SYSTEM"
        vm_size = "Standard_DS2_v2"
      }
    }
    nodepool {
      name = "userpool"
      spec {
        count = 1
        mode = "USER"
        vm_size = "Standard_DS2_v2"
      }
    }
  }
}`, fn.Name, fn.CredentialName, fn.SubscriptionID, fn.Name)
}

func testAKSClusterRemoveUserNodepool(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName) string {
	return testAKSClusterDisableCSI(fn)
}
