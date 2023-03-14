/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
	gitrepositoryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description       string
		input             *ScopedFullname
		expectedData      []interface{}
		expectedName      string
		expectedNamespace string
		expectedOrgID     string
	}{
		{
			description:       "check for nil scope",
			input:             nil,
			expectedData:      nil,
			expectedName:      "",
			expectedNamespace: "",
			expectedOrgID:     "",
		},
		{
			description: "normal scenario with complete cluster scope",
			input: &ScopedFullname{
				Scope: commonscope.ClusterScope,
				FullnameCluster: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName{
					Name:                  "n",
					NamespaceName:         "nn",
					ClusterName:           "c",
					ManagementClusterName: "m",
					ProvisionerName:       "p",
					OrgID:                 "o",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					commonscope.ClusterKey: []interface{}{
						map[string]interface{}{
							commonscope.ManagementClusterNameKey: "m",
							commonscope.ClusterNameKey:           "c",
							commonscope.ProvisionerNameKey:       "p",
						},
					},
				},
			},
			expectedName:      "n",
			expectedNamespace: "nn",
			expectedOrgID:     "o",
		},
		{
			description: "normal scenario with complete cluster group scope",
			input: &ScopedFullname{
				Scope: commonscope.ClusterGroupScope,
				FullnameClusterGroup: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName{
					Name:             "n",
					NamespaceName:    "nn",
					ClusterGroupName: "c",
					OrgID:            "o",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					commonscope.ClusterGroupKey: []interface{}{
						map[string]interface{}{
							commonscope.ClusterGroupNameKey: "c",
						},
					},
				},
			},
			expectedName:      "n",
			expectedNamespace: "nn",
			expectedOrgID:     "o",
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actualData, actualName, actualNamespace, actualOrgID := FlattenScope(test.input)
			require.Equal(t, test.expectedData, actualData)
			require.Equal(t, test.expectedName, actualName)
			require.Equal(t, test.expectedNamespace, actualNamespace)
			require.Equal(t, test.expectedOrgID, actualOrgID)
		})
	}
}
