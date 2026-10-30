// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzukubernetescluster

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	clusterclassclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/clusterclass"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	tanzukubernetesclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/tanzukubernetescluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clusterclassmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clusterclass"
	tkcmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
	tkccommonmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/common"
	tkcnodepoolmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/nodepool"
)

const (
	testClusterName        = "test-cluster"
	testManagementCluster  = "mgmt"
	testProvisionerName    = "prov"
	testWorkerPool1        = "worker-pool-1"
	testWorkerPool2        = "worker-pool-2"
	testWorkerPoolInit     = "worker-pool-init"
	testErrConnectionReset = "connection reset by peer"
	testK8sVersion         = "v1.28.0"
	testNodeClass          = "node-class"
	testPoolA              = "pool-a"
	testPoolB              = "pool-b"
	testClusterClass       = "test-class"
)

type mockNodePoolClientService struct {
	tanzukubernetesclusterclient.ClientService
	listFunc          func(clusterFn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error)
	createFunc        func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error)
	updateFunc        func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error)
	deleteFunc        func(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) error
	getFunc           func(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error)
	clusterCreateFunc func(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error)
	clusterGetFunc    func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error)
}

func (m *mockNodePoolClientService) TanzuKubernetesClusterResourceServiceCreate(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
	if m.clusterCreateFunc != nil {
		return m.clusterCreateFunc(req)
	}

	return nil, nil
}

func (m *mockNodePoolClientService) TanzuKubernetesClusterResourceServiceGet(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
	if m.clusterGetFunc != nil {
		return m.clusterGetFunc(fn)
	}

	return nil, nil
}

func (m *mockNodePoolClientService) TanzuKubernetesClusterNodePoolResourceServiceList(clusterFn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error) {
	if m.listFunc != nil {
		return m.listFunc(clusterFn)
	}

	return nil, nil
}

func (m *mockNodePoolClientService) TanzuKubernetesClusterNodePoolResourceServiceCreate(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
	if m.createFunc != nil {
		return m.createFunc(req)
	}

	return nil, nil
}

func (m *mockNodePoolClientService) TanzuKubernetesClusterNodePoolResourceServiceUpdate(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
	if m.updateFunc != nil {
		return m.updateFunc(req)
	}

	return nil, nil
}

func (m *mockNodePoolClientService) TanzuKubernetesClusterNodePoolResourceServiceDelete(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(fn)
	}

	return nil
}

func (m *mockNodePoolClientService) TanzuKubernetesClusterNodePoolResourceServiceGet(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
	if m.getFunc != nil {
		return m.getFunc(fn)
	}

	return nil, nil
}

// TestWaitNodePoolsReady_EmptyExpected verifies that waitNodePoolsReady returns immediately
// without error when no node pools are expected.
func TestWaitNodePoolsReady_EmptyExpected(t *testing.T) {
	ctx := context.Background()
	config := &authctx.TanzuContext{}
	clusterFn := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
		Name: testClusterName,
	}

	result, err := waitNodePoolsReady(ctx, config, clusterFn, nil)
	assert.NoError(t, err)
	assert.Nil(t, result)

	result, err = waitNodePoolsReady(ctx, config, clusterFn, []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{})
	assert.NoError(t, err)
	assert.Nil(t, result)
}

// TestEvaluateExpectedNodePools_EmptyName tests handling expected nodepool entries with nil or empty FullName.
func TestEvaluateExpectedNodePools_EmptyName(t *testing.T) {
	clusterFn := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
		Name: testClusterName,
	}

	expectedNps := []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
		{
			FullName: nil,
		},
		{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: "",
			},
		},
	}

	npStatusMap := make(map[string]string)
	matched, allReady, err := evaluateExpectedNodePools(clusterFn, expectedNps, nil, npStatusMap)
	assert.NoError(t, err)
	assert.False(t, allReady)
	assert.Empty(t, matched)
}

// TestWaitNodePoolsReady_SuccessReady tests successful polling transition from CREATING to READY.
func TestWaitNodePoolsReady_SuccessReady(t *testing.T) {
	origInterval := nodePoolPollInterval
	nodePoolPollInterval = 10 * time.Millisecond

	defer func() {
		nodePoolPollInterval = origInterval
	}()

	clusterFn := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
		ManagementClusterName: testManagementCluster,
		ProvisionerName:       testProvisionerName,
		Name:                  testClusterName,
	}

	expectedNps := []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
		{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: testWorkerPool1,
			},
		},
	}

	callCount := 0
	readyPhase := tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseREADY
	creatingPhase := tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseCREATING

	mockClient := &mockNodePoolClientService{
		listFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error) {
			callCount++

			if callCount == 1 {
				// First call: CREATING
				return &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData{
					Nodepools: []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
						{
							FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
								Name: testWorkerPool1,
							},
							Status: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatus{
								Phase: &creatingPhase,
							},
						},
					},
				}, nil
			}

			// Second call: READY
			return &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData{
				Nodepools: []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
					{
						FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
							Name: testWorkerPool1,
						},
						Status: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatus{
							Phase: &readyPhase,
						},
					},
				},
			}, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	matched, err := waitNodePoolsReady(ctx, config, clusterFn, expectedNps)
	require.NoError(t, err)
	require.Len(t, matched, 1)
	assert.Equal(t, testWorkerPool1, matched[0].FullName.Name)
	assert.Equal(t, readyPhase, *matched[0].Status.Phase)
	assert.GreaterOrEqual(t, callCount, 2)
}

// TestWaitNodePoolsReady_MissingNodePool_Timeout tests timeout error formatting when a node pool never appears remotely.
func TestWaitNodePoolsReady_MissingNodePool_Timeout(t *testing.T) {
	origInterval := nodePoolPollInterval
	nodePoolPollInterval = 10 * time.Millisecond

	defer func() {
		nodePoolPollInterval = origInterval
	}()

	clusterFn := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
		ManagementClusterName: testManagementCluster,
		ProvisionerName:       testProvisionerName,
		Name:                  testClusterName,
	}

	// Expected nodepool that never appears remotely (e.g. lost due to webhook denial)
	expectedNps := []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
		{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: "worker-pool-lost",
			},
		},
	}

	mockClient := &mockNodePoolClientService{
		listFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error) {
			// Remote returns empty nodepools
			return &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData{
				Nodepools: []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{},
			}, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	matched, err := waitNodePoolsReady(ctx, config, clusterFn, expectedNps)
	assert.Nil(t, matched)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Timeout exceeded while waiting for the cluster node pools to be ready")
	assert.Contains(t, err.Error(), `Node pool "worker-pool-lost" is in status: not found in cluster`)
}

// TestWaitNodePoolsReady_TerminalFailures tests fast-fail behavior when a node pool enters ERROR or UPGRADEFAILED phase.
func TestWaitNodePoolsReady_TerminalFailures(t *testing.T) {
	origInterval := nodePoolPollInterval
	nodePoolPollInterval = 10 * time.Millisecond

	defer func() {
		nodePoolPollInterval = origInterval
	}()

	testCases := []struct {
		name          string
		poolName      string
		phase         tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhase
		expectedError string
	}{
		{
			name:          "error phase fast fail",
			poolName:      "worker-pool-err",
			phase:         tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseERROR,
			expectedError: "TKG Cluster node pool errored",
		},
		{
			name:          "upgrade failed phase fast fail",
			poolName:      "worker-pool-upgrade-failed",
			phase:         tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseUPGRADEFAILED,
			expectedError: "TKG Cluster node pool upgrade failed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clusterFn := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
				ManagementClusterName: testManagementCluster,
				ProvisionerName:       testProvisionerName,
				Name:                  testClusterName,
			}

			expectedNps := []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
				{
					FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
						Name: tc.poolName,
					},
				},
			}

			phaseCopy := tc.phase
			mockClient := &mockNodePoolClientService{
				listFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error) {
					return &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData{
						Nodepools: []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
							{
								FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
									Name: tc.poolName,
								},
								Status: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatus{
									Phase: &phaseCopy,
								},
							},
						},
					}, nil
				},
			}

			config := &authctx.TanzuContext{
				TMCConnection: &client.TanzuMissionControl{
					TanzuKubernetesClusterResourceService: mockClient,
				},
			}

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			matched, err := waitNodePoolsReady(ctx, config, clusterFn, expectedNps)
			assert.Nil(t, matched)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.expectedError)
		})
	}
}

// TestWaitNodePoolsReady_NilStatusSafety tests safe handling of nil Status or nil Phase fields during cluster boot.
func TestWaitNodePoolsReady_NilStatusSafety(t *testing.T) {
	origInterval := nodePoolPollInterval
	nodePoolPollInterval = 10 * time.Millisecond

	defer func() {
		nodePoolPollInterval = origInterval
	}()

	clusterFn := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
		ManagementClusterName: testManagementCluster,
		ProvisionerName:       testProvisionerName,
		Name:                  testClusterName,
	}

	expectedNps := []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
		{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: testWorkerPoolInit,
			},
		},
	}

	callCount := 0
	readyPhase := tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseREADY

	mockClient := &mockNodePoolClientService{
		listFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error) {
			callCount++

			if callCount == 1 {
				// Status or Status.Phase is nil initially
				return &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData{
					Nodepools: []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
						{
							FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
								Name: testWorkerPoolInit,
							},
							Status: nil,
						},
					},
				}, nil
			}

			return &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData{
				Nodepools: []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
					{
						FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
							Name: testWorkerPoolInit,
						},
						Status: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatus{
							Phase: &readyPhase,
						},
					},
				},
			}, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	matched, err := waitNodePoolsReady(ctx, config, clusterFn, expectedNps)
	require.NoError(t, err)
	require.Len(t, matched, 1)
	assert.Equal(t, readyPhase, *matched[0].Status.Phase)
}

// TestWaitNodePoolsReady_ContextCanceled tests fast return when context is canceled rather than timed out.
func TestWaitNodePoolsReady_ContextCanceled(t *testing.T) {
	origInterval := nodePoolPollInterval
	nodePoolPollInterval = 10 * time.Millisecond

	defer func() {
		nodePoolPollInterval = origInterval
	}()

	clusterFn := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
		Name: testClusterName,
	}

	mockClient := &mockNodePoolClientService{
		listFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error) {
			return &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData{}, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	expectedNps := []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
		{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: "worker-pool-cancel",
			},
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	matched, err := waitNodePoolsReady(ctx, config, clusterFn, expectedNps)
	assert.Nil(t, matched)
	require.Error(t, err)
	assert.True(t, errors.Is(err, context.Canceled))
}

// TestWaitNodePoolsReady_UnauthorizedAndListErrorRecovery tests token refresh on 401 and logging transient list errors.
func TestWaitNodePoolsReady_UnauthorizedAndListErrorRecovery(t *testing.T) {
	origInterval := nodePoolPollInterval
	nodePoolPollInterval = 10 * time.Millisecond

	defer func() {
		nodePoolPollInterval = origInterval
	}()

	origRefresh := authctx.RefreshUserAuthContext
	authctx.RefreshUserAuthContext = func(config *authctx.TanzuContext, refreshCondition func(error) bool, err error) {}

	defer func() {
		authctx.RefreshUserAuthContext = origRefresh
	}()

	clusterFn := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
		ManagementClusterName: testManagementCluster,
		ProvisionerName:       testProvisionerName,
		Name:                  testClusterName,
	}

	expectedNps := []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
		{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: testWorkerPool1,
			},
		},
	}

	readyPhase := tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseREADY
	callCount := 0
	mockClient := &mockNodePoolClientService{
		listFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error) {
			callCount++

			if callCount == 1 {
				// Call 1: 401 Unauthorized error
				return nil, errors.New("status : 401 Unauthorized")
			}

			if callCount == 2 {
				// Call 2: Generic error (logged)
				return nil, errors.New("network timeout")
			}

			// Call 3: success
			return &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData{
				Nodepools: []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
					{
						FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
							Name: testWorkerPool1,
						},
						Status: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatus{
							Phase: &readyPhase,
						},
					},
				},
			}, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	matched, err := waitNodePoolsReady(ctx, config, clusterFn, expectedNps)
	require.NoError(t, err)
	require.Len(t, matched, 1)
	assert.Equal(t, testWorkerPool1, matched[0].FullName.Name)
	assert.GreaterOrEqual(t, callCount, 3)
}

// TestWaitNodePoolsReady_ListErrorTimeout tests formatting timeout error when list repeatedly failed.
func TestWaitNodePoolsReady_ListErrorTimeout(t *testing.T) {
	origInterval := nodePoolPollInterval
	nodePoolPollInterval = 10 * time.Millisecond

	defer func() {
		nodePoolPollInterval = origInterval
	}()

	clusterFn := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
		Name: testClusterName,
	}

	expectedNps := []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
		{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: "worker-pool-timeout-notfound",
			},
		},
	}

	mockClient := &mockNodePoolClientService{
		listFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error) {
			return nil, errors.New("backend error")
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Millisecond)
	defer cancel()

	matched, err := waitNodePoolsReady(ctx, config, clusterFn, expectedNps)
	assert.Nil(t, matched)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Timeout exceeded while waiting for the cluster node pools to be ready")
	assert.Contains(t, err.Error(), `Node pool "worker-pool-timeout-notfound" is in status: not found in cluster`)
}

// TestCalculateBackoffWithJitter verifies exponential scaling, caps, fallback defaults, and random distribution.
func TestCalculateBackoffWithJitter(t *testing.T) {
	// 1. Basic calculation and bounds
	base := 10 * time.Second
	max := 60 * time.Second
	factor := 2.0
	jitter := 0.2

	sleep, next := calculateBackoffWithJitter(base, max, factor, jitter)
	// sleep should be in [base * 0.8, base * 1.2] => [8s, 12s]
	assert.GreaterOrEqual(t, sleep, 8*time.Second)
	assert.LessOrEqual(t, sleep, 12*time.Second)
	assert.Equal(t, 20*time.Second, next)

	// 2. Cap at maxBackoff
	sleepMax, nextMax := calculateBackoffWithJitter(50*time.Second, max, factor, jitter)
	assert.LessOrEqual(t, sleepMax, max)
	assert.Equal(t, max, nextMax)

	sleepCapped, _ := calculateBackoffWithJitter(100*time.Second, 50*time.Second, 1.5, 0.2)
	assert.Equal(t, 50*time.Second, sleepCapped)

	// 3. Fallbacks on invalid parameters
	sleepDef, nextDef := calculateBackoffWithJitter(0, 0, 0, -1)
	assert.GreaterOrEqual(t, sleepDef, time.Duration(float64(nodePoolInitialBackoff)*0.79))
	assert.LessOrEqual(t, sleepDef, nodePoolMaxBackoff)
	assert.Equal(t, time.Duration(float64(nodePoolInitialBackoff)*nodePoolBackoffFactor), nextDef)

	// 4. Jitter randomness distribution across multiple invocations
	durations := make(map[time.Duration]bool)

	for i := 0; i < 50; i++ {
		s, _ := calculateBackoffWithJitter(base, max, factor, jitter)
		durations[s] = true
	}
	// With 50 runs and nanosecond random jitter, there should be multiple distinct durations
	assert.Greater(t, len(durations), 10, "jitter should produce random variations")
}

// TestCreateNodePoolWithRetry_SuccessFirstAttempt tests immediate success without retry attempts.
func TestCreateNodePoolWithRetry_SuccessFirstAttempt(t *testing.T) {
	callCount := 0
	mockClient := &mockNodePoolClientService{
		createFunc: func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
			callCount++
			return req, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{
		Nodepool: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: testWorkerPool1,
			},
		},
	}

	err := createNodePoolWithRetry(context.Background(), config, req, 2*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, 1, callCount)
}

func testNodePoolRetryRunner(
	t *testing.T,
	mockClient *mockNodePoolClientService,
	op func(ctx context.Context, config *authctx.TanzuContext, req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, timeout time.Duration) error,
) {
	origInit := nodePoolInitialBackoff
	origMax := nodePoolMaxBackoff

	nodePoolInitialBackoff = 5 * time.Millisecond
	nodePoolMaxBackoff = 20 * time.Millisecond

	defer func() {
		nodePoolInitialBackoff = origInit
		nodePoolMaxBackoff = origMax
	}()

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{
		Nodepool: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: testWorkerPool1,
			},
		},
	}

	err := op(context.Background(), config, req, 2*time.Second)
	assert.NoError(t, err)
}

// TestCreateNodePoolWithRetry_RetryOnTransientAndSucceed tests recovery after webhook rejection errors.
func TestCreateNodePoolWithRetry_RetryOnTransientAndSucceed(t *testing.T) {
	callCount := 0
	mockClient := &mockNodePoolClientService{
		createFunc: func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
			callCount++

			if callCount < 3 {
				return nil, &testError{msg: "admission webhook \"capi.validating.kubernetes.vmware.com\" denied the request: update cannot be initiated as cluster's SystemChecksSucceeded condition is not True"}
			}

			return req, nil
		},
	}

	testNodePoolRetryRunner(t, mockClient, createNodePoolWithRetry)

	assert.Equal(t, 3, callCount)
}

// TestCreateNodePoolWithRetry_NonRetryableError tests immediate failure on non-retryable 400 Bad Request error.
func TestCreateNodePoolWithRetry_NonRetryableError(t *testing.T) {
	callCount := 0
	mockClient := &mockNodePoolClientService{
		createFunc: func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
			callCount++

			return nil, &testError{msg: "status : 400 Bad Request, storage class not found"}
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{
		Nodepool: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: "worker-pool-invalid",
			},
		},
	}

	err := createNodePoolWithRetry(context.Background(), config, req, 2*time.Second)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "storage class not found")
	assert.Equal(t, 1, callCount)
}

// TestCreateNodePoolWithRetry_Timeout tests context deadline expiration with informative error reporting.
func TestCreateNodePoolWithRetry_Timeout(t *testing.T) {
	origInit, origMax := nodePoolInitialBackoff, nodePoolMaxBackoff
	nodePoolInitialBackoff = 10 * time.Millisecond
	nodePoolMaxBackoff = 20 * time.Millisecond

	defer func() {
		nodePoolInitialBackoff = origInit
		nodePoolMaxBackoff = origMax
	}()

	mockClient := &mockNodePoolClientService{
		createFunc: func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
			return nil, &testError{msg: testErrConnectionReset}
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{
		Nodepool: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: "worker-pool-timeout",
			},
		},
	}

	err := createNodePoolWithRetry(context.Background(), config, req, 35*time.Millisecond)
	require.Error(t, err)
	assert.Contains(t, err.Error(), `timed out waiting for nodepool "worker-pool-timeout" creation to be accepted after`)
}

// TestCreateNodePoolWithRetry_AlreadyExistsIdempotent tests idempotent success when the node pool already exists.
func TestCreateNodePoolWithRetry_AlreadyExistsIdempotent(t *testing.T) {
	mockClient := &mockNodePoolClientService{
		createFunc: func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
			return nil, &testError{msg: "resource already exists"}
		},
		getFunc: func(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
			return &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{
				Nodepool: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
					FullName: fn,
				},
			}, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{
		Nodepool: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: "worker-pool-exists",
			},
		},
	}

	err := createNodePoolWithRetry(context.Background(), config, req, 2*time.Second)
	assert.NoError(t, err)
}

// TestUpdateNodePoolWithRetry_RetryAndSucceed tests recovery after optimistic locking conflict during update.
func TestUpdateNodePoolWithRetry_RetryAndSucceed(t *testing.T) {
	callCount := 0
	mockClient := &mockNodePoolClientService{
		updateFunc: func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
			callCount++

			if callCount == 1 {
				return nil, &testError{msg: "the object has been modified; please apply your changes to the latest version"}
			}

			return req, nil
		},
	}

	testNodePoolRetryRunner(t, mockClient, updateNodePoolWithRetry)

	assert.Equal(t, 2, callCount)
}

// TestDeleteNodePoolWithRetry_RetryAndSucceed tests recovery after transient network drop during delete.
func TestDeleteNodePoolWithRetry_RetryAndSucceed(t *testing.T) {
	origInit, origMax := nodePoolInitialBackoff, nodePoolMaxBackoff
	nodePoolInitialBackoff = 5 * time.Millisecond
	nodePoolMaxBackoff = 20 * time.Millisecond

	defer func() {
		nodePoolInitialBackoff = origInit
		nodePoolMaxBackoff = origMax
	}()

	callCount := 0
	mockClient := &mockNodePoolClientService{
		deleteFunc: func(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) error {
			callCount++

			if callCount == 1 {
				return &testError{msg: testErrConnectionReset}
			}

			return nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	fn := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
		Name: testWorkerPool1,
	}

	err := deleteNodePoolWithRetry(context.Background(), config, fn, 2*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, 2, callCount)
}

func TestUpdateNodePoolWithRetry_NonRetryableError(t *testing.T) {
	callCount := 0
	mockClient := &mockNodePoolClientService{
		updateFunc: func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
			callCount++
			return nil, errors.New("status : 400 Bad Request, invalid field")
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{
		Nodepool: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: testWorkerPool1,
			},
		},
	}

	err := updateNodePoolWithRetry(context.Background(), config, req, 2*time.Second)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "400 Bad Request")
	assert.Equal(t, 1, callCount)
}

func TestUpdateNodePoolWithRetry_Timeout(t *testing.T) {
	origInit, origMax := nodePoolInitialBackoff, nodePoolMaxBackoff
	nodePoolInitialBackoff = 5 * time.Millisecond
	nodePoolMaxBackoff = 10 * time.Millisecond

	defer func() {
		nodePoolInitialBackoff = origInit
		nodePoolMaxBackoff = origMax
	}()

	mockClient := &mockNodePoolClientService{
		updateFunc: func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
			return nil, &testError{msg: testErrConnectionReset}
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{
		Nodepool: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: testWorkerPool1,
			},
		},
	}

	err := updateNodePoolWithRetry(context.Background(), config, req, 20*time.Millisecond)
	require.Error(t, err)
	assert.Contains(t, err.Error(), `timed out waiting for nodepool "worker-pool-1" update to be accepted after`)
}

func TestDeleteNodePoolWithRetry_NotFoundSuccess(t *testing.T) {
	mockClient := &mockNodePoolClientService{
		deleteFunc: func(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) error {
			return errors.New("status : 404 Not Found")
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	fn := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
		Name: testWorkerPool1,
	}

	err := deleteNodePoolWithRetry(context.Background(), config, fn, 2*time.Second)
	assert.NoError(t, err)
}

func TestDeleteNodePoolWithRetry_NonRetryableError(t *testing.T) {
	callCount := 0
	mockClient := &mockNodePoolClientService{
		deleteFunc: func(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) error {
			callCount++
			return errors.New("status : 403 Forbidden")
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	fn := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
		Name: testWorkerPool1,
	}

	err := deleteNodePoolWithRetry(context.Background(), config, fn, 2*time.Second)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "403 Forbidden")
	assert.Equal(t, 1, callCount)
}

func TestDeleteNodePoolWithRetry_Timeout(t *testing.T) {
	origInit, origMax := nodePoolInitialBackoff, nodePoolMaxBackoff
	nodePoolInitialBackoff = 5 * time.Millisecond
	nodePoolMaxBackoff = 10 * time.Millisecond

	defer func() {
		nodePoolInitialBackoff = origInit
		nodePoolMaxBackoff = origMax
	}()

	mockClient := &mockNodePoolClientService{
		deleteFunc: func(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) error {
			return &testError{msg: testErrConnectionReset}
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	fn := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
		Name: testWorkerPool1,
	}

	err := deleteNodePoolWithRetry(context.Background(), config, fn, 20*time.Millisecond)
	require.Error(t, err)
	assert.Contains(t, err.Error(), `timed out waiting for nodepool "worker-pool-1" deletion to be accepted after`)
}

func TestCreateClusterWithRetry_SuccessFirstAttempt(t *testing.T) {
	callCount := 0
	mockClient := &mockNodePoolClientService{
		clusterCreateFunc: func(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			callCount++
			return req, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
		TanzuKubernetesCluster: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
			FullName: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
				Name: testClusterName,
			},
		},
	}

	err := createClusterWithRetry(context.Background(), config, req, 2*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, 1, callCount)
}

func TestCreateClusterWithRetry_AlreadyExists_ActiveIdempotent(t *testing.T) {
	readyPhase := tkcmodels.VmwareTanzuManageV1alpha1CommonClusterStatusPhaseREADY
	mockClient := &mockNodePoolClientService{
		clusterCreateFunc: func(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			return nil, errors.New("status : 409 Conflict, cluster already exists")
		},
		clusterGetFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			return &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
				TanzuKubernetesCluster: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
					FullName: fn,
					Status: &tkcmodels.VmwareTanzuManageV1alpha1CommonClusterStatus{
						Phase: &readyPhase,
					},
				},
			}, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
		TanzuKubernetesCluster: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
			FullName: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
				Name: testClusterName,
			},
		},
	}

	err := createClusterWithRetry(context.Background(), config, req, 2*time.Second)
	assert.NoError(t, err)
}

func TestCreateClusterWithRetry_AlreadyExists_NilStatus(t *testing.T) {
	mockClient := &mockNodePoolClientService{
		clusterCreateFunc: func(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			return nil, errors.New("resource already exists")
		},
		clusterGetFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			return &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
				TanzuKubernetesCluster: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
					FullName: fn,
					Status:   nil,
				},
			}, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
		TanzuKubernetesCluster: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
			FullName: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
				Name: testClusterName,
			},
		},
	}

	err := createClusterWithRetry(context.Background(), config, req, 2*time.Second)
	assert.NoError(t, err)
}

func TestCreateClusterWithRetry_AlreadyExists_Deleting_ThenSucceed(t *testing.T) {
	origInit, origMax := nodePoolInitialBackoff, nodePoolMaxBackoff
	nodePoolInitialBackoff = 5 * time.Millisecond
	nodePoolMaxBackoff = 20 * time.Millisecond

	defer func() {
		nodePoolInitialBackoff = origInit
		nodePoolMaxBackoff = origMax
	}()

	deletingPhase := tkcmodels.VmwareTanzuManageV1alpha1CommonClusterStatusPhaseDELETING
	createCount := 0
	mockClient := &mockNodePoolClientService{
		clusterCreateFunc: func(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			createCount++

			if createCount == 1 {
				return nil, errors.New("status : 409 Conflict, already exists")
			}

			return req, nil
		},
		clusterGetFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			return &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
				TanzuKubernetesCluster: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
					FullName: fn,
					Status: &tkcmodels.VmwareTanzuManageV1alpha1CommonClusterStatus{
						Phase: &deletingPhase,
					},
				},
			}, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
		TanzuKubernetesCluster: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
			FullName: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
				Name: testClusterName,
			},
		},
	}

	err := createClusterWithRetry(context.Background(), config, req, 2*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, 2, createCount)
}

func TestCreateClusterWithRetry_TransientRateLimit_ThenSucceed(t *testing.T) {
	origInit, origMax := nodePoolInitialBackoff, nodePoolMaxBackoff
	nodePoolInitialBackoff = 5 * time.Millisecond
	nodePoolMaxBackoff = 20 * time.Millisecond

	defer func() {
		nodePoolInitialBackoff = origInit
		nodePoolMaxBackoff = origMax
	}()

	createCount := 0
	mockClient := &mockNodePoolClientService{
		clusterCreateFunc: func(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			createCount++

			if createCount == 1 {
				return nil, errors.New("status : 429 Too Many Requests")
			}

			return req, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
		TanzuKubernetesCluster: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
			FullName: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
				Name: testClusterName,
			},
		},
	}

	err := createClusterWithRetry(context.Background(), config, req, 2*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, 2, createCount)
}

func TestCreateClusterWithRetry_NonRetryableError(t *testing.T) {
	callCount := 0
	mockClient := &mockNodePoolClientService{
		clusterCreateFunc: func(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			callCount++

			return nil, errors.New("status : 400 Bad Request, invalid cluster topology")
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
		TanzuKubernetesCluster: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
			FullName: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
				Name: testClusterName,
			},
		},
	}

	err := createClusterWithRetry(context.Background(), config, req, 2*time.Second)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "400 Bad Request")
	assert.Equal(t, 1, callCount)
}

func TestCreateClusterWithRetry_Timeout(t *testing.T) {
	origInit, origMax := nodePoolInitialBackoff, nodePoolMaxBackoff
	nodePoolInitialBackoff = 5 * time.Millisecond
	nodePoolMaxBackoff = 10 * time.Millisecond

	defer func() {
		nodePoolInitialBackoff = origInit
		nodePoolMaxBackoff = origMax
	}()

	mockClient := &mockNodePoolClientService{
		clusterCreateFunc: func(req *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			return nil, errors.New("status : 503 Service Unavailable")
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	req := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
		TanzuKubernetesCluster: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
			FullName: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
				Name: testClusterName,
			},
		},
	}

	err := createClusterWithRetry(context.Background(), config, req, 20*time.Millisecond)
	require.Error(t, err)
	assert.Contains(t, err.Error(), `timed out waiting for cluster "test-cluster" creation to be accepted after`)
}

func TestSuppressNodePoolsOrderChanges_PreservesOrderAndAppendsNewPool(t *testing.T) {
	rawMap := map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
		SpecKey: []interface{}{
			map[string]interface{}{
				TopologyKey: []interface{}{
					map[string]interface{}{
						VersionKey:          testK8sVersion,
						ClusterVariablesKey: "{}",
						NodePoolKey: []interface{}{
							map[string]interface{}{
								NameKey: testWorkerPool2,
								SpecKey: []interface{}{
									map[string]interface{}{
										WorkerClassKey: testNodeClass,
									},
								},
							},
							map[string]interface{}{
								NameKey: testWorkerPool1,
								SpecKey: []interface{}{
									map[string]interface{}{
										WorkerClassKey: testNodeClass,
									},
								},
							},
							map[string]interface{}{
								NameKey: "worker-pool-3",
								SpecKey: []interface{}{
									map[string]interface{}{
										WorkerClassKey: testNodeClass,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	d := schema.TestResourceDataRaw(t, tanzuKubernetesClusterSchema, rawMap)

	oldNodePools := []interface{}{
		map[string]interface{}{
			NameKey: testWorkerPool1,
		},
		map[string]interface{}{
			NameKey: testWorkerPool2,
		},
	}

	suppressNodePoolsOrderChanges(oldNodePools, d)

	specData := d.Get(SpecKey).([]interface{})[0].(map[string]interface{})
	topologyData := specData[TopologyKey].([]interface{})[0].(map[string]interface{})
	orderedPools := topologyData[NodePoolKey].([]interface{})

	require.Len(t, orderedPools, 3)
	assert.Equal(t, testWorkerPool1, orderedPools[0].(map[string]interface{})[NameKey])
	assert.Equal(t, testWorkerPool2, orderedPools[1].(map[string]interface{})[NameKey])
	assert.Equal(t, "worker-pool-3", orderedPools[2].(map[string]interface{})[NameKey])
}

func TestSuppressNodePoolsOrderChanges_EmptyOldNodePools(t *testing.T) {
	rawMap := map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
		SpecKey: []interface{}{
			map[string]interface{}{
				TopologyKey: []interface{}{
					map[string]interface{}{
						VersionKey:          testK8sVersion,
						ClusterVariablesKey: "{}",
						NodePoolKey: []interface{}{
							map[string]interface{}{
								NameKey: testWorkerPool1,
								SpecKey: []interface{}{
									map[string]interface{}{
										WorkerClassKey: testNodeClass,
									},
								},
							},
							map[string]interface{}{
								NameKey: testWorkerPool2,
								SpecKey: []interface{}{
									map[string]interface{}{
										WorkerClassKey: testNodeClass,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	d := schema.TestResourceDataRaw(t, tanzuKubernetesClusterSchema, rawMap)

	suppressNodePoolsOrderChanges([]interface{}{}, d)

	specData := d.Get(SpecKey).([]interface{})[0].(map[string]interface{})
	topologyData := specData[TopologyKey].([]interface{})[0].(map[string]interface{})
	orderedPools := topologyData[NodePoolKey].([]interface{})

	require.Len(t, orderedPools, 2)
	assert.Equal(t, testWorkerPool1, orderedPools[0].(map[string]interface{})[NameKey])
	assert.Equal(t, testWorkerPool2, orderedPools[1].(map[string]interface{})[NameKey])
}

func TestSuppressNodePoolsOrderChanges_SafeNilHandling(t *testing.T) {
	rawMap := map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
	}

	d := schema.TestResourceDataRaw(t, tanzuKubernetesClusterSchema, rawMap)
	assert.NotPanics(t, func() {
		suppressNodePoolsOrderChanges(nil, d)
	})
}

func TestRemoveUnspecifiedNodePoolsOverrides_MatchesByNameRegardlessOfOrder(t *testing.T) {
	tfNodePools := []interface{}{
		map[string]interface{}{
			NameKey: testPoolB,
			SpecKey: []interface{}{
				map[string]interface{}{
					OverridesKey: `{"keyB": "valB"}`,
				},
			},
		},
		map[string]interface{}{
			NameKey: testPoolA,
			SpecKey: []interface{}{
				map[string]interface{}{
					OverridesKey: `{"keyA": "valA"}`,
				},
			},
		},
	}

	poolA := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
		FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
			Name: testPoolA,
		},
		Spec: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolSpec{
			Overrides: []*tkccommonmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterClusterVariable{
				{Name: "keyA", Value: "valA"},
				{Name: "droppedA", Value: "drop"},
			},
		},
	}
	poolB := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
		FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
			Name: testPoolB,
		},
		Spec: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolSpec{
			Overrides: []*tkccommonmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterClusterVariable{
				{Name: "keyB", Value: "valB"},
				{Name: "droppedB", Value: "drop"},
			},
		},
	}

	model := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
		Spec: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterSpec{
			Topology: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTopology{
				NodePools: []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
					poolA,
					poolB,
				},
			},
		},
	}

	removeUnspecifiedNodePoolsOverrides(tfNodePools, model)

	require.Len(t, poolA.Spec.Overrides, 1)
	assert.Equal(t, "keyA", poolA.Spec.Overrides[0].Name)
	assert.Equal(t, "valA", poolA.Spec.Overrides[0].Value)

	require.Len(t, poolB.Spec.Overrides, 1)
	assert.Equal(t, "keyB", poolB.Spec.Overrides[0].Name)
	assert.Equal(t, "valB", poolB.Spec.Overrides[0].Value)
}

func TestRemoveUnspecifiedNodePoolsOverrides_RemotePoolMissing_NoPanic(t *testing.T) {
	tfNodePools := []interface{}{
		map[string]interface{}{
			NameKey: testPoolA,
			SpecKey: []interface{}{
				map[string]interface{}{
					OverridesKey: `{"keyA": "valA"}`,
				},
			},
		},
		map[string]interface{}{
			NameKey: testPoolB,
			SpecKey: []interface{}{
				map[string]interface{}{
					OverridesKey: `{"keyB": "valB"}`,
				},
			},
		},
	}

	poolA := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
		FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
			Name: testPoolA,
		},
		Spec: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolSpec{
			Overrides: []*tkccommonmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterClusterVariable{
				{Name: "keyA", Value: "valA"},
			},
		},
	}

	model := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
		Spec: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterSpec{
			Topology: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTopology{
				NodePools: []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
					poolA,
				},
			},
		},
	}

	assert.NotPanics(t, func() {
		removeUnspecifiedNodePoolsOverrides(tfNodePools, model)
	})

	require.Len(t, poolA.Spec.Overrides, 1)
	assert.Equal(t, "keyA", poolA.Spec.Overrides[0].Name)
}

func TestRemoveUnspecifiedNodePoolsOverrides_NilSafety(t *testing.T) {
	assert.NotPanics(t, func() {
		removeUnspecifiedNodePoolsOverrides(nil, nil)
	})

	model := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
		Spec: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterSpec{
			Topology: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTopology{
				NodePools: []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{},
			},
		},
	}

	assert.NotPanics(t, func() {
		removeUnspecifiedNodePoolsOverrides([]interface{}{
			"invalid-format",
			map[string]interface{}{NameKey: ""},
			map[string]interface{}{NameKey: "pool-no-spec"},
			map[string]interface{}{NameKey: "pool-empty-spec", SpecKey: []interface{}{}},
			map[string]interface{}{NameKey: "pool-bad-spec", SpecKey: []interface{}{"not-a-map"}},
			map[string]interface{}{NameKey: "pool-empty-overrides", SpecKey: []interface{}{map[string]interface{}{OverridesKey: ""}}},
		}, model)
	})
}

func TestRemoveUnspecifiedControlPlaneOverrides_NilSafety(t *testing.T) {
	assert.NotPanics(t, func() {
		removeUnspecifiedControlPlaneOverrides(nil, nil)
	})

	model := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster{
		Spec: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterSpec{
			Topology: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTopology{
				ControlPlane: &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterControlPlane{
					Overrides: []*tkccommonmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterClusterVariable{
						{Name: "cpKey", Value: "cpVal"},
						{Name: "dropKey", Value: "dropVal"},
					},
				},
			},
		},
	}

	assert.NotPanics(t, func() {
		removeUnspecifiedControlPlaneOverrides("invalid", model)
	})

	// Valid control plane overrides
	cpData := map[string]interface{}{
		OverridesKey: `{"cpKey": "cpVal"}`,
	}
	removeUnspecifiedControlPlaneOverrides(cpData, model)
	require.Len(t, model.Spec.Topology.ControlPlane.Overrides, 1)
	assert.Equal(t, "cpKey", model.Spec.Topology.ControlPlane.Overrides[0].Name)
}

type mockClusterClassClientService struct {
	clusterclassclient.ClientService
	getFunc func(fn *clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassFullName) (*clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassListData, error)
}

func (m *mockClusterClassClientService) ClusterClassResourceServiceGet(fn *clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassFullName) (*clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassListData, error) {
	if m.getFunc != nil {
		return m.getFunc(fn)
	}

	return nil, nil
}

func TestValidateSchema_MissingSpec_NoPanic(t *testing.T) {
	r := ResourceTanzuKubernetesCluster()
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
	})
	auth := authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			ClusterClassResourceService: &mockClusterClassClientService{},
		},
	}
	diff, err := r.Diff(context.Background(), nil, cfg, auth)
	assert.NoError(t, err)
	assert.NotNil(t, diff)
}

func TestValidateSchema_EmptySpec_NoPanic(t *testing.T) {
	r := ResourceTanzuKubernetesCluster()
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
		SpecKey:                  []interface{}{},
	})
	auth := authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			ClusterClassResourceService: &mockClusterClassClientService{},
		},
	}
	diff, err := r.Diff(context.Background(), nil, cfg, auth)
	assert.NoError(t, err)
	assert.NotNil(t, diff)
}

func TestValidateSchema_SpecWithoutTopology_NoPanic(t *testing.T) {
	r := ResourceTanzuKubernetesCluster()
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
		SpecKey: []interface{}{
			map[string]interface{}{
				ClusterGroupNameKey: "default",
			},
		},
	})
	auth := authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			ClusterClassResourceService: &mockClusterClassClientService{},
		},
	}
	diff, err := r.Diff(context.Background(), nil, cfg, auth)
	assert.NoError(t, err)
	assert.NotNil(t, diff)
}

func TestValidateSchema_TopologyMissingClusterClass_ReturnsError(t *testing.T) {
	r := ResourceTanzuKubernetesCluster()
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
		SpecKey: []interface{}{
			map[string]interface{}{
				TopologyKey: []interface{}{
					map[string]interface{}{
						VersionKey:          testK8sVersion,
						ClusterVariablesKey: "{}",
					},
				},
			},
		},
	})
	auth := authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			ClusterClassResourceService: &mockClusterClassClientService{},
		},
	}
	_, err := r.Diff(context.Background(), nil, cfg, auth)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Cluster Class must be set during creation")
}

func TestValidateSchema_NilAuthContext_NoPanic(t *testing.T) {
	r := ResourceTanzuKubernetesCluster()
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
	})
	diff, err := r.Diff(context.Background(), nil, cfg, nil)
	assert.NoError(t, err)
	assert.NotNil(t, diff)
}

func TestValidateSchema_ValidClusterClassSuccess(t *testing.T) {
	r := ResourceTanzuKubernetesCluster()
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
		SpecKey: []interface{}{
			map[string]interface{}{
				TopologyKey: []interface{}{
					map[string]interface{}{
						VersionKey:          testK8sVersion,
						ClusterClassKey:     testClusterClass,
						ClusterVariablesKey: "{}",
					},
				},
			},
		},
	})
	mockCC := &mockClusterClassClientService{
		getFunc: func(fn *clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassFullName) (*clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassListData, error) {
			return &clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassListData{
				ClusterClasses: []*clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClass{
					{
						Spec: &clusterclassmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassSpec{},
					},
				},
			}, nil
		},
	}
	auth := authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			ClusterClassResourceService: mockCC,
		},
	}
	diff, err := r.Diff(context.Background(), nil, cfg, auth)
	assert.NoError(t, err)
	assert.NotNil(t, diff)
}

func TestValidateSchema_ClusterClassGetError(t *testing.T) {
	r := ResourceTanzuKubernetesCluster()
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
		SpecKey: []interface{}{
			map[string]interface{}{
				TopologyKey: []interface{}{
					map[string]interface{}{
						VersionKey:      testK8sVersion,
						ClusterClassKey: testClusterClass,
					},
				},
			},
		},
	})
	mockCC := &mockClusterClassClientService{
		getFunc: func(fn *clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassFullName) (*clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassListData, error) {
			return nil, errors.New("network error fetching cluster class")
		},
	}
	auth := authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			ClusterClassResourceService: mockCC,
		},
	}
	_, err := r.Diff(context.Background(), nil, cfg, auth)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Couldn't find cluster class")
}

func TestValidateSchema_ClusterClassesEmpty(t *testing.T) {
	r := ResourceTanzuKubernetesCluster()
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
		SpecKey: []interface{}{
			map[string]interface{}{
				TopologyKey: []interface{}{
					map[string]interface{}{
						VersionKey:      testK8sVersion,
						ClusterClassKey: testClusterClass,
					},
				},
			},
		},
	})
	mockCC := &mockClusterClassClientService{
		getFunc: func(fn *clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassFullName) (*clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassListData, error) {
			return &clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassListData{
				ClusterClasses: []*clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClass{},
			}, nil
		},
	}
	auth := authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			ClusterClassResourceService: mockCC,
		},
	}
	_, err := r.Diff(context.Background(), nil, cfg, auth)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Couldn't find cluster class")
}

func testClusterResourceData(t *testing.T) *schema.ResourceData {
	rawMap := map[string]interface{}{
		NameKey:                  testClusterName,
		ManagementClusterNameKey: testManagementCluster,
		ProvisionerNameKey:       testProvisionerName,
		SpecKey: []interface{}{
			map[string]interface{}{
				ClusterGroupNameKey: "default",
				TopologyKey: []interface{}{
					map[string]interface{}{
						ClusterClassKey: "tkg-cluster-class",
						VersionKey:      testK8sVersion,
						ControlPlaneKey: []interface{}{
							map[string]interface{}{
								OverridesKey: "{}",
							},
						},
						NodePoolKey: []interface{}{
							map[string]interface{}{
								NameKey: testWorkerPool1,
								SpecKey: []interface{}{
									map[string]interface{}{
										WorkerClassKey: "worker-class",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, tanzuKubernetesClusterSchema, rawMap)
	d.SetId(testManagementCluster + "/" + testProvisionerName + "/" + testClusterName)

	return d
}

func TestResourceTanzuKubernetesClusterRead_NotFound_Drift_ClearsStateID(t *testing.T) {
	d := testClusterResourceData(t)

	mockClient := &mockNodePoolClientService{
		clusterGetFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			return nil, clienterrors.ErrorWithHTTPCode(404, errors.New("status : 404 Not Found"))
		},
	}

	auth := authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	diags := resourceTanzuKubernetesClusterRead(context.Background(), d, auth)
	assert.False(t, diags.HasError())
	assert.Empty(t, d.Id())
}

func TestResourceTanzuKubernetesClusterRead_NotFound_DeleteState_ClearsStateID(t *testing.T) {
	d := testClusterResourceData(t)

	mockClient := &mockNodePoolClientService{
		clusterGetFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			return nil, clienterrors.ErrorWithHTTPCode(404, errors.New("status : 404 Not Found"))
		},
	}

	auth := authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	ctx := helper.GetContextWithCaller(context.Background(), helper.DeleteState)
	diags := resourceTanzuKubernetesClusterRead(ctx, d, auth)
	assert.False(t, diags.HasError())
	assert.Empty(t, d.Id())
}

func TestResourceTanzuKubernetesClusterRead_NotFound_CreateState_ReturnsError(t *testing.T) {
	d := testClusterResourceData(t)

	mockClient := &mockNodePoolClientService{
		clusterGetFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			return nil, clienterrors.ErrorWithHTTPCode(404, errors.New("status : 404 Not Found"))
		},
	}

	auth := authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	ctx := helper.GetContextWithCaller(context.Background(), helper.CreateState)
	diags := resourceTanzuKubernetesClusterRead(ctx, d, auth)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "Couldn't read TKG cluster")
}

func TestResourceTanzuKubernetesClusterRead_Non404Error_ReturnsError(t *testing.T) {
	d := testClusterResourceData(t)

	mockClient := &mockNodePoolClientService{
		clusterGetFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, error) {
			return nil, errors.New("500 Internal Server Error")
		},
	}

	auth := authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	diags := resourceTanzuKubernetesClusterRead(context.Background(), d, auth)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "Couldn't read TKG cluster")
}

type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}
