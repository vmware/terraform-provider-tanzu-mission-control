// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzukubernetescluster

import (
	"encoding/json"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	tanzukubernetesclustermodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var (
	nodePoolsArrayField  = tfModelConverterHelper.BuildArrayField("nodePools")
	coreAddonsArrayField = tfModelConverterHelper.BuildArrayField("coreAddons")
)

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey:                  tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
	ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("fullName", "provisionerName"),
	common.MetaKey:           common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	SpecKey: &tfModelConverterHelper.BlockToStruct{
		ClusterGroupNameKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "clusterGroupName"),
		ImageRegistryKey:    tfModelConverterHelper.BuildDefaultModelPath("spec", "imageRegistry"),
		ProxyNameKey:        tfModelConverterHelper.BuildDefaultModelPath("spec", "proxyName"),
		TMCManagedKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "tmcManaged"),
		KubeConfigKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "kubeconfig"),
		TopologyKey: &tfModelConverterHelper.BlockToStruct{
			ClusterClassKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", "clusterClass"),
			VersionKey:      tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", "version"),
			NodePoolKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					NameKey:               tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", nodePoolsArrayField, "fullName", "name"),
					common.DescriptionKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", nodePoolsArrayField, "meta", "description"),
					SpecKey: &tfModelConverterHelper.BlockToStruct{
						WorkerClassKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", nodePoolsArrayField, "spec", "class"),
						FailureDomainKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", nodePoolsArrayField, "spec", "failureDomain"),
						common.MetaKey: &tfModelConverterHelper.BlockToStruct{
							common.LabelsKey:      tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", nodePoolsArrayField, "spec", "metadata", "labels"),
							common.AnnotationsKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", nodePoolsArrayField, "spec", "metadata", "annotations"),
						},
						ReplicasKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", nodePoolsArrayField, "spec", "replicas"),
						OSImageKey: &tfModelConverterHelper.BlockToStruct{
							NameKey:    tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", nodePoolsArrayField, "spec", "osImage", "name"),
							OSArchKey:  tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", nodePoolsArrayField, "spec", "osImage", "arch"),
							VersionKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", nodePoolsArrayField, "spec", "osImage", "version"),
						},
						OverridesKey: &tfModelConverterHelper.EvaluatedField{
							Field:    tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", nodePoolsArrayField, "spec", tfModelConverterHelper.BuildArrayField("overrides")),
							EvalFunc: tfModelConverterHelper.EvaluationFunc(evaluateClusterVariables),
						},
					},
				},
			},
			CoreAddonKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					TypeKey:     tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", coreAddonsArrayField, "type"),
					ProviderKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", coreAddonsArrayField, "provider"),
				},
			},
			NetworkKey: &tfModelConverterHelper.BlockToStruct{
				PodCIDRBlocksKey:     tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", "network", "pods", "cidrBlocks"),
				ServiceCIDRBlocksKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", "network", "services", "cidrBlocks"),
				ServiceDomainKey:     tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", "network", "serviceDomain"),
			},
			ControlPlaneKey: &tfModelConverterHelper.BlockToStruct{
				ReplicasKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", "controlPlane", "replicas"),
				OSImageKey: &tfModelConverterHelper.BlockToStruct{
					NameKey:    tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", "controlPlane", "osImage", "name"),
					OSArchKey:  tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", "controlPlane", "osImage", "arch"),
					VersionKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", "controlPlane", "osImage", "version"),
				},
				common.MetaKey: &tfModelConverterHelper.BlockToStruct{
					common.LabelsKey:      tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", "controlPlane", "metadata", "labels"),
					common.AnnotationsKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", "controlPlane", "metadata", "annotations"),
				},
			},
			ClusterVariablesKey: &tfModelConverterHelper.EvaluatedField{
				Field:    tfModelConverterHelper.BuildDefaultModelPath("spec", "topology", tfModelConverterHelper.BuildArrayField("variables")),
				EvalFunc: tfModelConverterHelper.EvaluationFunc(evaluateClusterVariables),
			},
		},
	},
}

var tfModelResourceConverter = tfModelConverterHelper.TFSchemaModelConverter[*tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster]{
	TFModelMap: tfModelResourceMap,
}

func evaluateClusterVariables(mode tfModelConverterHelper.EvaluationMode, value interface{}) interface{} {
	var (
		variablesData interface{}
		err           error
	)

	if mode == tfModelConverterHelper.ConstructModel {
		variablesData = make([]interface{}, 0)
		overridesTFData := value.(string)
		overridesTFJSON := map[string]interface{}{}

		err = json.Unmarshal([]byte(overridesTFData), &overridesTFJSON)

		if err != nil {
			return nil
		}

		for k, v := range overridesTFJSON {
			ov := make(map[string]interface{})
			ov["name"] = k
			ov["value"] = v
			variablesData = append(variablesData.([]interface{}), ov)
		}
	} else {
		overridesTFJSON := map[string]interface{}{}
		overridesModelData := value.([]interface{})

		for _, ov := range overridesModelData {
			k := ov.(map[string]interface{})["name"].(string)
			v := ov.(map[string]interface{})["value"]
			overridesTFJSON[k] = v
		}

		overridesJSONBytes, err := json.Marshal(overridesTFJSON)
		variablesData = helper.ConvertToString(overridesJSONBytes, "")

		if err != nil {
			return ""
		}
	}

	return variablesData
}
