/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgservicevsphere

import (
	"testing"

	"github.com/stretchr/testify/require"

	tkgservicevspheremodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgservicevsphere"
)

func TestFlattenSettings(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSettings
		expected []interface{}
	}{
		{
			name:     "check for nil settings data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with TKGs workload settings network data",
			input: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSettings{
				Network: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkSettings{
					Pods: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkRanges{
						CidrBlocks: []string{"127.0.01"},
					},
					Services: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkRanges{
						CidrBlocks: []string{"192.0.0.3"},
					},
				},
				Storage: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereStorageSettings{
					Classes:      []string{"wcpglobal-storage-class"},
					DefaultClass: "default",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					networkKey: []interface{}{
						map[string]interface{}{
							podsKey: []interface{}{
								map[string]interface{}{
									cidrBlocksKey: []string{"127.0.01"},
								},
							},
							servicesKey: []interface{}{
								map[string]interface{}{
									cidrBlocksKey: []string{"192.0.0.3"},
								},
							},
						},
					},
					storageKey: []interface{}{
						map[string]interface{}{
							classesKey:      []string{"wcpglobal-storage-class"},
							defaultClassKey: "default",
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := flattenTKGSSettings(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
