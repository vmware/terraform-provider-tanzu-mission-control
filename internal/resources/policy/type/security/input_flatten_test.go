/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package security

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipesecuritymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/security"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/type/security/recipe"
)

func TestFlattenInput(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *inputRecipe
		expected    []interface{}
	}{
		{
			description: "check for nil input",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete input",
			input: &inputRecipe{
				recipe: baselineRecipe,
				inputBaseline: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Baseline{
					Audit:            true,
					DisableNativePsp: false,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					reciperesource.BaselineKey: []interface{}{
						map[string]interface{}{
							reciperesource.AuditKey:            true,
							reciperesource.DisableNativePspKey: false,
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenInput(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
