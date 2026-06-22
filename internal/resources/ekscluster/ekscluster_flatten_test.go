// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package ekscluster

import (
	"testing"

	"github.com/stretchr/testify/require"

	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
)

func TestFlattenClusterSpec(t *testing.T) {
	tests := []struct {
		description string
		getInput    func() (*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec, []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition)
		expected    []interface{}
	}{
		{
			description: "nil spec",
			getInput: func() (*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec, []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) {
				return nil, nil
			},
			expected: []interface{}{},
		},
		{
			description: "full cluster spec without addonsconfig",
			getInput: func() (*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec, []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) {
				spec, nps := getClusterSpec()
				spec.Config.AddonsConfig = nil

				return spec, nps
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey: testTestCg,
					proxyNameKey:    testTestProoxy,
					configKey: []interface{}{
						map[string]interface{}{
							kubernetesNetworkConfigKey: []interface{}{
								map[string]interface{}{
									serviceCidrKey: test1000010,
								},
							},

							kubernetesVersionKey: test112,
							loggingKey: []interface{}{
								map[string]interface{}{
									apiServerKey:         false,
									auditKey:             false,
									authenticatorKey:     false,
									controllerManagerKey: true,
									schedulerKey:         true,
								},
							},
							roleArnKey: testRoleArn,
							tagsKey: map[string]string{
								testTag1: testTag2,
							},
							vpcKey: []interface{}{
								map[string]interface{}{
									enablePrivateAccessKey: false,
									enablePublicAccessKey:  false,
									publicAccessCidrsKey: []string{
										test00001,
										test10001,
									},
									securityGroupsKey: []string{
										testSg1,
										testSg2,
									},
									subnetIdsKey: []string{
										testSubnet1,
										testSubnet2,
									},
								},
							},
						},
					},
					"nodepool": []interface{}{
						map[string]interface{}{
							infoKey: []interface{}{
								map[string]interface{}{
									testDescription: testTestNp,
									NameKey:         testTestNp2,
								},
							},
							specKey: []interface{}{
								map[string]interface{}{
									amiTypeKey: testCustom,
									amiInfoKey: []interface{}{
										map[string]interface{}{
											amiIDKey:                testAmi2qu8409oisdfj0qw,
											overrideBootstrapCmdKey: testBootstrapCmd,
										},
									},
									capacityTypeKey: testOnDemand,
									instanceTypesKey: []string{
										testT3medium,
										testM3large,
									},
									launchTemplateKey: []interface{}{
										map[string]interface{}{
											"id":       "",
											NameKey:    testTempl,
											versionKey: "7",
										},
									},
									nodeLabelsKey: map[string]string{
										testKey1: testVal1,
									},
									remoteAccessKey: []interface{}{
										map[string]interface{}{
											securityGroupsKey: []string{
												testSg0a6768722e9716768,
											},
											sshKeyKey: testTestKey,
										},
									},
									roleArnKey:      testArnAwsIam000000000000,
									rootDiskSizeKey: int32(20),
									scalingConfigKey: []interface{}{
										map[string]interface{}{
											desiredSizeKey: int32(8),
											maxSizeKey:     int32(16),
											minSizeKey:     int32(3),
										},
									},
									subnetIdsKey: []string{
										testSubnet0a184f9301ae39a86,
										testSubnet0b495d7c212fc92a1,
										testSubnet0c86ec9ecde7b9bf7,
										testSubnet06497e6063c209f4d,
									},
									tagsKey: map[string]string{
										testTg1: testTv1,
									},
									taintsKey: []interface{}{
										map[string]interface{}{
											effectKey: eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
											keyKey:    testTkey,
											valueKey:  testTvalue,
										},
									},
									updateConfigKey: []interface{}{
										map[string]interface{}{
											maxUnavailableNodesKey:      "10",
											maxUnavailablePercentageKey: "12",
										},
									},
									releaseVersionKey: test126420230703,
								},
							},
						},
					},
				},
			},
		},
		{
			description: "full cluster spec with addonsconfig",
			getInput: func() (*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec, []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) {
				spec, nps := getClusterSpec()
				return spec, nps
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey: testTestCg,
					proxyNameKey:    testTestProoxy,
					configKey: []interface{}{
						map[string]interface{}{
							kubernetesNetworkConfigKey: []interface{}{
								map[string]interface{}{
									serviceCidrKey: test1000010,
								},
							},

							kubernetesVersionKey: test112,
							loggingKey: []interface{}{
								map[string]interface{}{
									apiServerKey:         false,
									auditKey:             false,
									authenticatorKey:     false,
									controllerManagerKey: true,
									schedulerKey:         true,
								},
							},
							roleArnKey: testRoleArn,
							tagsKey: map[string]string{
								testTag1: testTag2,
							},
							vpcKey: []interface{}{
								map[string]interface{}{
									enablePrivateAccessKey: false,
									enablePublicAccessKey:  false,
									publicAccessCidrsKey: []string{
										test00001,
										test10001,
									},
									securityGroupsKey: []string{
										testSg1,
										testSg2,
									},
									subnetIdsKey: []string{
										testSubnet1,
										testSubnet2,
									},
								},
							},
							"addons_config": []interface{}{
								map[string]interface{}{
									"vpc_cni_config": []interface{}{
										map[string]interface{}{
											"eni_config": []interface{}{
												map[string]interface{}{
													"id": testSubnet1,
													securityGroupsKey: []string{
														testSg1,
														testSg2,
													},
												},
												map[string]interface{}{
													"id": testSubnet2,
													securityGroupsKey: []string{
														testSg3,
														testSg4,
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"nodepool": []interface{}{
						map[string]interface{}{
							infoKey: []interface{}{
								map[string]interface{}{
									testDescription: testTestNp,
									NameKey:         testTestNp2,
								},
							},
							specKey: []interface{}{
								map[string]interface{}{
									amiTypeKey: testCustom,
									amiInfoKey: []interface{}{
										map[string]interface{}{
											amiIDKey:                testAmi2qu8409oisdfj0qw,
											overrideBootstrapCmdKey: testBootstrapCmd,
										},
									},
									capacityTypeKey: testOnDemand,
									instanceTypesKey: []string{
										testT3medium,
										testM3large,
									},
									launchTemplateKey: []interface{}{
										map[string]interface{}{
											"id":       "",
											NameKey:    testTempl,
											versionKey: "7",
										},
									},
									nodeLabelsKey: map[string]string{
										testKey1: testVal1,
									},
									remoteAccessKey: []interface{}{
										map[string]interface{}{
											securityGroupsKey: []string{
												testSg0a6768722e9716768,
											},
											sshKeyKey: testTestKey,
										},
									},
									roleArnKey:      testArnAwsIam000000000000,
									rootDiskSizeKey: int32(20),
									scalingConfigKey: []interface{}{
										map[string]interface{}{
											desiredSizeKey: int32(8),
											maxSizeKey:     int32(16),
											minSizeKey:     int32(3),
										},
									},
									subnetIdsKey: []string{
										testSubnet0a184f9301ae39a86,
										testSubnet0b495d7c212fc92a1,
										testSubnet0c86ec9ecde7b9bf7,
										testSubnet06497e6063c209f4d,
									},
									tagsKey: map[string]string{
										testTg1: testTv1,
									},
									taintsKey: []interface{}{
										map[string]interface{}{
											effectKey: eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
											keyKey:    testTkey,
											valueKey:  testTvalue,
										},
									},
									updateConfigKey: []interface{}{
										map[string]interface{}{
											maxUnavailableNodesKey:      "10",
											maxUnavailablePercentageKey: "12",
										},
									},
									releaseVersionKey: test126420230703,
								},
							},
						},
					},
				},
			},
		},
		{
			description: "empty nodepools",
			getInput: func() (*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec, []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) {
				spec, _ := getClusterSpec()
				spec.Config.AddonsConfig = nil

				return spec, nil
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey: testTestCg,
					proxyNameKey:    testTestProoxy,
					configKey: []interface{}{
						map[string]interface{}{
							kubernetesNetworkConfigKey: []interface{}{
								map[string]interface{}{
									serviceCidrKey: test1000010,
								},
							},

							kubernetesVersionKey: test112,
							loggingKey: []interface{}{
								map[string]interface{}{
									apiServerKey:         false,
									auditKey:             false,
									authenticatorKey:     false,
									controllerManagerKey: true,
									schedulerKey:         true,
								},
							},
							roleArnKey: testRoleArn,
							tagsKey: map[string]string{
								testTag1: testTag2,
							},
							vpcKey: []interface{}{
								map[string]interface{}{
									enablePrivateAccessKey: false,
									enablePublicAccessKey:  false,
									publicAccessCidrsKey: []string{
										test00001,
										test10001,
									},
									securityGroupsKey: []string{
										testSg1,
										testSg2,
									},
									subnetIdsKey: []string{
										testSubnet1,
										testSubnet2,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			description: "empty proxy",
			getInput: func() (*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec, []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) {
				spec, _ := getClusterSpec()
				spec.ProxyName = ""
				spec.Config.AddonsConfig = nil

				return spec, nil
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey: testTestCg,
					configKey: []interface{}{
						map[string]interface{}{
							kubernetesNetworkConfigKey: []interface{}{
								map[string]interface{}{
									serviceCidrKey: test1000010,
								},
							},

							kubernetesVersionKey: test112,
							loggingKey: []interface{}{
								map[string]interface{}{
									apiServerKey:         false,
									auditKey:             false,
									authenticatorKey:     false,
									controllerManagerKey: true,
									schedulerKey:         true,
								},
							},
							roleArnKey: testRoleArn,
							tagsKey: map[string]string{
								testTag1: testTag2,
							},
							vpcKey: []interface{}{
								map[string]interface{}{
									enablePrivateAccessKey: false,
									enablePublicAccessKey:  false,
									publicAccessCidrsKey: []string{
										test00001,
										test10001,
									},
									securityGroupsKey: []string{
										testSg1,
										testSg2,
									},
									subnetIdsKey: []string{
										testSubnet1,
										testSubnet2,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			description: "empty config",
			getInput: func() (*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec, []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) {
				spec, _ := getClusterSpec()
				spec.ProxyName = ""
				spec.Config = nil

				return spec, nil
			},
			expected: []interface{}{
				map[string]interface{}{
					clusterGroupKey: testTestCg,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			spec, nps := test.getInput()
			output := flattenClusterSpec(spec, nps)

			print(test.expected)
			print(output)
			require.Equal(t, test.expected, output)
		})
	}
}

func TestFlattenConfig(t *testing.T) {
	tests := []struct {
		description string
		getInput    func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig
		expected    []interface{}
	}{
		{
			description: "nil config",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
				return nil
			},
			expected: []interface{}{},
		},
		{
			description: "full config",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
				config := getConfig()
				config.AddonsConfig = nil

				return config
			},
			expected: []interface{}{
				map[string]interface{}{
					kubernetesNetworkConfigKey: []interface{}{
						map[string]interface{}{
							serviceCidrKey: test1000010,
						},
					},
					kubernetesVersionKey: test112,
					loggingKey: []interface{}{
						map[string]interface{}{
							apiServerKey:         false,
							auditKey:             false,
							authenticatorKey:     false,
							controllerManagerKey: true,
							schedulerKey:         true,
						},
					},
					roleArnKey: testRoleArn,
					tagsKey: map[string]string{
						testTag1: testTag2,
					},
					vpcKey: []interface{}{
						map[string]interface{}{
							enablePrivateAccessKey: false,
							enablePublicAccessKey:  false,
							publicAccessCidrsKey: []string{
								test00001,
								test10001,
							},
							securityGroupsKey: []string{
								testSg1,
								testSg2,
							},
							subnetIdsKey: []string{
								testSubnet1,
								testSubnet2,
							},
						},
					},
				},
			},
		},
		{
			description: "k8s network config is nil config",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
				config := getConfig()
				config.KubernetesNetworkConfig = nil
				config.AddonsConfig = nil

				return config
			},
			expected: []interface{}{
				map[string]interface{}{
					kubernetesVersionKey: test112,
					loggingKey: []interface{}{
						map[string]interface{}{
							apiServerKey:         false,
							auditKey:             false,
							authenticatorKey:     false,
							controllerManagerKey: true,
							schedulerKey:         true,
						},
					},
					roleArnKey: testRoleArn,
					tagsKey: map[string]string{
						testTag1: testTag2,
					},
					vpcKey: []interface{}{
						map[string]interface{}{
							enablePrivateAccessKey: false,
							enablePublicAccessKey:  false,
							publicAccessCidrsKey: []string{
								test00001,
								test10001,
							},
							securityGroupsKey: []string{
								testSg1,
								testSg2,
							},
							subnetIdsKey: []string{
								testSubnet1,
								testSubnet2,
							},
						},
					},
				},
			},
		},
		{
			description: "logging is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
				config := getConfig()
				config.Logging = nil
				config.AddonsConfig = nil

				return config
			},
			expected: []interface{}{
				map[string]interface{}{
					kubernetesNetworkConfigKey: []interface{}{
						map[string]interface{}{
							serviceCidrKey: test1000010,
						},
					},
					roleArnKey:           testRoleArn,
					kubernetesVersionKey: test112,
					tagsKey: map[string]string{
						testTag1: testTag2,
					},
					vpcKey: []interface{}{
						map[string]interface{}{
							enablePrivateAccessKey: false,
							enablePublicAccessKey:  false,
							publicAccessCidrsKey: []string{
								test00001,
								test10001,
							},
							securityGroupsKey: []string{
								testSg1,
								testSg2,
							},
							subnetIdsKey: []string{
								testSubnet1,
								testSubnet2,
							},
						},
					},
				},
			},
		},
		{
			description: "vpc is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
				config := getConfig()
				config.Vpc = nil
				config.AddonsConfig = nil

				return config
			},
			expected: []interface{}{
				map[string]interface{}{
					kubernetesNetworkConfigKey: []interface{}{
						map[string]interface{}{
							serviceCidrKey: test1000010,
						},
					},
					kubernetesVersionKey: test112,
					loggingKey: []interface{}{
						map[string]interface{}{
							apiServerKey:         false,
							auditKey:             false,
							authenticatorKey:     false,
							controllerManagerKey: true,
							schedulerKey:         true,
						},
					},
					roleArnKey: testRoleArn,
					tagsKey: map[string]string{
						testTag1: testTag2,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output := flattenConfig(test.getInput())
			require.Equal(t, test.expected, output)
		})
	}
}

func getClusterSpec() (*eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec, []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) {
	spec := &eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec{
		ClusterGroupName: testTestCg,
		ProxyName:        testTestProoxy,
		Config:           getConfig(),
	}
	nodepool := []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
		{
			Info: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo{
				Description: testTestNp,
				Name:        testTestNp2,
			},
			Spec: getNodepoolSpec(),
		},
	}

	return spec, nodepool
}

func getConfig() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
	return &eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig{
		KubernetesNetworkConfig: &eksmodel.VmwareTanzuManageV1alpha1EksclusterKubernetesNetworkConfig{
			ServiceCidr: test1000010,
		},
		Logging: &eksmodel.VmwareTanzuManageV1alpha1EksclusterLogging{
			APIServer:         false,
			Audit:             false,
			Authenticator:     false,
			ControllerManager: true,
			Scheduler:         true,
		},
		RoleArn: testRoleArn,
		Tags: map[string]string{
			testTag1: testTag2,
		},
		Version: test112,
		Vpc: &eksmodel.VmwareTanzuManageV1alpha1EksclusterVPCConfig{
			EnablePrivateAccess: false,
			EnablePublicAccess:  false,
			PublicAccessCidrs: []string{
				test00001,
				test10001,
			},
			SecurityGroups: []string{
				testSg1,
				testSg2,
			},
			SubnetIds: []string{
				testSubnet1,
				testSubnet2,
			},
		},
		AddonsConfig: &eksmodel.VmwareTanzuManageV1alpha1EksclusterAddonsConfig{
			VpcCniAddonConfig: &eksmodel.VmwareTanzuManageV1alpha1EksclusterVpcCniAddonConfig{
				EniConfigs: []*eksmodel.VmwareTanzuManageV1alpha1EksclusterEniConfig{
					{
						SubnetID: testSubnet1,
						SecurityGroupIds: []string{
							testSg1,
							testSg2,
						},
					},
					{
						SubnetID: testSubnet2,
						SecurityGroupIds: []string{
							testSg3,
							testSg4,
						},
					},
				},
			},
		},
	}
}
