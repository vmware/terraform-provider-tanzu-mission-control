// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package iampolicy

import (
	"testing"

	"github.com/stretchr/testify/require"

	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

func TestFlattenRoleBindingList(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding
		expected    []interface{}
	}{
		{
			description: "check for nil role binding data list",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with empty role binding data list",
			input:       []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding{},
			expected:    []interface{}{},
		},
		{
			description: "normal scenario with all fields of role binding data list",
			input: []*iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding{
				{
					Role: testClustergroupadmin,
					Subjects: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
						{
							Name: subject1Name,
							Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindGROUP.Pointer(),
						},
					},
				},
				{
					Role: clusterRole,
					Subjects: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
						{
							Name: testTest2,
							Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindUSER.Pointer(),
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					roleKey: testClustergroupadmin,
					subjectsKey: []interface{}{
						map[string]interface{}{
							subjectNameKey: subject1Name,
							subjectKindKey: subject1Kind,
						},
					},
				},
				map[string]interface{}{
					roleKey: clusterRole,
					subjectsKey: []interface{}{
						map[string]interface{}{
							subjectNameKey: testTest2,
							subjectKindKey: "USER",
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenRoleBindingList(test.input)
			require.EqualValues(t, test.expected, actual)
		})
	}
}
