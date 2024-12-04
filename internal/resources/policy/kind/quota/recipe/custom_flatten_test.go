// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipequotamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/quota"
)

func TestFlattenCustomRecipe(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipequotamodel.VmwareTanzuManageV1alpha1CommonPolicySpecQuotaV1Custom
		expected    []interface{}
	}{
		{
			description: "check for nil namespace quota policy custom recipe",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete namespace quota policy custom recipe",
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
			expected: []interface{}{
				map[string]interface{}{
					LimitsCPUKey:                      "4",
					LimitsMemoryKey:                   "8Mi",
					PersistentVolumeClaimsKey:         int64(2),
					PersistentVolumeClaimsPerClassKey: map[string]int{"test-1": 2},
					RequestsCPUKey:                    "2",
					RequestsMemoryKey:                 "4Mi",
					RequestsStorageKey:                "2G",
					RequestsStoragePerClassKey:        map[string]string{"test-2": "2G"},
					ResourceCountsKey:                 map[string]int{"pods": 2},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenCustom(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
