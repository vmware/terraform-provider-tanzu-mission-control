/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindnetwork

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network/recipe"
)

func TestFlattenInput(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *inputRecipe
		expected    []interface{}
	}{
		{
			description: "check for nil input",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with input as allow-all recipe",
			input: &inputRecipe{
				recipe: AllowAllRecipe,
				inputAllowAll: &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAll{
					FromOwnNamespace: helper.BoolPointer(true),
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					reciperesource.AllowAllKey: []interface{}{
						map[string]interface{}{
							reciperesource.FromOwnNamespaceKey: true,
						},
					},
				},
			},
		},
		{
			description: "normal scenario with input as allow-all-to-pods recipe",
			input: &inputRecipe{
				recipe: AllowAllToPodsRecipe,
				inputAllowAllToPods: &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAllToPods{
					FromOwnNamespace: helper.BoolPointer(true),
					ToPodLabels: []*policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
						{
							Key:   "foo",
							Value: "bar",
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					reciperesource.AllowAllToPodsKey: []interface{}{
						map[string]interface{}{
							reciperesource.FromOwnNamespaceKey: true,
							reciperesource.ToPodLabelsKey: map[string]interface{}{
								"foo": "bar",
							},
						},
					},
				},
			},
		},
		{
			description: "normal scenario with input as allow-all-egress recipe",
			input: &inputRecipe{
				recipe: AllowAllEgressRecipe,
			},
			expected: []interface{}{
				map[string]interface{}{
					reciperesource.AllowAllEgressKey: []interface{}{
						map[string]interface{}{},
					},
				},
			},
		},
		{
			description: "normal scenario with input as deny-all recipe",
			input: &inputRecipe{
				recipe: DenyAllRecipe,
			},
			expected: []interface{}{
				map[string]interface{}{
					reciperesource.DenyAllKey: []interface{}{
						map[string]interface{}{},
					},
				},
			},
		},
		{
			description: "normal scenario with input as deny-all-to-pods recipe",
			input: &inputRecipe{
				recipe: DenyAllToPodsRecipe,
				inputDenyAllToPods: &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1DenyAllToPods{
					ToPodLabels: []*policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
						{
							Key:   "foo",
							Value: "bar",
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					reciperesource.DenyAllToPodsKey: []interface{}{
						map[string]interface{}{
							reciperesource.ToPodLabelsKey: map[string]interface{}{
								"foo": "bar",
							},
						},
					},
				},
			},
		},
		{
			description: "normal scenario with input as deny-all-egress recipe",
			input: &inputRecipe{
				recipe: DenyAllEgressRecipe,
			},
			expected: []interface{}{
				map[string]interface{}{
					reciperesource.DenyAllEgressKey: []interface{}{
						map[string]interface{}{},
					},
				},
			},
		},
		{
			description: "normal scenario with input as custom-egress recipe",
			input: &inputRecipe{
				recipe:            CustomEgressRecipe,
				inputCustomEgress: &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomEgress{},
			},
			expected: []interface{}{
				map[string]interface{}{
					reciperesource.CustomEgressKey: []interface{}{
						map[string]interface{}{},
					},
				},
			},
		},
		{
			description: "normal scenario with input as custom-ingress recipe",
			input: &inputRecipe{
				recipe:             CustomIngressRecipe,
				inputCustomIngress: &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomIngress{},
			},
			expected: []interface{}{
				map[string]interface{}{
					reciperesource.CustomIngressKey: []interface{}{
						map[string]interface{}{},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenInput(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
