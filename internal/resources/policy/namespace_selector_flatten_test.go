// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policy

import (
	"testing"

	"github.com/stretchr/testify/require"

	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
)

func TestFlattenNamespaceSelector(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policymodel.VmwareTanzuManageV1alpha1CommonPolicyLabelSelector
		expected    []interface{}
	}{
		{
			description: "check for nil policy label selector",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete policy label selector",
			input: &policymodel.VmwareTanzuManageV1alpha1CommonPolicyLabelSelector{
				MatchExpressions: []*policymodel.K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement{
					{
						Key:      "k1",
						Operator: "In",
						Values: []string{
							"v1",
							"v2",
						},
					},
					{
						Key:      "k2",
						Operator: "Exists",
						Values:   []string{},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					MatchExpressionsKey: []interface{}{
						map[string]interface{}{
							KeyKey:      "k1",
							OperatorKey: "In",
							ValuesKey: []string{
								"v1",
								"v2",
							},
						},
						map[string]interface{}{
							KeyKey:      "k2",
							OperatorKey: "Exists",
							ValuesKey:   []string{},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenNamespaceSelector(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
