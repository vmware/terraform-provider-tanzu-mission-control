/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	secretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
)

func TestFlattenClusterFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster policy full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete secret cluster full name",
			input: &secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName{
				ClusterName:           "c",
				ManagementClusterName: "m",
				ProvisionerName:       "p",
			},
			expected: []interface{}{
				map[string]interface{}{
					ClusterNameKey:           "c",
					ManagementClusterNameKey: "m",
					ProvisionerNameKey:       "p",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenClusterFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
