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
					getNodepoolDef(testTestNp2),
				}
			},
			expected: []interface{}{
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
		{
			description: "multiple nodepool",
			getInput: func() []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
				return []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
					getNodepoolDef(testTestNp2),
					getNodepoolDef("test-np-2"),
				}
			},
			expected: []interface{}{
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
				map[string]interface{}{
					infoKey: []interface{}{
						map[string]interface{}{
							testDescription: testTestNp,
							NameKey:         "test-np-2",
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
				return getNodepoolDef(testTestNp2)
			},
			expected: map[string]interface{}{
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
							"id":       "lt-id",
							NameKey:    "",
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
		{
			description: "remote access is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.RemoteAccess = nil

				return spec
			},
			expected: []interface{}{
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
		{
			description: "root disk size is 0",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.RootDiskSize = 0

				return spec
			},
			expected: []interface{}{
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
					roleArnKey: testArnAwsIam000000000000,
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
		{
			description: "scaling config is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.ScalingConfig = nil

				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					amiTypeKey: testCustom,
					amiInfoKey: []interface{}{
						map[string]interface{}{
							amiIDKey:                testAmi2qu8409oisdfj0qw,
							overrideBootstrapCmdKey: testBootstrapCmd,
						},
					},
					capacityTypeKey: testOnDemand,
					rootDiskSizeKey: int32(20),
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
					roleArnKey: testArnAwsIam000000000000,
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
		{
			description: "subnet ids are nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.SubnetIds = nil

				return spec
			},
			expected: []interface{}{
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
		{
			description: "taints are nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.Taints = nil

				return spec
			},
			expected: []interface{}{
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
		{
			description: "update config is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.UpdateConfig = nil

				return spec
			},
			expected: []interface{}{
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
					releaseVersionKey: test126420230703,
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
			Description: testTestNp,
			Name:        name,
		},
		Spec: getNodepoolSpec(),
	}
}
