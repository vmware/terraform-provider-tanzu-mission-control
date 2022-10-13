/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroup

import (
	"testing"

	"github.com/stretchr/testify/require"

	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
)

func TestFlattenClusterGroupFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster group full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group full name",
			input: &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
				Name: "default",
			},
			expected: []interface{}{
				map[string]interface{}{
					NameKey: "default",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenClusterGroupFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
