/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindcustom

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom/recipe"
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
				Input: constructTMCHTTPSIngressInput(),
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
				Recipe: string(tmcHTTPSIngressRecipe),
			},
			expected: []interface{}{
				map[string]interface{}{
					policy.InputKey: []interface{}{
						map[string]interface{}{
							reciperesource.TMCHTTPSIngressKey: []interface{}{
								map[string]interface{}{
									reciperesource.AuditKey: true,
									reciperesource.TargetKubernetesResourcesKey: []interface{}{
										map[string]interface{}{
											reciperesource.APIGroupsKey: []string{"policy"},
											reciperesource.KindsKey:     []string{"pod"},
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

func constructTMCHTTPSIngressInput() (tmcHTTPSIngressRecipeInput map[string]interface{}) {
	tmcHTTPSIngressInput := policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe{
		Audit: true,
		TargetKubernetesResources: []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources{
			{
				APIGroups: []string{"policy"},
				Kinds:     []string{"pod"},
			},
		},
	}

	binary, err := tmcHTTPSIngressInput.MarshalBinary()
	if err != nil {
		return nil
	}

	err = json.Unmarshal(binary, &tmcHTTPSIngressRecipeInput)
	if err != nil {
		return nil
	}

	return tmcHTTPSIngressRecipeInput
}
