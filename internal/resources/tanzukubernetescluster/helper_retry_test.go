// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzukubernetescluster

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIsRetryableNodePoolError verifies that transient admission webhook rejections,
// race conditions, optimistic locking conflicts, network errors, rate limits, and 5xx errors
// are classified as retryable, while permanent client errors (400, 401) are not.
func TestIsRetryableNodePoolError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "capi validating webhook SystemChecksSucceeded condition not true",
			err:      errors.New("admission webhook \"capi.validating.kubernetes.vmware.com\" denied the request: update cannot be initiated as cluster's SystemChecksSucceeded condition is not True. Message: . Reason: UnknownReason"),
			expected: true,
		},
		{
			name:     "generic update cannot be initiated error",
			err:      errors.New("update cannot be initiated at this stage"),
			expected: true,
		},
		{
			name:     "cluster initializing duplicate labels race condition",
			err:      errors.New("race condition when cluster is being initialized and duplicate nodePoolLabels"),
			expected: true,
		},
		{
			name:     "optimistic locking / object modified conflict",
			err:      errors.New("Operation cannot be fulfilled on clusters.cluster.x-k8s.io: the object has been modified; please apply your changes to the latest version"),
			expected: true,
		},
		{
			name:     "resourceversion conflict",
			err:      errors.New("conflict on resourceversion 12345"),
			expected: true,
		},
		{
			name:     testErrConnectionReset,
			err:      errors.New("read: " + testErrConnectionReset),
			expected: true,
		},
		{
			name:     "EOF error",
			err:      errors.New("unexpected EOF"),
			expected: true,
		},
		{
			name:     "service unavailable error",
			err:      errors.New("rpc error: code = Unavailable desc = transport is closing"),
			expected: true,
		},
		{
			name:     "HTTP 500 error",
			err:      errors.New("POST request failed with status : 500 Internal Server Error"),
			expected: true,
		},
		{
			name:     "HTTP 503 error",
			err:      errors.New("POST request failed with status : 503 Service Unavailable"),
			expected: true,
		},
		{
			name:     "HTTP 429 Too Many Requests rate limiting error",
			err:      errors.New("POST request failed with status : 429 Too Many Requests, response: {\"error\":\"/vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.nodepool.NodepoolResourceService/Create is rejected due to rate limiting, please retry later.\",\"code\":8}"),
			expected: true,
		},
		{
			name:     "gRPC ResourceExhausted rate limit error",
			err:      errors.New("rpc error: code = ResourceExhausted desc = rate limit exceeded, retry later"),
			expected: true,
		},
		{
			name:     "permanent client error: invalid storage class",
			err:      errors.New("POST request failed with status : 400 Bad Request, response: storage class 'invalid-sc' not found"),
			expected: false,
		},
		{
			name:     "permanent client error: unauthorized",
			err:      errors.New("POST request failed with status : 401 Unauthorized"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isRetryableNodePoolError(tt.err)
			assert.Equal(t, tt.expected, got, "mismatch for test case: %s", tt.name)
		})
	}
}

// TestIsRetryableClusterError verifies that transient rate limiting, conflict during teardown,
// transport drops, and 5xx errors are classified as retryable for cluster creation.
func TestIsRetryableClusterError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "HTTP 409 Conflict already exists error",
			err:      errors.New("POST request failed with status : 409 Conflict, response: {\"error\":\"clusters.cluster.x-k8s.io \\\"terraform-wc-photon-03\\\" already exists\"}"),
			expected: true,
		},
		{
			name:     "HTTP 429 Too Many Requests rate limiting error",
			err:      errors.New("POST request failed with status : 429 Too Many Requests, response: {\"error\":\"rejected due to rate limiting\",\"code\":8}"),
			expected: true,
		},
		{
			name:     testErrConnectionReset,
			err:      errors.New("read: " + testErrConnectionReset),
			expected: true,
		},
		{
			name:     "HTTP 503 Service Unavailable",
			err:      errors.New("POST request failed with status : 503 Service Unavailable"),
			expected: true,
		},
		{
			name:     "permanent client error: 400 Bad Request",
			err:      errors.New("POST request failed with status : 400 Bad Request"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isRetryableClusterError(tt.err)
			assert.Equal(t, tt.expected, got, "mismatch for test case: %s", tt.name)
		})
	}
}
