// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
)

func TestFlattenAllowAll(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAll
		expected    []interface{}
	}{
		{
			description: "check for nil allow-all recipe network policy",
			input:       nil,
			expected:    nil,
		},
		{
			description: "scenario with no FromOwnNamespace attribute for allow-all recipe network policy",
			input: &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAll{
				FromOwnNamespace: nil,
			},
			expected: []interface{}(nil),
		},
		{
			description: "normal scenario with valid allow-all recipe network policy",
			input: &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAll{
				FromOwnNamespace: helper.BoolPointer(true),
			},
			expected: []interface{}{
				map[string]interface{}{
					FromOwnNamespaceKey: true,
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenAllowAll(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
