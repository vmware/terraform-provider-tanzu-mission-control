/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

// nolint: dupl
package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
)

func TestFlattenTMCRequireLabels(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabels
		expected    []interface{}
	}{
		{
			description: "check for nil custom policy tmc_require_labels recipe",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with with complete custom policy tmc_require_labels recipe",
			input: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabels{
				Audit: true,
				Parameters: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParameters{
					Labels: []*policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParametersLabels{
						{
							Key:   "key-1",
							Value: "value-1",
						},
					},
				},
				TargetKubernetesResources: []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources{
					{
						APIGroups: []string{"policy"},
						Kinds:     []string{"pod"},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					AuditKey: true,
					ParametersKey: []interface{}{
						map[string]interface{}{
							parametersLabelKey: []interface{}{
								map[string]interface{}{
									labelKey:      "key-1",
									labelValueKey: "value-1",
								},
							},
						},
					},
					TargetKubernetesResourcesKey: []interface{}{
						map[string]interface{}{
							APIGroupsKey: []string{"policy"},
							KindsKey:     []string{"pod"},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenTMCRequireLabels(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenRequiredLabelsParameters(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParameters
		expected    []interface{}
	}{
		{
			description: "check for nil custom policy tmc_require_labels parameters",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with with complete custom policy tmc_require_labels parameters",
			input: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParameters{
				Labels: []*policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParametersLabels{
					{
						Key:   "key-1",
						Value: "value-1",
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					parametersLabelKey: []interface{}{
						map[string]interface{}{
							labelKey:      "key-1",
							labelValueKey: "value-1",
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenRequiredLabelsParameters(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenRequiredLabelsParametersLabels(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParametersLabels
		expected    interface{}
	}{
		{
			description: "check for nil custom policy tmc_require_labels parameters labels",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with with complete custom policy tmc_require_labels parameters labels",
			input: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParametersLabels{
				Key:   "key-1",
				Value: "value-1",
			},
			expected: map[string]interface{}{
				labelKey:      "key-1",
				labelValueKey: "value-1",
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenLabels(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
