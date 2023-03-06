/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	secretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
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
				FullnameCluster: &secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName{
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
