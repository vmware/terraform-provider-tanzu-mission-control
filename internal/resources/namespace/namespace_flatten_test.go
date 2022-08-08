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

func TestFlattenSpec(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceSpec
		expected []interface{}
	}{
		{
			name:     "check for nil data in namespace spec",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with attach set to false",
			input: &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceSpec{
				WorkspaceName: "default",
			},
			expected: []interface{}{
				map[string]interface{}{
					workspaceNameKey: "default",
					attachKey:        false,
				},
			},
		},
		{
			name: "normal scenario with attach set to true",
			input: &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceSpec{
				WorkspaceName: "workspace_name",
				Attach:        true,
			},
			expected: []interface{}{
				map[string]interface{}{
					workspaceNameKey: "workspace_name",
					attachKey:        true,
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
