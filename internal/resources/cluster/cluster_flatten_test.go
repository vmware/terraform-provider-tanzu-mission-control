/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"testing"

	"github.com/stretchr/testify/require"

	clustermodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/cluster"
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
			name: "normal scenario with cluster group and proxy",
			input: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				ProxyName:        "test",
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey: "default",
					proxyKey:        "test",
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
