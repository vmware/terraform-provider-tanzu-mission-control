/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkc

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tkccommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/common"
)

var commonClusterMetadataSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		labelsKey: {
			Type:        schema.TypeMap,
			Description: "The labels configuration",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		annotationsKey: {
			Type:        schema.TypeMap,
			Description: "The annotations configuration",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	},
}

var commonClusterOsImageSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		nameKey: {
			Type:        schema.TypeString,
			Description: "The name of the OS image",
			Required:    true,
		},
		versionKey: {
			Type:        schema.TypeString,
			Description: "The version of the OS image",
			Required:    true,
		},
		archKey: {
			Type:        schema.TypeString,
			Description: "The arch of the OS image",
			Required:    true,
		},
	},
}

var commonClusterClusterVariableSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		variablesKey: {
			Type:        schema.TypeList,
			Description: "The TKG cluster variable configuration",
			Required:    true,
			MinItems:    1,
			MaxItems:    1,
			// ExactlyOneOf: ,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					tkgVsphereV100Key: {
						Type:     schema.TypeString, //schema.TypeList,
						Optional: true,
						// Elem:     tkgVsphereV100Schema,
						ValidateDiagFunc: validateConfiguration,
					},
					// vsphereTanzuV100Key: {
					// 	Type: schema.TypeList,
					// 	Optional:    true,
					// 	Elem: vsphereTanzuV100Schema,
					// },
				},
			},
		},
	},
}

func constructCommonClusterMetadata(data []interface{}) *tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata {
	meta := &tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata{}
	if len(data) == 0 || data[0] == nil {
		return meta
	}

	metadataData, _ := data[0].(map[string]interface{})

	if v, ok := metadataData[labelsKey]; ok {
		data, _ := v.(map[string]interface{})
		meta.Labels = constructStringMap(data)
	}

	if v, ok := metadataData[annotationsKey]; ok {
		data, _ := v.(map[string]interface{})
		meta.Annotations = constructStringMap(data)
	}

	return meta
}

func constructCommonClusterOsImage(data []interface{}) *tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterOSImage {
	osi := &tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterOSImage{}
	if len(data) == 0 || data[0] == nil {
		return osi
	}

	osImageData, _ := data[0].(map[string]interface{})

	if v, ok := osImageData[nameKey]; ok {
		helper.SetPrimitiveValue(v, &osi.Name, nameKey)
	}

	if v, ok := osImageData[versionKey]; ok {
		helper.SetPrimitiveValue(v, &osi.Version, versionKey)
	}

	if v, ok := osImageData[archKey]; ok {
		helper.SetPrimitiveValue(v, &osi.Arch, archKey)
	}

	return osi
}

//

func constructCommonClusterClusterVariables(data []interface{}) []*tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable {
	clusterVariables := []*tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable{}
	if len(data) == 0 || data[0] == nil {
		return clusterVariables
	}

	varsData, _ := data[0].(map[string]interface{})

	clusterVariablesData := varsData[tkgVsphereV100Key].(string)

	json.Unmarshal([]byte(clusterVariablesData), &clusterVariables)

	return clusterVariables
}

// func flattenCommonClusterVariables(spec *integration.VmwareTanzuManageV1alpha1ClusterIntegrationSpec) interface{} {
// 	flattened := map[string]interface{}{}

// 	if spec != nil && spec.Configurations != nil {
// 		flattened[configurationKey] = toJSON(spec.Configurations)
// 	}

// 	return []map[string]interface{}{flattened}
// }

func validateConfiguration(i interface{}, p cty.Path) diag.Diagnostics {
	var (
		ok bool
		v  string
		m  map[string]interface{}
	)

	if v, ok = i.(string); !ok {
		return diag.Errorf("unexpected type for configuration: %T (%v)", i, p)
	}

	if err := json.Unmarshal([]byte(v), &m); err != nil {
		return diag.FromErr(fmt.Errorf("%w: %s", err, v))
	}

	return nil
}

//

// func constructCommonClusterClusterVariables(variablesData []interface{}) []*tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable {
// 	clusterVariables := []*tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable{}

// 	varsData, _ := variablesData[0].(map[string]interface{})

// 	for varKey, varValue := range varsData {
// 		cv := constructCommonClusterClusterVariable(varKey, varValue)
// 		clusterVariables = append(clusterVariables, cv)
// 	}

// 	return clusterVariables
// }

// func constructCommonClusterClusterVariable(varKey string, varValue interface{}) *tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable {
// 	clusterVariable := &tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable{}

// 	varData := make(map[string]interface{})
// 	varData[varKey] = varValue

// 	if varKey == "" || varValue == nil {
// 		return clusterVariable
// 	}

// 	helper.SetPrimitiveValue(varKey, &clusterVariable.Name, nameKey)

// 	clusterVariable.Value = varValue // TODO?

// 	return clusterVariable
// }

func flattenCommonClusterMetadata(item *tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[annotationsKey] = item.Annotations
	data[labelsKey] = item.Labels

	return []interface{}{data}
}

func flattenCommonClusterOsImage(item *tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterOSImage) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	if item.Arch != "" {
		data[archKey] = item.Arch
	}

	if item.Name != "" {
		data[nameKey] = item.Name
	}

	if item.Version != "" {
		data[versionKey] = item.Version
	}

	return []interface{}{data}
}

func flattenCommonClusterVariables(arr []*tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable) []interface{} {
	data := make([]interface{}, 0, len(arr))

	if len(arr) == 0 {
		return data
	}

	for _, item := range arr {
		data = append(data, flattenCommonClusterVariable(item))
	}

	return data
}

func flattenCommonClusterVariable(item *tkccommon.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable) []interface{} {
	data := make(map[string]interface{})

	if item.Name != "" {
		data[nameKey] = item.Name
	}

	data[valueKey] = item.Value

	return []interface{}{data}
}
