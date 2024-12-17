// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipeimagecommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image/common"
)

func TestFlattenTag(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag
		expected    interface{}
	}{
		{
			description: "check for nil tag data",
			input:       nil,
			expected:    []interface{}(nil),
		},
		{
			description: "scenario for nil negate value of tag",
			input: &policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag{
				Negate: nil,
				Value:  "test",
			},
			expected: []interface{}{
				map[string]interface{}{
					ValueKey: "test",
				},
			},
		},
		{
			description: "normal scenario with all values of tag data",
			input: &policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag{
				Negate: helper.BoolPointer(false),
				Value:  "test",
			},
			expected: []interface{}{
				map[string]interface{}{
					NegateKey: false,
					ValueKey:  "test",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenTag(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
