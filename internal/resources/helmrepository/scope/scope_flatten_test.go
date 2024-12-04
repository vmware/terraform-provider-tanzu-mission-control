// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrepository"
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
				FullnameCluster: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryFullName{
					ClusterName:           "c",
					ManagementClusterName: "m",
					ProvisionerName:       "p",
					Name:                  "n",
					NamespaceName:         "nn",
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
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actualData, name, namespace := FlattenScope(test.input)
			require.Equal(t, test.expectedName, name)
			require.Equal(t, test.expectedData, actualData)
			require.Equal(t, test.expectedNamespace, namespace)
		})
	}
}
