// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
	policyrecipemutationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation"
	policyrecipemutationcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation/common"
)

const (
	testAlways              = "Always"
	testCluster             = "Cluster"
	testIffielddoesnotexist = "IfFieldDoesNotExist"
	testIffieldexists       = "IfFieldExists"
	testKeyValue            = "key_value"
	testPod                 = "pod"
	testPolicy              = "policy"
	testPrune               = "prune"
	testValueValue          = "value_value"
)

func TestFlattenAnnotation(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation
		expected    []interface{}
	}{
		{
			description: "check for nil mutation annotation",
		},
		{
			description: "flatten normal annotation mutation policy struct",
			input: &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation{
				Annotation: &policyrecipemutationcommonmodel.KeyValue{
					Key:   testKeyValue,
					Value: testValueValue},
				Scope: policyrecipemutationcommonmodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope(policyrecipemutationcommonmodel.Cluster),
				TargetKubernetesResources: []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources{
					{
						APIGroups: []string{testPolicy},
						Kinds:     []string{testPod},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					AnnotationKey: []interface{}{
						map[string]interface{}{
							keyKey:   testKeyValue,
							valueKey: testValueValue,
						},
					},
					scopeKey: testCluster,
					targetKubernetesResourcesKey: []interface{}{
						map[string]interface{}{
							apiGroupsKey: []string{testPolicy},
							kindsKey:     []string{testPod},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenAnnotation(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
