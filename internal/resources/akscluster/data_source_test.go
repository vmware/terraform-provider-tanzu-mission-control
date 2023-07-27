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

type ReadDatasourceTestSuite struct {
	suite.Suite
	ctx        context.Context
	mocks      mocks
	datasource *schema.Resource
	config     authctx.TanzuContext
}

func TestDatasource(t *testing.T) {
	suite.Run(t, &ReadDatasourceTestSuite{})
}

func (s *ReadDatasourceTestSuite) SetupTest() {
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
	s.datasource = akscluster.DataSourceTMCAKSCluster()
	s.ctx = context.WithValue(context.Background(), akscluster.RetryInterval, 10*time.Millisecond)
}

func (s *ReadDatasourceTestSuite) Test_datasourceRead() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.datasource.ReadContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Equal(expectedFullName(), s.mocks.clusterClient.AksClusterResourceServiceGetCalledWith)
	s.Assert().Equal(expectedFullName(), s.mocks.nodepoolClient.AksNodePoolResourceServiceListCalledWith)
	s.Assert().Equal("test-uid", d.Id(), "expect id from REST request")
	s.Assert().NotNil(d.Get(common.MetaKey), "expected metadata from REST request")
	s.Assert().NotNil(d.Get("spec"), "expected cluster spec from REST request")
}

func (s *ReadDatasourceTestSuite) Test_datasourceRead_invalidConfig() {
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.datasource.ReadContext(s.ctx, d, "config")

	s.Assert().True(result.HasError())
}

func (s *ReadDatasourceTestSuite) Test_datasourceRead_getCluster_Err() {
	s.mocks.clusterClient.getErr = errors.New("failed to get cluster")
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.datasource.ReadContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *ReadDatasourceTestSuite) Test_datasourceRead_getCluster_NotFound() {
	s.mocks.clusterClient.getErr = clienterrors.ErrorWithHTTPCode(http.StatusNotFound, nil)
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.datasource.ReadContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Empty(d.Id())
}

func (s *ReadDatasourceTestSuite) Test_datasourceRead_getNodepools_Err() {
	s.mocks.nodepoolClient.listErr = errors.New("failed to get nodepools")
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.datasource.ReadContext(s.ctx, d, s.config)

	s.Assert().True(result.HasError())
}

func (s *ReadDatasourceTestSuite) Test_datasourceRead_getNodepools_NotFound() {
	s.mocks.nodepoolClient.listErr = clienterrors.ErrorWithHTTPCode(http.StatusNotFound, nil)
	d := schema.TestResourceDataRaw(s.T(), akscluster.ClusterSchema, aTestClusterDataMap())

	result := s.datasource.ReadContext(s.ctx, d, s.config)

	s.Assert().False(result.HasError())
	s.Assert().Equal("test-uid", d.Id(), "expect id from REST request")
}
