/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
)

func TestFlattenClusterPolicyFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster policy full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster policy full name",
			input: &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName{
				ClusterName:           "c",
				ManagementClusterName: "m",
				ProvisionerName:       "p",
			},
			expected: []interface{}{
				map[string]interface{}{
					ClusterNameKey:           "c",
					ManagementClusterNameKey: "m",
					ProvisionerNameKey:       "p",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenClusterPolicyFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
