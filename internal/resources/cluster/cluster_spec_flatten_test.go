/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"testing"

	"github.com/stretchr/testify/require"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
)

func TestFlattenSpec(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *clustermodel.VmwareTanzuManageV1alpha1ClusterSpec
		expected    []interface{}
	}{
		{
			description: "check for nil data in cluster spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with cluster group",
			input: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey: "default",
					proxyNameKey:    "",
				},
			},
		},
		{
			description: "normal scenario with cluster group and proxy",
			input: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				ProxyName:        "proxy",
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey: "default",
					proxyNameKey:    "proxy",
				},
			},
		},
		{
			description: "normal scenario with cluster group, proxy and image registry",
			input: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				ProxyName:        "proxy",
				ImageRegistry:    "image-registry",
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey:      "default",
					proxyNameKey:         "proxy",
					imageRegistryNameKey: "image-registry",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenSpec(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
