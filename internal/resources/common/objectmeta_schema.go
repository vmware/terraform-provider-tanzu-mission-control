/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package common

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

const (
	MetaKey            = "meta"
	LabelsKey          = "labels"
	DescriptionKey     = "description"
	AnnotationsKey     = "annotations"
	uidKey             = "uid"
	resourceVersionKey = "resource_version"
	CreatorLabelKey    = "tmc.cloud.vmware.com/creator"
)

var Meta = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Metadata for the resource",
	Computed:    true,
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			AnnotationsKey: {
				Type:        schema.TypeMap,
				Description: "Annotations for the resource",
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if strings.Contains(k, "tmc.cloud.vmware.com") ||
						strings.Contains(k, "x-customer-domain") ||
						strings.Contains(k, "GeneratedTemplateID") {
						return true
					}
					return false
				},
			},
			LabelsKey: {
				Type:        schema.TypeMap,
				Description: "Labels for the resource",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.Contains(k, "tmc.cloud.vmware.com")
				},
			},
			DescriptionKey: {
				Type:        schema.TypeString,
				Description: "Description of the resource",
				Optional:    true,
			},
			uidKey: {
				Type:        schema.TypeString,
				Description: "UID of the resource",
				Computed:    true,
			},
			resourceVersionKey: {
				Type:        schema.TypeString,
				Description: "Resource version of the resource",
				Computed:    true,
			},
		},
	},
}

func HasMetaChanged(d *schema.ResourceData) bool {
	updateRequired := false

	switch {
	case d.HasChange(helper.GetFirstElementOf(MetaKey, LabelsKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(MetaKey, DescriptionKey)):
		updateRequired = true
	}

	return updateRequired
}

func ConstructMeta(d *schema.ResourceData) (objectMeta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta) {
	objectMeta = &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
		Annotations: make(map[string]string),
		Labels:      make(map[string]string),
	}

	value, ok := d.GetOk(MetaKey)
	if !ok {
		return objectMeta
	}

	data := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return objectMeta
	}

	objectMetaData := data[0].(map[string]interface{})

	if v, ok := objectMetaData[AnnotationsKey]; ok {
		objectMeta.Annotations = GetTypeStringMapData(v.(map[string]interface{}))
	}

	if v, ok := objectMetaData[LabelsKey]; ok {
		objectMeta.Labels = GetTypeStringMapData(v.(map[string]interface{}))
	}

	if v, ok := objectMetaData[DescriptionKey]; ok {
		helper.SetPrimitiveValue(v, &objectMeta.Description, DescriptionKey)
	}

	if v, ok := objectMetaData[uidKey]; ok {
		helper.SetPrimitiveValue(v, &objectMeta.UID, uidKey)
	}

	if v, ok := objectMetaData[resourceVersionKey]; ok {
		helper.SetPrimitiveValue(v, &objectMeta.ResourceVersion, resourceVersionKey)
	}

	return objectMeta
}

func FlattenMeta(objectMeta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta) (data []interface{}) {
	if objectMeta == nil {
		return data
	}

	flattenMetaData := make(map[string]interface{})

	flattenMetaData[AnnotationsKey] = objectMeta.Annotations
	flattenMetaData[LabelsKey] = objectMeta.Labels
	flattenMetaData[DescriptionKey] = objectMeta.Description
	flattenMetaData[uidKey] = objectMeta.UID
	flattenMetaData[resourceVersionKey] = objectMeta.ResourceVersion

	return []interface{}{flattenMetaData}
}

func GetTypeStringMapData(data map[string]interface{}) map[string]string {
	convertedMapData := make(map[string]string)

	for key, value := range data {
		value := fmt.Sprintf("%v", value)
		convertedMapData[key] = value
	}

	return convertedMapData
}

func GetTypeIntMapData(data map[string]interface{}) map[string]int {
	convertedMapData := make(map[string]int)

	for key, value := range data {
		value, _ := value.(int)
		convertedMapData[key] = value
	}

	return convertedMapData
}

// GetMetaConverterMap returns mapping for converter.
func GetMetaConverterMap(modelPathSeparator string, modelPath ...string) *converter.BlockToStruct {
	var MetaConverterMap = &converter.BlockToStruct{
		AnnotationsKey: &converter.Map{
			converter.AllMapKeysFieldMarker: converter.BuildModelPath(modelPathSeparator, append(modelPath, "meta", "annotations", converter.AllMapKeysFieldMarker)...),
		},
		LabelsKey: &converter.Map{
			converter.AllMapKeysFieldMarker: converter.BuildModelPath(modelPathSeparator, append(modelPath, "meta", "labels", converter.AllMapKeysFieldMarker)...),
		},
		DescriptionKey:     converter.BuildModelPath(modelPathSeparator, append(modelPath, "meta", "description")...),
		resourceVersionKey: converter.BuildModelPath(modelPathSeparator, append(modelPath, "meta", "resourceVersion")...),
		uidKey:             converter.BuildModelPath(modelPathSeparator, append(modelPath, "meta", "uid")...),
	}

	return MetaConverterMap
}
