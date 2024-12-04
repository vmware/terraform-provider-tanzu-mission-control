// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policykindquota

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyrecipequotamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/quota"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/quota/recipe"
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
				Input: constructCustomInput(),
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
				Recipe: string(CustomRecipe),
			},
			expected: []interface{}{
				map[string]interface{}{
					policy.InputKey: []interface{}{
						map[string]interface{}{
							reciperesource.CustomKey: []interface{}{
								map[string]interface{}{
									reciperesource.LimitsCPUKey:                      "4",
									reciperesource.LimitsMemoryKey:                   "8Mi",
									reciperesource.PersistentVolumeClaimsKey:         int64(2),
									reciperesource.PersistentVolumeClaimsPerClassKey: map[string]int{"test-1": 2},
									reciperesource.RequestsCPUKey:                    "2",
									reciperesource.RequestsMemoryKey:                 "4Mi",
									reciperesource.RequestsStorageKey:                "2G",
									reciperesource.RequestsStoragePerClassKey:        map[string]string{"test-2": "2G"},
									reciperesource.ResourceCountsKey:                 map[string]int{"pods": 2},
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

func constructCustomInput() (customRecipeInput map[string]interface{}) {
	customInput := policyrecipequotamodel.VmwareTanzuManageV1alpha1CommonPolicySpecQuotaV1Custom{
		LimitsCPU:                      "4",
		LimitsMemory:                   "8Mi",
		Persistentvolumeclaims:         2,
		PersistentvolumeclaimsPerClass: map[string]int{"test-1": 2},
		RequestsCPU:                    "2",
		RequestsMemory:                 "4Mi",
		RequestsStorage:                "2G",
		RequestsStoragePerClass:        map[string]string{"test-2": "2G"},
		ResourceCounts:                 map[string]int{"pods": 2},
	}

	binary, err := customInput.MarshalBinary()
	if err != nil {
		return nil
	}

	err = json.Unmarshal(binary, &customRecipeInput)
	if err != nil {
		return nil
	}

	return customRecipeInput
}
