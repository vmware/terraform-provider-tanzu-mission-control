// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
)

func TestFlattenClusterGroupPolicyFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster group policy full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group policy full name",
			input: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName{
				ClusterGroupName: "c",
			},
			expected: []interface{}{
				map[string]interface{}{
					ClusterGroupNameKey: "c",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenClusterGroupPolicyFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
