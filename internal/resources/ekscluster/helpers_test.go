/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"testing"

	"github.com/stretchr/testify/require"

	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
)

func TestNodepoolSpecEqual(t *testing.T) {
	tests := []struct {
		name        string
		modifySpec1 func(*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec)
		modifySpec2 func(*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec)
		result      bool
	}{
		{
			name:        "both are equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			result:      true,
		},
		{
			name:        "ami types are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.AmiType = "new-ami"
			},
			result: false,
		},
		{
			name:        "capacity types are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.CapacityType = "new-cp"
			},
			result: false,
		},
		{
			name:        "instance types are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.InstanceTypes = []string{
					"type1",
					"type2",
				}
			},
			result: false,
		},
		{
			name: "instance types are not set",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec1.InstanceTypes = nil
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.InstanceTypes = nil
			},
			result: true,
		},
		{
			name: "instance types are set equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.InstanceTypes = []string{
					"m3.large",
					"t3.medium",
				}
			},
			result: true,
		},
		{
			name:        "launch templates are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.LaunchTemplate = &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate{
					ID:      "lt-kjasdf9ui",
					Version: "6",
				}
			},
			result: false,
		},
		{
			name: "launch template for 1 is not set and the other is empty",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec1.LaunchTemplate = nil
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.LaunchTemplate = &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate{}
			},
			result: true,
		},
		{
			name: "launch template for 1 is empty and the other is not set",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec1.LaunchTemplate = &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate{}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.LaunchTemplate = nil
			},
			result: true,
		},
		{
			name:        "node labels are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.NodeLabels = map[string]string{
					"key2": "val2",
				}
			},
			result: false,
		},
		{
			name:        "remote accesses' SSH keys are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.RemoteAccess.SSHKey = "new-key"
			},
			result: false,
		},
		{
			name:        "remote accesses' security grps are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.RemoteAccess.SecurityGroups = []string{"sg-2"}
			},
			result: false,
		},
		{
			name: "remote accesses' security grps are set equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec1.RemoteAccess.SecurityGroups = []string{"sg-1", "sg-2", "sg-3"}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.RemoteAccess.SecurityGroups = []string{"sg-3", "sg-1", "sg-2"}
			},
			result: true,
		},
		{
			name: "first remote access is nil and the other empty",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec1.RemoteAccess = nil
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.RemoteAccess = &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess{}
			},
			result: true,
		},
		{
			name: "first remote access is empty and the other nil",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec1.RemoteAccess = &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess{}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.RemoteAccess = nil
			},
			result: true,
		},
		{
			name:        "role arns are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.RoleArn = "arn-2"
			},
			result: false,
		},
		{
			name:        "root disk sizes are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.RootDiskSize = 200
			},
			result: false,
		},
		{
			name:        "scaling config are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.ScalingConfig.MaxSize = 20
			},
			result: false,
		},
		{
			name:        "subnet IDs are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.SubnetIds = []string{
					"subnet-1",
					"subnet-2",
				}
			},
			result: false,
		},
		{
			name: "subnet IDs are not set",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec1.SubnetIds = nil
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.SubnetIds = nil
			},
			result: true,
		},
		{
			name: "subnet IDs are set equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec1.SubnetIds = []string{
					"subnet-1",
					"subnet-2",
					"subnet-3",
				}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.SubnetIds = []string{
					"subnet-3",
					"subnet-1",
					"subnet-2",
				}
			},
			result: true,
		},
		{
			name:        "tags are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.Tags = map[string]string{
					"tag2": "val2",
				}
			},
			result: false,
		},
		{
			name:        "taints are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.Taints = []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint{
					{
						Effect: eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
						Key:    "tkey2",
						Value:  "tvalue2",
					},
				}
			},
			result: false,
		},
		{
			name: "taints are not set",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec1.Taints = nil
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.Taints = nil
			},
			result: true,
		},
		{
			name: "taints are set equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec1.Taints = []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint{
					{
						Effect: eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOEXECUTE),
						Key:    "tkey1",
						Value:  "tvalue1",
					},
					{
						Effect: eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
						Key:    "tkey2",
						Value:  "tvalue2",
					},
					{
						Effect: eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOSCHEDULE),
						Key:    "tkey3",
						Value:  "tvalue3",
					},
				}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.Taints = []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint{
					{
						Effect: eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOSCHEDULE),
						Key:    "tkey3",
						Value:  "tvalue3",
					},
					{
						Effect: eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOEXECUTE),
						Key:    "tkey1",
						Value:  "tvalue1",
					},
					{
						Effect: eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
						Key:    "tkey2",
						Value:  "tvalue2",
					},
				}
			},
			result: true,
		},
		{
			name:        "updated config are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.UpdateConfig.MaxUnavailableNodes = "20"
			},
			result: false,
		},
		{
			name:        "ami release versions are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
				spec2.ReleaseVersion = "1.26.4-20230728"
			},
			result: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			spec1 := getNodepoolSpec()
			spec2 := getNodepoolSpec()
			test.modifySpec1(spec1)
			test.modifySpec2(spec2)

			require.Equal(t, test.result, nodepoolSpecEqual(spec1, spec2), "return didn't match the expected output")
		})
	}
}

func TestClusterSpecEqual(t *testing.T) {
	tests := []struct {
		name        string
		modifySpec1 func(*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec)
		modifySpec2 func(*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec)
		result      bool
	}{
		{
			name:        "both are equal",
			modifySpec1: func(*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {},
			modifySpec2: func(*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {},
			result:      true,
		},
		{
			name:        "cluster group names are not equal",
			modifySpec1: func(*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.ClusterGroupName = "cg-2"
			},
		},
		{
			name:        "proxies are not equal",
			modifySpec1: func(*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.ProxyName = "p2"
			},
			result: false,
		},
		{
			name: "config1 is nil",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config = nil
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {},
			result:      false,
		},
		{
			name:        "config2 is nil",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config = nil
			},
			result: false,
		},
		{
			name: "both config are nil",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config = nil
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config = nil
			},
			result: true,
		},
		{
			name: "k8s network configs are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.KubernetesNetworkConfig.ServiceCidr = "192.100.0.0/16"
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
			},
			result: false,
		},
		{
			name: "loggings are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Logging.Audit = true
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Logging.Audit = false
			},
			result: false,
		},
		{
			name: "role arns are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.RoleArn = "arn:arn1"
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.RoleArn = "arn:arn2"
			},
			result: false,
		},
		{
			name: "tags are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Tags = map[string]string{
					"abc": "def",
					"pqr": "stu",
				}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Tags = map[string]string{
					"ghi": "jkl",
					"mno": "pqr",
				}
			},
			result: false,
		},
		{
			name: "versions are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Version = "1.23"
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Version = "1.24"
			},
			result: false,
		},
		{
			name: "vpc1 is nil",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Vpc = nil
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {},
			result:      false,
		},
		{
			name:        "vpc2 is nil",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Vpc = nil
			},
			result: false,
		},
		{
			name: "both vpcs are nil",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Vpc = nil
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Vpc = nil
			},
			result: true,
		},
		{
			name: "vpc private accesses are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Vpc.EnablePrivateAccess = true
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Vpc.EnablePrivateAccess = false
			},
			result: false,
		},
		{
			name: "vpc public accesses are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Vpc.EnablePublicAccess = true
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Vpc.EnablePublicAccess = false
			},
			result: false,
		},
		{
			name: "public cirds are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Vpc.PublicAccessCidrs = []string{
					"100.10.1.0/24",
					"196.19.1.0/24",
				}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Vpc.PublicAccessCidrs = []string{
					"100.10.2.0/24",
					"196.19.2.0/24",
				}
			},
			result: false,
		},
		{
			name: "public cirds are set equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Vpc.PublicAccessCidrs = []string{
					"100.10.1.0/24",
					"196.19.1.0/24",
				}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Vpc.PublicAccessCidrs = []string{
					"196.19.1.0/24",
					"100.10.1.0/24",
				}
			},
			result: true,
		},
		{
			name: "security groups are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Vpc.SecurityGroups = []string{
					"sg-1",
					"sg-2",
				}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Vpc.SecurityGroups = []string{
					"sg-3",
					"sg-4",
				}
			},
			result: false,
		},
		{
			name: "security groups are set equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Vpc.SecurityGroups = []string{
					"sg-1",
					"sg-2",
				}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Vpc.SecurityGroups = []string{
					"sg-2",
					"sg-1",
				}
			},
			result: true,
		},
		{
			name: "subnets are not equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Vpc.SubnetIds = []string{
					"subnet-1",
					"subnet-2",
				}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Vpc.SubnetIds = []string{
					"subnet-3",
					"subnet-4",
				}
			},
			result: false,
		},
		{
			name: "subnets are set equal",
			modifySpec1: func(spec1 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec1.Config.Vpc.SubnetIds = []string{
					"subnet-1",
					"subnet-2",
				}
			},
			modifySpec2: func(spec2 *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
				spec2.Config.Vpc.SubnetIds = []string{
					"subnet-2",
					"subnet-1",
				}
			},
			result: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			spec1, _ := getClusterSpec()
			spec2, _ := getClusterSpec()
			test.modifySpec1(spec1)
			test.modifySpec2(spec2)

			require.Equal(t, test.result, clusterSpecEqual(spec1, spec2), "return didn't match the expected output")
		})
	}
}

func getNodepoolSpec() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
	return &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec{
		AmiType: "CUSTOM",
		AmiInfo: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo{
			AmiID:                "ami-2qu8409oisdfj0qw",
			OverrideBootstrapCmd: "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
		},
		CapacityType: "ON_DEMAND",
		InstanceTypes: []string{
			"t3.medium",
			"m3.large",
		},
		LaunchTemplate: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate{
			Name:    "templ",
			Version: "7",
		},
		NodeLabels: map[string]string{
			"key1": "val1",
		},
		RemoteAccess: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess{
			SecurityGroups: []string{
				"sg-0a6768722e9716768",
			},
			SSHKey: "test-key",
		},
		RoleArn:      "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
		RootDiskSize: 20,
		ScalingConfig: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig{
			DesiredSize: 8,
			MaxSize:     16,
			MinSize:     3,
		},
		SubnetIds: []string{
			"subnet-0a184f9301ae39a86",
			"subnet-0b495d7c212fc92a1",
			"subnet-0c86ec9ecde7b9bf7",
			"subnet-06497e6063c209f4d",
		},
		Tags: map[string]string{
			"tg1": "tv1",
		},
		Taints: []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint{
			{
				Effect: eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
				Key:    "tkey",
				Value:  "tvalue",
			},
		},
		UpdateConfig: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig{
			MaxUnavailableNodes:      "10",
			MaxUnavailablePercentage: "12",
		},
		ReleaseVersion: "1.26.4-20230703",
	}
}
