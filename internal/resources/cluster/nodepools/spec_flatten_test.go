// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package nodepools

import (
	"testing"

	"github.com/stretchr/testify/require"

	nodepoolmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
)

const (
	testABCD           = "/a/b/c/d"
	testKey            = "key"
	testPvStorageClass = "pv_storage_class"
	testVolume1        = "volume1"

	testValue = "value"
)

func TestFlattenNodePoolSpec(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec
		expected    []interface{}
	}{
		{
			description: "check for nil TKGs/TKGm vsphere spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete nodepool spec",
			input: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
				CloudLabels: map[string]string{
					testKey: testValue,
				},
				NodeLabels: map[string]string{
					testKey: testValue,
				},
				TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
					Class:         classKey,
					StorageClass:  storageClassKey,
					FailureDomain: failureDomainKey,
					Volumes: []*nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume{
						{
							Capacity:     4,
							MountPath:    testABCD,
							Name:         testVolume1,
							StorageClass: testPvStorageClass,
						},
					},
				},
				TkgVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool{
					VMConfig: &nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig{
						CPU:       "2",
						DiskGib:   "20",
						MemoryMib: "4096",
					},
				},
				WorkerNodeCount: "1",
			},
			expected: []interface{}{
				map[string]interface{}{
					cloudLabelsKey: map[string]string{
						testKey: testValue,
					},
					nodeLabelsKey: map[string]string{
						testKey: testValue,
					},
					workerNodeCountKey: "1",
					tkgServiceVsphereKey: []interface{}{
						map[string]interface{}{
							classKey:         classKey,
							storageClassKey:  storageClassKey,
							failureDomainKey: failureDomainKey,
							volumesKey: []interface{}{
								map[string]interface{}{
									capacityKey:        float32(4),
									mountPathKey:       testABCD,
									volumeNameKey:      testVolume1,
									pvcStorageClassKey: testPvStorageClass,
								},
							},
						},
					},
					tkgVsphereKey: []interface{}{
						map[string]interface{}{
							vmConfigKey: []interface{}{
								map[string]interface{}{
									cpuKey:    "2",
									diskKey:   "20",
									memoryKey: "4096",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "normal scenario with empty TKGs Vsphere data",
			input: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
				CloudLabels: map[string]string{
					testKey: testValue,
				},
				NodeLabels: map[string]string{
					testKey: testValue,
				},
				TkgVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool{
					VMConfig: &nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig{
						CPU:       "2",
						DiskGib:   "20",
						MemoryMib: "4096",
					},
				},
				TkgServiceVsphere: nil,
				WorkerNodeCount:   "1",
			},
			expected: []interface{}{
				map[string]interface{}{
					cloudLabelsKey: map[string]string{
						testKey: testValue,
					},
					nodeLabelsKey: map[string]string{
						testKey: testValue,
					},
					workerNodeCountKey: "1",
					tkgVsphereKey: []interface{}{
						map[string]interface{}{
							vmConfigKey: []interface{}{
								map[string]interface{}{
									cpuKey:    "2",
									diskKey:   "20",
									memoryKey: "4096",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "normal scenario with empty VM config",
			input: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
				CloudLabels: map[string]string{
					testKey: testValue,
				},
				NodeLabels: map[string]string{
					testKey: testValue,
				},
				TkgVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool{
					VMConfig: nil,
				},
				TkgServiceVsphere: nil,
				WorkerNodeCount:   "1",
			},
			expected: []interface{}{
				map[string]interface{}{
					cloudLabelsKey: map[string]string{
						testKey: testValue,
					},
					nodeLabelsKey: map[string]string{
						testKey: testValue,
					},
					workerNodeCountKey: "1",
					tkgVsphereKey: []interface{}{
						map[string]interface{}{},
					},
				},
			},
		},
		{
			description: "normal scenario with empty TKGm Vsphere data",
			input: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
				CloudLabels: map[string]string{
					testKey: testValue,
				},
				NodeLabels: map[string]string{
					testKey: testValue,
				},
				TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
					Class:         classKey,
					StorageClass:  storageClassKey,
					FailureDomain: failureDomainKey,
					Volumes: []*nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume{
						{
							Capacity:     4,
							MountPath:    testABCD,
							Name:         testVolume1,
							StorageClass: testPvStorageClass,
						},
					},
				},
				TkgVsphere:      nil,
				WorkerNodeCount: "1",
			},
			expected: []interface{}{
				map[string]interface{}{
					cloudLabelsKey: map[string]string{
						testKey: testValue,
					},
					nodeLabelsKey: map[string]string{
						testKey: testValue,
					},
					workerNodeCountKey: "1",
					tkgServiceVsphereKey: []interface{}{
						map[string]interface{}{
							classKey:         classKey,
							storageClassKey:  storageClassKey,
							failureDomainKey: failureDomainKey,
							volumesKey: []interface{}{
								map[string]interface{}{
									capacityKey:        float32(4),
									mountPathKey:       testABCD,
									volumeNameKey:      testVolume1,
									pvcStorageClassKey: testPvStorageClass,
								},
							},
						},
					},
				},
			},
		},
		{
			description: "normal scenario with empty TKGs and TKGm Vsphere data",
			input: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
				CloudLabels: map[string]string{
					testKey: testValue,
				},
				NodeLabels: map[string]string{
					testKey: testValue,
				},
				WorkerNodeCount: "1",
			},
			expected: []interface{}{
				map[string]interface{}{
					cloudLabelsKey: map[string]string{
						testKey: testValue,
					},
					nodeLabelsKey: map[string]string{
						testKey: testValue,
					},
					workerNodeCountKey: "1",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenSpec(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
