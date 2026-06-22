// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
)

func TestFlattenTMCExternalIPS(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPS
		expected    []interface{}
	}{
		{
			description: "check for nil custom policy tmc_external_ips recipe",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete custom policy tmc_external_ips recipe",
			input: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPS{
				Audit: true,
				Parameters: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPSParameters{
					AllowedIPs: []string{test127001},
				},
				TargetKubernetesResources: []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources{
					{
						APIGroups: []string{testPolicy},
						Kinds:     []string{testPod},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					AuditKey: true,
					ParametersKey: []interface{}{
						map[string]interface{}{
							allowedIPsKey: []string{test127001},
						},
					},
					TargetKubernetesResourcesKey: []interface{}{
						map[string]interface{}{
							APIGroupsKey: []string{testPolicy},
							KindsKey:     []string{testPod},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenTMCExternalIPS(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenExternalIPSParameters(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPSParameters
		expected    []interface{}
	}{
		{
			description: "check for nil custom policy tmc_external_ips parameters",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete custom policy tmc_external_ips parameters",
			input: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPSParameters{
				AllowedIPs: []string{test127001},
			},
			expected: []interface{}{
				map[string]interface{}{
					allowedIPsKey: []string{test127001},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenExternalIPSParameters(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
