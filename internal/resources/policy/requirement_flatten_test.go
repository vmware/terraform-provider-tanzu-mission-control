/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policy

import (
	"testing"

	"github.com/stretchr/testify/require"

	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
)

func TestFlattenLabelSelectorRequirement(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policymodel.K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement
		expected    interface{}
	}{
		{
			description: "check for nil label selector requirement",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete label selector requirement",
			input: &policymodel.K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement{
				Key:      "k",
				Operator: "In",
				Values: []string{
					"v1",
					"v2",
				},
			},
			expected: map[string]interface{}{
				KeyKey:      "k",
				OperatorKey: "In",
				ValuesKey: []string{
					"v1",
					"v2",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenLabelSelectorRequirement(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
