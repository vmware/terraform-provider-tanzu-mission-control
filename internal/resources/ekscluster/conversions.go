/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func constructSpec(d *schema.ResourceData) (spec *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
	spec = &eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec{
		ClusterGroupName: clusterGroupDefaultValue,
	}

	value, ok := d.GetOk(specKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})
	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData, _ := data[0].(map[string]interface{})

	if v, ok := specData[clusterGroupKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ClusterGroupName, clusterGroupKey)
	}

	if v, ok := specData[proxyNameKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ProxyName, proxyNameKey)
	}

	if v, ok := specData[configKey]; ok {
		configData, _ := v.([]interface{})
		spec.Config = constructConfig(configData)
	}

	if v, ok := specData[nodepoolKey]; ok {
		nodepoolListData, _ := v.([]interface{})
		spec.NodePools = constructNodepools(nodepoolListData)
	}

	return spec
}

func constructConfig(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
	config := &eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig{}

	if len(data) == 0 || data[0] == nil {
		return config
	}

	configData, _ := data[0].(map[string]interface{})

	if v, ok := configData[roleArnKey]; ok {
		helper.SetPrimitiveValue(v, &config.RoleArn, roleArnKey)
	}

	if v, ok := configData[kubernetesVersionKey]; ok {
		helper.SetPrimitiveValue(v, &config.Version, kubernetesVersionKey)
	}

	if v, ok := configData[tagsKey]; ok {
		data, _ := v.(map[string]interface{})
		config.Tags = constructStringMap(data)
	}

	if v, ok := configData[kubernetesNetworkConfigKey]; ok {
		data, _ := v.([]interface{})
		config.KubernetesNetworkConfig = constructNetworkData(data)
	}

	if v, ok := configData[loggingKey]; ok {
		data, _ := v.([]interface{})
		config.Logging = constructLogging(data)
	}

	if v, ok := configData[vpcKey]; ok {
		data, _ := v.([]interface{})
		config.Vpc = constructVpc(data)
	}

	return config
}

func constructNetworkData(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterKubernetesNetworkConfig {
	nc := &eksmodel.VmwareTanzuManageV1alpha1EksclusterKubernetesNetworkConfig{}
	if len(data) == 0 || data[0] == nil {
		return nc
	}

	networkData, _ := data[0].(map[string]interface{})
	if v, ok := networkData[serviceCidrKey]; ok {
		helper.SetPrimitiveValue(v, &nc.ServiceCidr, serviceCidrKey)
	}

	return nc
}

func constructLogging(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterLogging {
	logging := &eksmodel.VmwareTanzuManageV1alpha1EksclusterLogging{}

	if len(data) == 0 || data[0] == nil {
		return logging
	}

	loggingData, _ := data[0].(map[string]interface{})

	if v, ok := loggingData[apiServerKey]; ok {
		helper.SetPrimitiveValue(v, &logging.APIServer, apiServerKey)
	}

	if v, ok := loggingData[auditKey]; ok {
		helper.SetPrimitiveValue(v, &logging.Audit, auditKey)
	}

	if v, ok := loggingData[authenticatorKey]; ok {
		helper.SetPrimitiveValue(v, &logging.Authenticator, authenticatorKey)
	}

	if v, ok := loggingData[controllerManagerKey]; ok {
		helper.SetPrimitiveValue(v, &logging.ControllerManager, controllerManagerKey)
	}

	if v, ok := loggingData[schedulerKey]; ok {
		helper.SetPrimitiveValue(v, &logging.Scheduler, schedulerKey)
	}

	return logging
}

func constructVpc(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterVPCConfig {
	vpc := &eksmodel.VmwareTanzuManageV1alpha1EksclusterVPCConfig{}

	if len(data) == 0 || data[0] == nil {
		return vpc
	}

	vpcData, _ := data[0].(map[string]interface{})

	if v, ok := vpcData[enablePrivateAccessKey]; ok {
		helper.SetPrimitiveValue(v, &vpc.EnablePrivateAccess, enablePrivateAccessKey)
	}

	if v, ok := vpcData[enablePublicAccessKey]; ok {
		helper.SetPrimitiveValue(v, &vpc.EnablePublicAccess, enablePublicAccessKey)
	}

	if v, ok := vpcData[publicAccessCidrsKey]; ok {
		data, _ := v.([]interface{})
		vpc.PublicAccessCidrs = constructStringList(data)
	}

	if v, ok := vpcData[securityGroupsKey]; ok {
		data, _ := v.([]interface{})
		vpc.SecurityGroups = constructStringList(data)
	}

	if v, ok := vpcData[subnetIdsKey]; ok {
		data, _ := v.([]interface{})
		vpc.SubnetIds = constructStringList(data)
	}

	return vpc
}

func constructNodepools(nodepoolsDefData []interface{}) []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
	nodepools := []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{}

	for _, npDefData := range nodepoolsDefData {
		data, _ := npDefData.(map[string]interface{})
		np := constructNodepoolDef(data)
		nodepools = append(nodepools, np)
	}

	return nodepools
}

func constructNodepoolDef(nodepoolDefData map[string]interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
	definition := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{}

	if v, ok := nodepoolDefData[infoKey]; ok {
		data, _ := v.([]interface{})
		definition.Info = constructNodepoolInfo(data)
	}

	if v, ok := nodepoolDefData[specKey]; ok {
		data, _ := v.([]interface{})
		definition.Spec = constructNodepoolSpec(data)
	}

	return definition
}

func constructNodepoolInfo(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo {
	info := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo{}

	if len(data) == 0 || data[0] == nil {
		return info
	}

	nodepoolInfoData, _ := data[0].(map[string]interface{})

	if v, ok := nodepoolInfoData[nameKey]; ok {
		helper.SetPrimitiveValue(v, &info.Name, nameKey)
	}

	if v, ok := nodepoolInfoData[common.DescriptionKey]; ok {
		helper.SetPrimitiveValue(v, &info.Description, common.DescriptionKey)
	}

	return info
}

func constructNodepoolSpec(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
	spec := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec{}

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData, _ := data[0].(map[string]interface{})

	if v, ok := specData[roleArnKey]; ok {
		helper.SetPrimitiveValue(v, &spec.RoleArn, roleArnKey)
	}

	if v, ok := specData[amiTypeKey]; ok {
		helper.SetPrimitiveValue(v, &spec.AmiType, amiTypeKey)
	}

	if v, ok := specData[capacityTypeKey]; ok {
		helper.SetPrimitiveValue(v, &spec.CapacityType, capacityTypeKey)
	}

	if v, ok := specData[rootDiskSizeKey]; ok {
		helper.SetPrimitiveValue(v, &spec.RootDiskSize, nameKey)
	}

	if v, ok := specData[tagsKey]; ok {
		data, _ := v.(map[string]interface{})
		spec.Tags = constructStringMap(data)
	}

	if v, ok := specData[nodeLabelsKey]; ok {
		data, _ := v.(map[string]interface{})
		spec.NodeLabels = constructStringMap(data)
	}

	if v, ok := specData[subnetIdsKey]; ok {
		data, _ := v.([]interface{})
		spec.SubnetIds = constructStringList(data)
	}

	if v, ok := specData[launchTemplateKey]; ok {
		data, _ := v.([]interface{})
		spec.LaunchTemplate = constructLaunchTemplate(data)
	}

	if v, ok := specData[remoteAccessKey]; ok {
		data, _ := v.([]interface{})
		spec.RemoteAccess = constructRemoteAccess(data)
	}

	if v, ok := specData[scalingConfigKey]; ok {
		data, _ := v.([]interface{})
		spec.ScalingConfig = constructScalingConfig(data)
	}

	if v, ok := specData[updateConfigKey]; ok {
		data, _ := v.([]interface{})
		spec.UpdateConfig = constructUpdateConfig(data)
	}

	if v, ok := specData[taintsKey]; ok {
		data, _ := v.([]interface{})
		spec.Taints = constructTaints(data)
	}

	if v, ok := specData[instanceTypesKey]; ok {
		data, _ := v.([]interface{})
		spec.InstanceTypes = constructStringList(data)
	}

	return spec
}

func constructLaunchTemplate(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate {
	lt := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate{}

	if len(data) == 0 || data[0] == nil {
		return lt
	}

	ltData, _ := data[0].(map[string]interface{})

	if v, ok := ltData[idKey]; ok {
		helper.SetPrimitiveValue(v, &lt.ID, idKey)
	}

	if v, ok := ltData[nameKey]; ok {
		helper.SetPrimitiveValue(v, &lt.Name, nameKey)
	}

	if v, ok := ltData[versionKey]; ok {
		helper.SetPrimitiveValue(v, &lt.Version, versionKey)
	}

	return lt
}

func constructRemoteAccess(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess {
	ra := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess{}

	if len(data) == 0 || data[0] == nil {
		return ra
	}

	raData, _ := data[0].(map[string]interface{})

	if v, ok := raData[sshKeyKey]; ok {
		helper.SetPrimitiveValue(v, &ra.SSHKey, sshKeyKey)
	}

	if v, ok := raData[securityGroupsKey]; ok {
		data, _ := v.([]interface{})
		ra.SecurityGroups = constructStringList(data)
	}

	return ra
}

func constructScalingConfig(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig {
	sc := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig{}

	if len(data) == 0 || data[0] == nil {
		return sc
	}

	scData, _ := data[0].(map[string]interface{})

	if v, ok := scData[desiredSizeKey]; ok {
		helper.SetPrimitiveValue(v, &sc.DesiredSize, desiredSizeKey)
	}

	if v, ok := scData[maxSizeKey]; ok {
		helper.SetPrimitiveValue(v, &sc.MaxSize, maxSizeKey)
	}

	if v, ok := scData[minSizeKey]; ok {
		helper.SetPrimitiveValue(v, &sc.MinSize, minSizeKey)
	}

	return sc
}

func constructUpdateConfig(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig {
	uc := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig{}

	if len(data) == 0 || data[0] == nil {
		return uc
	}

	ucData, _ := data[0].(map[string]interface{})

	if v, ok := ucData[maxUnavailableNodesKey]; ok {
		helper.SetPrimitiveValue(v, &uc.MaxUnavailableNodes, maxUnavailableNodesKey)
	}

	if v, ok := ucData[maxUnavailablePercentageKey]; ok {
		helper.SetPrimitiveValue(v, &uc.MaxUnavailablePercentage, maxUnavailablePercentageKey)
	}

	return uc
}

func constructTaints(taintsData []interface{}) []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint {
	taints := []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint{}

	for _, data := range taintsData {
		taint := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaint{}
		tdata, _ := data.(map[string]interface{})

		if v, ok := tdata[effectKey]; ok {
			data, _ := v.(string)
			switch eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(data) {
			case eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectEFFECTUNSPECIFIED:
				taint.Effect = eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectEFFECTUNSPECIFIED.Pointer()
			case eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOEXECUTE:
				taint.Effect = eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOEXECUTE.Pointer()
			case eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOSCHEDULE:
				taint.Effect = eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectNOSCHEDULE.Pointer()
			case eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE:
				taint.Effect = eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE.Pointer()
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

func constructStringMap(data map[string]interface{}) map[string]string {
	out := make(map[string]string)

	for k, v := range data {
		var value string

		helper.SetPrimitiveValue(v, &value, valueKey)

		out[k] = value
	}

	return out
}

func constructStringList(data []interface{}) []string {
	out := make([]string, 0, len(data))

	for _, v := range data {
		var value string

		helper.SetPrimitiveValue(v, &value, "")

		out = append(out, value)
	}

	return out
}
