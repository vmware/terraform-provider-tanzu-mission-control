/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgvsphere

import (
	"testing"

	"github.com/stretchr/testify/require"

	tkgvspheremodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgvsphere"
)

func TestFlattenDistribution(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereDistribution
		expected []interface{}
	}{
		{
			name:     "check for nil distribution data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with all fields of distribution data",
			input: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereDistribution{
				Version: "v1.20",
				Workspace: &tkgvspheremodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereWorkspace{
					Datacenter:   "/dc0",
					Datastore:    "/dc0/datastore/local-0",
					Folder:       "/dc0/vm",
					Network:      "/dc0/network/Avi Internal",
					ResourcePool: "/dc0/host/cluster0/Resources",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					versionKey: "v1.20",
					workspaceKey: []interface{}{
						map[string]interface{}{
							datacenterKey:       "/dc0",
							datastoreKey:        "/dc0/datastore/local-0",
							folderKey:           "/dc0/vm",
							workspaceNetworkKey: "/dc0/network/Avi Internal",
							resourcePoolKey:     "/dc0/host/cluster0/Resources",
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := flattenTKGVsphereDistribution(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
