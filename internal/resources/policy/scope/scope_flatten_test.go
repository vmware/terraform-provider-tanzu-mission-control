/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
	policyworkspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/workspace"
)

func TestFlattenScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description  string
		input        *ScopedFullname
		allowedScope []string
		expectedData []interface{}
		expectedName string
	}{
		{
			description:  "check for nil scope",
			input:        nil,
			expectedData: nil,
			expectedName: "",
		},
		{
			description: "normal scenario with complete cluster scope",
			input: &ScopedFullname{
				Scope: ClusterScope,
				FullnameCluster: &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName{
					Name:                  "n",
					ClusterName:           "c",
					ManagementClusterName: "m",
					ProvisionerName:       "p",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					ClusterKey: []interface{}{
						map[string]interface{}{
							ManagementClusterNameKey: "m",
							ClusterNameKey:           "c",
							ProvisionerNameKey:       "p",
						},
					},
				},
			},
			expectedName: "n",
		},
		{
			description: "normal scenario with complete cluster group scope",
			input: &ScopedFullname{
				Scope: ClusterGroupScope,
				FullnameClusterGroup: &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName{
					Name:             "n",
					ClusterGroupName: "c",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					ClusterGroupKey: []interface{}{
						map[string]interface{}{
							ClusterGroupNameKey: "c",
						},
					},
				},
			},
			expectedName: "n",
		},
		{
			description: "normal scenario with complete workspace scope",
			input: &ScopedFullname{
				Scope: WorkspaceScope,
				FullnameWorkspace: &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName{
					Name:          "n",
					WorkspaceName: "w",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					WorkspaceKey: []interface{}{
						map[string]interface{}{
							WorkspaceNameKey: "w",
						},
					},
				},
			},
			expectedName: "n",
		},
		{
			description: "normal scenario with complete organization scope",
			input: &ScopedFullname{
				Scope: OrganizationScope,
				FullnameOrganization: &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{
					Name:  "n",
					OrgID: "o",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					OrganizationKey: []interface{}{
						map[string]interface{}{
							OrganizationIDKey: "o",
						},
					},
				},
			},
			expectedName: "n",
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actualData, actualName := FlattenScope(test.input, test.allowedScope)
			require.Equal(t, test.expectedData, actualData)
			require.Equal(t, test.expectedName, actualName)
		})
	}
}
