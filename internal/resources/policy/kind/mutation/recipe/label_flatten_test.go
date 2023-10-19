/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
	policyrecipemutationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation"
	policyrecipemutationcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation/common"
)

func TestFlattenLabel(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label
		expected    []interface{}
	}{
		{
			description: "check for nil mutation label",
		},
		{
			description: "flatten normal label mutation policy struct",
			input: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label{
				Label: &policyrecipemutationcommonmodel.KeyValue{
					Key:   "key_value",
					Value: "value_value"},
				Scope: policyrecipemutationcommonmodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope(policyrecipemutationcommonmodel.Cluster),
				TargetKubernetesResources: []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources{
					{
						APIGroups: []string{"policy"},
						Kinds:     []string{"pod"},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					LabelKey: []interface{}{
						map[string]interface{}{
							keyKey:   "key_value",
							valueKey: "value_value",
						},
					},
					scopeKey: "Cluster",
					targetKubernetesResourcesKey: []interface{}{
						map[string]interface{}{
							apiGroupsKey: []string{"policy"},
							kindsKey:     []string{"pod"},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenLabel(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
