/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindquota

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipequotamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/quota"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/quota/recipe"
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
				recipe: CustomRecipe,
				input: &policyrecipequotamodel.VmwareTanzuManageV1alpha1CommonPolicySpecQuotaV1Custom{
					LimitsCPU:                      "4",
					LimitsMemory:                   "8Mi",
					Persistentvolumeclaims:         2,
					PersistentvolumeclaimsPerClass: map[string]int{"test-1": 2},
					RequestsCPU:                    "2",
					RequestsMemory:                 "4Mi",
					RequestsStorage:                "2G",
					RequestsStoragePerClass:        map[string]string{"test-2": "2G"},
					ResourceCounts:                 map[string]int{"pods": 2},
				},
			},

			expected: []interface{}{
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
