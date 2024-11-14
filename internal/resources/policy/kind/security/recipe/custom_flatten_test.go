// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	policyrecipesecuritymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/security"
)

func TestFlattenCustom(t *testing.T) {
	t.Parallel()

	var (
		forbiddenSysctlsTestInput1      = "kernel.msgmax"
		forbiddenSysctlsTestInput2      = "kernel.sem"
		allowedProfilesTestInput        = "Localhost"
		allowedLocalhostFilesTestInput1 = "profiles/audit.json"
		allowedLocalhostFilesTestInput2 = "profiles/violation.json"
	)

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom
		expected    []interface{}
	}{
		{
			description: "check for nil security policy custom recipe",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete security policy custom recipe",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom{
				Audit:                     true,
				DisableNativePsp:          false,
				AllowPrivilegedContainers: true,
				AllowPrivilegeEscalation:  true,
				AllowHostNamespaceSharing: true,
				AllowHostNetwork:          true,
				ReadOnlyRootFileSystem:    true,
				AllowedHostPortRange: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
					Min: 3000,
					Max: 5000,
				},
				AllowedVolumes: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume{
					policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeConfigMap),
					policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeNfs),
					policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeVsphereVolume),
				},
				RunAsUser: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser{
					Rule: policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleRunAsAny,
					Ranges: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
						{
							Min: 3,
							Max: 5,
						},
						{
							Min: 7,
							Max: 12,
						},
					},
				},
				RunAsGroup: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup{
					Rule: policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny,
					Ranges: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
						{
							Min: 3,
							Max: 5,
						},
						{
							Min: 7,
							Max: 12,
						},
					},
				},
				SupplementalGroups: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup{
					Rule: policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny,
					Ranges: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
						{
							Min: 3,
							Max: 5,
						},
						{
							Min: 7,
							Max: 12,
						},
					},
				},
				FsGroup: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup{
					Rule: policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny,
					Ranges: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
						{
							Min: 3,
							Max: 5,
						},
						{
							Min: 7,
							Max: 12,
						},
					},
				},
				LinuxCapabilities: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities{
					AllowedCapabilities: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability{
						policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityChown),
						policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityIpcLock),
					},
					RequiredDropCapabilities: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability{
						policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysTime),
					},
				},
				AllowedHostPaths: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath{
					{
						ReadOnly:   true,
						PathPrefix: "p1",
					},
					{
						ReadOnly:   false,
						PathPrefix: "p2",
					},
					{
						ReadOnly:   true,
						PathPrefix: "p3",
					},
				},
				AllowedSELinuxOptions: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption{
					{
						Level: "s0",
						Role:  "sysadm_r",
						Type:  "httpd_sys_content_t",
						User:  "root",
					},
				},
				Sysctls: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls{
					ForbiddenSysctls: []*string{&forbiddenSysctlsTestInput1, &forbiddenSysctlsTestInput2},
				},
				Seccomp: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp{
					AllowedProfiles:       []*string{&allowedProfilesTestInput},
					AllowedLocalhostFiles: []*string{&allowedLocalhostFilesTestInput1, &allowedLocalhostFilesTestInput2},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					AuditKey:                     true,
					DisableNativePspKey:          false,
					allowPrivilegedContainersKey: true,
					allowPrivilegeEscalationKey:  true,
					allowHostNamespaceSharingKey: true,
					allowHostNetworkKey:          true,
					readOnlyRootFileSystemKey:    true,
					allowedHostPortRangeKey: []interface{}{
						map[string]interface{}{
							minKey: 3000,
							maxKey: 5000,
						},
					},
					allowedVolumesKey: []interface{}{
						string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeConfigMap),
						string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeNfs),
						string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolumeVsphereVolume),
					},
					runAsUserKey: []interface{}{
						map[string]interface{}{
							ruleKey: string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleRunAsAny),
							rangesKey: []interface{}{
								map[string]interface{}{
									minKey: 3,
									maxKey: 5,
								},
								map[string]interface{}{
									minKey: 7,
									maxKey: 12,
								},
							},
						},
					},
					runAsGroupKey: []interface{}{
						map[string]interface{}{
							ruleKey: string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny),
							rangesKey: []interface{}{
								map[string]interface{}{
									minKey: 3,
									maxKey: 5,
								},
								map[string]interface{}{
									minKey: 7,
									maxKey: 12,
								},
							},
						},
					},
					supplementalGroupsKey: []interface{}{
						map[string]interface{}{
							ruleKey: string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny),
							rangesKey: []interface{}{
								map[string]interface{}{
									minKey: 3,
									maxKey: 5,
								},
								map[string]interface{}{
									minKey: 7,
									maxKey: 12,
								},
							},
						},
					},
					fsGroupKey: []interface{}{
						map[string]interface{}{
							ruleKey: string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny),
							rangesKey: []interface{}{
								map[string]interface{}{
									minKey: 3,
									maxKey: 5,
								},
								map[string]interface{}{
									minKey: 7,
									maxKey: 12,
								},
							},
						},
					},
					linuxCapabilitiesKey: []interface{}{
						map[string]interface{}{
							allowedCapabilitiesKey: []interface{}{
								string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityChown),
								string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityIpcLock),
							},
							requiredDropCapabilitiesKey: []interface{}{
								string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysTime),
							},
						},
					},
					allowedHostPathsKey: []interface{}{
						map[string]interface{}{
							readOnlyKey:   true,
							pathPrefixKey: "p1",
						},
						map[string]interface{}{
							readOnlyKey:   false,
							pathPrefixKey: "p2",
						},
						map[string]interface{}{
							readOnlyKey:   true,
							pathPrefixKey: "p3",
						},
					},
					allowedSELinuxOptionsKey: []interface{}{
						map[string]interface{}{
							levelKey: "s0",
							roleKey:  "sysadm_r",
							typeKey:  "httpd_sys_content_t",
							userKey:  "root",
						},
					},
					sysctlsKey: []interface{}{
						map[string]interface{}{
							forbiddenSysctlsKey: []interface{}{
								forbiddenSysctlsTestInput1,
								forbiddenSysctlsTestInput2,
							},
						},
					},
					seccompKey: []interface{}{
						map[string]interface{}{
							allowedProfilesKey: []interface{}{
								allowedProfilesTestInput,
							},
							allowedLocalhostFilesKey: []interface{}{
								allowedLocalhostFilesTestInput1,
								allowedLocalhostFilesTestInput2,
							},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenCustom(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenAllowedHostPortRange(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange
		expected    []interface{}
	}{
		{
			description: "check for nil allowed host port range",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete allowed host port range",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
				Min: 3000,
				Max: 5000,
			},
			expected: []interface{}{
				map[string]interface{}{
					minKey: 3000,
					maxKey: 5000,
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenAllowedHostPortRange(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

// nolint: dupl
func TestFlattenRunAsUser(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser
		expected    []interface{}
	}{
		{
			description: "check for nil run as user",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete run as user",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser{
				Rule: policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleRunAsAny,
				Ranges: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
					{
						Min: 3,
						Max: 5,
					},
					{
						Min: 7,
						Max: 12,
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					ruleKey: string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRuleRunAsAny),
					rangesKey: []interface{}{
						map[string]interface{}{
							minKey: 3,
							maxKey: 5,
						},
						map[string]interface{}{
							minKey: 7,
							maxKey: 12,
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenRunAsUser(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenUserIDRange(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange
		expected    interface{}
	}{
		{
			description: "check for nil user id range",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete user id range",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
				Min: 3,
				Max: 5,
			},
			expected: map[string]interface{}{
				minKey: 3,
				maxKey: 5,
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenUserIDRange(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

// nolint: dupl
func TestFlattenRunAsGroup(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup
		expected    []interface{}
	}{
		{
			description: "check for nil run as group",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete run as group",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup{
				Rule: policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny,
				Ranges: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
					{
						Min: 3,
						Max: 5,
					},
					{
						Min: 7,
						Max: 12,
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					ruleKey: string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny),
					rangesKey: []interface{}{
						map[string]interface{}{
							minKey: 3,
							maxKey: 5,
						},
						map[string]interface{}{
							minKey: 7,
							maxKey: 12,
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenRunAsGroup(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenGroupIDRange(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange
		expected    interface{}
	}{
		{
			description: "check for nil group id range",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete group id range",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
				Min: 3,
				Max: 5,
			},
			expected: map[string]interface{}{
				minKey: 3,
				maxKey: 5,
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenGroupIDRange(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

// nolint: dupl
func TestFlattenSupplementalGroups(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup
		expected    []interface{}
	}{
		{
			description: "check for nil supplemental groups",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete supplemental groups",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup{
				Rule: policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny,
				Ranges: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
					{
						Min: 3,
						Max: 5,
					},
					{
						Min: 7,
						Max: 12,
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					ruleKey: string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny),
					rangesKey: []interface{}{
						map[string]interface{}{
							minKey: 3,
							maxKey: 5,
						},
						map[string]interface{}{
							minKey: 7,
							maxKey: 12,
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenSupplementalGroups(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

// nolint: dupl
func TestFlattenFsGroup(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup
		expected    []interface{}
	}{
		{
			description: "check for nil fs group",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete fs group",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup{
				Rule: policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny,
				Ranges: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{
					{
						Min: 3,
						Max: 5,
					},
					{
						Min: 7,
						Max: 12,
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					ruleKey: string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRuleRunAsAny),
					rangesKey: []interface{}{
						map[string]interface{}{
							minKey: 3,
							maxKey: 5,
						},
						map[string]interface{}{
							minKey: 7,
							maxKey: 12,
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenFsGroup(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenLinuxCapabilities(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities
		expected    []interface{}
	}{
		{
			description: "check for nil linux capabilities",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete linux capabilities",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities{
				AllowedCapabilities: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability{
					policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityChown),
					policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityIpcLock),
				},
				RequiredDropCapabilities: []*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability{
					policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysTime),
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					allowedCapabilitiesKey: []interface{}{
						string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityChown),
						string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapabilityIpcLock),
					},
					requiredDropCapabilitiesKey: []interface{}{
						string(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapabilitySysTime),
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenLinuxCapabilities(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenAllowedHostPath(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath
		expected    interface{}
	}{
		{
			description: "check for nil allowed host path",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete allowed host path",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath{
				ReadOnly:   true,
				PathPrefix: "p1",
			},
			expected: map[string]interface{}{
				readOnlyKey:   true,
				pathPrefixKey: "p1",
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenAllowedHostPath(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenAllowedSELinuxOption(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption
		expected    interface{}
	}{
		{
			description: "check for nil allowed se linux option",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete allowed se linux option",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption{
				Level: "s0",
				Role:  "sysadm_r",
				Type:  "httpd_sys_content_t",
				User:  "root",
			},
			expected: map[string]interface{}{
				levelKey: "s0",
				roleKey:  "sysadm_r",
				typeKey:  "httpd_sys_content_t",
				userKey:  "root",
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenAllowedSELinuxOption(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenSysctls(t *testing.T) {
	t.Parallel()

	var (
		forbiddenSysctlsTestInput1 = "kernel.msgmax"
		forbiddenSysctlsTestInput2 = "kernel.sem"
	)

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls
		expected    []interface{}
	}{
		{
			description: "check for nil sysctls",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete sysctls",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls{
				ForbiddenSysctls: []*string{&forbiddenSysctlsTestInput1, &forbiddenSysctlsTestInput2},
			},
			expected: []interface{}{
				map[string]interface{}{
					forbiddenSysctlsKey: []interface{}{
						forbiddenSysctlsTestInput1,
						forbiddenSysctlsTestInput2,
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenSysctls(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenSeccomp(t *testing.T) {
	t.Parallel()

	var (
		allowedProfilesTestInput        = "Localhost"
		allowedLocalhostFilesTestInput1 = "profiles/audit.json"
		allowedLocalhostFilesTestInput2 = "profiles/violation.json"
	)

	cases := []struct {
		description string
		input       *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp
		expected    []interface{}
	}{
		{
			description: "check for nil seccomp",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete seccomp",
			input: &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp{
				AllowedProfiles:       []*string{&allowedProfilesTestInput},
				AllowedLocalhostFiles: []*string{&allowedLocalhostFilesTestInput1, &allowedLocalhostFilesTestInput2},
			},
			expected: []interface{}{
				map[string]interface{}{
					allowedProfilesKey: []interface{}{
						allowedProfilesTestInput,
					},
					allowedLocalhostFilesKey: []interface{}{
						allowedLocalhostFilesTestInput1,
						allowedLocalhostFilesTestInput2,
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenSeccomp(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
