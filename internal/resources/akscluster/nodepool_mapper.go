/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
)

// ConstructNodepools extracts all nodepool sections from schema data and converts them to a list of Nodepool Objects.
func ConstructNodepools(data *schema.ResourceData) []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool {
	cfn := extractClusterFullName(data)
	specData := extractClusterSpec(data)

	v := specData[nodepoolKey]
	nodepoolsData := v.([]any)

	nodepools := make([]*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool, 0, len(nodepoolsData))

	for _, d := range nodepoolsData {
		nodepoolData := d.(map[string]any)
		nodepools = append(nodepools, constructNodepool(cfn, nodepoolData))
	}

	return nodepools
}

func constructNodepool(cfn *models.VmwareTanzuManageV1alpha1AksclusterFullName, data map[string]any) *models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool {
	nodepool := &models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool{}

	nodepool.FullName = constructNodepoolFullName(cfn, data)
	nodepool.Spec = constructNodepoolSpec(data)

	return nodepool
}

func constructNodepoolFullName(cfn *models.VmwareTanzuManageV1alpha1AksclusterFullName, data map[string]any) *models.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName {
	fn := &models.VmwareTanzuManageV1alpha1AksclusterNodepoolFullName{}

	fn.OrgID = cfn.OrgID
	fn.CredentialName = cfn.CredentialName
	fn.SubscriptionID = cfn.SubscriptionID
	fn.ResourceGroupName = cfn.ResourceGroupName
	fn.AksClusterName = cfn.Name
	fn.Name = data[NameKey].(string)

	return fn
}

func constructNodepoolSpec(data map[string]any) *models.VmwareTanzuManageV1alpha1AksclusterNodepoolSpec {
	npSpec := &models.VmwareTanzuManageV1alpha1AksclusterNodepoolSpec{}
	npSpecData := extractNodepoolSpec(data)

	if v, ok := npSpecData[modeKey]; ok {
		mode := models.VmwareTanzuManageV1alpha1AksclusterNodepoolMode(v.(string))
		npSpec.Mode = &mode
	}

	if v, ok := npSpecData[typeKey]; ok {
		npType := models.VmwareTanzuManageV1alpha1AksclusterNodepoolType(v.(string))
		npSpec.Type = &npType
	}

	if v, ok := npSpecData[availabilityZonesKey]; ok {
		npSpec.AvailabilityZones = helper.SetPrimitiveList[string](v)
	}

	if v, ok := npSpecData[countKey]; ok {
		helper.SetPrimitiveValue(v, &npSpec.Count, countKey)
	}

	if v, ok := npSpecData[vmSizeKey]; ok {
		helper.SetPrimitiveValue(v, &npSpec.VMSize, vmSizeKey)
	}

	if v, ok := npSpecData[osTypeKey]; ok {
		osType := models.VmwareTanzuManageV1alpha1AksclusterNodepoolOsType(v.(string))
		npSpec.OsType = &osType
	}

	if v, ok := npSpecData[osDiskTypeKey]; ok && v != "" {
		osDiskType := models.VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType(v.(string))
		npSpec.OsDiskType = &osDiskType
	}

	if v, ok := npSpecData[osDiskSizeKey]; ok {
		helper.SetPrimitiveValue(v, &npSpec.OsDiskSizeGb, osDiskSizeKey)
	}

	if v, ok := npSpecData[maxPodsKey]; ok {
		helper.SetPrimitiveValue(v, &npSpec.MaxPods, maxPodsKey)
	}

	if v, ok := npSpecData[enableNodePublicIPKey]; ok {
		helper.SetPrimitiveValue(v, &npSpec.EnableNodePublicIP, enableNodePublicIPKey)
	}

	if v, ok := npSpecData[taintsKey]; ok {
		data, _ := v.([]interface{})
		npSpec.NodeTaints = constructTaints(data)
	}

	if v, ok := npSpecData[vnetSubnetKey]; ok {
		helper.SetPrimitiveValue(v, &npSpec.VnetSubnetID, vnetSubnetKey)
	}

	if v, ok := npSpecData[nodeLabelsKey]; ok {
		data, _ := v.(map[string]interface{})
		npSpec.NodeLabels = constructStringMap[string](data)
	}

	if v, ok := npSpecData[tagsKey]; ok {
		data, _ := v.(map[string]interface{})
		npSpec.Tags = constructStringMap[string](data)
	}

	if v, ok := npSpecData[autoscalingConfigKey]; ok {
		data, _ := v.([]interface{})
		npSpec.AutoScaling = constructAutoscalingConfig(data)
	}

	if v, ok := npSpecData[upgradeConfigKey]; ok {
		data, _ := v.([]interface{})
		npSpec.UpgradeConfig = constructUpgradeConfig(data)
	}

	return npSpec
}

func extractNodepoolSpec(data map[string]any) map[string]any {
	value, ok := data[nodepoolSpecKey]
	if !ok {
		return nil
	}

	dataSpec := value.([]any)
	if len(dataSpec) < 1 {
		return nil
	}

	// Spec schema defines max 1
	return dataSpec[0].(map[string]any)
}

func constructTaints(taintsData []interface{}) []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaint {
	taints := make([]*models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaint, 0, len(taintsData))

	for _, data := range taintsData {
		taint := &models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaint{}
		tdata, _ := data.(map[string]interface{})

		if v, ok := tdata[effectKey]; ok {
			data, _ := v.(string)
			switch models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect(data) {
			case models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectEFFECTUNSPECIFIED:
				taint.Effect = models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectEFFECTUNSPECIFIED.Pointer()
			case models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectNOEXECUTE:
				taint.Effect = models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectNOEXECUTE.Pointer()
			case models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectNOSCHEDULE:
				taint.Effect = models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectNOSCHEDULE.Pointer()
			case models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectPREFERNOSCHEDULE:
				taint.Effect = models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffectPREFERNOSCHEDULE.Pointer()
			default:
				panic("unknown taint effect")
			}
		}

		if v, ok := tdata[keyKey]; ok {
			helper.SetPrimitiveValue(v, &taint.Key, keyKey)
		}

		if v, ok := tdata[valueKey]; ok {
			helper.SetPrimitiveValue(v, &taint.Value, valueKey)
		}

		taints = append(taints, taint)
	}

	return taints
}

func constructAutoscalingConfig(data []interface{}) *models.VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig {
	if len(data) < 1 {
		return nil
	}

	// AutoscalingConfig schema defines max 1
	autoScalingData, _ := data[0].(map[string]any)
	autoscalingConfig := &models.VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig{}

	if v, ok := autoScalingData[enableKey]; ok {
		helper.SetPrimitiveValue(v, &autoscalingConfig.Enabled, enableKey)
	}

	if v, ok := autoScalingData[minCountKey]; ok {
		helper.SetPrimitiveValue(v, &autoscalingConfig.MinCount, minCountKey)
	}

	if v, ok := autoScalingData[maxCountKey]; ok {
		helper.SetPrimitiveValue(v, &autoscalingConfig.MaxCount, maxCountKey)
	}

	if v, ok := autoScalingData[scaleSetPriorityKey]; ok {
		scaleSetPriority := models.VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority(v.(string))
		autoscalingConfig.ScaleSetPriority = &scaleSetPriority
	}

	if v, ok := autoScalingData[scaleSetEvictionPolicyKey]; ok {
		scaleSetEvictionPolicy := models.VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy(v.(string))
		autoscalingConfig.ScaleSetEvictionPolicy = &scaleSetEvictionPolicy
	}

	if v, ok := autoScalingData[maxSpotPriceKey]; ok {
		helper.SetPrimitiveValue(v, &autoscalingConfig.SpotMaxPrice, maxSpotPriceKey)
	}

	return autoscalingConfig
}

func constructUpgradeConfig(data []interface{}) *models.VmwareTanzuManageV1alpha1AksclusterNodepoolUpgradeConfig {
	if len(data) < 1 {
		return nil
	}

	// UpgradeConfigData schema defines max 1
	upgradeConfigData, _ := data[0].(map[string]any)
	upgradeConfig := &models.VmwareTanzuManageV1alpha1AksclusterNodepoolUpgradeConfig{}

	if v, ok := upgradeConfigData[maxSurgeKey]; ok {
		helper.SetPrimitiveValue(v, &upgradeConfig.MaxSurge, maxSurgeKey)
	}

	return upgradeConfig
}

func ToNodepoolMap(np *models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) map[string]any {
	if np == nil {
		return nil
	}

	data := make(map[string]any)
	data[NameKey] = np.FullName.Name
	data[nodepoolSpecKey] = toNodepoolSpecMap(np.Spec)

	return data
}

func toNodepoolSpecMap(spec *models.VmwareTanzuManageV1alpha1AksclusterNodepoolSpec) []any {
	data := make(map[string]any)
	if spec == nil {
		return []any{data}
	}

	data[modeKey] = helper.PtrString(spec.Mode)
	data[typeKey] = helper.PtrString(spec.Type)
	data[availabilityZonesKey] = toInterfaceArray(spec.AvailabilityZones)
	data[countKey] = int(spec.Count)
	data[vmSizeKey] = spec.VMSize
	data[osTypeKey] = helper.PtrString(spec.OsType)
	data[osDiskTypeKey] = helper.PtrString(spec.OsDiskType)
	data[osDiskSizeKey] = int(spec.OsDiskSizeGb)
	data[maxPodsKey] = int(spec.MaxPods)
	data[enableNodePublicIPKey] = spec.EnableNodePublicIP
	data[taintsKey] = toTaintList(spec.NodeTaints)
	data[vnetSubnetKey] = spec.VnetSubnetID
	data[nodeLabelsKey] = toInterfaceMap(spec.NodeLabels)
	data[tagsKey] = toInterfaceMap(spec.Tags)
	data[autoscalingConfigKey] = toAutoscalingConfigMap(spec.AutoScaling)
	data[upgradeConfigKey] = toUpgradeConfigMap(spec.UpgradeConfig)

	return []any{data}
}

func toTaintList(t []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaint) []any {
	data := make([]any, 0, len(t))
	for _, item := range t {
		data = append(data, toTaintMap(item))
	}

	return data
}

func toTaintMap(item *models.VmwareTanzuManageV1alpha1AksclusterNodepoolTaint) map[string]any {
	data := make(map[string]any)
	if item == nil {
		return data
	}

	data[effectKey] = helper.PtrString(item.Effect)
	data[keyKey] = item.Key
	data[valueKey] = item.Value

	return data
}

func toAutoscalingConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig) []any {
	if config == nil {
		return nil
	}

	data := make(map[string]any)
	data[enableKey] = config.Enabled
	data[minCountKey] = int(config.MinCount)
	data[maxCountKey] = int(config.MaxCount)
	data[scaleSetPriorityKey] = helper.PtrString(config.ScaleSetPriority)
	data[scaleSetEvictionPolicyKey] = helper.PtrString(config.ScaleSetEvictionPolicy)
	data[maxSpotPriceKey] = float64(config.SpotMaxPrice)

	return []any{data}
}

func toUpgradeConfigMap(config *models.VmwareTanzuManageV1alpha1AksclusterNodepoolUpgradeConfig) []any {
	if config == nil {
		return nil
	}

	data := make(map[string]any)
	data[maxSurgeKey] = config.MaxSurge

	return []any{data}
}
