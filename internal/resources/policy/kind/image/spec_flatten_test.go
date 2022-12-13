/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindimage

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyrecipeimagemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image"
	policyrecipeimagecommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/image/recipe"
)

func TestFlattenSpec(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec
		expected    []interface{}
	}{
		{
			description: "check for nil spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete spec",
			input: &policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec{
				Input: constructAllowedNameTag(),
				NamespaceSelector: &policymodel.VmwareTanzuManageV1alpha1CommonPolicyLabelSelector{
					MatchExpressions: []*policymodel.K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement{
						{
							Key:      "k1",
							Operator: "In",
							Values: []string{
								"v1",
								"v2",
							},
						},
						{
							Key:      "k2",
							Operator: "Exists",
							Values:   []string{},
						},
					},
				},
				Recipe: string(AllowedNameTagRecipe),
			},
			expected: []interface{}{
				map[string]interface{}{
					policy.InputKey: []interface{}{
						map[string]interface{}{
							reciperesource.AllowedNameTagKey: []interface{}{
								map[string]interface{}{
									reciperesource.AuditKey: true,
									reciperesource.RulesKey: []interface{}{
										map[string]interface{}{
											reciperesource.ImageNameKey: "foo",
											reciperesource.TagKey: []interface{}{
												map[string]interface{}{
													reciperesource.NegateKey: true,
													reciperesource.ValueKey:  "test",
												},
											},
										},
									},
								},
							},
						},
					},
					policy.NamespaceSelectorKey: []interface{}{
						map[string]interface{}{
							policy.MatchExpressionsKey: []interface{}{
								map[string]interface{}{
									policy.KeyKey:      "k1",
									policy.OperatorKey: "In",
									policy.ValuesKey: []string{
										"v1",
										"v2",
									},
								},
								map[string]interface{}{
									policy.KeyKey:      "k2",
									policy.OperatorKey: "Exists",
									policy.ValuesKey:   []string{},
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
			actual := FlattenSpec(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func constructAllowedNameTag() (allowedNameTagRecipeInput map[string]interface{}) {
	allowedNameTagInput := policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag{
		Audit: helper.BoolPointer(true),
		Rules: []*policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules{
			{
				ImageName: "foo",
				Tag: &policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag{
					Negate: helper.BoolPointer(true),
					Value:  "test",
				},
			},
		},
	}

	binary, err := allowedNameTagInput.MarshalBinary()
	if err != nil {
		return nil
	}

	err = json.Unmarshal(binary, &allowedNameTagRecipeInput)
	if err != nil {
		return nil
	}

	return allowedNameTagRecipeInput
}
