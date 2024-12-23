// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipesecuritymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/security"
)

func TestFlattenStrict(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict
		expected    []interface{}
	}{
		{
			description: "check for nil security policy strict recipe",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete security policy strict recipe",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict{
				Audit:            true,
				DisableNativePsp: false,
			},
			expected: []interface{}{
				map[string]interface{}{
					AuditKey:            true,
					DisableNativePspKey: false,
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenStrict(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
