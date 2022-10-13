/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

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
								CidrBlocks: "100.96.0.0/11",
							},
						},
						Services: []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange{
							{
								CidrBlocks: "100.64.0.0/13",
							},
						},
					},
					Provider: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork{
						Subnets: []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet{
							{
								AvailabilityZone: "us-west-2a",
								CidrBlock:        "10.0.0.0/16",
								ID:               "default",
								IsPublic:         false,
							},
						},
						Vpc: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsVPC{
							CidrBlock: "10.0.0.0/16",
							ID:        "default",
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
											cidrBlocksKey: "100.96.0.0/11",
										},
									},
									servicesKey: []interface{}{
										map[string]interface{}{
											cidrBlocksKey: "100.64.0.0/13",
										},
									},
								},
							},
							providerKey: []interface{}{
								map[string]interface{}{
									subnetsKey: []interface{}{
										map[string]interface{}{
											availabilityZoneKey: "us-west-2a",
											subnetCIDRBlockKey:  "10.0.0.0/16",
											subnetIDKey:         "default",
											isPublicKey:         false,
										},
									},
									vpcKey: []interface{}{
										map[string]interface{}{
											vpcCIDRBlockKey: "10.0.0.0/16",
											vpcIDKey:        "default",
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
					SSHKey: "jumper_ssh_key-sh-1529663-220321-074908",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					securityKey: []interface{}{
						map[string]interface{}{
							sshKey: "jumper_ssh_key-sh-1529663-220321-074908",
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
								CidrBlocks: "100.96.0.0/11",
							},
						},
						Services: []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange{
							{
								CidrBlocks: "100.64.0.0/13",
							},
						},
					},
					Provider: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork{
						Subnets: []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet{
							{
								AvailabilityZone: "us-west-2a",
								CidrBlock:        "10.0.0.0/16",
								ID:               "default",
								IsPublic:         false,
							},
						},
						Vpc: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsVPC{
							CidrBlock: "10.0.0.0/16",
							ID:        "default",
						},
					},
				},
				Security: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSecuritySettings{
					SSHKey: "jumper_ssh_key-sh-1529663-220321-074908",
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
											cidrBlocksKey: "100.96.0.0/11",
										},
									},
									servicesKey: []interface{}{
										map[string]interface{}{
											cidrBlocksKey: "100.64.0.0/13",
										},
									},
								},
							},
							providerKey: []interface{}{
								map[string]interface{}{
									subnetsKey: []interface{}{
										map[string]interface{}{
											availabilityZoneKey: "us-west-2a",
											subnetCIDRBlockKey:  "10.0.0.0/16",
											subnetIDKey:         "default",
											isPublicKey:         false,
										},
									},
									vpcKey: []interface{}{
										map[string]interface{}{
											vpcCIDRBlockKey: "10.0.0.0/16",
											vpcIDKey:        "default",
										},
									},
								},
							},
						},
					},
					securityKey: []interface{}{
						map[string]interface{}{
							sshKey: "jumper_ssh_key-sh-1529663-220321-074908",
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
