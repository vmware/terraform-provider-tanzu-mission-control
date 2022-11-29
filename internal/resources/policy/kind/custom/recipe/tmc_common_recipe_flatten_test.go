/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
)

func TestFlattenTMCCommonRecipe(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe
		expected    []interface{}
	}{
		{
			description: "check for nil custom policy tmc-block-nodeport-service/ tmc-block-resources/ tmc-https-ingress recipes",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete custom policy tmc-block-nodeport-service/ tmc-block-resources/ tmc-https-ingress recipes",
			input: &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe{
				Audit: true,
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
			actual := FlattenTMCCommonRecipe(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
