/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"testing"

	"github.com/stretchr/testify/require"

	clustermodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster"
)

func TestFlattenSpec(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *clustermodel.VmwareTanzuManageV1alpha1ClusterSpec
		expected []interface{}
	}{
		{
			name:     "check for nil data in cluster spec",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with cluster group",
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
			name: "normal scenario with cluster group and proxy",
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
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := flattenSpec(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
