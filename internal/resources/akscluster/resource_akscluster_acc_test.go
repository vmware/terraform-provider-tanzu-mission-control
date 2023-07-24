/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster_test

import (
	"context"
	"fmt"
	"net/http"
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
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	aksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
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

func initTestProvider(t *testing.T, clusterClient *mockClusterService, nodepoolClient *mockNodepoolService) *schema.Provider {
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

func getConfigContext(clusterClient *mockClusterService, nodepoolClient *mockNodepoolService) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		if _, found := os.LookupEnv("ENABLE_AKS_ENV_TEST"); !found {
			return authctx.TanzuContext{
				TMCConnection: &client.TanzuMissionControl{
					AKSClusterResourceService:  clusterClient,
					AKSNodePoolResourceService: nodepoolClient,
				},
			}, nil
		}

		return authctx.ProviderConfigureContext(ctx, d)
	}
}

func getTanzuConfig(clusterClient aksclients.ClientService, nodepoolClient aksnodepool.ClientService) (authctx.TanzuContext, error) {
	if _, found := os.LookupEnv("ENABLE_AKS_ENV_TEST"); !found {
		return authctx.TanzuContext{
			TMCConnection: &client.TanzuMissionControl{
				AKSClusterResourceService:  clusterClient,
				AKSNodePoolResourceService: nodepoolClient,
			},
		}, nil
	}

	config := authctx.TanzuContext{
		ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
		Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
		VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
		TLSConfig:        &proxy.TLSConfig{},
	}

	err := config.Setup()

	return config, err
}

func initMocks(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName, clusterClient *mockClusterService, nodepoolClient *mockNodepoolService) {
	clusterClient.createResponses = []*aksmodel.VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse{}
	clusterClient.updateResponse = []*aksmodel.VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse{}

	nodepoolClient.createResponses = []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse{}
	nodepoolClient.listResponses = []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{}
	nodepoolClient.getResponses = []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolGetNodepoolResponse{}

	// ===== Mocks for create Cluster ======
	// Get initial state.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster()})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp"))},
	})

	// Create cluster and system nodepool.
	clusterClient.createResponses = append(clusterClient.createResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse{AksCluster: mockCluster()})
	nodepoolClient.createResponses = append(nodepoolClient.createResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse{
		Nodepool: mockNodepool(forCluster(fn), withNodepoolName("systemnp")),
	})

	// Wait for cluster to be Ready.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster()})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp"))},
	})

	// Read new cluster state.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster()})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp"))},
	})

	// Verify Cluster Exists.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster()})

	// Terraform test state check.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster()})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp"))},
	})

	// ===== Mocks for update Cluster ======
	// Get initial state.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster()})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp"))},
	})

	// Update Cluster.
	clusterClient.updateResponse = append(clusterClient.updateResponse, &aksmodel.VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse{AksCluster: mockCluster(enableCSI)})

	// Wait for cluster to be Ready.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster(enableCSI)})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp"))},
	})

	// Read new state.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster(enableCSI)})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp"))},
	})

	// Check cluster exists
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster(enableCSI)})

	// Terraform test state check.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster(enableCSI)})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp"))},
	})

	// ===== Mocks for create Nodepool ======
	// Get initial state.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster(enableCSI)})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp"))},
	})

	// Create nodepool.
	nodepoolClient.createResponses = append(nodepoolClient.createResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse{Nodepool: mockNodepool(withNodepoolName("userpool"), withUserMode)})
	nodepoolClient.getResponses = append(nodepoolClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolGetNodepoolResponse{Nodepool: mockNodepool(forCluster(fn), withNodepoolName("userpool"), withUserMode)})

	// Save new state.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster(enableCSI)})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp")),
		mockNodepool(forCluster(fn), withNodepoolName("userpool"), withUserMode)},
	})

	// Check nodepool exists.
	nodepoolClient.getResponses = append(nodepoolClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolGetNodepoolResponse{Nodepool: mockNodepool(forCluster(fn), withNodepoolName("userpool"), withUserMode)})

	// Terraform test state validation.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster(enableCSI)})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp")),
		mockNodepool(forCluster(fn), withNodepoolName("userpool"), withUserMode)},
	})

	// ===== Mocks for delete Nodepool ======
	// Get initial state.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster(enableCSI)})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp")),
		mockNodepool(forCluster(fn), withNodepoolName("userpool"), withUserMode)},
	})

	// Delete nodepool.
	nodepoolClient.deleteResponse = append(nodepoolClient.deleteResponse, nil)
	nodepoolClient.getResponses = append(nodepoolClient.getResponses, nil)

	// Save new state.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster(enableCSI)})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp"))},
	})

	// Check nodepool does not exist.
	nodepoolClient.getResponses = append(nodepoolClient.getResponses, nil)

	// Terraform test state validation.
	clusterClient.getResponses = append(clusterClient.getResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{AksCluster: mockCluster(enableCSI)})
	nodepoolClient.listResponses = append(nodepoolClient.listResponses, &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse{Nodepools: []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		mockNodepool(forCluster(fn), withNodepoolName("systemnp"))},
	})

	// ===== Mocks for cleanup ======
	// Cleanup.
	clusterClient.getResponses = append(clusterClient.getResponses, nil) // poll until deleted
	clusterClient.getResponses = append(clusterClient.getResponses, nil) // verify deleted
}

func TestAccAksCluster_basics(t *testing.T) {
	clusterClient := &mockClusterService{}
	nodepoolClient := &mockNodepoolService{}

	fn := getFullName()
	rname := fmt.Sprintf("tanzu-mission-control_akscluster.%v", fn.Name)

	if _, found := os.LookupEnv("ENABLE_AKS_ENV_TEST"); !found {
		initMocks(fn, clusterClient, nodepoolClient)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			validateSetup(t)
			testhelper.TestPreCheck(t)
		},
		ProviderFactories: testhelper.GetTestProviderFactories(initTestProvider(t, clusterClient, nodepoolClient)),
		CheckDestroy: func(state *terraform.State) error {
			config, err := getTanzuConfig(clusterClient, nodepoolClient)
			if err != nil {
				return err
			}
			_, err = config.TMCConnection.AKSClusterResourceService.AksClusterResourceServiceGet(fn)
			if !clienterrors.IsNotFoundError(err) {
				return err
			}

			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAKSCluster(fn),
				Check: resource.ComposeTestCheckFunc(
					testAksClusterExists(rname, fn, clusterClient),
				),
			},
			{
				Config: testAKSClusterEnableCSI(fn),
				Check: resource.ComposeTestCheckFunc(
					testAksClusterExists(rname, fn, clusterClient),
					resource.TestCheckResourceAttr(rname, "spec.0.config.0.storage_config.0.enable_disk_csi_driver", "true"),
					resource.TestCheckResourceAttr(rname, "spec.0.config.0.storage_config.0.enable_file_csi_driver", "true"),
				),
			},
			{
				Config: testAKSClusterAddUserNodepool(fn),
				Check: resource.ComposeTestCheckFunc(
					testAksNodepoolExists(getNodepoolFullName(fn, "userpool"), nodepoolClient),
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

func testAksNodepoolDoesNotExists(npfn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName, nodepoolClient aksnodepool.ClientService) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		config, err := getTanzuConfig(nil, nodepoolClient)
		if err != nil {
			return err
		}

		np, err := config.TMCConnection.AKSNodePoolResourceService.AksNodePoolResourceServiceGet(npfn)
		if clienterrors.IsNotFoundError(err) {
			return nil
		}

		if err != nil {
			return fmt.Errorf("expected nodepool to not exist: %v", np.Nodepool.Status.Conditions)
		}

		return err
	}
}

func testAksNodepoolExists(npfn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName, nodepoolClient aksnodepool.ClientService) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		config, err := getTanzuConfig(nil, nodepoolClient)
		if err != nil {
			return err
		}

		_, err = config.TMCConnection.AKSNodePoolResourceService.AksNodePoolResourceServiceGet(npfn)
		if clienterrors.IsNotFoundError(err) {
			return err
		}

		return err
	}
}

func testAksClusterExists(name string, fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName, clusterClient aksclients.ClientService) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", fn.Name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Cluster ID not set: %s", fn.Name)
		}

		config, err := getTanzuConfig(clusterClient, nil)
		if err != nil {
			return err
		}

		_, err = config.TMCConnection.AKSClusterResourceService.AksClusterResourceServiceGet(fn)
		if err != nil {
			return err
		}

		return nil
	}
}

func getFullName() *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName {
	return &aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName{
		CredentialName:    os.Getenv("AKS_CREDENTIAL_NAME"),
		Name:              fmt.Sprintf("tf-acc-test-%s", acctest.RandString(5)),
		ResourceGroupName: "test-group",
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

func testAKSClusterEnableCSI(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName) string {
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
	return testAKSClusterEnableCSI(fn)
}

func mockCluster(w ...clusterWither) *aksmodel.VmwareTanzuManageV1alpha1AksCluster {
	c := &aksmodel.VmwareTanzuManageV1alpha1AksCluster{
		FullName: getFullName(),
		Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
			UID: "test-uid",
		},
		Spec: &aksmodel.VmwareTanzuManageV1alpha1AksclusterSpec{
			AgentName:        "",
			ClusterGroupName: "default",
			Config: &aksmodel.VmwareTanzuManageV1alpha1AksclusterClusterConfig{
				Location: "eastus",
				NetworkConfig: &aksmodel.VmwareTanzuManageV1alpha1AksclusterNetworkConfig{
					DNSPrefix: "dns-tf-test",
				},
				StorageConfig: &aksmodel.VmwareTanzuManageV1alpha1AksclusterStorageConfig{
					EnableDiskCsiDriver: false,
					EnableFileCsiDriver: false,
				},
				Version: "1.24.10",
			},
		},
		Status: &aksmodel.VmwareTanzuManageV1alpha1AksclusterStatus{
			Phase: aksmodel.VmwareTanzuManageV1alpha1AksclusterPhaseREADY.Pointer(),
		},
	}

	for _, f := range w {
		f(c)
	}

	return c
}

func mockNodepool(w ...nodepoolWither) *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool {
	np := &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		FullName: &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName{},
		Spec: &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolSpec{
			Count:  1,
			Mode:   aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolModeSYSTEM.Pointer(),
			OsType: aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolOsTypeLINUX.Pointer(),
			VMSize: "Standard_DS2_v2",
			Type:   aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolTypeVIRTUALMACHINESCALESETS.Pointer(),
		},
		Status: &aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolStatus{
			Phase: aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseREADY.Pointer(),
		},
	}

	for _, f := range w {
		f(np)
	}

	return np
}

type mockClusterService struct {
	createResponses []*aksmodel.VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse
	createCall      int
	updateResponse  []*aksmodel.VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse
	updateCall      int
	getResponses    []*aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse
	getCall         int
}

func (m *mockClusterService) AksClusterResourceServiceCreate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse, error) {
	resp := m.createResponses[m.createCall]
	m.createCall += 1

	return resp, nil
}

func (m *mockClusterService) AksClusterResourceServiceGet(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse, error) {
	resp := m.getResponses[m.getCall]
	m.getCall += 1

	if resp == nil {
		return nil, clienterrors.ErrorWithHTTPCode(http.StatusNotFound, nil)
	}

	return resp, nil
}

func (m *mockClusterService) AksClusterResourceServiceGetByID(id string) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse, error) {
	panic("not implemented")
}

func (m *mockClusterService) AksClusterResourceServiceUpdate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse, error) {
	resp := m.updateResponse[m.updateCall]
	m.updateCall += 1

	return resp, nil
}

func (m *mockClusterService) AksClusterResourceServiceDelete(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName, force string) error {
	return nil
}

type mockNodepoolService struct {
	createResponses []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse
	createCall      int
	listResponses   []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse
	listCall        int
	getResponses    []*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolGetNodepoolResponse
	getCall         int
	deleteResponse  []error
	deleteCall      int
}

func (m *mockNodepoolService) AksNodePoolResourceServiceCreate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse, error) {
	resp := m.createResponses[m.createCall]
	m.createCall += 1

	return resp, nil
}

func (m *mockNodepoolService) AksNodePoolResourceServiceList(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse, error) {
	resp := m.listResponses[m.listCall]
	m.listCall += 1

	return resp, nil
}

func (m *mockNodepoolService) AksNodePoolResourceServiceGet(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolGetNodepoolResponse, error) {
	resp := m.getResponses[m.getCall]
	m.getCall += 1

	if resp == nil {
		return nil, clienterrors.ErrorWithHTTPCode(http.StatusNotFound, nil)
	}

	return resp, nil
}

func (m *mockNodepoolService) AksNodePoolResourceServiceUpdate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolResponse, error) {
	return nil, nil
}

func (m *mockNodepoolService) AksNodePoolResourceServiceDelete(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName) error {
	resp := m.deleteResponse[m.deleteCall]
	m.deleteCall += 1

	return resp
}
