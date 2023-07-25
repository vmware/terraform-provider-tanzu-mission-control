/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	gitrepositoryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenClusterGroupGitRepositoryFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster group git repository full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group git repository full name",
			input: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName{
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
			actual := FlattenClusterGroupGitRepositoryFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
