// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policykindquota

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipequotamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/quota"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/quota/recipe"
)

const (
	test4mi   = "4Mi"
	test8mi   = "8Mi"
	testPods  = "pods"
	testTest1 = "test-1"
	testTest2 = "test-2"
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
					LimitsMemory:                   test8mi,
					Persistentvolumeclaims:         2,
					PersistentvolumeclaimsPerClass: map[string]int{testTest1: 2},
					RequestsCPU:                    "2",
					RequestsMemory:                 test4mi,
					RequestsStorage:                "2G",
					RequestsStoragePerClass:        map[string]string{"test-2": "2G"},
					ResourceCounts:                 map[string]int{testPods: 2},
				},
			},

			expected: []interface{}{
				map[string]interface{}{
					reciperesource.CustomKey: []interface{}{
						map[string]interface{}{
							reciperesource.LimitsCPUKey:                      "4",
							reciperesource.LimitsMemoryKey:                   test8mi,
							reciperesource.PersistentVolumeClaimsKey:         int64(2),
							reciperesource.PersistentVolumeClaimsPerClassKey: map[string]int{testTest1: 2},
							reciperesource.RequestsCPUKey:                    "2",
							reciperesource.RequestsMemoryKey:                 test4mi,
							reciperesource.RequestsStorageKey:                "2G",
							reciperesource.RequestsStoragePerClassKey:        map[string]string{"test-2": "2G"},
							reciperesource.ResourceCountsKey:                 map[string]int{testPods: 2},
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
