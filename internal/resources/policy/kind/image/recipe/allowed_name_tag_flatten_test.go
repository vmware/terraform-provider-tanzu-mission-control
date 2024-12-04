// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipeimagemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image"
	policyrecipeimagecommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image/common"
)

func TestFlattenAllowedNameTag(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag
		expected    []interface{}
	}{
		{
			description: "check for nil image policy allowed-name-tag recipe",
			input:       nil,
			expected:    nil,
		},
		{
			description: "scenario with nil value of Audit value in allowed-name-tag recipe",
			input: &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag{
				Audit: nil,
				Rules: []*policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules{
					{
						ImageName: "foo",
						Tag: &policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag{
							Negate: helper.BoolPointer(false),
							Value:  "test",
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					RulesKey: []interface{}{
						map[string]interface{}{
							ImageNameKey: "foo",
							TagKey: []interface{}{
								map[string]interface{}{
									NegateKey: false,
									ValueKey:  "test",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "normal scenario with complete image policy allowed-name-tag recipe",
			input: &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag{
				Audit: helper.BoolPointer(true),
				Rules: []*policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules{
					{
						ImageName: "foo",
						Tag: &policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag{
							Negate: helper.BoolPointer(false),
							Value:  "test",
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					AuditKey: true,
					RulesKey: []interface{}{
						map[string]interface{}{
							ImageNameKey: "foo",
							TagKey: []interface{}{
								map[string]interface{}{
									NegateKey: false,
									ValueKey:  "test",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenAllowedNameTag(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenNameTagRules(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules
		expected    interface{}
	}{
		{
			description: "check for nil rules value of allowed-name-tag recipe",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with all fields of rules spec of allowed-name-tag recipe",
			input: &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules{
				ImageName: "foo",
				Tag: &policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag{
					Negate: helper.BoolPointer(false),
					Value:  "test",
				},
			},
			expected: map[string]interface{}{
				ImageNameKey: "foo",
				TagKey: []interface{}{
					map[string]interface{}{
						NegateKey: false,
						ValueKey:  "test",
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenNameTagRules(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
