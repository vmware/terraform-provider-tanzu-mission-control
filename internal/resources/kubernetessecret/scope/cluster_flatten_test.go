/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	secretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenClusterSecretFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster secret full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster secret full name",
			input: &secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName{
				ClusterName:           "c",
				ManagementClusterName: "m",
				ProvisionerName:       "p",
			},
			expected: []interface{}{
				map[string]interface{}{
					commonscope.NameKey:                  "c",
					commonscope.ManagementClusterNameKey: "m",
					commonscope.ProvisionerNameKey:       "p",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenClusterSecretFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
