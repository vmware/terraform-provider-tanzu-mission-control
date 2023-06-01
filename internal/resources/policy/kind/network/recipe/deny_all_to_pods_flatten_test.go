/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

func TestFlattenDenyAllToPods(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1DenyAllToPods
		expected    []interface{}
	}{
		{
			description: "check for nil deny-all-to-pods recipe network policy",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with valid deny-all-to-pods recipe network policy",
			input: &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1DenyAllToPods{
				ToPodLabels: []*policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
					{
						Key:   "foo",
						Value: "bar",
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					ToPodLabelsKey: map[string]interface{}{"foo": "bar"},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenDenyAllToPods(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
