// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policykindimage

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipeimagemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/image/recipe"
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
				recipe: BlockLatestTagRecipe,
				inputBlockLatestTag: &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CommonRecipe{
					Audit: helper.BoolPointer(true),
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					reciperesource.BlockLatestTagKey: []interface{}{
						map[string]interface{}{
							reciperesource.AuditKey: true,
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
