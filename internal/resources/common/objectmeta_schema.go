/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package common

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/helper"
	objectmetamodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

const (
	MetaKey            = "meta"
	LabelsKey          = "labels"
	DescriptionKey     = "description"
	annotationsKey     = "annotations"
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
			annotationsKey: {
				Type:        schema.TypeMap,
				Description: "Annotations for the resource",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			LabelsKey: {
				Type:        schema.TypeMap,
				Description: "Labels for the resource",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
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

	if v, ok := objectMetaData[annotationsKey]; ok {
		objectMeta.Annotations = GetTypeMapData(v.(map[string]interface{}))
	}

	if v, ok := objectMetaData[LabelsKey]; ok {
		objectMeta.Labels = GetTypeMapData(v.(map[string]interface{}))
	}

	if v, ok := objectMetaData[DescriptionKey]; ok {
		objectMeta.Description = v.(string)
	}

	if v, ok := objectMetaData[uidKey]; ok {
		objectMeta.UID = v.(string)
	}

	return objectMeta
}

func FlattenMeta(objectMeta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta) (data []interface{}) {
	if objectMeta == nil {
		return data
	}

	flattenMetaData := make(map[string]interface{})

	flattenMetaData[annotationsKey] = objectMeta.Annotations
	flattenMetaData[LabelsKey] = objectMeta.Labels
	flattenMetaData[DescriptionKey] = objectMeta.Description
	flattenMetaData[uidKey] = objectMeta.UID

	return []interface{}{flattenMetaData}
}

func GetTypeMapData(data map[string]interface{}) map[string]string {
	convertedMapData := make(map[string]string)

	for key, value := range data {
		value := fmt.Sprintf("%v", value)
		convertedMapData[key] = value
	}

	return convertedMapData
}
