/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetescluster

import (
	"encoding/json"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	tanzukubernetesclustermodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey:                  tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "fullName", "name"),
	ManagementClusterNameKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "fullName", "managementClusterName"),
	ProvisionerNameKey:       tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "fullName", "provisionerName"),
	common.MetaKey:           common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	SpecKey: &tfModelConverterHelper.BlockToStruct{
		ClusterGroupNameKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "clusterGroupName"),
		ImageRegistryKey:    tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "imageRegistry"),
		ProxyNameKey:        tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "proxyName"),
		TMCManagedKey:       tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "tmcManaged"),
		TopologyKey: &tfModelConverterHelper.BlockToStruct{
			ClusterClassKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", "clusterClass"),
			VersionKey:      tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", "version"),
			NodePoolKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					NameKey:        tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("nodePools"), "info", "name"),
					DescriptionKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("nodePools"), "info", "description"),
					SpecKey: &tfModelConverterHelper.BlockToStruct{
						WorkerClassKey:   tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("nodePools"), "spec", "class"),
						FailureDomainKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("nodePools"), "spec", "failureDomain"),
						common.MetaKey: &tfModelConverterHelper.BlockToStruct{
							common.LabelsKey:      tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("nodePools"), "spec", "metadata", "labels"),
							common.AnnotationsKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("nodePools"), "spec", "metadata", "annotations"),
						},
						ReplicasKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("nodePools"), "spec", "replicas"),
						OSImageKey: &tfModelConverterHelper.BlockToStruct{
							NameKey:    tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("nodePools"), "spec", "osImage", "name"),
							OSArchKey:  tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("nodePools"), "spec", "osImage", "arch"),
							VersionKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("nodePools"), "spec", "osImage", "version"),
						},
						OverridesKey: &tfModelConverterHelper.EvaluatedField{
							Field:    tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("nodePools"), "spec", tfModelConverterHelper.BuildArrayField("overrides")),
							EvalFunc: tfModelConverterHelper.EvaluationFunc(evaluateClusterVariables),
						},
					},
				},
			},
			CoreAddonKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					TypeKey:     tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("coreAddons"), "type"),
					ProviderKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("coreAddons"), "provider"),
				},
			},
			NetworkKey: &tfModelConverterHelper.BlockToStruct{
				PodCIDRBlocksKey:     tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", "network", "pods", "cidrBlocks"),
				ServiceCIDRBlocksKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", "network", "services", "cidrBlocks"),
				ServiceDomainKey:     tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", "network", "serviceDomain"),
			},
			ControlPlaneKey: &tfModelConverterHelper.BlockToStruct{
				ReplicasKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", "controlPlane", "replicas"),
				OSImageKey: &tfModelConverterHelper.BlockToStruct{
					NameKey:    tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", "controlPlane", "osImage", "name"),
					OSArchKey:  tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", "controlPlane", "osImage", "arch"),
					VersionKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", "controlPlane", "osImage", "version"),
				},
				common.MetaKey: &tfModelConverterHelper.BlockToStruct{
					common.LabelsKey:      tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", "controlPlane", "metadata", "labels"),
					common.AnnotationsKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", "controlPlane", "metadata", "annotations"),
				},
			},
			ClusterVariablesKey: &tfModelConverterHelper.EvaluatedField{
				Field:    tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "topology", tfModelConverterHelper.BuildArrayField("variables")),
				EvalFunc: tfModelConverterHelper.EvaluationFunc(evaluateClusterVariables),
			},
		},
	},
}

var tfModelResourceConverter = tfModelConverterHelper.TFSchemaModelConverter[*tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster]{
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
