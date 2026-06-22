// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

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
				ClusterGroupName: clusterGroupDefaultValue,
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey:      clusterGroupDefaultValue,
					proxyNameKey:         "",
					imageRegistryNameKey: "",
				},
			},
		},
		{
			description: "normal scenario with cluster group and proxy",
			input: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
				ProxyName:        proxyNameKey,
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey:      clusterGroupDefaultValue,
					proxyNameKey:         proxyNameKey,
					imageRegistryNameKey: "",
				},
			},
		},
		{
			description: "normal scenario with cluster group, proxy and image registry",
			input: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
				ProxyName:        proxyNameKey,
				ImageRegistry:    "image-registry",
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey:      clusterGroupDefaultValue,
					proxyNameKey:         proxyNameKey,
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
