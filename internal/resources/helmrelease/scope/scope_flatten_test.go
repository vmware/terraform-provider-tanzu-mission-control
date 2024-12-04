// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	releaseclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/cluster"
	releaseclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/clustergroup"
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
	}{
		{
			description:       "check for nil scope",
			input:             nil,
			expectedData:      nil,
			expectedName:      "",
			expectedNamespace: "",
		},
		{
			description: "normal scenario with complete cluster scope",
			input: &ScopedFullname{
				Scope: commonscope.ClusterScope,
				FullnameCluster: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName{
					Name:                  "n",
					NamespaceName:         "nn",
					ClusterName:           "c",
					ManagementClusterName: "m",
					ProvisionerName:       "p",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					commonscope.ClusterKey: []interface{}{
						map[string]interface{}{
							commonscope.ManagementClusterNameKey: "m",
							commonscope.NameKey:                  "c",
							commonscope.ProvisionerNameKey:       "p",
						},
					},
				},
			},
			expectedName:      "n",
			expectedNamespace: "nn",
		},
		{
			description: "normal scenario with complete cluster group scope",
			input: &ScopedFullname{
				Scope: commonscope.ClusterGroupScope,
				FullnameClusterGroup: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseFullName{
					Name:             "n",
					NamespaceName:    "nn",
					ClusterGroupName: "c",
				},
			},
			expectedData: []interface{}{
				map[string]interface{}{
					commonscope.ClusterGroupKey: []interface{}{
						map[string]interface{}{
							commonscope.NameKey: "c",
						},
					},
				},
			},
			expectedName:      "n",
			expectedNamespace: "nn",
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actualData, actualName, actualNamespace := FlattenScope(test.input)
			require.Equal(t, test.expectedData, actualData)
			require.Equal(t, test.expectedName, actualName)
			require.Equal(t, test.expectedNamespace, actualNamespace)
		})
	}
}
