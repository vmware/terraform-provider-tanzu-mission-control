/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"testing"

	"github.com/stretchr/testify/require"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
	namespacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/namespace"
	organizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/organization"
	workspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/workspace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/namespace"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/workspace"
)

func TestFlattenScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description  string
		input        *scopedFullname
		expectedData []interface{}
	}{
		{
			description:  "check for nil scope",
			input:        nil,
			expectedData: nil,
		},
		{
			description: "normal scenario with complete cluster scope",
			input: &scopedFullname{
				scope: clusterScope,
				fullnameCluster: &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{
					Name:                  "dummy",
					ManagementClusterName: "attached",
					ProvisionerName:       "attached",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					clusterKey: []interface{}{
						map[string]interface{}{
							cluster.NameKey:                  "dummy",
							cluster.ManagementClusterNameKey: "attached",
							cluster.ProvisionerNameKey:       "attached",
						},
					},
				},
			},
		},
		{
			description: "normal scenario with complete cluster group scope",
			input: &scopedFullname{
				scope: clusterGroupScope,
				fullnameClusterGroup: &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{
					Name: "default",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					clusterGroupKey: []interface{}{
						map[string]interface{}{
							clustergroup.NameKey: "default",
						},
					},
				},
			},
		},
		{
			description: "normal scenario with complete organization scope",
			input: &scopedFullname{
				scope: organizationScope,
				fullnameOrganization: &organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName{
					OrgID: "playground",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					organizationKey: []interface{}{
						map[string]interface{}{
							organizationIDKey: "playground",
						},
					},
				},
			},
		},
		{
			description: "normal scenario with complete workspace scope",
			input: &scopedFullname{
				scope: workspaceScope,
				fullnameWorkspace: &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
					Name: "default",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					workspaceKey: []interface{}{
						map[string]interface{}{
							workspace.NameKey: "default",
						},
					},
				},
			},
		},
		{
			description: "normal scenario with complete namespace scope",
			input: &scopedFullname{
				scope: namespaceScope,
				fullnameNamespace: &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName{
					Name:                  "n-1",
					ClusterName:           "dummy",
					ManagementClusterName: "attached",
					ProvisionerName:       "attached",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					namespaceKey: []interface{}{
						map[string]interface{}{
							namespace.NameKey:                  "n-1",
							namespace.ClusterNameKey:           "dummy",
							namespace.ManagementClusterNameKey: "attached",
							namespace.ProvisionerNameKey:       "attached",
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actualData := flattenScope(test.input)
			require.Equal(t, test.expectedData, actualData)
		})
	}
}
