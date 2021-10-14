/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgservicevsphere

import (
	"testing"

	"github.com/stretchr/testify/require"

	tkgservicevspheremodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgservicevsphere"
)

func TestFlattenTopology(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology
		expected []interface{}
	}{
		{
			name:     "check for nil topology data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with control plane data of topology",
			input: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
				ControlPlane: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane{
					Class:            "test-class",
					HighAvailability: false,
					StorageClass:     "test-storage-class",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}{
						map[string]interface{}{
							classKey:            "test-class",
							highAvailabilityKey: false,
							storageClassKey:     "test-storage-class",
						},
					},
					nodePoolsKey: []interface{}{},
				},
			},
		},
		{
			name: "normal scenario with node pool data of topology",
			input: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
				NodePools: []*tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
					{
						Info: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
							Name:        "test",
							Description: "testing topology flatten function",
						},
						Spec: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
							CloudLabels:     map[string]string{"cloud-key": "cloud-value"},
							NodeLabels:      map[string]string{"node-key": "node-value"},
							WorkerNodeCount: "1",
							TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
								Class:        "test-class-spec",
								StorageClass: "test-storage-spec",
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
							infoKey: []interface{}{
								map[string]interface{}{
									clusterNameKey: "test",
									descriptionKey: "testing topology flatten function",
								},
							},
							specKey: []interface{}{
								map[string]interface{}{
									cloudLabelKey:      map[string]string{"cloud-key": "cloud-value"},
									nodeLabelKey:       map[string]string{"node-key": "node-value"},
									workerNodeCountKey: "1",
									tkgServiceVsphereKey: []interface{}{
										map[string]interface{}{
											classKey:        "test-class-spec",
											storageClassKey: "test-storage-spec",
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
			input: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
				ControlPlane: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane{
					Class:            "test-class",
					HighAvailability: false,
					StorageClass:     "test-storage-class",
				},
				NodePools: []*tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
					{
						Info: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
							Name:        "test",
							Description: "testing topology flatten function",
						},
						Spec: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
							CloudLabels:     map[string]string{"cloud-key": "cloud-value"},
							NodeLabels:      map[string]string{"node-key": "node-value"},
							WorkerNodeCount: "1",
							TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
								Class:        "test-class-spec",
								StorageClass: "test-storage-spec",
							},
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}{
						map[string]interface{}{
							classKey:            "test-class",
							highAvailabilityKey: false,
							storageClassKey:     "test-storage-class",
						},
					},
					nodePoolsKey: []interface{}{
						map[string]interface{}{
							infoKey: []interface{}{
								map[string]interface{}{
									clusterNameKey: "test",
									descriptionKey: "testing topology flatten function",
								},
							},
							specKey: []interface{}{
								map[string]interface{}{
									cloudLabelKey:      map[string]string{"cloud-key": "cloud-value"},
									nodeLabelKey:       map[string]string{"node-key": "node-value"},
									workerNodeCountKey: "1",
									tkgServiceVsphereKey: []interface{}{
										map[string]interface{}{
											classKey:        "test-class-spec",
											storageClassKey: "test-storage-spec",
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
			actual := flattenTKGSTopology(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
