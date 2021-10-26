/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepools

import (
	"testing"

	"github.com/stretchr/testify/require"

	nodepoolmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
)

func TestFlattenNodePoolSpec(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec
		expected    []interface{}
	}{
		{
			description: "check for nil TKGs spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete nodepool spec",
			input: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
				CloudLabels: map[string]string{
					"key": "value",
				},
				NodeLabels: map[string]string{
					"key": "value",
				},
				TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
					Class:        "class",
					StorageClass: "storage_class",
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
						"key": "value",
					},
					nodeLabelsKey: map[string]string{
						"key": "value",
					},
					workerNodeCountKey: "1",
					tkgServiceVsphereKey: []interface{}{
						map[string]interface{}{
							classKey:        "class",
							storageClassKey: "storage_class",
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
					"key": "value",
				},
				NodeLabels: map[string]string{
					"key": "value",
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
						"key": "value",
					},
					nodeLabelsKey: map[string]string{
						"key": "value",
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
					"key": "value",
				},
				NodeLabels: map[string]string{
					"key": "value",
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
						"key": "value",
					},
					nodeLabelsKey: map[string]string{
						"key": "value",
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
					"key": "value",
				},
				NodeLabels: map[string]string{
					"key": "value",
				},
				TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
					Class:        "class",
					StorageClass: "storage_class",
				},
				TkgVsphere:      nil,
				WorkerNodeCount: "1",
			},
			expected: []interface{}{
				map[string]interface{}{
					cloudLabelsKey: map[string]string{
						"key": "value",
					},
					nodeLabelsKey: map[string]string{
						"key": "value",
					},
					workerNodeCountKey: "1",
					tkgServiceVsphereKey: []interface{}{
						map[string]interface{}{
							classKey:        "class",
							storageClassKey: "storage_class",
						},
					},
				},
			},
		},
		{
			description: "normal scenario with empty TKGs and TKGm Vsphere data",
			input: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
				CloudLabels: map[string]string{
					"key": "value",
				},
				NodeLabels: map[string]string{
					"key": "value",
				},
				WorkerNodeCount: "1",
			},
			expected: []interface{}{
				map[string]interface{}{
					cloudLabelsKey: map[string]string{
						"key": "value",
					},
					nodeLabelsKey: map[string]string{
						"key": "value",
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
