/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"testing"

	"github.com/stretchr/testify/require"

	iammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

func TestFlattenRoleBinding(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding
		expected []interface{}
	}{
		{
			name:     "check for nil role binding data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with all fields of role binding data",
			input: &iammodel.VmwareTanzuCoreV1alpha1PolicyRoleBinding{
				Role: "cluster-group.admin",
				Subjects: []*iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
					{
						Name: "test-new",
						Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindGROUP.Pointer(),
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					roleKey: "cluster-group.admin",
					subjectsKey: []interface{}{
						map[string]interface{}{
							subjectNameKey: "test-new",
							subjectKindKey: "GROUP",
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := flattenRoleBinding(test.input)
			require.EqualValues(t, test.expected, actual)
		})
	}
}
