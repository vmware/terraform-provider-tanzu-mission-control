/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyworkspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/workspace"
)

func TestFlattenWorkspacePolicyFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName
		expected    []interface{}
	}{
		{
			description: "check for nil workspace policy full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete workspace policy full name",
			input: &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName{
				WorkspaceName: "w",
			},
			expected: []interface{}{
				map[string]interface{}{
					WorkspaceNameKey: "w",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenWorkspacePolicyFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
