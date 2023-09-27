/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkc

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tkcnodepool "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/nodepool"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var nodepoolsDefinitionSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		NameKey: {
			Type:        schema.TypeString,
			Description: "Name of the nodepool, immutable",
			Required:    true,
		},
		common.DescriptionKey: {
			Type:        schema.TypeString,
			Description: "Description for the nodepool",
			Optional:    true,
		},
		specKey: {
			Type:        schema.TypeList,
			Description: "Spec for the cluster nodepool",
			Required:    true,
			MinItems:    1,
			Elem:        nodepoolSpecSchema,
		},
	},
}

var nodepoolSpecSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		classKey: {
			Type:        schema.TypeString,
			Description: "The name of the machine deployment class used to create the nodepool",
			Required:    true,
		},
		replicasKey: {
			Type:        schema.TypeInt,
			Description: "The replicas of the nodepool",
			Required:    true,
		},
		failureDomainKey: {
			Type:        schema.TypeString,
			Description: "The failure domain the machines will be created in",
			Optional:    true,
			Computed:    true,
		},
		overridesKey: {
			Type:        schema.TypeList,
			Description: "Overrides can be used to override cluster level variables",
			Optional:    true,
			Elem:        overridesSchema,
		},
		metadataKey: {
			Type:        schema.TypeList,
			Description: "The labels and annotations configurations of the control plane",
			Optional:    true,
			Elem:        commonClusterMetadataSchema,
		},
		osImageKey: {
			Type:        schema.TypeList,
			Description: "The OS image configuration of the control plane",
			Required:    true,
			Elem:        commonClusterOsImageSchema,
		},
	},
}

var overridesSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		tkgVsphereKey: {
			Type:        schema.TypeList,
			Description: "The TKG cluster variable configuration",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeList}, //tkgVsphereSchema,
		},
		vsphereTanzuKey: {
			Type:        schema.TypeList,
			Description: "The TKG cluster variable configuration",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeList}, //vsphereTanzuSchema,
		},
	},
}

func flattenNodePools(arr []*tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition) []interface{} {
	data := make([]interface{}, 0, len(arr))

	if len(arr) == 0 {
		return data
	}

	for _, item := range arr {
		data = append(data, flattenNodePool(item))
	}

	return data
}

func flattenNodePool(item *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition) map[string]interface{} {
	data := make(map[string]interface{})

	if item == nil {
		return data
	}

	if item.Info.Name != "" {
		data[nameKey] = item.Info.Name
	}

	if item.Info.Description != "" {
		data[common.DescriptionKey] = item.Info.Description
	}

	if item.Spec != nil {
		data[specKey] = flattenSpec(item.Spec)
	}

	return data
}

func flattenSpec(item *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	if item.Class != "" {
		data[classKey] = item.Class
	}

	if item.FailureDomain != "" {
		data[failureDomainKey] = item.FailureDomain
	}

	data[replicasKey] = item.Replicas

	if item.Metadata != nil {
		data[metadataKey] = flattenCommonClusterMetadata(item.Metadata)
	}

	if item.OsImage != nil {
		data[osImageKey] = flattenCommonClusterOsImage(item.OsImage)
	}

	if item.Overrides != nil {
		data[overridesKey] = flattenCommonClusterVariables(item.Overrides)
	}

	return []interface{}{data}
}

func constructNodePools(nodepoolsDefData []interface{}) []*tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition {
	nodepools := []*tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition{}

	for _, npDefData := range nodepoolsDefData {
		data, _ := npDefData.(map[string]interface{})
		np := constructNodepoolDef(data)
		nodepools = append(nodepools, np)
	}

	return nodepools
}

func constructNodepoolDef(nodepoolDefData map[string]interface{}) *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition {
	definition := &tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition{}

	definition.Info = constructNodepoolInfo(nodepoolDefData)

	if v, ok := nodepoolDefData[specKey]; ok {
		data, _ := v.([]interface{})
		definition.Spec = constructNodepoolSpec(data)
	}

	return definition
}

func constructNodepoolInfo(data interface{}) *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolInfo {
	info := &tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolInfo{}

	nodepoolInfoData, _ := data.(map[string]interface{})

	if v, ok := nodepoolInfoData[nameKey]; ok {
		helper.SetPrimitiveValue(v, &info.Name, nameKey)
	}

	if v, ok := nodepoolInfoData[common.DescriptionKey]; ok {
		helper.SetPrimitiveValue(v, &info.Description, common.DescriptionKey)
	}

	return info
}

func constructNodepoolSpec(data []interface{}) *tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec {
	spec := &tkcnodepool.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec{}

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData, _ := data[0].(map[string]interface{})

	if v, ok := specData[classKey]; ok {
		helper.SetPrimitiveValue(v, &spec.Class, classKey)
	}

	if v, ok := specData[replicasKey]; ok {
		helper.SetPrimitiveValue(v, &spec.Replicas, replicasKey)
	}

	if v, ok := specData[failureDomainKey]; ok {
		helper.SetPrimitiveValue(v, &spec.FailureDomain, failureDomainKey)
	}

	if v, ok := specData[overridesKey]; ok {
		data, _ := v.([]interface{})
		spec.Overrides = constructCommonClusterClusterVariables(data)
	}

	if v, ok := specData[metadataKey]; ok {
		data, _ := v.([]interface{})
		spec.Metadata = constructCommonClusterMetadata(data)
	}

	if v, ok := specData[osImageKey]; ok {
		data, _ := v.([]interface{})
		spec.OsImage = constructCommonClusterOsImage(data)
	}

	return spec
}
