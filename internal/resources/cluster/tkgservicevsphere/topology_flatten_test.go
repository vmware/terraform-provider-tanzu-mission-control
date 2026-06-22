// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgservicevsphere

import (
	"testing"

	"github.com/stretchr/testify/require"

	nodepoolmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	tkgservicevspheremodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgservicevsphere"
)

const (
	testCloudKey                       = "cloud-key"
	testEtcd0                          = "etcd-0"
	testNodeKey                        = "node-key"
	testTest                           = "test"
	testTestClass                      = "test-class"
	testTestClassSpec                  = "test-class-spec"
	testTestStorageClass               = "test-storage-class"
	testTestStorageSpec                = "test-storage-spec"
	testTestingTopologyFlattenFunction = "testing topology flatten function"
	testTkgsK8sObjPolicy               = "tkgs-k8s-obj-policy"
	testVarLibEtcd                     = "/var/lib/etcd"

	testCloudValue = "cloud-value"
	testNodeValue  = "node-value"
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
					Class:            testTestClass,
					HighAvailability: false,
					StorageClass:     testTestStorageClass,
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
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}{
						map[string]interface{}{
							classKey:            testTestClass,
							highAvailabilityKey: false,
							storageClassKey:     testTestStorageClass,
							volumesKey: []interface{}{
								map[string]interface{}{
									capacityKey:        float32(4),
									mountPathKey:       testVarLibEtcd,
									volumeNameKey:      testEtcd0,
									pvcStorageClassKey: testTkgsK8sObjPolicy,
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
							Name:        testTest,
							Description: testTestingTopologyFlattenFunction,
						},
						Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
							CloudLabels:     map[string]string{testCloudKey: testCloudValue},
							NodeLabels:      map[string]string{testNodeKey: testNodeValue},
							WorkerNodeCount: "1",
							TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
								Class:         testTestClassSpec,
								StorageClass:  testTestStorageSpec,
								FailureDomain: "",
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
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}(nil),
					nodePoolsKey: []interface{}{
						map[string]interface{}{
							infoKey: []interface{}{
								map[string]interface{}{
									nodepoolNameKey: testTest,
									descriptionKey:  testTestingTopologyFlattenFunction,
								},
							},
							specKey: []interface{}{
								map[string]interface{}{
									cloudLabelKey:      map[string]string{testCloudKey: testCloudValue},
									nodeLabelKey:       map[string]string{testNodeKey: testNodeValue},
									workerNodeCountKey: "1",
									tkgServiceVsphereKey: []interface{}{
										map[string]interface{}{
											classKey:         testTestClassSpec,
											storageClassKey:  testTestStorageSpec,
											failureDomainKey: "",
											volumesKey: []interface{}{
												map[string]interface{}{
													capacityKey:        float32(4),
													mountPathKey:       testVarLibEtcd,
													volumeNameKey:      testEtcd0,
													pvcStorageClassKey: testTkgsK8sObjPolicy,
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
					Class:            testTestClass,
					HighAvailability: false,
					StorageClass:     testTestStorageClass,
					Volumes: []*nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGServiceVsphereVolume{
						{
							Capacity:     4,
							MountPath:    testVarLibEtcd,
							Name:         testEtcd0,
							StorageClass: testTkgsK8sObjPolicy,
						},
					},
				},
				NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
					{
						Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
							Name:        testTest,
							Description: testTestingTopologyFlattenFunction,
						},
						Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
							CloudLabels:     map[string]string{testCloudKey: testCloudValue},
							NodeLabels:      map[string]string{testNodeKey: testNodeValue},
							WorkerNodeCount: "1",
							TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{
								Class:         testTestClassSpec,
								StorageClass:  testTestStorageSpec,
								FailureDomain: "domain-x50",
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
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}{
						map[string]interface{}{
							classKey:            testTestClass,
							highAvailabilityKey: false,
							storageClassKey:     testTestStorageClass,
							volumesKey: []interface{}{
								map[string]interface{}{
									capacityKey:        float32(4),
									mountPathKey:       testVarLibEtcd,
									volumeNameKey:      testEtcd0,
									pvcStorageClassKey: testTkgsK8sObjPolicy,
								},
							},
						},
					},
					nodePoolsKey: []interface{}{
						map[string]interface{}{
							infoKey: []interface{}{
								map[string]interface{}{
									nodepoolNameKey: testTest,
									descriptionKey:  testTestingTopologyFlattenFunction,
								},
							},
							specKey: []interface{}{
								map[string]interface{}{
									cloudLabelKey:      map[string]string{testCloudKey: testCloudValue},
									nodeLabelKey:       map[string]string{testNodeKey: testNodeValue},
									workerNodeCountKey: "1",
									tkgServiceVsphereKey: []interface{}{
										map[string]interface{}{
											classKey:         testTestClassSpec,
											storageClassKey:  testTestStorageSpec,
											failureDomainKey: "domain-x50",
											volumesKey: []interface{}{
												map[string]interface{}{
													capacityKey:        float32(4),
													mountPathKey:       testVarLibEtcd,
													volumeNameKey:      testEtcd0,
													pvcStorageClassKey: testTkgsK8sObjPolicy,
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
