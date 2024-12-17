// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/cluster"
	helmclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"
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
				FullnameCluster: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmFullName{
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
				FullnameClusterGroup: &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmFullName{
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
			actualData := FlattenScope(test.input)
			require.Equal(t, test.expectedData, actualData)
		})
	}
}
