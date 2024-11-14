// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	"testing"

	"github.com/stretchr/testify/require"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
)

func TestFlattenClusterFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster full name",
			input: &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{
				Name:                  "dummy",
				ManagementClusterName: "attached",
				ProvisionerName:       "attached",
			},
			expected: []interface{}{
				map[string]interface{}{
					NameKey:                  "dummy",
					ManagementClusterNameKey: "attached",
					ProvisionerNameKey:       "attached",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenClusterFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
