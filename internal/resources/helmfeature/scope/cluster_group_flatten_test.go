// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	releaseclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenClusterGroupHelmFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster group git repository full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group git repository full name",
			input: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmFullName{
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
			actual := FlattenClusterGroupHelmFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
