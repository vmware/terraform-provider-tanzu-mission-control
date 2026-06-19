// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

// nolint: dupl
package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
)

const (
	customRecipeTemplateName = "some-custom-template"
	test127001               = "127.0.0.1"
	testKey1                 = "key-1"
	testNodes                = "nodes"
	testPod                  = "pod"
	testPolicy               = "policy"
	testTest                 = "test"
	testUser                 = "User"
	testValue1               = "value-1"
)

func TestFlattenTMCBlockRolebindingSubjects(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjects
		expected    []interface{}
	}{
		{
			description: "check for nil custom policy tmc_block_rolebinding_subjects recipe",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete custom policy tmc_block_rolebinding_subjects recipe",
			input: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjects{
				Audit: true,
				Parameters: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParameters{
					DisallowedSubjects: []*policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParametersDisallowedSubjects{
						{
							Kind: testUser,
							Name: testTest,
						},
					},
				},
				TargetKubernetesResources: []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources{
					{
						APIGroups: []string{testPolicy},
						Kinds:     []string{testPod},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					AuditKey: true,
					ParametersKey: []interface{}{
						map[string]interface{}{
							disallowedSubjectsKey: []interface{}{
								map[string]interface{}{
									kindKey: testUser,
									nameKey: testTest,
								},
							},
						},
					},
					TargetKubernetesResourcesKey: []interface{}{
						map[string]interface{}{
							APIGroupsKey: []string{testPolicy},
							KindsKey:     []string{testPod},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenTMCBlockRolebindingSubjects(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenBlockRoleBindingParameters(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParameters
		expected    []interface{}
	}{
		{
			description: "check for nil custom policy tmc_block_rolebinding_subjects parameters",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete custom policy tmc_block_rolebinding_subjects parameters",
			input: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParameters{
				DisallowedSubjects: []*policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParametersDisallowedSubjects{
					{
						Kind: testNodes,
						Name: testTest,
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					disallowedSubjectsKey: []interface{}{
						map[string]interface{}{
							kindKey: testNodes,
							nameKey: testTest,
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenBlockRoleBindingParameters(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenDisallowedSubjects(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParametersDisallowedSubjects
		expected    interface{}
	}{
		{
			description: "check for nil custom policy tmc_block_rolebinding_subjects parameters disallowed subjects",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete custom policy tmc_block_rolebinding_subjects parameters disallowed subjects",
			input: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParametersDisallowedSubjects{
				Kind: testNodes,
				Name: testTest,
			},
			expected: map[string]interface{}{
				kindKey: testNodes,
				nameKey: testTest,
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenDisallowedSubjects(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
