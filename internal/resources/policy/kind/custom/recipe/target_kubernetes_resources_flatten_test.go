/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
)

func TestFlattenTargetKubernetesResources(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources
		expected    interface{}
	}{
		{
			description: "check for nil Target Kubernetes resources",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with all values of Target Kubernetes resources",
			input: &policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources{
				APIGroups: []string{"policy"},
				Kinds:     []string{"pod"},
			},
			expected: map[string]interface{}{
				APIGroupsKey: []string{"policy"},
				KindsKey:     []string{"pod"},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenTargetKubernetesResources(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
