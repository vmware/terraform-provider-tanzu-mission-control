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

func TestFlattenCustomEgress(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomEgress
		expected    []interface{}
	}{
		{
			description: "check for nil custom-egress recipe network policy ",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with valid custom-egress recipe network policy ",
			input: &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomEgress{
				Rules: []policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRules{
					{
						Ports: &[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesPorts{
							{
								Port: func() *string {
									port := "8443"
									return &port
								}(),
								Protocol: policyrecipenetworkcommonmodel.NewV1alpha1CommonPolicySpecNetworkV1CustomRulesPortsProtocol(policyrecipenetworkcommonmodel.TCP),
							},
						},
					},
				},
				ToPodLabels: func() *[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels {
					return &[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
						{
							Key:   "foo",
							Value: "bar",
						},
					}
				}(),
			},
			expected: []interface{}{
				map[string]interface{}{
					rulesKey: []interface{}{
						map[string]interface{}{
							portsKey: []interface{}{
								map[string]interface{}{
									portKey:     "8443",
									protocolKey: "TCP",
								},
							},
						},
					},
					ToPodLabelsKey: map[string]interface{}{"foo": "bar"},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenCustomEgress(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
