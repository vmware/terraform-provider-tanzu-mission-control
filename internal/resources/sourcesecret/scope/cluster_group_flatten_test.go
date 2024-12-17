// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	sourcesecretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenClusterGroupSourcesecretFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster group source secret full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group source secret full name",
			input: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName{
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
			actual := FlattenClusterGroupSourcesecretFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
