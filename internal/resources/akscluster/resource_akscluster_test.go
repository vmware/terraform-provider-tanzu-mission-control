/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/suite"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/akscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type mocks struct {
	clusterClient  *mockClusterClient
	nodepoolClient *mockNodepoolClient
}

func TestAKSClusterResource(t *testing.T) {
	suite.Run(t, &CreatClusterTestSuite{})
	suite.Run(t, &ReadClusterTestSuite{})
	suite.Run(t, &UpdateClusterTestSuite{})
	suite.Run(t, &DeleteClusterTestSuite{})
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
	s.config = authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			AKSClusterResourceService:  s.mocks.clusterClient,
			AKSNodePoolResourceService: s.mocks.nodepoolClient,
		},
	}
	s.aksClusterResource = akscluster.ResourceTMCAKSCluster()
	s.ctx = context.WithValue(context.Background(), akscluster.RetryInterval, 10*time.Millisecond)
}

func (s *CreatClusterTestSuite) Test_resourceClusterCreate() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())
	expectedNP := aTestNodePool(forCluster(aTestCluster()))

	result := s.aksClusterResource.CreateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().True(s.mocks.clusterClient.AksCreateClusterWasCalled, "cluster create was not called")                   //TODO: verify payload
	s.Assert().Equal(s.mocks.nodepoolClient.CreateNodepoolWasCalledWith, expectedNP, "nodepool create was not called ") //TODO: verify payload
	s.Assert().Equal(expectedFullName(), s.mocks.clusterClient.AksClusterResourceServiceGetCalledWith)
	s.Assert().Equal(expectedFullName(), s.mocks.nodepoolClient.AksNodePoolResourceServiceListCalledWith)
	s.Assert().Equal("test-uid", d.Id())
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
		nodepoolListResp: []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{aTestNodePool()},
		nodepoolGetResp:  aTestNodePool(withNodepoolStatusSuccess),
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
	expected := aTestNodePool(forCluster(aTestCluster()))
	expected.Spec.Count = 5

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Nil(s.mocks.clusterClient.AksUpdateClusterWasCalledWith)
	s.Assert().Equal(expected, s.mocks.nodepoolClient.UpdatedNodepoolWasCalledWith)
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_addNodepool() {
	originalNodepools := []any{aTestNodepoolDataMap()}
	updatedNodepools := []any{aTestNodepoolDataMap(), aTestNodepoolDataMap(withName("np2"))}
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))
	expected := aTestNodePool(forCluster(aTestCluster()), withNodepoolName("np2"))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Nil(s.mocks.clusterClient.AksUpdateClusterWasCalledWith)
	s.Assert().Equal(expected, s.mocks.nodepoolClient.CreateNodepoolWasCalledWith)
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_deleteNodepool() {
	originalNodepools := []any{aTestNodepoolDataMap(), aTestNodepoolDataMap(withName("np2"))}
	updatedNodepools := []any{aTestNodepoolDataMap(), nil}
	s.mocks.nodepoolClient.getErr = clienterrors.ErrorWithHTTPCode(http.StatusNotFound, nil)

	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))
	expected := aTestNodePool(forCluster(aTestCluster()), withNodepoolName("np2")).FullName

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
	originalNodepools := []any{aTestNodepoolDataMap(withNodepoolCount(1))}
	updatedNodepools := []any{aTestNodepoolDataMap(withNodepoolCount(5))}
	s.mocks.nodepoolClient.updateErr = errors.New("failed to update nodepool")
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_updateNodepoolTimeout() {
	originalNodepools := []any{aTestNodepoolDataMap(withNodepoolCount(1))}
	updatedNodepools := []any{aTestNodepoolDataMap(withNodepoolCount(5))}
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
	originalNodepools := []any{aTestNodepoolDataMap(), aTestNodepoolDataMap(withName("np2"))}
	updatedNodepools := []any{aTestNodepoolDataMap(), nil}
	s.mocks.nodepoolClient.DeleteErr = errors.New("failed to delete nodepool")
	d := dataDiffFrom(s.T(), aTestClusterDataMap(withNodepools(originalNodepools)), aTestClusterDataMap(withNodepools(updatedNodepools)))

	result := s.aksClusterResource.UpdateContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *UpdateClusterTestSuite) Test_resourceClusterUpdate_deleteNodepoolTimeout() {
	originalNodepools := []any{aTestNodepoolDataMap(), aTestNodepoolDataMap(withName("np2"))}
	updatedNodepools := []any{aTestNodepoolDataMap(), nil}

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
