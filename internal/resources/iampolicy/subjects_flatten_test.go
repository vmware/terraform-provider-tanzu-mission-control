/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"testing"

	"github.com/stretchr/testify/require"

	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

func TestFlattenSubjects(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *iammodel.VmwareTanzuCoreV1alpha1PolicySubject
		expected    interface{}
	}{
		{
			description: "check for nil subject",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with all fields of subject",
			input: &iammodel.VmwareTanzuCoreV1alpha1PolicySubject{
				Name: "test",
				Kind: iammodel.VmwareTanzuCoreV1alpha1PolicySubjectKindGROUP.Pointer(),
			},
			expected: map[string]interface{}{
				subjectNameKey: "test",
				subjectKindKey: "GROUP",
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenSubject(test.input)
			require.EqualValues(t, test.expected, actual)
		})
	}
}
