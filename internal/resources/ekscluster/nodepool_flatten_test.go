package ekscluster

import (
	"testing"

	"github.com/stretchr/testify/require"

	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
)

func TestFlattenNodepools(t *testing.T) {
	tests := []struct {
		description string
		getInput    func() []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition
		expected    []interface{}
	}{
		{
			description: "nil list",
			getInput: func() []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
				return nil
			},
			expected: []interface{}{},
		},
		{
			description: "single nodepool",
			getInput: func() []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
				return []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
					getNodepoolDef("test-np"),
				}
			},
			expected: []interface{}{
				map[string]interface{}{
					"info": []interface{}{
						map[string]interface{}{
							"description": "test np",
							"name":        "test-np",
						},
					},
					"spec": []interface{}{
						map[string]interface{}{
							"ami_type": "CUSTOM",
							"ami_info": []interface{}{
								map[string]interface{}{
									"ami_id":                 "ami-2qu8409oisdfj0qw",
									"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
								},
							},
							"capacity_type": "ON_DEMAND",
							"instance_types": []string{
								"t3.medium",
								"m3.large",
							},
							"launch_template": []interface{}{
								map[string]interface{}{
									"id":      "",
									"name":    "templ",
									"version": "7",
								},
							},
							"node_labels": map[string]string{
								"key1": "val1",
							},
							"remote_access": []interface{}{
								map[string]interface{}{
									"security_groups": []string{
										"sg-0a6768722e9716768",
									},
									"ssh_key": "test-key",
								},
							},
							"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
							"root_disk_size": int32(20),
							"scaling_config": []interface{}{
								map[string]interface{}{
									"desired_size": int32(8),
									"max_size":     int32(16),
									"min_size":     int32(3),
								},
							},
							"subnet_ids": []string{
								"subnet-0a184f9301ae39a86",
								"subnet-0b495d7c212fc92a1",
								"subnet-0c86ec9ecde7b9bf7",
								"subnet-06497e6063c209f4d",
							},
							"tags": map[string]string{
								"tg1": "tv1",
							},
							"taints": []interface{}{
								map[string]interface{}{
									"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
									"key":    "tkey",
									"value":  "tvalue",
								},
							},
							"update_config": []interface{}{
								map[string]interface{}{
									"max_unavailable_nodes":      "10",
									"max_unavailable_percentage": "12",
								},
							},
							"release_version": "1.26.4-20230703",
						},
					},
				},
			},
		},
		{
			description: "multiple nodepool",
			getInput: func() []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
				return []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
					getNodepoolDef("test-np"),
					getNodepoolDef("test-np-2"),
				}
			},
			expected: []interface{}{
				map[string]interface{}{
					"info": []interface{}{
						map[string]interface{}{
							"description": "test np",
							"name":        "test-np",
						},
					},
					"spec": []interface{}{
						map[string]interface{}{
							"ami_type": "CUSTOM",
							"ami_info": []interface{}{
								map[string]interface{}{
									"ami_id":                 "ami-2qu8409oisdfj0qw",
									"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
								},
							},
							"capacity_type": "ON_DEMAND",
							"instance_types": []string{
								"t3.medium",
								"m3.large",
							},
							"launch_template": []interface{}{
								map[string]interface{}{
									"id":      "",
									"name":    "templ",
									"version": "7",
								},
							},
							"node_labels": map[string]string{
								"key1": "val1",
							},
							"remote_access": []interface{}{
								map[string]interface{}{
									"security_groups": []string{
										"sg-0a6768722e9716768",
									},
									"ssh_key": "test-key",
								},
							},
							"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
							"root_disk_size": int32(20),
							"scaling_config": []interface{}{
								map[string]interface{}{
									"desired_size": int32(8),
									"max_size":     int32(16),
									"min_size":     int32(3),
								},
							},
							"subnet_ids": []string{
								"subnet-0a184f9301ae39a86",
								"subnet-0b495d7c212fc92a1",
								"subnet-0c86ec9ecde7b9bf7",
								"subnet-06497e6063c209f4d",
							},
							"tags": map[string]string{
								"tg1": "tv1",
							},
							"taints": []interface{}{
								map[string]interface{}{
									"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
									"key":    "tkey",
									"value":  "tvalue",
								},
							},
							"update_config": []interface{}{
								map[string]interface{}{
									"max_unavailable_nodes":      "10",
									"max_unavailable_percentage": "12",
								},
							},
							"release_version": "1.26.4-20230703",
						},
					},
				},
				map[string]interface{}{
					"info": []interface{}{
						map[string]interface{}{
							"description": "test np",
							"name":        "test-np-2",
						},
					},
					"spec": []interface{}{
						map[string]interface{}{
							"ami_type": "CUSTOM",
							"ami_info": []interface{}{
								map[string]interface{}{
									"ami_id":                 "ami-2qu8409oisdfj0qw",
									"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
								},
							},
							"capacity_type": "ON_DEMAND",
							"instance_types": []string{
								"t3.medium",
								"m3.large",
							},
							"launch_template": []interface{}{
								map[string]interface{}{
									"id":      "",
									"name":    "templ",
									"version": "7",
								},
							},
							"node_labels": map[string]string{
								"key1": "val1",
							},
							"remote_access": []interface{}{
								map[string]interface{}{
									"security_groups": []string{
										"sg-0a6768722e9716768",
									},
									"ssh_key": "test-key",
								},
							},
							"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
							"root_disk_size": int32(20),
							"scaling_config": []interface{}{
								map[string]interface{}{
									"desired_size": int32(8),
									"max_size":     int32(16),
									"min_size":     int32(3),
								},
							},
							"subnet_ids": []string{
								"subnet-0a184f9301ae39a86",
								"subnet-0b495d7c212fc92a1",
								"subnet-0c86ec9ecde7b9bf7",
								"subnet-06497e6063c209f4d",
							},
							"tags": map[string]string{
								"tg1": "tv1",
							},
							"taints": []interface{}{
								map[string]interface{}{
									"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
									"key":    "tkey",
									"value":  "tvalue",
								},
							},
							"update_config": []interface{}{
								map[string]interface{}{
									"max_unavailable_nodes":      "10",
									"max_unavailable_percentage": "12",
								},
							},
							"release_version": "1.26.4-20230703",
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output := flattenNodePools(test.getInput())
			require.Equal(t, test.expected, output)
		})
	}
}

func TestFlattenNodepool(t *testing.T) {
	tests := []struct {
		description string
		getInput    func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition
		expected    map[string]interface{}
	}{
		{
			description: "nil np",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
				return nil
			},
			expected: map[string]interface{}{},
		},
		{
			description: "full nodepool",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
				return getNodepoolDef("test-np")
			},
			expected: map[string]interface{}{
				"info": []interface{}{
					map[string]interface{}{
						"description": "test np",
						"name":        "test-np",
					},
				},
				"spec": []interface{}{
					map[string]interface{}{
						"ami_type": "CUSTOM",
						"ami_info": []interface{}{
							map[string]interface{}{
								"ami_id":                 "ami-2qu8409oisdfj0qw",
								"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
							},
						},
						"capacity_type": "ON_DEMAND",
						"instance_types": []string{
							"t3.medium",
							"m3.large",
						},
						"launch_template": []interface{}{
							map[string]interface{}{
								"id":      "",
								"name":    "templ",
								"version": "7",
							},
						},
						"node_labels": map[string]string{
							"key1": "val1",
						},
						"remote_access": []interface{}{
							map[string]interface{}{
								"security_groups": []string{
									"sg-0a6768722e9716768",
								},
								"ssh_key": "test-key",
							},
						},
						"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
						"root_disk_size": int32(20),
						"scaling_config": []interface{}{
							map[string]interface{}{
								"desired_size": int32(8),
								"max_size":     int32(16),
								"min_size":     int32(3),
							},
						},
						"subnet_ids": []string{
							"subnet-0a184f9301ae39a86",
							"subnet-0b495d7c212fc92a1",
							"subnet-0c86ec9ecde7b9bf7",
							"subnet-06497e6063c209f4d",
						},
						"tags": map[string]string{
							"tg1": "tv1",
						},
						"taints": []interface{}{
							map[string]interface{}{
								"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
								"key":    "tkey",
								"value":  "tvalue",
							},
						},
						"update_config": []interface{}{
							map[string]interface{}{
								"max_unavailable_nodes":      "10",
								"max_unavailable_percentage": "12",
							},
						},
						"release_version": "1.26.4-20230703",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output := flattenNodePool(test.getInput())
			require.Equal(t, test.expected, output)
		})
	}
}

func TestFlattenNodepoolSpec(t *testing.T) {
	tests := []struct {
		description string
		getInput    func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec
		expected    []interface{}
	}{
		{
			description: "nil spec",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				return nil
			},
			expected: []interface{}{},
		},
		{
			description: "full spec",
			getInput:    getNodepoolSpec,
			expected: []interface{}{
				map[string]interface{}{
					"ami_type": "CUSTOM",
					"ami_info": []interface{}{
						map[string]interface{}{
							"ami_id":                 "ami-2qu8409oisdfj0qw",
							"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
						},
					},
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
					"release_version": "1.26.4-20230703",
				},
			},
		},
		{
			description: "launch template with id",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.LaunchTemplate = &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate{
					ID:      "lt-id",
					Name:    "",
					Version: "7",
				}
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type": "CUSTOM",
					"ami_info": []interface{}{
						map[string]interface{}{
							"ami_id":                 "ami-2qu8409oisdfj0qw",
							"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
						},
					},
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "lt-id",
							"name":    "",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
					"release_version": "1.26.4-20230703",
				},
			},
		},
		{
			description: "remote access is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.RemoteAccess = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type": "CUSTOM",
					"ami_info": []interface{}{
						map[string]interface{}{
							"ami_id":                 "ami-2qu8409oisdfj0qw",
							"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
						},
					},
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
					"release_version": "1.26.4-20230703",
				},
			},
		},
		{
			description: "root disk size is 0",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.RootDiskSize = 0
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type": "CUSTOM",
					"ami_info": []interface{}{
						map[string]interface{}{
							"ami_id":                 "ami-2qu8409oisdfj0qw",
							"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
						},
					},
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn": "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
					"release_version": "1.26.4-20230703",
				},
			},
		},
		{
			description: "scaling config is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.ScalingConfig = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type": "CUSTOM",
					"ami_info": []interface{}{
						map[string]interface{}{
							"ami_id":                 "ami-2qu8409oisdfj0qw",
							"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
						},
					},
					"capacity_type":  "ON_DEMAND",
					"root_disk_size": int32(20),
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn": "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
					"release_version": "1.26.4-20230703",
				},
			},
		},
		{
			description: "subnet ids are nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.SubnetIds = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type": "CUSTOM",
					"ami_info": []interface{}{
						map[string]interface{}{
							"ami_id":                 "ami-2qu8409oisdfj0qw",
							"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
						},
					},
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
					"release_version": "1.26.4-20230703",
				},
			},
		},
		{
			description: "taints are nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.Taints = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type": "CUSTOM",
					"ami_info": []interface{}{
						map[string]interface{}{
							"ami_id":                 "ami-2qu8409oisdfj0qw",
							"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
						},
					},
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
					"release_version": "1.26.4-20230703",
				},
			},
		},
		{
			description: "update config is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.UpdateConfig = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type": "CUSTOM",
					"ami_info": []interface{}{
						map[string]interface{}{
							"ami_id":                 "ami-2qu8409oisdfj0qw",
							"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
						},
					},
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"release_version": "1.26.4-20230703",
				},
			},
		},
		{
			description: "sg in remote access is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.RemoteAccess.SecurityGroups = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type": "CUSTOM",
					"ami_info": []interface{}{
						map[string]interface{}{
							"ami_id":                 "ami-2qu8409oisdfj0qw",
							"override_bootstrap_cmd": "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami",
						},
					},
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
					"release_version": "1.26.4-20230703",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output := flattenSpec(test.getInput())
			require.Equal(t, test.expected, output)
		})
	}
}

func getNodepoolDef(name string) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
	return &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
		Info: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo{
			Description: "test np",
			Name:        name,
		},
		Spec: getNodepoolSpec(),
	}
}
