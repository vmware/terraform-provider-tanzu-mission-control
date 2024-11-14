// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"testing"

	"github.com/stretchr/testify/require"

	clustercommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/common"
)

func TestFlattenAdvancedConfig(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *clustercommon.VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig
		expected interface{}
	}{
		{
			name:     "check for nil advanced configs data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with all fields of advanced configs data",
			input: &clustercommon.VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig{
				Key:   "key-1",
				Value: "val-1",
			},
			expected: map[string]interface{}{
				advancedConfigurationKey:      "key-1",
				advancedConfigurationValueKey: "val-1",
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := FlattenAdvancedConfig(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
