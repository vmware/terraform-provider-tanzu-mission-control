// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgaws

import (
	"testing"

	"github.com/stretchr/testify/require"

	tkgawsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgaws"
)

func TestFlattenSettings(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings
		expected []interface{}
	}{
		{
			name:     "check for nil settings data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with network settings data",
			input: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings{
				Network: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkSettings{
					Cluster: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork{
						APIServerPort: 6443,
						Pods: []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange{
							{
								CidrBlocks: testPodsCIDRBlock,
							},
						},
						Services: []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange{
							{
								CidrBlocks: testServicesCIDRBlock,
							},
						},
					},
					Provider: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork{
						Subnets: []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet{
							{
								AvailabilityZone: testUsWest2aAvailabilityZone,
								CidrBlock:        testVpcCIDRBlock,
								ID:               testDefaultID,
								IsPublic:         false,
							},
						},
						Vpc: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsVPC{
							CidrBlock: testVpcCIDRBlock,
							ID:        testDefaultID,
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					networkKey: []interface{}{
						map[string]interface{}{
							clusterKey: []interface{}{
								map[string]interface{}{
									apiServerPortKey: int32(6443),
									podsKey: []interface{}{
										map[string]interface{}{
											cidrBlocksKey: testPodsCIDRBlock,
										},
									},
									servicesKey: []interface{}{
										map[string]interface{}{
											cidrBlocksKey: testServicesCIDRBlock,
										},
									},
								},
							},
							providerKey: []interface{}{
								map[string]interface{}{
									subnetsKey: []interface{}{
										map[string]interface{}{
											availabilityZoneKey: testUsWest2aAvailabilityZone,
											subnetCIDRBlockKey:  testVpcCIDRBlock,
											subnetIDKey:         testDefaultID,
											isPublicKey:         false,
										},
									},
									vpcKey: []interface{}{
										map[string]interface{}{
											vpcCIDRBlockKey: testVpcCIDRBlock,
											vpcIDKey:        testDefaultID,
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
			name: "normal scenario with security settings data",
			input: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings{
				Security: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSecuritySettings{
					SSHKey: testJumperSSHKey,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					securityKey: []interface{}{
						map[string]interface{}{
							sshKey: testJumperSSHKey,
						},
					},
				},
			},
		},
		{
			name: "normal scenario with TKGm AWS workload settings data",
			input: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings{
				Network: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkSettings{
					Cluster: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork{
						APIServerPort: 6443,
						Pods: []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange{
							{
								CidrBlocks: testPodsCIDRBlock,
							},
						},
						Services: []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange{
							{
								CidrBlocks: testServicesCIDRBlock,
							},
						},
					},
					Provider: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork{
						Subnets: []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet{
							{
								AvailabilityZone: testUsWest2aAvailabilityZone,
								CidrBlock:        testVpcCIDRBlock,
								ID:               testDefaultID,
								IsPublic:         false,
							},
						},
						Vpc: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsVPC{
							CidrBlock: testVpcCIDRBlock,
							ID:        testDefaultID,
						},
					},
				},
				Security: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSecuritySettings{
					SSHKey: testJumperSSHKey,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					networkKey: []interface{}{
						map[string]interface{}{
							clusterKey: []interface{}{
								map[string]interface{}{
									apiServerPortKey: int32(6443),
									podsKey: []interface{}{
										map[string]interface{}{
											cidrBlocksKey: testPodsCIDRBlock,
										},
									},
									servicesKey: []interface{}{
										map[string]interface{}{
											cidrBlocksKey: testServicesCIDRBlock,
										},
									},
								},
							},
							providerKey: []interface{}{
								map[string]interface{}{
									subnetsKey: []interface{}{
										map[string]interface{}{
											availabilityZoneKey: testUsWest2aAvailabilityZone,
											subnetCIDRBlockKey:  testVpcCIDRBlock,
											subnetIDKey:         testDefaultID,
											isPublicKey:         false,
										},
									},
									vpcKey: []interface{}{
										map[string]interface{}{
											vpcCIDRBlockKey: testVpcCIDRBlock,
											vpcIDKey:        testDefaultID,
										},
									},
								},
							},
						},
					},
					securityKey: []interface{}{
						map[string]interface{}{
							sshKey: testJumperSSHKey,
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := flattenTKGAWSSettings(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
