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

func TestFlattenSettings(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSettings
		expected []interface{}
	}{
		{
			name:     "check for nil settings data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with TKGm Vsphere workload settings network data",
			input: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSettings{
				Network: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkSettings{
					APIServerPort:        6443,
					ControlPlaneEndpoint: "10.185.107.47",
					Pods: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkRanges{
						CidrBlocks: []string{"127.0.0.1"},
					},
					Services: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkRanges{
						CidrBlocks: []string{"192.0.0.3"},
					},
				},
				Security: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSecuritySettings{
					SSHKey: "default",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					networkKey: []interface{}{
						map[string]interface{}{
							apiServerPortKey:        int32(6443),
							controlPlaneEndPointKey: "10.185.107.47",
							podsKey: []interface{}{
								map[string]interface{}{
									cidrBlockKey: []string{"127.0.0.1"},
								},
							},
							servicesKey: []interface{}{
								map[string]interface{}{
									cidrBlockKey: []string{"192.0.0.3"},
								},
							},
						},
					},
					securityKey: []interface{}{
						map[string]interface{}{
							sshKey: "default",
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := flattenTKGVsphereSettings(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
