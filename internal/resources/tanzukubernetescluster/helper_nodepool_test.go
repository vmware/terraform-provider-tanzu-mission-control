// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzukubernetescluster

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client"
	tanzukubernetesclusterclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/tanzukubernetescluster"
	tkcmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
	tkcnodepoolmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/nodepool"
)

const (
	testClusterName        = "test-cluster"
	testManagementCluster  = "mgmt"
	testProvisionerName    = "prov"
	testWorkerPool1        = "worker-pool-1"
	testWorkerPoolInit     = "worker-pool-init"
	testErrConnectionReset = "connection reset by peer"
)

type mockNodePoolClientService struct {
	tanzukubernetesclusterclient.ClientService
	listFunc   func(clusterFn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error)
	createFunc func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error)
	updateFunc func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error)
	deleteFunc func(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) error
	getFunc    func(fn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error)
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

// TestWaitNodePoolsReady_SuccessReady tests successful polling transition from CREATING to READY.
func TestWaitNodePoolsReady_SuccessReady(t *testing.T) {
	origInterval := nodePoolPollInterval
	nodePoolPollInterval = 10 * time.Millisecond
	defer func() { nodePoolPollInterval = origInterval }()

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
	defer func() { nodePoolPollInterval = origInterval }()

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
	assert.Contains(t, err.Error(), "Node pool 'worker-pool-lost' is in status: not found in cluster")
}

// TestWaitNodePoolsReady_NodePoolError tests fast-fail behavior when a node pool enters ERROR phase.
func TestWaitNodePoolsReady_NodePoolError(t *testing.T) {
	origInterval := nodePoolPollInterval
	nodePoolPollInterval = 10 * time.Millisecond
	defer func() { nodePoolPollInterval = origInterval }()

	clusterFn := &tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
		ManagementClusterName: testManagementCluster,
		ProvisionerName:       testProvisionerName,
		Name:                  testClusterName,
	}

	expectedNps := []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
		{
			FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				Name: "worker-pool-err",
			},
		},
	}

	errPhase := tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseERROR
	mockClient := &mockNodePoolClientService{
		listFunc: func(fn *tkcmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData, error) {
			return &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolListNodepoolsData{
				Nodepools: []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool{
					{
						FullName: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
							Name: "worker-pool-err",
						},
						Status: &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatus{
							Phase: &errPhase,
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
	assert.Contains(t, err.Error(), "TKG Cluster node pool errored")
}

// TestWaitNodePoolsReady_NilStatusSafety tests safe handling of nil Status or nil Phase fields during cluster boot.
func TestWaitNodePoolsReady_NilStatusSafety(t *testing.T) {
	origInterval := nodePoolPollInterval
	nodePoolPollInterval = 10 * time.Millisecond
	defer func() { nodePoolPollInterval = origInterval }()

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

// TestCreateNodePoolWithRetry_RetryOnTransientAndSucceed tests recovery after webhook rejection errors.
func TestCreateNodePoolWithRetry_RetryOnTransientAndSucceed(t *testing.T) {
	origInit, origMax := nodePoolInitialBackoff, nodePoolMaxBackoff
	nodePoolInitialBackoff = 5 * time.Millisecond
	nodePoolMaxBackoff = 20 * time.Millisecond
	defer func() {
		nodePoolInitialBackoff = origInit
		nodePoolMaxBackoff = origMax
	}()

	callCount := 0
	mockClient := &mockNodePoolClientService{
		createFunc: func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
			callCount++
			if callCount < 3 {
				// Simulating CAPI validating webhook SystemChecksSucceeded condition rejection
				return nil, assert.AnError
			}
			return req, nil
		},
	}

	config := &authctx.TanzuContext{
		TMCConnection: &client.TanzuMissionControl{
			TanzuKubernetesClusterResourceService: mockClient,
		},
	}

	// Make sure the error matches isRetryableNodePoolError
	transientErr := &struct {
		mockNodePoolClientService
	}{}
	_ = transientErr

	mockClient.createFunc = func(req *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData, error) {
		callCount++
		if callCount < 3 {
			return nil, &testError{msg: "admission webhook \"capi.validating.kubernetes.vmware.com\" denied the request: update cannot be initiated as cluster's SystemChecksSucceeded condition is not True"}
		}
		return req, nil
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
	assert.Contains(t, err.Error(), "timed out waiting for nodepool 'worker-pool-timeout' creation to be accepted after")
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
	origInit, origMax := nodePoolInitialBackoff, nodePoolMaxBackoff
	nodePoolInitialBackoff = 5 * time.Millisecond
	nodePoolMaxBackoff = 20 * time.Millisecond
	defer func() {
		nodePoolInitialBackoff = origInit
		nodePoolMaxBackoff = origMax
	}()

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
	assert.NoError(t, err)
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

type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}
