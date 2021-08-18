/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package common

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	objectmetamodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/objectmeta"
)

const (
	MetaKey        = "meta"
	annotationsKey = "annotations"
	labelsKey      = "labels"
	descriptionKey = "description"
	uidKey         = "uid"
)

var Meta = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			annotationsKey: {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			labelsKey: {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			descriptionKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			uidKey: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	},
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
		objectMeta.Annotations = getTypeMapData(v.(map[string]interface{}))
	}

	if v, ok := objectMetaData[labelsKey]; ok {
		objectMeta.Labels = getTypeMapData(v.(map[string]interface{}))
	}

	if v, ok := objectMetaData[descriptionKey]; ok {
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
	flattenMetaData[labelsKey] = objectMeta.Labels
	flattenMetaData[descriptionKey] = objectMeta.Description
	flattenMetaData[uidKey] = objectMeta.UID

	return []interface{}{flattenMetaData}
}

func getTypeMapData(data map[string]interface{}) map[string]string {
	convertedMapData := make(map[string]string)

	for key, value := range data {
		value := fmt.Sprintf("%v", value)
		convertedMapData[key] = value
	}

	return convertedMapData
}
