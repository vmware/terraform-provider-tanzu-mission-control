/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func flattenClusterSpec(item *models.VmwareTanzuManageV1alpha1EksclusterSpec) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[clusterGroupKey] = item.ClusterGroupName

	if item.Config != nil {
		data[configKey] = flattenConfig(item.Config)
	}

	if len(item.NodePools) > 0 {
		data[nodepoolKey] = flattenNodePools(item.NodePools)
	}

	if item.ProxyName != "" {
		data[proxyNameKey] = item.ProxyName
	}

	return []interface{}{data}
}

func flattenConfig(item *models.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	if item.KubernetesNetworkConfig != nil {
		data[kubernetesNetworkConfigKey] = flattenKubernetesNetworkConfig(item.KubernetesNetworkConfig)
	}

	if item.Logging != nil {
		data[loggingKey] = flattenLogging(item.Logging)
	}

	data[roleArnKey] = item.RoleArn
	data[tagsKey] = item.Tags
	data[kubernetesVersionKey] = item.Version

	if item.Vpc != nil {
		data[vpcKey] = flattenVpc(item.Vpc)
	}

	return []interface{}{data}
}

func flattenKubernetesNetworkConfig(item *models.VmwareTanzuManageV1alpha1EksclusterKubernetesNetworkConfig) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[serviceCidrKey] = item.ServiceCidr

	return []interface{}{data}
}

func flattenLogging(item *models.VmwareTanzuManageV1alpha1EksclusterLogging) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[apiServerKey] = item.APIServer
	data[auditKey] = item.Audit
	data[authenticatorKey] = item.Authenticator
	data[controllerManagerKey] = item.ControllerManager
	data[schedulerKey] = item.Scheduler

	return []interface{}{data}
}

func flattenVpc(item *models.VmwareTanzuManageV1alpha1EksclusterVPCConfig) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[enablePrivateAccessKey] = item.EnablePrivateAccess
	data[enablePublicAccessKey] = item.EnablePublicAccess

	if len(item.PublicAccessCidrs) > 0 {
		data[publicAccessCidrsKey] = item.PublicAccessCidrs
	}

	if len(item.SecurityGroups) > 0 {
		data[securityGroupsKey] = item.SecurityGroups
	}

	if len(item.SubnetIds) > 0 {
		data[subnetIdsKey] = item.SubnetIds
	}

	return []interface{}{data}
}

func flattenNodePools(arr []*models.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) []interface{} {
	data := make([]interface{}, 0, len(arr))

	if len(arr) == 0 {
		return data
	}

	for _, item := range arr {
		data = append(data, flattenNodePool(item))
	}

	return data
}

func flattenNodePool(item *models.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) map[string]interface{} {
	data := make(map[string]interface{})

	if item == nil {
		return data
	}

	if item.Info != nil {
		data[infoKey] = flattenInfo(item.Info)
	}

	if item.Spec != nil {
		data[specKey] = flattenSpec(item.Spec)
	}

	return data
}

func flattenInfo(item *models.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[common.DescriptionKey] = item.Description
	data[nameKey] = item.Name

	return []interface{}{data}
}

func flattenSpec(item *models.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[amiTypeKey] = item.AmiType
	data[capacityTypeKey] = item.CapacityType

	if len(item.InstanceTypes) > 0 {
		data[instanceTypesKey] = item.InstanceTypes
	}

	if item.LaunchTemplate != nil &&
		(item.LaunchTemplate.ID != "" || item.LaunchTemplate.Name != "" || item.LaunchTemplate.Version != "") {
		data[launchTemplateKey] = flattenLaunchTemplate(item.LaunchTemplate)
	}

	data[nodeLabelsKey] = item.NodeLabels

	if item.RemoteAccess != nil && (item.RemoteAccess.SSHKey != "" || len(item.RemoteAccess.SecurityGroups) > 0) {
		data[remoteAccessKey] = flattenRemoteAccess(item.RemoteAccess)
	}

	data[roleArnKey] = item.RoleArn

	if item.RootDiskSize != 0 {
		data[rootDiskSizeKey] = item.RootDiskSize
	}

	if item.ScalingConfig != nil {
		data[scalingConfigKey] = flattenScalingConfig(item.ScalingConfig)
	}

	if len(item.SubnetIds) > 0 {
		data[subnetIdsKey] = item.SubnetIds
	}

	data[tagsKey] = item.Tags

	if len(item.Taints) > 0 {
		data[taintsKey] = flattenTaints(item.Taints)
	}

	if item.UpdateConfig != nil {
		data[updateConfigKey] = flattenUpdateConfig(item.UpdateConfig)
	}

	return []interface{}{data}
}

func flattenLaunchTemplate(item *models.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[idKey] = item.ID
	data[nameKey] = item.Name
	data[versionKey] = item.Version

	return []interface{}{data}
}

func flattenRemoteAccess(item *models.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	if len(item.SecurityGroups) > 0 {
		data[securityGroupsKey] = item.SecurityGroups
	}

	data[sshKeyKey] = item.SSHKey

	return []interface{}{data}
}

func flattenScalingConfig(item *models.VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[desiredSizeKey] = item.DesiredSize
	data[maxSizeKey] = item.MaxSize
	data[minSizeKey] = item.MinSize

	return []interface{}{data}
}

func flattenTaints(arr []*models.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint) []interface{} {
	data := make([]interface{}, 0, len(arr))

	if len(arr) == 0 {
		return data
	}

	for _, item := range arr {
		data = append(data, flattenTaint(item))
	}

	return data
}

func flattenTaint(item *models.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint) map[string]interface{} {
	data := make(map[string]interface{})

	if item == nil {
		return data
	}

	data[effectKey] = item.Effect
	data[keyKey] = item.Key
	data[valueKey] = item.Value

	return data
}

func flattenUpdateConfig(item *models.VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	if item.MaxUnavailableNodes != "" {
		data[maxUnavailableNodesKey] = item.MaxUnavailableNodes
	}

	if item.MaxUnavailablePercentage != "" {
		data[maxUnavailablePercentageKey] = item.MaxUnavailablePercentage
	}

	return []interface{}{data}
}
