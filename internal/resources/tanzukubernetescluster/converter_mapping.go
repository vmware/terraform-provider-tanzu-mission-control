/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetescluster

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	tanzukubernetesclustermodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey:                  "fullName.name",
	ManagementClusterNameKey: "fullName.managementClusterName",
	ProvisionerNameKey:       "fullName.provisionerName",
	common.MetaKey:           common.MetaConverterMap,
	SpecKey: &tfModelConverterHelper.BlockToStruct{
		ClusterGroupNameKey: "spec.clusterGroupName",
		ImageRegistryKey:    "spec.imageRegistry",
		ProxyNameKey:        "spec.proxyName",
		TMCManagedKey:       "spec.tmcManaged",
		TopologyKey: &tfModelConverterHelper.BlockToStruct{
			ClusterClassKey: "spec.topology.clusterClass",
			VersionKey:      "spec.topology.version",
			NodePoolKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					NameKey:        "spec.topology.nodePools[].info.name",
					DescriptionKey: "spec.topology.nodePools[].info.description",
					SpecKey: &tfModelConverterHelper.BlockToStruct{
						WorkerClassKey:   "spec.topology.nodePools[].spec.class",
						FailureDomainKey: "spec.topology.nodePools[].spec.failureDomain",
						common.MetaKey: &tfModelConverterHelper.BlockToStruct{
							common.LabelsKey:      "spec.topology.nodePools[].spec.metadata.labels",
							common.AnnotationsKey: "spec.topology.nodePools[].spec.metadata.annotations",
						},
						ReplicasKey: "spec.topology.nodePools[].spec.replicas",
						OSImageKey: &tfModelConverterHelper.BlockToStruct{
							NameKey:    "spec.topology.nodePools[].spec.osImage.name",
							OSArchKey:  "spec.topology.nodePools[].spec.osImage.arch",
							VersionKey: "spec.topology.nodePools[].spec.osImage.version",
						},
						OverridesKey: &tfModelConverterHelper.EvaluatedField{
							Field:    "spec.topology.nodePools[].spec.overrides[]",
							EvalFunc: tfModelConverterHelper.EvaluationFunc(evaluateClusterVariables),
						},
					},
				},
			},
			CoreAddonKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					TypeKey:     "spec.topology.coreAddons[].type",
					ProviderKey: "spec.topology.coreAddons[].provider",
				},
			},
			NetworkKey: &tfModelConverterHelper.BlockToStruct{
				PodCIDRBlocksKey:     "spec.topology.network.pods.cidrBlocks",
				ServiceCIDRBlocksKey: "spec.topology.network.services.cidrBlocks",
				ServiceDomainKey:     "spec.topology.network.serviceDomain",
			},
			ControlPlaneKey: &tfModelConverterHelper.BlockToStruct{
				ReplicasKey: "spec.topology.controlPlane.replicas",
				OSImageKey: &tfModelConverterHelper.BlockToStruct{
					NameKey:    "spec.topology.controlPlane.osImage.name",
					OSArchKey:  "spec.topology.controlPlane.osImage.arch",
					VersionKey: "spec.topology.controlPlane.osImage.version",
				},
				common.MetaKey: &tfModelConverterHelper.BlockToStruct{
					common.LabelsKey:      "spec.topology.controlPlane.metadata.Labels",
					common.AnnotationsKey: "spec.topology.controlPlane.metadata.annotations",
				},
			},
			ClusterVariablesKey: &tfModelConverterHelper.EvaluatedField{
				Field:    "spec.topology.variables[]",
				EvalFunc: tfModelConverterHelper.EvaluationFunc(evaluateClusterVariables),
			},
		},
	},
}

var tfModelResourceConverter = tfModelConverterHelper.TFSchemaModelConverter[*tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster]{
	TFModelMap: tfModelResourceMap,
}

func evaluateClusterVariables(mode tfModelConverterHelper.EvaluationMode, value interface{}) interface{} {
	return nil
}
