// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	secretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenClusterGroupSecretFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster group cluster secret full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group cluster secret full name",
			input: &secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName{
				ClusterGroupName: "c",
			},
			expected: []interface{}{
				map[string]interface{}{
					commonscope.NameKey: "c",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenClusterGroupSecretFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
