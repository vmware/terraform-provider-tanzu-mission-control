// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package iampolicy

import (
	"testing"

	"github.com/stretchr/testify/require"

	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

func TestFlattenRoleBinding(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding
		expected    interface{}
	}{
		{
			description: "check for nil role binding data",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with all fields of role binding data",
			input: &iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding{
				Role: testClustergroupadmin,
				Subjects: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
					{
						Name: subject1Name,
						Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindGROUP.Pointer(),
					},
				},
			},
			expected: map[string]interface{}{
				roleKey: testClustergroupadmin,
				subjectsKey: []interface{}{
					map[string]interface{}{
						subjectNameKey: subject1Name,
						subjectKindKey: subject1Kind,
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenRoleBinding(test.input)
			require.EqualValues(t, test.expected, actual)
		})
	}
}
