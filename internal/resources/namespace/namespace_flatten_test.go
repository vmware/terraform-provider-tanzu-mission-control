/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespace

import (
	"testing"

	"github.com/stretchr/testify/require"

	namespacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/namespace"
)

func TestFlattenNamespaceFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName
		expected    []interface{}
	}{
		{
			description: "check for nil namespace full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete namespace full name",
			input: &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName{
				Name:                  "n-1",
				ClusterName:           "dummy",
				ManagementClusterName: "attached",
				ProvisionerName:       "attached",
			},
			expected: []interface{}{
				map[string]interface{}{
					NameKey:                  "n-1",
					ClusterNameKey:           "dummy",
					ManagementClusterNameKey: "attached",
					ProvisionerNameKey:       "attached",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenNamespaceFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
