/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgvsphere

import (
	"testing"

	"github.com/stretchr/testify/require"

	nodepoolmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	tkgvspheremodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgvsphere"
)

func TestFlattenTopology(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology
		expected []interface{}
	}{
		{
			name:     "check for nil topology data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with control plane data of topology",
			input: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{
				ControlPlane: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereControlPlane{
					VMConfig: &nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig{
						CPU:       "2",
						DiskGib:   "20",
						MemoryMib: "4096",
					},
					HighAvailability: false,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}{
						map[string]interface{}{
							vmConfigKey: []interface{}{
								map[string]interface{}{
									cpuKey:    "2",
									diskKey:   "20",
									memoryKey: "4096",
								},
							},
							highAvailabilityKey: false,
						},
					},
					nodePoolsKey: []interface{}{},
				},
			},
		},
		{
			name: "normal scenario with node pool data of topology",
			input: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{
				NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
					{
						Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
							Name:        "test",
							Description: "testing topology flatten function",
						},
						Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
							WorkerNodeCount: "1",
							TkgVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool{
								VMConfig: &nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig{
									CPU:       "2",
									DiskGib:   "40",
									MemoryMib: "8192",
								},
							},
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}(nil),
					nodePoolsKey: []interface{}{
						map[string]interface{}{
							nodePoolInfoKey: []interface{}{
								map[string]interface{}{
									nodePoolNameKey:        "test",
									nodePoolDescriptionKey: "testing topology flatten function",
								},
							},
							nodePoolSpecKey: []interface{}{
								map[string]interface{}{
									workerNodeCountKey: "1",
									tkgVsphereKey: []interface{}{
										map[string]interface{}{
											vmConfigKey: []interface{}{
												map[string]interface{}{
													cpuKey:    "2",
													diskKey:   "40",
													memoryKey: "8192",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "normal scenario with all fields of topology data",
			input: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{
				ControlPlane: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereControlPlane{
					VMConfig: &nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig{
						CPU:       "2",
						DiskGib:   "20",
						MemoryMib: "4096",
					},
					HighAvailability: false,
				},
				NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
					{
						Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
							Name:        "test",
							Description: "testing topology flatten function",
						},
						Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
							WorkerNodeCount: "1",
							TkgVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool{
								VMConfig: &nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig{
									CPU:       "2",
									DiskGib:   "40",
									MemoryMib: "8192",
								},
							},
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}{
						map[string]interface{}{
							vmConfigKey: []interface{}{
								map[string]interface{}{
									cpuKey:    "2",
									diskKey:   "20",
									memoryKey: "4096",
								},
							},
							highAvailabilityKey: false,
						},
					},
					nodePoolsKey: []interface{}{
						map[string]interface{}{
							nodePoolInfoKey: []interface{}{
								map[string]interface{}{
									nodePoolNameKey:        "test",
									nodePoolDescriptionKey: "testing topology flatten function",
								},
							},
							nodePoolSpecKey: []interface{}{
								map[string]interface{}{
									workerNodeCountKey: "1",
									tkgVsphereKey: []interface{}{
										map[string]interface{}{
											vmConfigKey: []interface{}{
												map[string]interface{}{
													cpuKey:    "2",
													diskKey:   "40",
													memoryKey: "8192",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := flattenTKGVsphereTopology(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
