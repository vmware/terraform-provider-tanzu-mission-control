// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipemutationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation"
)

var PodSecuritySchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The pod security schema",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			allowPrivilegeEscalationKey: booleanConditionSchema,
			capabilitiesAddKey:          operationSchema,
			capabilitiesDropKey:         operationSchema,
			fsGroupKey:                  ipConditionSchema,
			privilegedKey:               booleanConditionSchema,
			readOnlyRootFilesystemKey:   booleanConditionSchema,
			runAsGroupKey:               ipConditionSchema,
			runAsNonRootKey:             booleanConditionSchema,
			runAsUserKey:                ipConditionSchema,
			seLinuxOptionsKey:           allowedSELinuxOptions,
			supplementalGroupsKey:       conditionSchemaArrays,
		},
	},
}

var booleanConditionSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			conditionKey: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"Always", "IfFieldDoesNotExist", "IfFieldExists"}, false),
			},
			valueKey: {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	},
}
var ipConditionSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			conditionKey: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"Always", "IfFieldDoesNotExist", "IfFieldExists"}, false),
			},
			valueKey: {
				Type:         schema.TypeFloat,
				Required:     true,
				ValidateFunc: validation.FloatBetween(0, 65535),
			},
		},
	},
}

var conditionSchemaArrays = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			conditionKey: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Always", "IfFieldDoesNotExist", "IfFieldExists"}, false),
			},
			valuesKey: {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
		},
	},
}
var operationSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Run as user",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			operationKey: {
				Type:         schema.TypeString,
				Description:  "Rule",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"merge", "override", "prune"}, false),
			},
			valuesKey: {
				Type:        schema.TypeList,
				Description: "Values is an array of string values",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	},
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
			conditionKey: {
				Type:         schema.TypeString,
				Description:  "SELinux condition",
				Optional:     true,
				Default:      "",
				ValidateFunc: validation.StringInSlice([]string{"Always", "IfFieldDoesNotExist", "IfFieldExists"}, false),
			},
		},
	},
}

func ConstructPodSecurity(data []interface{}) (podSecuritySpec *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity) {
	if len(data) == 0 || data[0] == nil {
		return podSecuritySpec
	}

	podSecurityData, _ := data[0].(map[string]interface{})

	podSecuritySpec = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity{}

	if v, ok := podSecurityData[allowPrivilegeEscalationKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			podSecuritySpec.AllowPrivilegeEscalation = expandAllowPrivilegeEscalation(v1)
		}
	}

	if v, ok := podSecurityData[capabilitiesAddKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			podSecuritySpec.CapabilitiesAdd = expandCapabilitiesAddKey(v1)
		}
	}

	if v, ok := podSecurityData[capabilitiesDropKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			podSecuritySpec.CapabilitiesDrop = expandCapabilitiesDrop(v1)
		}
	}

	if v, ok := podSecurityData[fsGroupKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			podSecuritySpec.FsGroup = expandFsGroup(v1)
		}
	}

	if v, ok := podSecurityData[privilegedKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			podSecuritySpec.Privileged = expandPrivileged(v1)
		}
	}

	if v, ok := podSecurityData[readOnlyRootFilesystemKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			podSecuritySpec.ReadOnlyRootFilesystem = expandReadOnlyRootFilesystem(v1)
		}
	}

	if v, ok := podSecurityData[runAsGroupKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			podSecuritySpec.RunAsGroup = expandReadOnlyRunAsGroup(v1)
		}
	}

	if v, ok := podSecurityData[runAsNonRootKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			podSecuritySpec.RunAsNonRoot = expandRunAsNonRoot(v1)
		}
	}

	if v, ok := podSecurityData[runAsUserKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			podSecuritySpec.RunAsUser = expandRunAsUser(v1)
		}
	}

	if v, ok := podSecurityData[seLinuxOptionsKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			podSecuritySpec.SeLinuxOptions = expandSeLinuxOptions(v1)
		}
	}

	if v, ok := podSecurityData[supplementalGroupsKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			podSecuritySpec.SupplementalGroups = expandSupplementalGroups(v1)
		}
	}

	return podSecuritySpec
}

func FlattenPodSecurity(podSecurity *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity) (data []interface{}) {
	if podSecurity == nil {
		return data
	}

	flattenPodSecurity := make(map[string]interface{})

	if podSecurity.AllowPrivilegeEscalation != nil {
		flattenPodSecurity[allowPrivilegeEscalationKey] = flattenAllowPrivilegeEscalation(podSecurity.AllowPrivilegeEscalation)
	}

	if podSecurity.CapabilitiesAdd != nil {
		flattenPodSecurity[capabilitiesAddKey] = flattenCapabilitiesAdd(podSecurity.CapabilitiesAdd)
	}

	if podSecurity.CapabilitiesDrop != nil {
		flattenPodSecurity[capabilitiesDropKey] = flattenCapabilitiesDrop(podSecurity.CapabilitiesDrop)
	}

	if podSecurity.FsGroup != nil {
		flattenPodSecurity[fsGroupKey] = flattenFsGroup(podSecurity.FsGroup)
	}

	if podSecurity.Privileged != nil {
		flattenPodSecurity[privilegedKey] = flattenPrivileged(podSecurity.Privileged)
	}

	if podSecurity.ReadOnlyRootFilesystem != nil {
		flattenPodSecurity[readOnlyRootFilesystemKey] = flattenReadOnlyRootFilesystem(podSecurity.ReadOnlyRootFilesystem)
	}

	if podSecurity.RunAsGroup != nil {
		flattenPodSecurity[runAsGroupKey] = flattenRunAsGroup(podSecurity.RunAsGroup)
	}

	if podSecurity.RunAsNonRoot != nil {
		flattenPodSecurity[runAsNonRootKey] = flattenRunAsNonRoot(podSecurity.RunAsNonRoot)
	}

	if podSecurity.RunAsUser != nil {
		flattenPodSecurity[runAsUserKey] = flattenRunAsUser(podSecurity.RunAsUser)
	}

	if podSecurity.SeLinuxOptions != nil {
		flattenPodSecurity[seLinuxOptionsKey] = flattenSeLinuxOptions(podSecurity.SeLinuxOptions)
	}

	if podSecurity.SupplementalGroups != nil {
		flattenPodSecurity[supplementalGroupsKey] = flattenSupplementalGroups(podSecurity.SupplementalGroups)
	}

	return []interface{}{flattenPodSecurity}
}

func expandAllowPrivilegeEscalation(data []interface{}) (privilegeEscalation *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation) {
	if data == nil {
		return privilegeEscalation
	}

	privilegeEscalationData, _ := data[0].(map[string]interface{})

	privilegeEscalation = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation{}

	if v, ok := privilegeEscalationData[conditionKey]; ok {
		privilegeEscalation.Condition = helper.StringPointer(v.(string))
	}

	if v, ok := privilegeEscalationData[valueKey]; ok {
		privilegeEscalation.Value = helper.BoolPointer(v.(bool))
	}

	return privilegeEscalation
}

func flattenAllowPrivilegeEscalation(allowPrivilegeEscalation *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation) (data []interface{}) {
	if allowPrivilegeEscalation == nil {
		return data
	}

	flattenAllowPrivilegeEscalation := make(map[string]interface{})

	flattenAllowPrivilegeEscalation[conditionKey] = *allowPrivilegeEscalation.Condition
	flattenAllowPrivilegeEscalation[valueKey] = *allowPrivilegeEscalation.Value

	return []interface{}{flattenAllowPrivilegeEscalation}
}

func expandCapabilitiesAddKey(data []interface{}) (capabilitiesAdd *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd) {
	if data == nil {
		return capabilitiesAdd
	}

	capabilitiesAddData, _ := data[0].(map[string]interface{})

	capabilitiesAdd = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd{}

	if v, ok := capabilitiesAddData[operationKey]; ok {
		capabilitiesAdd.Operation = helper.StringPointer(v.(string))
	}

	if v, ok := capabilitiesAddData[valuesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			rawString := raw.(string)
			capabilitiesAdd.Values = append(capabilitiesAdd.Values, rawString)
		}
	}

	return capabilitiesAdd
}

func flattenCapabilitiesAdd(capabilitiesAdd *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd) (data []interface{}) {
	if capabilitiesAdd == nil {
		return data
	}

	flattenCapabilitiesAdd := make(map[string]interface{})

	flattenCapabilitiesAdd[operationKey] = *capabilitiesAdd.Operation

	values := make([]interface{}, 0)

	for _, value := range capabilitiesAdd.Values {
		values = append(values, value)
	}

	flattenCapabilitiesAdd[valuesKey] = values

	return []interface{}{flattenCapabilitiesAdd}
}

func expandCapabilitiesDrop(data []interface{}) (capabilitiesDrop *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop) {
	if data == nil {
		return capabilitiesDrop
	}

	capabilitiesDropData, _ := data[0].(map[string]interface{})

	capabilitiesDrop = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop{}

	if v, ok := capabilitiesDropData[operationKey]; ok {
		capabilitiesDrop.Operation = helper.StringPointer(v.(string))
	}

	if v, ok := capabilitiesDropData[valuesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			rawString := raw.(string)
			capabilitiesDrop.Values = append(capabilitiesDrop.Values, rawString)
		}
	}

	return capabilitiesDrop
}

func flattenCapabilitiesDrop(capabilitiesDrop *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop) (data []interface{}) {
	if capabilitiesDrop == nil {
		return data
	}

	flattenCapabilitiesDrop := make(map[string]interface{})

	flattenCapabilitiesDrop[operationKey] = *capabilitiesDrop.Operation

	values := make([]interface{}, 0)

	for _, value := range capabilitiesDrop.Values {
		values = append(values, value)
	}

	flattenCapabilitiesDrop[valuesKey] = values

	return []interface{}{flattenCapabilitiesDrop}
}

func expandFsGroup(data []interface{}) (fsGroup *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup) {
	if data == nil {
		return fsGroup
	}

	fsGroupData, _ := data[0].(map[string]interface{})

	fsGroup = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup{}

	if v, ok := fsGroupData[conditionKey]; ok {
		fsGroup.Condition = helper.StringPointer(v.(string))
	}

	if v, ok := fsGroupData[valueKey]; ok {
		fsGroup.Value = helper.Float64Pointer(v.(float64))
	}

	return fsGroup
}

func flattenFsGroup(fsGroup *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup) (data []interface{}) {
	if fsGroup == nil {
		return data
	}

	flattenFsGroup := make(map[string]interface{})

	flattenFsGroup[conditionKey] = *fsGroup.Condition
	flattenFsGroup[valueKey] = *fsGroup.Value

	return []interface{}{flattenFsGroup}
}

func expandPrivileged(data []interface{}) (privileged *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged) {
	if data == nil {
		return privileged
	}

	privilegedData, _ := data[0].(map[string]interface{})

	privileged = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged{}

	if v, ok := privilegedData[conditionKey]; ok {
		privileged.Condition = helper.StringPointer(v.(string))
	}

	if v, ok := privilegedData[valueKey]; ok {
		privileged.Value = helper.BoolPointer(v.(bool))
	}

	return privileged
}

func flattenPrivileged(privileged *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged) (data []interface{}) {
	if privileged == nil {
		return data
	}

	flattenPrivileged := make(map[string]interface{})

	flattenPrivileged[conditionKey] = *privileged.Condition
	flattenPrivileged[valueKey] = *privileged.Value

	return []interface{}{flattenPrivileged}
}

func expandReadOnlyRootFilesystem(data []interface{}) (readonlyRootFilesystem *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem) {
	if data == nil {
		return readonlyRootFilesystem
	}

	readonlyRootFilesystemData, _ := data[0].(map[string]interface{})

	readonlyRootFilesystem = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem{}

	if v, ok := readonlyRootFilesystemData[conditionKey]; ok {
		readonlyRootFilesystem.Condition = helper.StringPointer(v.(string))
	}

	if v, ok := readonlyRootFilesystemData[valueKey]; ok {
		readonlyRootFilesystem.Value = helper.BoolPointer(v.(bool))
	}

	return readonlyRootFilesystem
}

func flattenReadOnlyRootFilesystem(readOnlyRootFilesystem *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem) (data []interface{}) {
	if readOnlyRootFilesystem == nil {
		return data
	}

	flattenReadOnlyRootFilesystem := make(map[string]interface{})

	flattenReadOnlyRootFilesystem[conditionKey] = *readOnlyRootFilesystem.Condition
	flattenReadOnlyRootFilesystem[valueKey] = *readOnlyRootFilesystem.Value

	return []interface{}{flattenReadOnlyRootFilesystem}
}

func expandReadOnlyRunAsGroup(data []interface{}) (runAsGroup *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup) {
	if data == nil {
		return runAsGroup
	}

	runAsGroupData, _ := data[0].(map[string]interface{})

	runAsGroup = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup{}

	if v, ok := runAsGroupData[conditionKey]; ok {
		runAsGroup.Condition = helper.StringPointer(v.(string))
	}

	if v, ok := runAsGroupData[valueKey]; ok {
		runAsGroup.Value = helper.Float64Pointer(v.(float64))
	}

	return runAsGroup
}

func flattenRunAsGroup(runAsGroup *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup) (data []interface{}) {
	if runAsGroup == nil {
		return data
	}

	flattenRunAsGroup := make(map[string]interface{})

	flattenRunAsGroup[conditionKey] = *runAsGroup.Condition
	flattenRunAsGroup[valueKey] = *runAsGroup.Value

	return []interface{}{flattenRunAsGroup}
}

func expandRunAsNonRoot(data []interface{}) (runAsNonRoot *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot) {
	if data == nil {
		return runAsNonRoot
	}

	runAsNonRootData, _ := data[0].(map[string]interface{})

	runAsNonRoot = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot{}

	if v, ok := runAsNonRootData[conditionKey]; ok {
		runAsNonRoot.Condition = helper.StringPointer(v.(string))
	}

	if v, ok := runAsNonRootData[valueKey]; ok {
		runAsNonRoot.Value = helper.BoolPointer(v.(bool))
	}

	return runAsNonRoot
}

func flattenRunAsNonRoot(runAsNonRoot *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot) (data []interface{}) {
	if runAsNonRoot == nil {
		return data
	}

	flattenRunAsNonRoot := make(map[string]interface{})

	flattenRunAsNonRoot[conditionKey] = *runAsNonRoot.Condition
	flattenRunAsNonRoot[valueKey] = *runAsNonRoot.Value

	return []interface{}{flattenRunAsNonRoot}
}

func expandRunAsUser(data []interface{}) (runAsUser *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser) {
	if data == nil {
		return runAsUser
	}

	runAsUserData, _ := data[0].(map[string]interface{})

	runAsUser = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser{}

	if v, ok := runAsUserData[conditionKey]; ok {
		runAsUser.Condition = helper.StringPointer(v.(string))
	}

	if v, ok := runAsUserData[valueKey]; ok {
		runAsUser.Value = helper.Float64Pointer(v.(float64))
	}

	return runAsUser
}

func flattenRunAsUser(runAsUser *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser) (data []interface{}) {
	if runAsUser == nil {
		return data
	}

	flattenRunAsUser := make(map[string]interface{})

	flattenRunAsUser[conditionKey] = *runAsUser.Condition
	flattenRunAsUser[valueKey] = *runAsUser.Value

	return []interface{}{flattenRunAsUser}
}

func expandSeLinuxOptions(data []interface{}) (seLinuxOptions *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions) {
	if data == nil {
		return seLinuxOptions
	}

	seLinuxOptionsData, _ := data[0].(map[string]interface{})

	seLinuxOptions = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions{}

	if v, ok := seLinuxOptionsData[conditionKey]; ok {
		seLinuxOptions.Condition = helper.StringPointer(v.(string))
	}

	value := &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptionsValue{}

	if v, ok := seLinuxOptionsData[levelKey]; ok {
		value.Level = v.(string)
	}

	if v, ok := seLinuxOptionsData[roleKey]; ok {
		value.Role = v.(string)
	}

	if v, ok := seLinuxOptionsData[typeKey]; ok {
		value.Type = v.(string)
	}

	if v, ok := seLinuxOptionsData[userKey]; ok {
		value.User = v.(string)
	}

	seLinuxOptions.Value = value

	return seLinuxOptions
}

func flattenSeLinuxOptions(seLinuxOptions *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions) (data []interface{}) {
	if seLinuxOptions == nil {
		return data
	}

	flattenSeLinuxOptions := make(map[string]interface{})

	flattenSeLinuxOptions[conditionKey] = *seLinuxOptions.Condition

	if seLinuxOptions.Value != nil {
		flattenSeLinuxOptions[levelKey] = seLinuxOptions.Value.Level
		flattenSeLinuxOptions[roleKey] = seLinuxOptions.Value.Role
		flattenSeLinuxOptions[typeKey] = seLinuxOptions.Value.Type
		flattenSeLinuxOptions[userKey] = seLinuxOptions.Value.User
	}

	return []interface{}{flattenSeLinuxOptions}
}

func expandSupplementalGroups(data []interface{}) (supplementalGroups *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups) {
	if data == nil {
		return supplementalGroups
	}

	supplementalGroupsData, _ := data[0].(map[string]interface{})

	supplementalGroups = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups{}

	if v, ok := supplementalGroupsData[conditionKey]; ok {
		supplementalGroups.Condition = helper.StringPointer(v.(string))
	}

	if v, ok := supplementalGroupsData[valuesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			rawFloat := raw.(float64)
			supplementalGroups.Values = append(supplementalGroups.Values, &rawFloat)
		}
	}

	return supplementalGroups
}

func flattenSupplementalGroups(supplementalGroups *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups) (data []interface{}) {
	if supplementalGroups == nil {
		return data
	}

	flattenSupplementalGroups := make(map[string]interface{})

	flattenSupplementalGroups[conditionKey] = *supplementalGroups.Condition

	values := make([]interface{}, 0)

	for _, value := range supplementalGroups.Values {
		values = append(values, *value)
	}

	flattenSupplementalGroups[valuesKey] = values

	return []interface{}{flattenSupplementalGroups}
}
