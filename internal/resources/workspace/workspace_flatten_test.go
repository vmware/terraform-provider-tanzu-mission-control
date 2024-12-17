// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package workspace

import (
	"testing"

	"github.com/stretchr/testify/require"

	workspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/workspace"
)

func TestFlattenWorkspaceFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName
		expected    []interface{}
	}{
		{
			description: "check for nil workspace full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete workspace full name",
			input: &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
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
			actual := FlattenWorkspaceFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
