// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	"testing"

	"github.com/stretchr/testify/require"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	nodepoolmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	tkgawsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgaws"
	tkgservicevspheremodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgservicevsphere"
	tkgvspheremodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgvsphere"
)

func TestGetNodepoolForCluster(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description      string
		inputRespSpec    *clustermodel.VmwareTanzuManageV1alpha1ClusterSpec
		inputClusterSpec *clustermodel.VmwareTanzuManageV1alpha1ClusterSpec
		expectedSpec     *clustermodel.VmwareTanzuManageV1alpha1ClusterSpec
	}{
		{
			description:      "check for nil data in cluster spec",
			inputClusterSpec: nil,
			inputRespSpec:    nil,
			expectedSpec:     nil,
		},
		{
			description: "scenario for attach cluster",
			inputClusterSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
			},
			inputRespSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
			},
		},
		{
			description: "scenario for TKG AWS spec",
			inputClusterSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									WorkerNodeCount: "1",
									TkgAws: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool{
										AvailabilityZone: "us-west-2a",
										InstanceType:     "m5.large",
										NodePlacement: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement{
											{
												AvailabilityZone: "us-west-2a",
											},
										},
										SubnetID: "default",
										Version:  "v1.21.2+vmware.1",
									},
								},
							},
						},
					},
				},
			},
			inputRespSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									WorkerNodeCount: "1",
									TkgAws: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool{
										AvailabilityZone: "us-east-2a",
										InstanceType:     "m5.large",
										NodePlacement: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement{
											{
												AvailabilityZone: "us-west-2a",
											},
										},
										SubnetID: "default",
										Version:  "v1.21.2+vmware.1",
									},
								},
							},
						},
					},
				},
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									WorkerNodeCount: "1",
									TkgAws: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool{
										AvailabilityZone: "us-west-2a",
										InstanceType:     "m5.large",
										NodePlacement: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement{
											{
												AvailabilityZone: "us-west-2a",
											},
										},
										SubnetID: "default",
										Version:  "v1.21.2+vmware.1",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			description: "scenario for TKG AWS spec topology is nil",
			inputClusterSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									WorkerNodeCount: "1",
									TkgAws: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool{
										AvailabilityZone: "us-west-2a",
										InstanceType:     "m5.large",
										NodePlacement: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement{
											{
												AvailabilityZone: "us-west-2a",
											},
										},
										SubnetID: "default",
										Version:  "v1.21.2+vmware.1",
									},
								},
							},
						},
					},
				},
			},
			inputRespSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: nil,
				},
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: nil,
				},
			},
		},
		{
			description: "scenario for TKG vSphere spec",
			inputClusterSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
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
				},
			},
			inputRespSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
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
				},
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
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
				},
			},
		},
		{
			description: "scenario for TKG vSphere spec topology is nil",
			inputClusterSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
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
				},
			},
			inputRespSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: nil,
				},
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: nil,
				},
			},
		},
		{
			description: "scenario for TKGs spec",
			inputClusterSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{
					Topology: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
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
				},
			},
			inputRespSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{
					Topology: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
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
				},
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{
					Topology: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
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
				},
			},
		},
		{
			description: "scenario for TKGs spec topology is nil",
			inputClusterSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{
					Topology: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        "test",
									Description: "testing get nodepool function",
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
				},
			},
			inputRespSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{
					Topology: nil,
				},
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: "default",
				TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{
					Topology: nil,
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			setNodepoolForClusterResource(test.inputRespSpec, test.inputClusterSpec)
			require.Equal(t, test.expectedSpec, test.inputRespSpec)
		})
	}
}
