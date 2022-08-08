/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgaws

import (
	"testing"

	"github.com/stretchr/testify/require"

	nodepoolmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	tkgawsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgaws"
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
					AvailabilityZones: []string{"us-east"},
					InstanceType:      "m5.large",
					HighAvailability:  false,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}{
						map[string]interface{}{
							availabilityZonesKey: []string{"us-east"},
							instanceTypeKey:      "m5.large",
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
							Name:        "test",
							Description: "testing topology flatten function",
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
									tkgAWSKey: []interface{}{
										map[string]interface{}{
											nodepoolAvailabilityZoneKey: "us-west-2a",
											nodepoolInstanceTypeKey:     "m5.large",
											nodePlacementKey: []interface{}{
												map[string]interface{}{
													awsAvailabilityZoneKey: "us-west-2a",
												},
											},
											nodePoolSubnetIDKey: "default",
											nodepoolVersionKey:  "v1.21.2+vmware.1",
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
					AvailabilityZones: []string{"us-east"},
					InstanceType:      "m5.large",
					HighAvailability:  false,
				},
				NodePools: []*nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
					{
						Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{
							Name:        "test",
							Description: "testing topology flatten function",
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
			expected: []interface{}{
				map[string]interface{}{
					controlPlaneKey: []interface{}{
						map[string]interface{}{
							availabilityZonesKey: []string{"us-east"},
							instanceTypeKey:      "m5.large",
							highAvailabilityKey:  false,
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
									tkgAWSKey: []interface{}{
										map[string]interface{}{
											nodepoolAvailabilityZoneKey: "us-west-2a",
											nodepoolInstanceTypeKey:     "m5.large",
											nodePlacementKey: []interface{}{
												map[string]interface{}{
													awsAvailabilityZoneKey: "us-west-2a",
												},
											},
											nodePoolSubnetIDKey: "default",
											nodepoolVersionKey:  "v1.21.2+vmware.1",
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
