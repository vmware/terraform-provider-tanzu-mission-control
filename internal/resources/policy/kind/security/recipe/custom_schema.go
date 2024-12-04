// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipesecuritymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/security"
)

var Custom = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for security policy custom recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			AuditKey: {
				Type:        schema.TypeBool,
				Description: "Audit (dry-run)",
				Optional:    true,
				Default:     false,
			},
			DisableNativePspKey: {
				Type:        schema.TypeBool,
				Description: "Disable native pod security policy",
				Optional:    true,
				Default:     false,
			},
			allowPrivilegedContainersKey: {
				Type:        schema.TypeBool,
				Description: "Allow privileged containers",
				Optional:    true,
				Default:     false,
			},
			allowPrivilegeEscalationKey: {
				Type:        schema.TypeBool,
				Description: "Allow privilege escalation",
				Optional:    true,
				Default:     false,
			},
			allowHostNamespaceSharingKey: {
				Type:        schema.TypeBool,
				Description: "Allow host namespace sharing",
				Optional:    true,
				Default:     false,
			},
			allowHostNetworkKey: {
				Type:        schema.TypeBool,
				Description: "Allow host network",
				Optional:    true,
				Default:     false,
			},
			readOnlyRootFileSystemKey: {
				Type:        schema.TypeBool,
				Description: "Read only root file system",
				Optional:    true,
				Default:     false,
			},
			allowedHostPortRangeKey:  allowedHostPortRange,
			allowedVolumesKey:        allowedVolumes,
			runAsUserKey:             runAsUser,
			runAsGroupKey:            runAsGroup,
			supplementalGroupsKey:    supplementalGroups,
			fsGroupKey:               fsGroup,
			linuxCapabilitiesKey:     linuxCapabilities,
			allowedHostPathsKey:      allowedHostPaths,
			allowedSELinuxOptionsKey: allowedSELinuxOptions,
			sysctlsKey:               sysctls,
			seccompKey:               seccomp,
		},
	},
}

// ConstructCustom constructs a security policy input with custom recipe.
// nolint: gocognit, gocyclo
func ConstructCustom(data []interface{}) (custom *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom) {
	if len(data) == 0 || data[0] == nil {
		return custom
	}

	customData, _ := data[0].(map[string]interface{})

	custom = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom{}

	if v, ok := customData[AuditKey]; ok {
		helper.SetPrimitiveValue(v, &custom.Audit, AuditKey)
	}

	if v, ok := customData[DisableNativePspKey]; ok {
		helper.SetPrimitiveValue(v, &custom.DisableNativePsp, DisableNativePspKey)
	}

	if v, ok := customData[allowPrivilegedContainersKey]; ok {
		helper.SetPrimitiveValue(v, &custom.AllowPrivilegedContainers, allowPrivilegedContainersKey)
	}

	if v, ok := customData[allowPrivilegeEscalationKey]; ok {
		helper.SetPrimitiveValue(v, &custom.AllowPrivilegeEscalation, allowPrivilegeEscalationKey)
	}

	if v, ok := customData[allowHostNamespaceSharingKey]; ok {
		helper.SetPrimitiveValue(v, &custom.AllowHostNamespaceSharing, allowHostNamespaceSharingKey)
	}

	if v, ok := customData[allowHostNetworkKey]; ok {
		helper.SetPrimitiveValue(v, &custom.AllowHostNetwork, allowHostNetworkKey)
	}

	if v, ok := customData[readOnlyRootFileSystemKey]; ok {
		helper.SetPrimitiveValue(v, &custom.ReadOnlyRootFileSystem, readOnlyRootFileSystemKey)
	}

	if v, ok := customData[allowedHostPortRangeKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			custom.AllowedHostPortRange = expandAllowedHostPortRange(v1)
		}
	}

	if v, ok := customData[allowedVolumesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				custom.AllowedVolumes = make([]*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume, 0)

				for _, raw := range vs {
					custom.AllowedVolumes = append(custom.AllowedVolumes, policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedVolume(raw.(string))))
				}
			}
		}
	}

	if v, ok := customData[runAsUserKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			custom.RunAsUser = expandRunAsUser(v1)
		}
	}

	if v, ok := customData[runAsGroupKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			custom.RunAsGroup = expandRunAsGroup(v1)
		}
	}

	if v, ok := customData[supplementalGroupsKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			custom.SupplementalGroups = expandSupplementalGroups(v1)
		}
	}

	if v, ok := customData[fsGroupKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			custom.FsGroup = expandFsGroup(v1)
		}
	}

	if v, ok := customData[linuxCapabilitiesKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			custom.LinuxCapabilities = expandLinuxCapabilities(v1)
		}
	}

	if v, ok := customData[allowedHostPathsKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				custom.AllowedHostPaths = make([]*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath, 0)

				for _, raw := range vs {
					custom.AllowedHostPaths = append(custom.AllowedHostPaths, expandAllowedHostPath(raw))
				}
			}
		}
	}

	if v, ok := customData[allowedSELinuxOptionsKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				custom.AllowedSELinuxOptions = make([]*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption, 0)

				for _, raw := range vs {
					custom.AllowedSELinuxOptions = append(custom.AllowedSELinuxOptions, expandAllowedSELinuxOption(raw))
				}
			}
		}
	}

	if v, ok := customData[sysctlsKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			custom.Sysctls = expandSysctls(v1)
		}
	}

	if v, ok := customData[seccompKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			custom.Seccomp = expandSeccomp(v1)
		}
	}

	return custom
}

func FlattenCustom(custom *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom) (data []interface{}) {
	if custom == nil {
		return data
	}

	flattenCustom := make(map[string]interface{})

	flattenCustom[AuditKey] = custom.Audit
	flattenCustom[DisableNativePspKey] = custom.DisableNativePsp
	flattenCustom[allowPrivilegedContainersKey] = custom.AllowPrivilegedContainers
	flattenCustom[allowPrivilegeEscalationKey] = custom.AllowPrivilegeEscalation
	flattenCustom[allowHostNamespaceSharingKey] = custom.AllowHostNamespaceSharing
	flattenCustom[allowHostNetworkKey] = custom.AllowHostNetwork
	flattenCustom[readOnlyRootFileSystemKey] = custom.ReadOnlyRootFileSystem

	if custom.AllowedHostPortRange != nil {
		flattenCustom[allowedHostPortRangeKey] = flattenAllowedHostPortRange(custom.AllowedHostPortRange)
	}

	if custom.AllowedVolumes != nil {
		avs := make([]interface{}, 0)

		for _, av := range custom.AllowedVolumes {
			avs = append(avs, string(*av))
		}

		flattenCustom[allowedVolumesKey] = avs
	}

	if custom.RunAsUser != nil {
		flattenCustom[runAsUserKey] = flattenRunAsUser(custom.RunAsUser)
	}

	if custom.RunAsGroup != nil {
		flattenCustom[runAsGroupKey] = flattenRunAsGroup(custom.RunAsGroup)
	}

	if custom.SupplementalGroups != nil {
		flattenCustom[supplementalGroupsKey] = flattenSupplementalGroups(custom.SupplementalGroups)
	}

	if custom.FsGroup != nil {
		flattenCustom[fsGroupKey] = flattenFsGroup(custom.FsGroup)
	}

	if custom.LinuxCapabilities != nil {
		flattenCustom[linuxCapabilitiesKey] = flattenLinuxCapabilities(custom.LinuxCapabilities)
	}

	if custom.AllowedHostPaths != nil {
		ahps := make([]interface{}, 0)

		for _, ahp := range custom.AllowedHostPaths {
			ahps = append(ahps, flattenAllowedHostPath(ahp))
		}

		flattenCustom[allowedHostPathsKey] = ahps
	}

	if custom.AllowedSELinuxOptions != nil {
		aselos := make([]interface{}, 0)

		for _, aselo := range custom.AllowedSELinuxOptions {
			aselos = append(aselos, flattenAllowedSELinuxOption(aselo))
		}

		flattenCustom[allowedSELinuxOptionsKey] = aselos
	}

	if custom.Sysctls != nil {
		flattenCustom[sysctlsKey] = flattenSysctls(custom.Sysctls)
	}

	if custom.Seccomp != nil {
		flattenCustom[seccompKey] = flattenSeccomp(custom.Seccomp)
	}

	return []interface{}{flattenCustom}
}

var allowedHostPortRange = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Allowed host port range",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			minKey: {
				Type:         schema.TypeInt,
				Description:  "Minimum allowed port",
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntBetween(0, 65535),
			},
			maxKey: {
				Type:         schema.TypeInt,
				Description:  "Maximum allowed port",
				Optional:     true,
				Default:      65535,
				ValidateFunc: validation.IntBetween(0, 65535),
			},
		},
	},
}

func expandAllowedHostPortRange(data []interface{}) (allowedHostPortRange *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange) {
	if len(data) == 0 || data[0] == nil {
		return allowedHostPortRange
	}

	allowedHostPortRangeData, _ := data[0].(map[string]interface{})

	allowedHostPortRange = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{}

	if v, ok := allowedHostPortRangeData[minKey]; ok {
		helper.SetPrimitiveValue(v, &allowedHostPortRange.Min, minKey)
	}

	if v, ok := allowedHostPortRangeData[maxKey]; ok {
		helper.SetPrimitiveValue(v, &allowedHostPortRange.Max, maxKey)
	}

	return allowedHostPortRange
}

func flattenAllowedHostPortRange(allowedHostPortRange *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange) (data []interface{}) {
	if allowedHostPortRange == nil {
		return data
	}

	flattenAllowedHostPortRange := make(map[string]interface{})

	flattenAllowedHostPortRange[minKey] = allowedHostPortRange.Min
	flattenAllowedHostPortRange[maxKey] = allowedHostPortRange.Max

	return []interface{}{flattenAllowedHostPortRange}
}

var allowedVolumes = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Allowed volumes",
	Optional:    true,
	DefaultFunc: func() (interface{}, error) {
		return []interface{}{"*"}, nil
	},
	Elem: &schema.Schema{
		Type: schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"*", "configMap", "downwardAPI", "emptyDir", "persistentVolumeClaim", "secret", "projected", "hostPath", "flexVolume",
			"awsElasticBlockStore", "azureDisk", "azureFile", "cephfs", "cinder", "csi", "fc", "flocker", "gcePersistentDisk",
			"gitRepo", "glusterfs", "iscsi", "local", "nfs", "portworxVolume", "quobyte", "rbd", "scaleIO", "storageos", "vsphereVolume"}, false),
	},
}

var runAsUser = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Run as user",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ruleKey: {
				Type:         schema.TypeString,
				Description:  "Rule",
				Optional:     true,
				Default:      "RunAsAny",
				ValidateFunc: validation.StringInSlice([]string{"RunAsAny", "MustRunAsNonRoot", "MustRunAs"}, false),
			},
			rangesKey: userIDRanges,
		},
	},
}

func expandRunAsUser(data []interface{}) (runAsUser *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser) {
	if len(data) == 0 || data[0] == nil {
		return runAsUser
	}

	runAsUserData, _ := data[0].(map[string]interface{})

	runAsUser = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser{
		Ranges: make([]*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange, 0),
	}

	if v, ok := runAsUserData[ruleKey]; ok {
		runAsUser.Rule = *policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUserRule(v.(string)))
	}

	if v, ok := runAsUserData[rangesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			runAsUser.Ranges = append(runAsUser.Ranges, expandUserIDRange(raw))
		}
	}

	return runAsUser
}

func flattenRunAsUser(runAsUser *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsUser) (data []interface{}) {
	if runAsUser == nil {
		return data
	}

	flattenRunAsUser := make(map[string]interface{})

	flattenRunAsUser[ruleKey] = string(runAsUser.Rule)

	rs := make([]interface{}, 0)

	for _, r := range runAsUser.Ranges {
		rs = append(rs, flattenUserIDRange(r))
	}

	flattenRunAsUser[rangesKey] = rs

	return []interface{}{flattenRunAsUser}
}

var userIDRanges = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Allowed user id ranges",
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			minKey: {
				Type:         schema.TypeInt,
				Description:  "Minimum user ID",
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntBetween(0, 65535),
			},
			maxKey: {
				Type:         schema.TypeInt,
				Description:  "Maximum user ID",
				Optional:     true,
				Default:      65535,
				ValidateFunc: validation.IntBetween(0, 65535),
			},
		},
	},
}

func expandUserIDRange(data interface{}) (userIDRange *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange) {
	if data == nil {
		return userIDRange
	}

	userIDRangeData, _ := data.(map[string]interface{})

	userIDRange = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{}

	if v, ok := userIDRangeData[minKey]; ok {
		helper.SetPrimitiveValue(v, &userIDRange.Min, minKey)
	}

	if v, ok := userIDRangeData[maxKey]; ok {
		helper.SetPrimitiveValue(v, &userIDRange.Max, maxKey)
	}

	return userIDRange
}

func flattenUserIDRange(userIDRange *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange) (data interface{}) {
	if userIDRange == nil {
		return data
	}

	flattenUserIDRange := make(map[string]interface{})

	flattenUserIDRange[minKey] = userIDRange.Min
	flattenUserIDRange[maxKey] = userIDRange.Max

	return flattenUserIDRange
}

var runAsGroup = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Run as group",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ruleKey: {
				Type:         schema.TypeString,
				Description:  "Rule",
				Optional:     true,
				Default:      "RunAsAny",
				ValidateFunc: validation.StringInSlice([]string{"RunAsAny", "MayRunAs", "MustRunAs"}, false),
			},
			rangesKey: groupIDRanges,
		},
	},
}

func expandRunAsGroup(data []interface{}) (runAsGroup *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup) {
	if len(data) == 0 || data[0] == nil {
		return runAsGroup
	}

	runAsGroupData, _ := data[0].(map[string]interface{})

	runAsGroup = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup{
		Ranges: make([]*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange, 0),
	}

	if v, ok := runAsGroupData[ruleKey]; ok {
		runAsGroup.Rule = *policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule(v.(string)))
	}

	if v, ok := runAsGroupData[rangesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			runAsGroup.Ranges = append(runAsGroup.Ranges, expandGroupIDRange(raw))
		}
	}

	return runAsGroup
}

func flattenRunAsGroup(runAsGroup *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup) (data []interface{}) {
	if runAsGroup == nil {
		return data
	}

	flattenRunAsGroup := make(map[string]interface{})

	flattenRunAsGroup[ruleKey] = string(runAsGroup.Rule)

	rs := make([]interface{}, 0)

	for _, r := range runAsGroup.Ranges {
		rs = append(rs, flattenGroupIDRange(r))
	}

	flattenRunAsGroup[rangesKey] = rs

	return []interface{}{flattenRunAsGroup}
}

var groupIDRanges = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Allowed group id ranges",
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			minKey: {
				Type:         schema.TypeInt,
				Description:  "Minimum group ID",
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntBetween(0, 65535),
			},
			maxKey: {
				Type:         schema.TypeInt,
				Description:  "Maximum group ID",
				Optional:     true,
				Default:      65535,
				ValidateFunc: validation.IntBetween(0, 65535),
			},
		},
	},
}

func expandGroupIDRange(data interface{}) (groupIDRange *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange) {
	if data == nil {
		return groupIDRange
	}

	groupIDRangeData, _ := data.(map[string]interface{})

	groupIDRange = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange{}

	if v, ok := groupIDRangeData[minKey]; ok {
		helper.SetPrimitiveValue(v, &groupIDRange.Min, minKey)
	}

	if v, ok := groupIDRangeData[maxKey]; ok {
		helper.SetPrimitiveValue(v, &groupIDRange.Max, maxKey)
	}

	return groupIDRange
}

func flattenGroupIDRange(groupIDRange *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange) (data interface{}) {
	if groupIDRange == nil {
		return data
	}

	flattenGroupIDRange := make(map[string]interface{})

	flattenGroupIDRange[minKey] = groupIDRange.Min
	flattenGroupIDRange[maxKey] = groupIDRange.Max

	return flattenGroupIDRange
}

var supplementalGroups = &schema.Schema{
	Type:        schema.TypeList,
	Description: "supplemental groups",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ruleKey: {
				Type:         schema.TypeString,
				Description:  "Rule",
				Optional:     true,
				Default:      "RunAsAny",
				ValidateFunc: validation.StringInSlice([]string{"RunAsAny", "MayRunAs", "MustRunAs"}, false),
			},
			rangesKey: groupIDRanges,
		},
	},
}

func expandSupplementalGroups(data []interface{}) (supplementalGroups *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup) {
	if len(data) == 0 || data[0] == nil {
		return supplementalGroups
	}

	supplementalGroupsData, _ := data[0].(map[string]interface{})

	supplementalGroups = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup{
		Ranges: make([]*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange, 0),
	}

	if v, ok := supplementalGroupsData[ruleKey]; ok {
		supplementalGroups.Rule = *policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule(v.(string)))
	}

	if v, ok := supplementalGroupsData[rangesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			supplementalGroups.Ranges = append(supplementalGroups.Ranges, expandGroupIDRange(raw))
		}
	}

	return supplementalGroups
}

func flattenSupplementalGroups(supplementalGroups *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup) (data []interface{}) {
	if supplementalGroups == nil {
		return data
	}

	flattenSupplementalGroups := make(map[string]interface{})

	flattenSupplementalGroups[ruleKey] = string(supplementalGroups.Rule)

	rs := make([]interface{}, 0)

	for _, r := range supplementalGroups.Ranges {
		rs = append(rs, flattenGroupIDRange(r))
	}

	flattenSupplementalGroups[rangesKey] = rs

	return []interface{}{flattenSupplementalGroups}
}

var fsGroup = &schema.Schema{
	Type:        schema.TypeList,
	Description: "fsGroup",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ruleKey: {
				Type:         schema.TypeString,
				Description:  "Rule",
				Optional:     true,
				Default:      "RunAsAny",
				ValidateFunc: validation.StringInSlice([]string{"RunAsAny", "MayRunAs", "MustRunAs"}, false),
			},
			rangesKey: groupIDRanges,
		},
	},
}

func expandFsGroup(data []interface{}) (fsGroup *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup) {
	if len(data) == 0 || data[0] == nil {
		return fsGroup
	}

	fsGroupData, _ := data[0].(map[string]interface{})

	fsGroup = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup{
		Ranges: make([]*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRange, 0),
	}

	if v, ok := fsGroupData[ruleKey]; ok {
		fsGroup.Rule = *policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroupRule(v.(string)))
	}

	if v, ok := fsGroupData[rangesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			fsGroup.Ranges = append(fsGroup.Ranges, expandGroupIDRange(raw))
		}
	}

	return fsGroup
}

func flattenFsGroup(fsGroup *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomRunAsGroup) (data []interface{}) {
	if fsGroup == nil {
		return data
	}

	flattenFsGroup := make(map[string]interface{})

	flattenFsGroup[ruleKey] = string(fsGroup.Rule)

	rs := make([]interface{}, 0)

	for _, r := range fsGroup.Ranges {
		rs = append(rs, flattenGroupIDRange(r))
	}

	flattenFsGroup[rangesKey] = rs

	return []interface{}{flattenFsGroup}
}

var linuxCapabilities = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Linux capabilities",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			allowedCapabilitiesKey: {
				Type:        schema.TypeList,
				Description: "Allowed capabilities",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					return []interface{}{"*"}, nil
				},
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"*", "AUDIT_CONTROL", "AUDIT_READ", "AUDIT_WRITE", "BLOCK_SUSPEND", "CHOWN", "DAC_OVERRIDE", "DAC_READ_SEARCH",
						"FOWNER", "FSETID", "IPC_LOCK", "IPC_OWNER", "KILL", "LEASE", "LINUX_IMMUTABLE", "MAC_ADMIN", "MAC_OVERRIDE",
						"MKNOD", "NET_ADMIN", "NET_BIND_SERVICE", "NET_BROADCAST", "NET_RAW", "SETGID", "SETFCAP", "SETPCAP", "SETUID",
						"SYS_ADMIN", "SYS_BOOT", "SYS_CHROOT", "SYS_MODULE", "SYS_NICE", "SYS_PACCT", "SYS_PTRACE", "SYS_RAWIO",
						"SYS_RESOURCE", "SYS_TIME", "SYS_TTY_CONFIG", "SYSLOG", "WAKE_ALARM"}, false),
				},
			},
			requiredDropCapabilitiesKey: {
				Type:        schema.TypeList,
				Description: "Required drop capabilities",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					return []interface{}{}, nil
				},
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"ALL", "AUDIT_CONTROL", "AUDIT_READ", "AUDIT_WRITE", "BLOCK_SUSPEND", "CHOWN", "DAC_OVERRIDE", "DAC_READ_SEARCH",
						"FOWNER", "FSETID", "IPC_LOCK", "IPC_OWNER", "KILL", "LEASE", "LINUX_IMMUTABLE", "MAC_ADMIN", "MAC_OVERRIDE",
						"MKNOD", "NET_ADMIN", "NET_BIND_SERVICE", "NET_BROADCAST", "NET_RAW", "SETGID", "SETFCAP", "SETPCAP", "SETUID",
						"SYS_ADMIN", "SYS_BOOT", "SYS_CHROOT", "SYS_MODULE", "SYS_NICE", "SYS_PACCT", "SYS_PTRACE", "SYS_RAWIO",
						"SYS_RESOURCE", "SYS_TIME", "SYS_TTY_CONFIG", "SYSLOG", "WAKE_ALARM"}, false),
				},
			},
		},
	},
}

func expandLinuxCapabilities(data []interface{}) (linuxCapabilities *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities) {
	if len(data) == 0 || data[0] == nil {
		return linuxCapabilities
	}

	linuxCapabilitiesData, _ := data[0].(map[string]interface{})

	linuxCapabilities = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities{
		AllowedCapabilities:      make([]*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability, 0),
		RequiredDropCapabilities: make([]*policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability, 0),
	}

	if v, ok := linuxCapabilitiesData[allowedCapabilitiesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			linuxCapabilities.AllowedCapabilities = append(linuxCapabilities.AllowedCapabilities, policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesAllowedCapability(raw.(string))))
		}
	}

	if v, ok := linuxCapabilitiesData[requiredDropCapabilitiesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			linuxCapabilities.RequiredDropCapabilities = append(linuxCapabilities.RequiredDropCapabilities, policyrecipesecuritymodel.NewVmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability(policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilitiesRequiredDropCapability(raw.(string))))
		}
	}

	return linuxCapabilities
}

func flattenLinuxCapabilities(linuxCapabilities *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomLinuxCapabilities) (data []interface{}) {
	if linuxCapabilities == nil {
		return data
	}

	flattenLinuxCapabilities := make(map[string]interface{})

	if linuxCapabilities.AllowedCapabilities != nil {
		acs := make([]interface{}, 0)

		for _, ac := range linuxCapabilities.AllowedCapabilities {
			acs = append(acs, string(*ac))
		}

		flattenLinuxCapabilities[allowedCapabilitiesKey] = acs
	}

	if linuxCapabilities.RequiredDropCapabilities != nil {
		rdcs := make([]interface{}, 0)

		for _, rdc := range linuxCapabilities.RequiredDropCapabilities {
			rdcs = append(rdcs, string(*rdc))
		}

		flattenLinuxCapabilities[requiredDropCapabilitiesKey] = rdcs
	}

	return []interface{}{flattenLinuxCapabilities}
}

var allowedHostPaths = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Allowed host paths",
	Optional:    true,
	DefaultFunc: func() (interface{}, error) {
		return []interface{}{}, nil
	},
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			readOnlyKey: {
				Type:        schema.TypeBool,
				Description: "Read only flag",
				Optional:    true,
				Default:     false,
			},
			pathPrefixKey: {
				Type:        schema.TypeString,
				Description: "Path prefix",
				Optional:    true,
				Default:     "",
			},
		},
	},
}

func expandAllowedHostPath(data interface{}) (allowedHostPath *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath) {
	if data == nil {
		return allowedHostPath
	}

	allowedHostPathData, _ := data.(map[string]interface{})

	allowedHostPath = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath{}

	if v, ok := allowedHostPathData[readOnlyKey]; ok {
		helper.SetPrimitiveValue(v, &allowedHostPath.ReadOnly, readOnlyKey)
	}

	if v, ok := allowedHostPathData[pathPrefixKey]; ok {
		helper.SetPrimitiveValue(v, &allowedHostPath.PathPrefix, pathPrefixKey)
	}

	return allowedHostPath
}

func flattenAllowedHostPath(allowedHostPath *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedHostPath) (data interface{}) {
	if allowedHostPath == nil {
		return data
	}

	flattenAllowedHostPath := make(map[string]interface{})

	flattenAllowedHostPath[readOnlyKey] = allowedHostPath.ReadOnly
	flattenAllowedHostPath[pathPrefixKey] = allowedHostPath.PathPrefix

	return flattenAllowedHostPath
}

var allowedSELinuxOptions = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Allowed selinux options",
	Optional:    true,
	DefaultFunc: func() (interface{}, error) {
		return []interface{}{}, nil
	},
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			levelKey: {
				Type:        schema.TypeString,
				Description: "SELinux level",
				Optional:    true,
				Default:     "",
			},
			roleKey: {
				Type:        schema.TypeString,
				Description: "SELinux role",
				Optional:    true,
				Default:     "",
			},
			typeKey: {
				Type:        schema.TypeString,
				Description: "SELinux type",
				Optional:    true,
				Default:     "",
			},
			userKey: {
				Type:        schema.TypeString,
				Description: "SELinux user",
				Optional:    true,
				Default:     "",
			},
		},
	},
}

func expandAllowedSELinuxOption(data interface{}) (allowedSELinuxOption *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption) {
	if data == nil {
		return allowedSELinuxOption
	}

	allowedSELinuxOptionData, _ := data.(map[string]interface{})

	allowedSELinuxOption = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption{}

	if v, ok := allowedSELinuxOptionData[levelKey]; ok {
		helper.SetPrimitiveValue(v, &allowedSELinuxOption.Level, levelKey)
	}

	if v, ok := allowedSELinuxOptionData[roleKey]; ok {
		helper.SetPrimitiveValue(v, &allowedSELinuxOption.Role, roleKey)
	}

	if v, ok := allowedSELinuxOptionData[typeKey]; ok {
		helper.SetPrimitiveValue(v, &allowedSELinuxOption.Type, typeKey)
	}

	if v, ok := allowedSELinuxOptionData[userKey]; ok {
		helper.SetPrimitiveValue(v, &allowedSELinuxOption.User, userKey)
	}

	return allowedSELinuxOption
}

func flattenAllowedSELinuxOption(allowedSELinuxOption *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomAllowedSELinuxOption) (data interface{}) {
	if allowedSELinuxOption == nil {
		return data
	}

	flattenAllowedSELinuxOption := make(map[string]interface{})

	flattenAllowedSELinuxOption[levelKey] = allowedSELinuxOption.Level
	flattenAllowedSELinuxOption[roleKey] = allowedSELinuxOption.Role
	flattenAllowedSELinuxOption[typeKey] = allowedSELinuxOption.Type
	flattenAllowedSELinuxOption[userKey] = allowedSELinuxOption.User

	return flattenAllowedSELinuxOption
}

var sysctls = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Sysctls",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			forbiddenSysctlsKey: {
				Type:        schema.TypeList,
				Description: "Forbidden sysctls",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					return []interface{}{}, nil
				},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	},
}

func expandSysctls(data []interface{}) (sysctls *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls) {
	if len(data) == 0 || data[0] == nil {
		return sysctls
	}

	sysctlsData, _ := data[0].(map[string]interface{})

	sysctls = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls{
		ForbiddenSysctls: make([]*string, 0),
	}

	if v, ok := sysctlsData[forbiddenSysctlsKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			rawString := raw.(string)
			sysctls.ForbiddenSysctls = append(sysctls.ForbiddenSysctls, &rawString)
		}
	}

	return sysctls
}

func flattenSysctls(sysctls *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSysctls) (data []interface{}) {
	if sysctls == nil {
		return data
	}

	flattenSysctls := make(map[string]interface{})

	if sysctls.ForbiddenSysctls != nil {
		fss := make([]interface{}, 0)

		for _, fs := range sysctls.ForbiddenSysctls {
			fss = append(fss, *fs)
		}

		flattenSysctls[forbiddenSysctlsKey] = fss
	}

	return []interface{}{flattenSysctls}
}

var seccomp = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Seccomp",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			allowedProfilesKey: {
				Type:        schema.TypeList,
				Description: "Allowed profiles",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					return []interface{}{"*"}, nil
				},
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			allowedLocalhostFilesKey: {
				Type:        schema.TypeList,
				Description: "Allowed local host files",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					return []interface{}{"*"}, nil
				},
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	},
}

func expandSeccomp(data []interface{}) (seccomp *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp) {
	if len(data) == 0 || data[0] == nil {
		return seccomp
	}

	seccompData, _ := data[0].(map[string]interface{})

	seccomp = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp{
		AllowedProfiles:       make([]*string, 0),
		AllowedLocalhostFiles: make([]*string, 0),
	}

	if v, ok := seccompData[allowedProfilesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			rawString := raw.(string)
			seccomp.AllowedProfiles = append(seccomp.AllowedProfiles, &rawString)
		}
	}

	if v, ok := seccompData[allowedLocalhostFilesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			rawString := raw.(string)
			seccomp.AllowedLocalhostFiles = append(seccomp.AllowedLocalhostFiles, &rawString)
		}
	}

	return seccomp
}

func flattenSeccomp(seccomp *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1CustomSeccomp) (data []interface{}) {
	if seccomp == nil {
		return data
	}

	flattenSeccomp := make(map[string]interface{})

	if seccomp.AllowedProfiles != nil {
		aps := make([]interface{}, 0)

		for _, ap := range seccomp.AllowedProfiles {
			aps = append(aps, *ap)
		}

		flattenSeccomp[allowedProfilesKey] = aps
	}

	if seccomp.AllowedLocalhostFiles != nil {
		alfs := make([]interface{}, 0)

		for _, alf := range seccomp.AllowedLocalhostFiles {
			alfs = append(alfs, *alf)
		}

		flattenSeccomp[allowedLocalhostFilesKey] = alfs
	}

	return []interface{}{flattenSeccomp}
}
