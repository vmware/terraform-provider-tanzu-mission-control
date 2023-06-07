/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindnetwork

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network/recipe"
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
				Input: constructAllowAll(),
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
				Recipe: string(AllowAllRecipe),
			},
			expected: []interface{}{
				map[string]interface{}{
					policy.InputKey: []interface{}{
						map[string]interface{}{
							reciperesource.AllowAllKey: []interface{}{
								map[string]interface{}{
									reciperesource.FromOwnNamespaceKey: true,
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

func constructAllowAll() (allowedAllRecipeInput map[string]interface{}) {
	allowAllInput := policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAll{
		FromOwnNamespace: helper.BoolPointer(true),
	}

	binary, err := allowAllInput.MarshalBinary()
	if err != nil {
		return nil
	}

	err = json.Unmarshal(binary, &allowedAllRecipeInput)
	if err != nil {
		return nil
	}

	return allowedAllRecipeInput
}
