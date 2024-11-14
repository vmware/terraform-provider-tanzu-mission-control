// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package akscluster_test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
	configModels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubeconfig"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/akscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type mocks struct {
	clusterClient    *mockClusterClient
	nodepoolClient   *mockNodepoolClient
	kubeConfigClient *mockKubeConfigClient
}

func TestAKSClusterResource(t *testing.T) {
	suite.Run(t, &CreatClusterTestSuite{})
	suite.Run(t, &ReadClusterTestSuite{})
	suite.Run(t, &UpdateClusterTestSuite{})
	suite.Run(t, &DeleteClusterTestSuite{})
	suite.Run(t, &ImportClusterTestSuite{})
}

type CreatClusterTestSuite struct {
	suite.Suite
	ctx                context.Context
	mocks              mocks
	aksClusterResource *schema.Resource
	config             authctx.TanzuContext
}

func (s *CreatClusterTestSuite) SetupTest() {
	s.mocks.clusterClient = &mockClusterClient{
		createClusterResp: aTestCluster(),
		getClusterResp:    aTestCluster(withStatusSuccess),
	}
	s.mocks.nodepoolClient = &mockNodepoolClient{
		nodepoolListResp: []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{aTestNodePool()},
	}
	s.mocks.kubeConfigClient = &mockKubeConfigClient{
		kubeConfigResponse: &configModels.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponse{
			Status:     configModels.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusREADY.Pointer(),
			Kubeconfig: "base64_kubeconfig",
		},
	}
	s.config = authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			AKSClusterResourceService:  s.mocks.clusterClient,
			AKSNodePoolResourceService: s.mocks.nodepoolClient,
			KubeConfigResourceService:  s.mocks.kubeConfigClient,
		},
	}
	s.aksClusterResource = akscluster.ResourceTMCAKSCluster()
	s.ctx = context.WithValue(context.Background(), akscluster.RetryInterval, 10*time.Millisecond)
}

func (s *CreatClusterTestSuite) Test_datasource_spec_should_be_required() {
	resourceSchema := s.aksClusterResource

	s.Assert().True(resourceSchema.Schema["spec"].Required)
	s.Assert().False(resourceSchema.Schema["spec"].Optional)
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())
	expectedNP := aTestNodePool(forCluster(aTestCluster().FullName))

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().True(s.mocks.clusterClient.AksCreateClusterWasCalled, "cluster create was not called")
	s.Assert().Equal(s.mocks.nodepoolClient.CreateNodepoolWasCalledWith, expectedNP, "nodepool create was not called")
	s.Assert().Equal(expectedFullName(), s.mocks.clusterClient.AksClusterResourceServiceGetCalledWith)
	s.Assert().Equal(expectedFullName(), s.mocks.nodepoolClient.AksNodePoolResourceServiceListCalledWith)
	s.Assert().Equal("test-uid", d.Id())

	s.Assert().False(s.mocks.kubeConfigClient.KubeConfigServicedWasCalled, "kubeconfig client was called when not expected")
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_withPodCIDR() {
	nodepools := []any{aTestNodepoolDataMap(withPodSubnetID(""))}
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap(withNodepools(nodepools),
		withNetworkPlugin("kubenet"), withPodCIDR([]any{"10.1.0.0/16"})))
	expectedNP := aTestNodePool(forCluster(aTestCluster(withTestPodCIDR).FullName))
	expectedNP.Spec.PodSubnetID = ""

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().True(s.mocks.clusterClient.AksCreateClusterWasCalled, "cluster create was not called")
	s.Assert().Equal(s.mocks.nodepoolClient.CreateNodepoolWasCalledWith, expectedNP, "nodepool create was not called")
	s.Assert().Equal(expectedFullName(), s.mocks.clusterClient.AksClusterResourceServiceGetCalledWith)
	s.Assert().Equal(expectedFullName(), s.mocks.nodepoolClient.AksNodePoolResourceServiceListCalledWith)
	s.Assert().Equal("test-uid", d.Id())

	s.Assert().False(s.mocks.kubeConfigClient.KubeConfigServicedWasCalled, "kubeconfig client was called when not expected")
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_waitFor_KubConfig() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap(withWaitForHealthy))
	expectedNP := aTestNodePool(forCluster(aTestCluster().FullName))

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().True(s.mocks.clusterClient.AksCreateClusterWasCalled, "cluster create was not called")
	s.Assert().Equal(s.mocks.nodepoolClient.CreateNodepoolWasCalledWith, expectedNP, "nodepool create was not called ")
	s.Assert().Equal(expectedFullName(), s.mocks.clusterClient.AksClusterResourceServiceGetCalledWith)
	s.Assert().Equal(expectedFullName(), s.mocks.nodepoolClient.AksNodePoolResourceServiceListCalledWith)
	s.Assert().True(s.mocks.kubeConfigClient.KubeConfigServicedWasCalled, "kubeconfig client was not called")
	s.Assert().Equal("my-agent-name", s.mocks.kubeConfigClient.KubeConfigServiceCalledWith.Name)
	s.Assert().Equal("test-uid", d.Id())
	s.Assert().Equal("base64_kubeconfig", d.Get("kubeconfig"))
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_waitFor_KubConfig_Timeout() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap(withWaitForHealthy, with5msTimeout))
	s.mocks.kubeConfigClient.kubeConfigError = errors.New("timeout")

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_invalidConfig() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.CreateContext(s.ctx, d, nil)

	s.Assert().True(result.HasError())
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_ClusterCreate_fails() {
	s.mocks.clusterClient.createErr = errors.New("create cluster failed")
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_NodepoolCreate_fails() {
	s.mocks.nodepoolClient.createErr = errors.New("create nodepool failed")
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_ClusterCreate_timeout() {
	s.mocks.clusterClient.getClusterResp = aTestCluster()
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap(with5msTimeout))

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_ClusterCreate_has_error_status() {
	s.mocks.clusterClient.getClusterResp = aTestCluster(withStatusError)
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_ClusterCreate_alreadyExists() {
	s.mocks.clusterClient.createErr = clienterrors.ErrorWithHTTPCode(http.StatusConflict, nil)
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Equal("test-uid", d.Id())
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_ClusterCreate_alreadyExists_but_notFound() {
	s.mocks.clusterClient.createErr = clienterrors.ErrorWithHTTPCode(http.StatusConflict, nil)
	s.mocks.clusterClient.getErr = clienterrors.ErrorWithHTTPCode(http.StatusNotFound, nil)
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Empty(d.Id())
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_ClusterCreate_succeeded_but_cluster_notFound() {
	s.mocks.clusterClient.getErr = clienterrors.ErrorWithHTTPCode(http.StatusNotFound, nil)
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Empty(d.Id())
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_ClusterCreate_succeeded_but_nodepools_notFound() {
	s.mocks.nodepoolClient.listErr = clienterrors.ErrorWithHTTPCode(http.StatusNotFound, nil)
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_ClusterCreate_all_system_pools_fail() {
	nodepools := []any{aTestNodepoolDataMap(withNodepoolMode("USER")), aTestNodepoolDataMap(withNodepoolMode("SYSTEM"))}
	cluster := aTestClusterDataMap(withNodepools(nodepools))
	s.mocks.nodepoolClient.failSystemPools = true
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, cluster)

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal("no system nodepools were successfully created. [failed to create system nodepool]", result[0].Summary)
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_ClusterCreate_overlay_with_kubenet_fail() {
	cluster := aTestClusterDataMap(withNetworkPlugin("kubenet"), withNetworkPluginMode("overlay"))
	s.mocks.nodepoolClient.failSystemPools = true
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, cluster)

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal("network_plugin_mode 'overlay' can only be used if network_plugin is set to 'azure'", result[0].Summary)
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_ClusterCreate_podCIDR_azure_fail() {
	cluster := aTestClusterDataMap(withNetworkPlugin("azure"), withPodCIDR([]any{"127.0.0.3"}))
	s.mocks.nodepoolClient.failSystemPools = true
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, cluster)

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal("podCIDR cannot be set if network-plugin is 'azure' without 'overlay'", result[0].Summary)
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_NodePool_overlay_with_pod_subnet_fail() {
	nodepools := []any{aTestNodepoolDataMap(withNodepoolMode("SYSTEM"), withPodSubnetID("vnet-1/subnet-1"))}
	cluster := aTestClusterDataMap(withNetworkPluginMode("overlay"), withNodepools(nodepools))
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, cluster)

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal("can not set pod_subnet_id when network_plugin is set to 'kubenet' or to 'azure' with network_plugin_mode set to 'overlay'", result[0].Summary)
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_NodePool_kubenet_with_pod_subnet_fail() {
	nodepools := []any{aTestNodepoolDataMap(withNodepoolMode("SYSTEM"), withPodSubnetID("vnet-1/subnet-1"))}
	cluster := aTestClusterDataMap(withNetworkPlugin("kubenet"), withNetworkPluginMode(""), withNodepools(nodepools))
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, cluster)

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal("can not set pod_subnet_id when network_plugin is set to 'kubenet' or to 'azure' with network_plugin_mode set to 'overlay'", result[0].Summary)
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_NodePool_wrong_subnet_format_fail() {
	nodepools := []any{aTestNodepoolDataMap(withNodepoolMode("SYSTEM"), withPodSubnetID("vnet-1/subnet-1"))}
	cluster := aTestClusterDataMap(withNodepools(nodepools))
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, cluster)

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal("cannot read vNet Id from subnet with Id 'vnet-1/subnet-1'", result[0].Summary)
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_NodePool_different_vNets_fail() {
	nodepools := []any{aTestNodepoolDataMap(withNodepoolMode("SYSTEM"), withPodSubnetID("vnet-2/subnets/subnet-2"))}
	cluster := aTestClusterDataMap(withNodepools(nodepools))
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, cluster)

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal("node (Vnet-subnet) and pod subNets should belong to the same vNet", result[0].Summary)
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_NodePool_pod_subnet_without_node_subnet_fail() {
	nodepools := []any{aTestNodepoolDataMap(withNodepoolMode("SYSTEM"),
		withNodeSubnetID(""), withPodSubnetID("vnet-1/subnets/subnet-2"))}
	cluster := aTestClusterDataMap(withNodepools(nodepools))
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, cluster)

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal("pod subNet cannot be specified if node subNet is not defined", result[0].Summary)
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_NodePool_pod_and_pod_submets_same_fail() {
	nodepools := []any{aTestNodepoolDataMap(withNodepoolMode("SYSTEM"), withPodSubnetID("vnet-1/subnets/subnet-1"))}
	cluster := aTestClusterDataMap(withNodepools(nodepools))
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, cluster)

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal("node (Vnet-subnet) and pod subNets cannot be the same", result[0].Summary)
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate_ClusterCreate_no_system_nodepool() {
	userpool := []any{aTestNodepoolDataMap(withNodepoolMode("USER"))}
	cluster := aTestClusterDataMap(withNodepools(userpool))
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, cluster)

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

type ReadClusterTestSuite struct {
	suite.Suite
	ctx                context.Context
	mocks              mocks
	aksClusterResource *schema.Resource
	config             authctx.TanzuContext
}

func (s *ReadClusterTestSuite) SetupTest() {
	s.mocks.clusterClient = &mockClusterClient{
		createClusterResp: aTestCluster(),
		getClusterResp:    aTestCluster(withStatusSuccess),
	}
	s.mocks.nodepoolClient = &mockNodepoolClient{
		nodepoolListResp: []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{aTestNodePool()},
	}
	s.config = authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			AKSClusterResourceService:  s.mocks.clusterClient,
			AKSNodePoolResourceService: s.mocks.nodepoolClient,
		},
	}
	s.aksClusterResource = akscluster.ResourceTMCAKSCluster()
	s.ctx = context.WithValue(context.Background(), akscluster.RetryInterval, 10*time.Millisecond)
}

func (s *ReadClusterTestSuite) Test_resourceClusterRead() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.ReadContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Equal(expectedFullName(), s.mocks.clusterClient.AksClusterResourceServiceGetCalledWith)
	s.Assert().Equal(expectedFullName(), s.mocks.nodepoolClient.AksNodePoolResourceServiceListCalledWith)
	s.Assert().Equal("test-uid", d.Id(), "expect id from REST request")
	s.Assert().NotNil(d.Get(common.MetaKey), "expected metadata from REST request")
	s.Assert().NotNil(d.Get("spec"), "expected cluster spec from REST request")
}

type UpdateClusterTestSuite struct {
	suite.Suite
	ctx                context.Context
	mocks              mocks
	aksClusterResource *schema.Resource
	config             authctx.TanzuContext
}

func (s *UpdateClusterTestSuite) SetupTest() {
	s.mocks.clusterClient = &mockClusterClient{
		createClusterResp: aTestCluster(),
		getClusterResp:    aTestCluster(withStatusSuccess),
	}
	s.mocks.nodepoolClient = &mockNodepoolClient{
		nodepoolListResp: []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{aTestNodePool(forCluster(aTestCluster().FullName))},
		nodepoolGetResp:  aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolStatusSuccess),
	}
	s.config = authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			AKSClusterResourceService:  s.mocks.clusterClient,
			AKSNodePoolResourceService: s.mocks.nodepoolClient,
		},
	}
	s.aksClusterResource = akscluster.ResourceTMCAKSCluster()
	s.ctx = context.WithValue(context.Background(), akscluster.RetryInterval, 10*time.Millisecond)
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_updateClusterConfig() {
	originalCluster := aTestClusterDataMap(withDNSPrefix("new-prefix1"))
	updatedCluster := aTestClusterDataMap(withDNSPrefix("new-prefix2"))
	d := dataDiffFrom(s.T(), originalCluster, updatedCluster)
	expected := aTestCluster()
	expected.Spec.Config.NetworkConfig.DNSPrefix = "new-prefix2"

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Equal(expected, s.mocks.clusterClient.AksUpdateClusterWasCalledWith)
	s.Assert().Nil(s.mocks.nodepoolClient.UpdatedNodepoolWasCalledWith)
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_updateNodepool() {
	originalNodepools := []any{aTestNodepoolDataMap(withNodepoolCount(1))}
	updatedNodepools := []any{aTestNodepoolDataMap(withNodepoolCount(5))}
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))
	expected := aTestNodePool(forCluster(aTestCluster().FullName))
	expected.Spec.Count = 5

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Nil(s.mocks.clusterClient.AksUpdateClusterWasCalledWith)
	s.Assert().Equal(expected, s.mocks.nodepoolClient.UpdatedNodepoolWasCalledWith)
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_addNodepool() {
	s.mocks.nodepoolClient.nodepoolListResp = []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np1")),
	}
	originalNodepools := []any{aTestNodepoolDataMap(withName("np1"))}
	updatedNodepools := []any{aTestNodepoolDataMap(withName("np1")), aTestNodepoolDataMap(withName("np2"))}
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))
	expected := aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np2"))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Nil(s.mocks.clusterClient.AksUpdateClusterWasCalledWith)
	s.Assert().Equal(expected, s.mocks.nodepoolClient.CreateNodepoolWasCalledWith)
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_deleteNodepool() {
	s.mocks.nodepoolClient.nodepoolListResp = []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np1")),
		aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np2"))}
	originalNodepools := []any{aTestNodepoolDataMap(withName("np1")), aTestNodepoolDataMap(withName("np2"))}
	updatedNodepools := []any{aTestNodepoolDataMap(withName("np1")), nil}
	s.mocks.nodepoolClient.getErr = clienterrors.ErrorWithHTTPCode(http.StatusNotFound, nil)
	expected := aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np2")).FullName

	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Nil(s.mocks.clusterClient.AksUpdateClusterWasCalledWith)
	s.Assert().Equal(expected, s.mocks.nodepoolClient.DeleteNodepoolWasCalledWith)
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_invalidConfig() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.UpdateContext(s.ctx, d, "config")

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_updateClusterFails() {
	originalCluster := aTestClusterDataMap(withDNSPrefix("new-prefix1"))
	updatedCluster := aTestClusterDataMap(withDNSPrefix("new-prefix2"))
	s.mocks.clusterClient.updateErr = errors.New("failed to update cluster")
	d := dataDiffFrom(s.T(), originalCluster, updatedCluster)

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_updateClusterTimeout() {
	originalCluster := aTestClusterDataMap(withDNSPrefix("new-prefix1"))
	updatedCluster := aTestClusterDataMap(withDNSPrefix("new-prefix2"), with5msTimeout)
	s.mocks.clusterClient.getClusterResp = aTestCluster() // without success status
	d := dataDiffFrom(s.T(), originalCluster, updatedCluster)

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_updateNodepoolFails() {
	s.mocks.nodepoolClient.nodepoolListResp = []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np1"), withCount(1)),
	}
	originalNodepools := []any{aTestNodepoolDataMap(withName("np1"), withNodepoolCount(1))}
	updatedNodepools := []any{aTestNodepoolDataMap(withName("np1"), withNodepoolCount(5))}
	s.mocks.nodepoolClient.updateErr = errors.New("failed to update nodepool")
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_updateNodepoolTimeout() {
	s.mocks.nodepoolClient.nodepoolListResp = []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np1"), withCount(1)),
	}
	originalNodepools := []any{aTestNodepoolDataMap(withName("np1"), withNodepoolCount(1))}
	updatedNodepools := []any{aTestNodepoolDataMap(withName("np1"), withNodepoolCount(5))}
	s.mocks.nodepoolClient.nodepoolGetResp = aTestNodePool() // Without success
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools), with5msTimeout))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_createNodepoolFails() {
	originalNodepools := []any{aTestNodepoolDataMap()}
	updatedNodepools := []any{aTestNodepoolDataMap(), aTestNodepoolDataMap(withName("np2"))}
	s.mocks.nodepoolClient.createErr = errors.New("failed to create nodepool")
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_createNodepoolError() {
	originalNodepools := []any{aTestNodepoolDataMap()}
	updatedNodepools := []any{aTestNodepoolDataMap(), aTestNodepoolDataMap(withName("np2"))}
	s.mocks.nodepoolClient.nodepoolGetResp = aTestNodePool(withNodepoolStatusError())
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_createNodepoolTimeout() {
	originalNodepools := []any{aTestNodepoolDataMap()}
	updatedNodepools := []any{aTestNodepoolDataMap(), aTestNodepoolDataMap(withName("np2"))}
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools), with5msTimeout))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_deleteNodepoolFails() {
	s.mocks.nodepoolClient.nodepoolListResp = []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np1")),
		aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np2"))}
	originalNodepools := []any{aTestNodepoolDataMap(withName("np1")), aTestNodepoolDataMap(withName("np2"))}
	updatedNodepools := []any{aTestNodepoolDataMap(withName("np1")), nil}
	s.mocks.nodepoolClient.DeleteErr = errors.New("failed to delete nodepool")
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_deleteNodepoolTimeout() {
	s.mocks.nodepoolClient.nodepoolListResp = []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np1")),
		aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np2"))}
	originalNodepools := []any{aTestNodepoolDataMap(withName("np1")), aTestNodepoolDataMap(withName("np2"))}
	updatedNodepools := []any{aTestNodepoolDataMap(withName("np1")), nil}

	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools), with5msTimeout))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_nodepoolOrderChange() {
	s.mocks.nodepoolClient.nodepoolListResp = []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		aTestNodePool(withNodepoolName("np1")),
		aTestNodePool(withNodepoolName("np2")),
	}
	originalNodepools := []any{aTestNodepoolDataMap(withName("np1")), aTestNodepoolDataMap(withName("np2"))}
	updatedNodepools := []any{aTestNodepoolDataMap(withName("np2")), aTestNodepoolDataMap(withName("np1"))}
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Nil(s.mocks.nodepoolClient.DeleteNodepoolWasCalledWith)
	s.Assert().Nil(s.mocks.nodepoolClient.CreateNodepoolWasCalledWith)
	s.Assert().Nil(s.mocks.nodepoolClient.UpdatedNodepoolWasCalledWith)
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_nodepoolImmutableChange_recreate() {
	s.mocks.nodepoolClient.nodepoolListResp = []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{
		aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np1")),
	}
	s.mocks.nodepoolClient.getErr = clienterrors.ErrorWithHTTPCode(http.StatusNotFound, nil)

	originalNodepools := []any{aTestNodepoolDataMap(withName("np1"))}
	updatedNodepools := []any{aTestNodepoolDataMap(withName("np1"), withNodepoolVMSize("STANDARD_DS2v3"))}
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))
	expected := aTestNodePool(forCluster(aTestCluster().FullName), withNodepoolName("np1"))
	expected.Spec.VMSize = "STANDARD_DS2v3"

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal(s.mocks.nodepoolClient.DeleteNodepoolWasCalledWith, expected.FullName)
	s.Assert().Equal(s.mocks.nodepoolClient.CreateNodepoolWasCalledWith, expected)
	s.Assert().Nil(s.mocks.nodepoolClient.UpdatedNodepoolWasCalledWith)
}

type DeleteClusterTestSuite struct {
	suite.Suite
	ctx                context.Context
	mocks              mocks
	aksClusterResource *schema.Resource
	config             authctx.TanzuContext
}

func (s *DeleteClusterTestSuite) SetupTest() {
	s.mocks.clusterClient = &mockClusterClient{
		createClusterResp: aTestCluster(),
		getClusterResp:    aTestCluster(withStatusSuccess),
	}
	s.config = authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			AKSClusterResourceService: s.mocks.clusterClient,
		},
	}
	s.aksClusterResource = akscluster.ResourceTMCAKSCluster()
	s.ctx = context.WithValue(context.Background(), akscluster.RetryInterval, 10*time.Millisecond)
}

func (s *DeleteClusterTestSuite) Test_resourceClusterDelete() {
	s.mocks.clusterClient.getErr = clienterrors.ErrorWithHTTPCode(http.StatusNotFound, nil)
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.DeleteContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Equal(expectedFullName(), s.mocks.clusterClient.AksClusterResourceServiceDeleteCalledWith)
}

func (s *DeleteClusterTestSuite) Test_resourceClusterDelete_invalidConfig() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.DeleteContext(s.ctx, d, "config")

	s.Assert().True(result.HasError())
}

func (s *DeleteClusterTestSuite) Test_resourceClusterDelete_fails() {
	s.mocks.clusterClient.deleteErr = errors.New("cluster delete failed")
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.aksClusterResource.DeleteContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal(expectedFullName(), s.mocks.clusterClient.AksClusterResourceServiceDeleteCalledWith)
}

func (s *DeleteClusterTestSuite) Test_resourceClusterDelete_timeout() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap(with5msTimeout))

	result := s.aksClusterResource.DeleteContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
	s.Assert().Equal(expectedFullName(), s.mocks.clusterClient.AksClusterResourceServiceDeleteCalledWith)
}

type ImportClusterTestSuite struct {
	suite.Suite
	ctx                context.Context
	mocks              mocks
	aksClusterResource *schema.Resource
	config             authctx.TanzuContext
}

func (s *ImportClusterTestSuite) SetupTest() {
	s.mocks.clusterClient = &mockClusterClient{
		getClusterByIDResp: aTestCluster(withStatusSuccess),
	}
	s.mocks.nodepoolClient = &mockNodepoolClient{
		nodepoolListResp: []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{aTestNodePool(forCluster(aTestCluster().FullName))},
	}
	s.config = authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			AKSClusterResourceService:  s.mocks.clusterClient,
			AKSNodePoolResourceService: s.mocks.nodepoolClient,
		},
	}
	s.aksClusterResource = akscluster.ResourceTMCAKSCluster()
	s.ctx = context.WithValue(context.Background(), akscluster.RetryInterval, 10*time.Millisecond)
}

func (s *ImportClusterTestSuite) Test_resourceClusterImport() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, nil)
	d.SetId("test-id")

	result, err := s.aksClusterResource.Importer.StateContext(s.ctx, d, s.config)

	s.Assert().NoError(err)
	s.Assert().Len(result, 1)
	cluster, _ := akscluster.ConstructCluster(result[0])
	s.Assert().Equal(cluster.FullName.Name, "test-cluster")
	s.Assert().Equal(cluster.FullName.CredentialName, "test-cred")
	s.Assert().Equal(cluster.FullName.SubscriptionID, "sub-id")
	s.Assert().Equal(cluster.FullName.ResourceGroupName, "resource-group")
	s.Assert().NotNil(cluster.Spec)
	s.Assert().NotNil(cluster.Meta)
}

func (s *ImportClusterTestSuite) Test_resourceClusterImport_GetClusterFails() {
	s.mocks.clusterClient.getErr = errors.New("failed to get cluster by ID")
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, nil)
	d.SetId("test-id")

	_, err := s.aksClusterResource.Importer.StateContext(s.ctx, d, s.config)

	s.Assert().Error(err)
}

func (s *ImportClusterTestSuite) Test_resourceClusterImport_GetNodepoolsFails() {
	s.mocks.nodepoolClient.listErr = errors.New("failed to get nodepools")
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, nil)
	d.SetId("test-id")

	_, err := s.aksClusterResource.Importer.StateContext(s.ctx, d, s.config)

	s.Assert().Error(err)
}

func Test_pollUntilReady(t *testing.T) {
	type args struct {
		timeOut  time.Duration
		data     *schema.ResourceData
		mc       *mocks
		interval time.Duration
	}

	tests := []struct {
		name       string
		args       args
		wantError  error
		validation func(t *testing.T, args args)
	}{
		{
			name: "success",
			args: args{
				timeOut: 2 * time.Second,
				data:    schema.TestResourceDataRaw(t, akscluster.ClusterSchema, aTestClusterDataMap()),
				mc: &mocks{
					clusterClient: &mockClusterClient{
						createClusterResp: aTestCluster(),
						getClusterResp:    aTestCluster(withStatusSuccess),
					},
					nodepoolClient: &mockNodepoolClient{
						nodepoolListResp: []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{aTestNodePool()},
					},
				},
				interval: 1 * time.Second,
			},
			wantError: nil,
			validation: func(t *testing.T, args args) {
				require.Equal(t, 1, args.mc.clusterClient.AksClusterResourceServiceGetCallCount, "wrong number of calls")
			},
		},
		{
			name: "success on a second call",
			args: args{
				timeOut: 3 * time.Second,
				data:    schema.TestResourceDataRaw(t, akscluster.ClusterSchema, aTestClusterDataMap()),
				mc: &mocks{
					clusterClient: &mockClusterClient{
						createClusterResp:                        aTestCluster(),
						AksClusterResourceServiceGetPendingFirst: true,
						getClusterResp:                           aTestCluster(withStatusSuccess),
					},
					nodepoolClient: &mockNodepoolClient{
						nodepoolListResp: []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{aTestNodePool()},
					},
				},
				interval: 1 * time.Second,
			},
			wantError: nil,
			validation: func(t *testing.T, args args) {
				require.Equal(t, 2, args.mc.clusterClient.AksClusterResourceServiceGetCallCount, "wrong number of calls")
			},
		},
		{
			name: "time out",
			args: args{
				timeOut: 2 * time.Second,
				data:    schema.TestResourceDataRaw(t, akscluster.ClusterSchema, aTestClusterDataMap()),
				mc: &mocks{
					clusterClient: &mockClusterClient{
						createClusterResp: aTestCluster(),
						getClusterResp:    aTestCluster(withStatusSuccess),
					},
					nodepoolClient: &mockNodepoolClient{
						nodepoolListResp: []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{aTestNodePool()},
					},
				},
				interval: 3 * time.Second,
			},
			wantError: errors.New("Timed out waiting for READY"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare TMC
			tmc := &client.TanzuMissionControl{
				AKSClusterResourceService:  tt.args.mc.clusterClient,
				AKSNodePoolResourceService: tt.args.mc.nodepoolClient,
			}

			ctx, cancel := context.WithTimeout(context.Background(), tt.args.timeOut)
			defer cancel()

			gotErr := akscluster.PollUntilReady(ctx, tt.args.data, tmc, tt.args.interval)

			// Compare errors, ordinary reflect.DeepEqual does work here because errors have different stack field data
			if !strings.Contains(fmt.Sprintf("%v", gotErr), fmt.Sprintf("%v", tt.wantError)) {
				if tt.wantError == nil {
					t.Errorf("PollUntilReady() with duration: %v got unexpected error: %v", tt.args.interval, gotErr)
				} else {
					t.Errorf("PollUntilReady()with duration: %v. Error should be: %v, got: %v\"",
						tt.args.interval, tt.wantError, gotErr)
				}
			}
			if tt.validation != nil {
				tt.validation(t, tt.args)
			}
		})
	}
}
