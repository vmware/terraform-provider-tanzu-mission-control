// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgaws

import (
	"testing"

	"github.com/stretchr/testify/require"

	nodepoolmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	tkgawsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgaws"
)

const (
	testM5LargeInstanceType      = "m5.large"
	testUSEastAvailabilityZone   = "us-east"
	testUsWest2aAvailabilityZone = "us-west-2a"
	testDefaultID                = "default"
	testVpcCIDRBlock             = "10.0.0.0/16"
	testPodsCIDRBlock            = "100.96.0.0/11"
	testServicesCIDRBlock        = "100.64.0.0/13"
	testJumperSSHKey             = "jumper_ssh_key-sh-1529663-220321-074908"

	testTest                           = "test"
	testTestingTopologyFlattenFunction = "testing topology flatten function"
	testV1212vmware1                   = "v1.21.2+vmware.1"
)

func TestFlattenTopology(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology
		expected []interface{}
	}{
		{
			name:     "check for nil topology data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with empty node pool data of topology",
			input: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
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
			input: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
				ControlPlane: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsControlPlane{
					AvailabilityZones: []string{testUSEastAvailabilityZone},
					InstanceType:      testM5LargeInstanceType,
					HighAvailability:  false,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}{
						map[string]interface{}{
							availabilityZonesKey: []string{testUSEastAvailabilityZone},
							instanceTypeKey:      testM5LargeInstanceType,
							highAvailabilityKey:  false,
						},
					},
					nodePoolsKey: []interface{}{},
				},
			},
		},
		{
			name: "normal scenario with node pool data of topology",
			input: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
				NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
					{
						Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
							Name:        testTest,
							Description: testTestingTopologyFlattenFunction,
						},
						Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
							WorkerNodeCount: "1",
							TkgAws: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool{
								AvailabilityZone: testUsWest2aAvailabilityZone,
								InstanceType:     testM5LargeInstanceType,
								NodePlacement: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement{
									{
										AvailabilityZone: testUsWest2aAvailabilityZone,
									},
								},
								SubnetID: testDefaultID,
								Version:  testV1212vmware1,
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
									nodePoolNameKey:        testTest,
									nodePoolDescriptionKey: testTestingTopologyFlattenFunction,
								},
							},
							nodePoolSpecKey: []interface{}{
								map[string]interface{}{
									workerNodeCountKey: "1",
									tkgAWSKey: []interface{}{
										map[string]interface{}{
											nodepoolAvailabilityZoneKey: testUsWest2aAvailabilityZone,
											nodepoolInstanceTypeKey:     testM5LargeInstanceType,
											nodePlacementKey: []interface{}{
												map[string]interface{}{
													awsAvailabilityZoneKey: testUsWest2aAvailabilityZone,
												},
											},
											nodePoolSubnetIDKey: testDefaultID,
											nodepoolVersionKey:  testV1212vmware1,
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
			input: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{
				ControlPlane: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsControlPlane{
					AvailabilityZones: []string{testUSEastAvailabilityZone},
					InstanceType:      testM5LargeInstanceType,
					HighAvailability:  false,
				},
				NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
					{
						Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
							Name:        testTest,
							Description: testTestingTopologyFlattenFunction,
						},
						Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
							WorkerNodeCount: "1",
							TkgAws: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool{
								AvailabilityZone: testUsWest2aAvailabilityZone,
								InstanceType:     testM5LargeInstanceType,
								NodePlacement: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement{
									{
										AvailabilityZone: testUsWest2aAvailabilityZone,
									},
								},
								SubnetID: testDefaultID,
								Version:  testV1212vmware1,
							},
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}{
						map[string]interface{}{
							availabilityZonesKey: []string{testUSEastAvailabilityZone},
							instanceTypeKey:      testM5LargeInstanceType,
							highAvailabilityKey:  false,
						},
					},
					nodePoolsKey: []interface{}{
						map[string]interface{}{
							nodePoolInfoKey: []interface{}{
								map[string]interface{}{
									nodePoolNameKey:        testTest,
									nodePoolDescriptionKey: testTestingTopologyFlattenFunction,
								},
							},
							nodePoolSpecKey: []interface{}{
								map[string]interface{}{
									workerNodeCountKey: "1",
									tkgAWSKey: []interface{}{
										map[string]interface{}{
											nodepoolAvailabilityZoneKey: testUsWest2aAvailabilityZone,
											nodepoolInstanceTypeKey:     testM5LargeInstanceType,
											nodePlacementKey: []interface{}{
												map[string]interface{}{
													awsAvailabilityZoneKey: testUsWest2aAvailabilityZone,
												},
											},
											nodePoolSubnetIDKey: testDefaultID,
											nodepoolVersionKey:  testV1212vmware1,
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
			actual := flattenTKGAWSTopology(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
