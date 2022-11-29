/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipeimagemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image"
)

func TestFlattenCommonRecipe(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CommonRecipe
		expected    []interface{}
	}{
		{
			description: "check for nil image policy block-latest-tag or require-digest recipes",
			input:       nil,
			expected:    nil,
		},
		{
			description: "scenario with nil data for Audit",
			input: &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CommonRecipe{
				Audit: nil,
			},
			expected: nil,
		},
		{
			description: "normal scenario with complete image policy block-latest-tag or require-digest recipes",
			input: &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CommonRecipe{
				Audit: helper.BoolPointer(true),
			},
			expected: []interface{}{
				map[string]interface{}{
					AuditKey: true,
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenCommonRecipe(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
