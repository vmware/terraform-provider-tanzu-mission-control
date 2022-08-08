/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgservicevsphere

import (
	"testing"

	"github.com/stretchr/testify/require"

	nodepoolmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	tkgservicevspheremodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgservicevsphere"
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
			name: "normal scenario with empty node pool data of topology",
			input: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
				NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
					nil,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}(nil),
					nodePoolsKey: []interface{}{
						nil,
					},
				},
			},
		},
		{
			name: "normal scenario with control plane data of topology",
			input: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
				ControlPlane: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane{
					Class:            "test-class",
					HighAvailability: false,
					StorageClass:     "test-storage-class",
					Volumes: []*nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume{
						{
							Capacity:     4,
							MountPath:    "/var/lib/etcd",
							Name:         "etcd-0",
							StorageClass: "tkgs-k8s-obj-policy",
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
							volumesKey: []interface{}{
								map[string]interface{}{
									capacityKey:        float32(4),
									mountPathKey:       "/var/lib/etcd",
									volumeNameKey:      "etcd-0",
									pvcStorageClassKey: "tkgs-k8s-obj-policy",
								},
							},
						},
					},
					nodePoolsKey: []interface{}{},
				},
			},
		},
		{
			name: "normal scenario with node pool data of topology",
			input: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
				NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
					{
						Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
							Name:        "test",
							Description: "testing topology flatten function",
						},
						Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
							CloudLabels:     map[string]string{"cloud-key": "cloud-value"},
							NodeLabels:      map[string]string{"node-key": "node-value"},
							WorkerNodeCount: "1",
							TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
								Class:        "test-class-spec",
								StorageClass: "test-storage-spec",
								Volumes: []*nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume{
									{
										Capacity:     4,
										MountPath:    "/var/lib/etcd",
										Name:         "etcd-0",
										StorageClass: "tkgs-k8s-obj-policy",
									},
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
							infoKey: []interface{}{
								map[string]interface{}{
									nodepoolNameKey: "test",
									descriptionKey:  "testing topology flatten function",
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
											volumesKey: []interface{}{
												map[string]interface{}{
													capacityKey:        float32(4),
													mountPathKey:       "/var/lib/etcd",
													volumeNameKey:      "etcd-0",
													pvcStorageClassKey: "tkgs-k8s-obj-policy",
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
			input: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
				ControlPlane: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane{
					Class:            "test-class",
					HighAvailability: false,
					StorageClass:     "test-storage-class",
					Volumes: []*nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume{
						{
							Capacity:     4,
							MountPath:    "/var/lib/etcd",
							Name:         "etcd-0",
							StorageClass: "tkgs-k8s-obj-policy",
						},
					},
				},
				NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
					{
						Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
							Name:        "test",
							Description: "testing topology flatten function",
						},
						Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
							CloudLabels:     map[string]string{"cloud-key": "cloud-value"},
							NodeLabels:      map[string]string{"node-key": "node-value"},
							WorkerNodeCount: "1",
							TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
								Class:        "test-class-spec",
								StorageClass: "test-storage-spec",
								Volumes: []*nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume{
									{
										Capacity:     4,
										MountPath:    "/var/lib/etcd",
										Name:         "etcd-0",
										StorageClass: "tkgs-k8s-obj-policy",
									},
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
							classKey:            "test-class",
							highAvailabilityKey: false,
							storageClassKey:     "test-storage-class",
							volumesKey: []interface{}{
								map[string]interface{}{
									capacityKey:        float32(4),
									mountPathKey:       "/var/lib/etcd",
									volumeNameKey:      "etcd-0",
									pvcStorageClassKey: "tkgs-k8s-obj-policy",
								},
							},
						},
					},
					nodePoolsKey: []interface{}{
						map[string]interface{}{
							infoKey: []interface{}{
								map[string]interface{}{
									nodepoolNameKey: "test",
									descriptionKey:  "testing topology flatten function",
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
											volumesKey: []interface{}{
												map[string]interface{}{
													capacityKey:        float32(4),
													mountPathKey:       "/var/lib/etcd",
													volumeNameKey:      "etcd-0",
													pvcStorageClassKey: "tkgs-k8s-obj-policy",
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
			actual := flattenTKGSTopology(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
