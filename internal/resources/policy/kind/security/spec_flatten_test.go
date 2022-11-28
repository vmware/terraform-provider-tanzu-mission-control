/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindsecurity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyrecipesecuritymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/security"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/security/recipe"
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
				Input: constructBaselineRecipeInput(),
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
				Recipe: string(BaselineRecipe),
			},
			expected: []interface{}{
				map[string]interface{}{
					policy.InputKey: []interface{}{
						map[string]interface{}{
							reciperesource.BaselineKey: []interface{}{
								map[string]interface{}{
									reciperesource.AuditKey:            true,
									reciperesource.DisableNativePspKey: false,
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

func constructBaselineRecipeInput() (baselineRecipeInput map[string]interface{}) {
	baselineInput := policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Baseline{
		Audit:            true,
		DisableNativePsp: false,
	}

	binary, err := baselineInput.MarshalBinary()
	if err != nil {
		return nil
	}

	err = json.Unmarshal(binary, &baselineRecipeInput)
	if err != nil {
		return nil
	}

	return baselineRecipeInput
}
