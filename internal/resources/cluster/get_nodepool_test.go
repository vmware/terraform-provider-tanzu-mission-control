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

const (
	testCloudKey                   = "cloud-key"
	testEtcd0                      = "etcd-0"
	testM5large                    = "m5.large"
	testNodeKey                    = "node-key"
	testTest                       = "test"
	testTestClassSpec              = "test-class-spec"
	testTestStorageSpec            = "test-storage-spec"
	testTestingGetNodepoolFunction = "testing get nodepool function"
	testTkgsK8sObjPolicy           = "tkgs-k8s-obj-policy"
	testUsWest2a                   = "us-west-2a"
	testV1212vmware1               = "v1.21.2+vmware.1"
	testVarLibEtcd                 = "/var/lib/etcd"

	testCloudValue = "cloud-value"
	testNodeValue  = "node-value"
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
				ClusterGroupName: clusterGroupDefaultValue,
			},
			inputRespSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
			},
		},
		{
			description: "scenario for TKG AWS spec",
			inputClusterSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									WorkerNodeCount: "1",
									TkgAws: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool{
										AvailabilityZone: testUsWest2a,
										InstanceType:     testM5large,
										NodePlacement: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement{
											{
												AvailabilityZone: testUsWest2a,
											},
										},
										SubnetID: clusterGroupDefaultValue,
										Version:  testV1212vmware1,
									},
								},
							},
						},
					},
				},
			},
			inputRespSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									WorkerNodeCount: "1",
									TkgAws: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool{
										AvailabilityZone: "us-east-2a",
										InstanceType:     testM5large,
										NodePlacement: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement{
											{
												AvailabilityZone: testUsWest2a,
											},
										},
										SubnetID: clusterGroupDefaultValue,
										Version:  testV1212vmware1,
									},
								},
							},
						},
					},
				},
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									WorkerNodeCount: "1",
									TkgAws: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool{
										AvailabilityZone: testUsWest2a,
										InstanceType:     testM5large,
										NodePlacement: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement{
											{
												AvailabilityZone: testUsWest2a,
											},
										},
										SubnetID: clusterGroupDefaultValue,
										Version:  testV1212vmware1,
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
				ClusterGroupName: clusterGroupDefaultValue,
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									WorkerNodeCount: "1",
									TkgAws: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool{
										AvailabilityZone: testUsWest2a,
										InstanceType:     testM5large,
										NodePlacement: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement{
											{
												AvailabilityZone: testUsWest2a,
											},
										},
										SubnetID: clusterGroupDefaultValue,
										Version:  testV1212vmware1,
									},
								},
							},
						},
					},
				},
			},
			inputRespSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: nil,
				},
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
				TkgAws: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{
					Topology: nil,
				},
			},
		},
		{
			description: "scenario for TKG vSphere spec",
			inputClusterSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
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
				ClusterGroupName: clusterGroupDefaultValue,
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
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
				ClusterGroupName: clusterGroupDefaultValue,
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
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
				ClusterGroupName: clusterGroupDefaultValue,
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
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
				ClusterGroupName: clusterGroupDefaultValue,
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: nil,
				},
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
				TkgVsphere: &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{
					Topology: nil,
				},
			},
		},
		{
			description: "scenario for TKGs spec",
			inputClusterSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
				TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{
					Topology: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									CloudLabels:     map[string]string{testCloudKey: testCloudValue},
									NodeLabels:      map[string]string{testNodeKey: testNodeValue},
									WorkerNodeCount: "1",
									TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
										Class:        testTestClassSpec,
										StorageClass: testTestStorageSpec,
										Volumes: []*nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume{
											{
												Capacity:     4,
												MountPath:    testVarLibEtcd,
												Name:         testEtcd0,
												StorageClass: testTkgsK8sObjPolicy,
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
				ClusterGroupName: clusterGroupDefaultValue,
				TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{
					Topology: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									CloudLabels:     map[string]string{testCloudKey: testCloudValue},
									NodeLabels:      map[string]string{testNodeKey: testNodeValue},
									WorkerNodeCount: "1",
									TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
										Class:        testTestClassSpec,
										StorageClass: testTestStorageSpec,
										Volumes: []*nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume{
											{
												Capacity:     4,
												MountPath:    testVarLibEtcd,
												Name:         testEtcd0,
												StorageClass: testTkgsK8sObjPolicy,
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
				ClusterGroupName: clusterGroupDefaultValue,
				TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{
					Topology: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									CloudLabels:     map[string]string{testCloudKey: testCloudValue},
									NodeLabels:      map[string]string{testNodeKey: testNodeValue},
									WorkerNodeCount: "1",
									TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
										Class:        testTestClassSpec,
										StorageClass: testTestStorageSpec,
										Volumes: []*nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume{
											{
												Capacity:     4,
												MountPath:    testVarLibEtcd,
												Name:         testEtcd0,
												StorageClass: testTkgsK8sObjPolicy,
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
				ClusterGroupName: clusterGroupDefaultValue,
				TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{
					Topology: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{
						NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
							{
								Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
									Name:        testTest,
									Description: testTestingGetNodepoolFunction,
								},
								Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
									CloudLabels:     map[string]string{testCloudKey: testCloudValue},
									NodeLabels:      map[string]string{testNodeKey: testNodeValue},
									WorkerNodeCount: "1",
									TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
										Class:        testTestClassSpec,
										StorageClass: testTestStorageSpec,
										Volumes: []*nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume{
											{
												Capacity:     4,
												MountPath:    testVarLibEtcd,
												Name:         testEtcd0,
												StorageClass: testTkgsK8sObjPolicy,
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
				ClusterGroupName: clusterGroupDefaultValue,
				TkgServiceVsphere: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{
					Topology: nil,
				},
			},
			expectedSpec: &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
				ClusterGroupName: clusterGroupDefaultValue,
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
