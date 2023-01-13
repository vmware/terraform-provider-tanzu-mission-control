/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipequotamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/quota"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var Custom = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for namespace quota policy custom recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			limitsCPUKey: {
				Type:        schema.TypeString,
				Description: "The sum of CPU limits across all pods in a non-terminal state cannot exceed this value",
				Optional:    true,
			},
			limitsMemoryKey: {
				Type:        schema.TypeString,
				Description: "The sum of memory limits across all pods in a non-terminal state cannot exceed this value",
				Optional:    true,
			},
			persistentVolumeClaimsKey: {
				Type:        schema.TypeInt,
				Description: "The total number of PersistentVolumeClaims that can exist in a namespace",
				Optional:    true,
			},
			persistentVolumeClaimsPerClassKey: {
				Type:        schema.TypeMap,
				Description: "Across all persistent volume claims associated with each storage class, the total number of persistent volume claims that can exist in the namespace",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			requestsCPUKey: {
				Type:        schema.TypeString,
				Description: "The sum of CPU requests across all pods in a non-terminal state cannot exceed this value",
				Optional:    true,
			},
			requestsMemoryKey: {
				Type:        schema.TypeString,
				Description: "The sum of memory requests across all pods in a non-terminal state cannot exceed this value",
				Optional:    true,
			},
			requestsStorageKey: {
				Type:        schema.TypeString,
				Description: "The sum of storage requests across all persistent volume claims cannot exceed this value",
				Optional:    true,
			},
			requestsStoragePerClassKey: {
				Type:        schema.TypeMap,
				Description: "Across all persistent volume claims associated with each storage class, the sum of storage requests cannot exceed this value",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			resourceCountsKey: {
				Type:        schema.TypeMap,
				Description: "The total number of Services of the given type that can exist in a namespace",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	},
}

func ConstructCustom(data []interface{}) (custom *policyrecipequotamodel.VmwareTanzuManageV1alpha1CommonPolicySpecQuotaV1Custom) {
	if len(data) == 0 || data[0] == nil {
		return custom
	}

	customData, _ := data[0].(map[string]interface{})

	custom = &policyrecipequotamodel.VmwareTanzuManageV1alpha1CommonPolicySpecQuotaV1Custom{}

	if v, ok := customData[limitsCPUKey]; ok {
		helper.SetPrimitiveValue(v, &custom.LimitsCPU, limitsCPUKey)
	}

	if v, ok := customData[limitsMemoryKey]; ok {
		helper.SetPrimitiveValue(v, &custom.LimitsMemory, limitsMemoryKey)
	}

	if v, ok := customData[persistentVolumeClaimsKey]; ok {
		helper.SetPrimitiveValue(v, &custom.Persistentvolumeclaims, persistentVolumeClaimsKey)
	}

	if v, ok := customData[persistentVolumeClaimsPerClassKey]; ok {
		custom.PersistentvolumeclaimsPerClass = common.GetTypeIntMapData(v.(map[string]interface{}))
	}

	if v, ok := customData[requestsCPUKey]; ok {
		helper.SetPrimitiveValue(v, &custom.RequestsCPU, requestsCPUKey)
	}

	if v, ok := customData[requestsMemoryKey]; ok {
		helper.SetPrimitiveValue(v, &custom.RequestsMemory, requestsMemoryKey)
	}

	if v, ok := customData[requestsStorageKey]; ok {
		helper.SetPrimitiveValue(v, &custom.RequestsStorage, requestsStorageKey)
	}

	if v, ok := customData[requestsStoragePerClassKey]; ok {
		custom.RequestsStoragePerClass = common.GetTypeStringMapData(v.(map[string]interface{}))
	}

	if v, ok := customData[resourceCountsKey]; ok {
		custom.ResourceCounts = common.GetTypeIntMapData(v.(map[string]interface{}))
	}

	return custom
}

func FlattenCustom(custom *policyrecipequotamodel.VmwareTanzuManageV1alpha1CommonPolicySpecQuotaV1Custom) (data []interface{}) {
	if custom == nil {
		return data
	}

	flattenCustom := make(map[string]interface{})

	flattenCustom[limitsCPUKey] = custom.LimitsCPU
	flattenCustom[limitsMemoryKey] = custom.LimitsMemory
	flattenCustom[persistentVolumeClaimsKey] = custom.Persistentvolumeclaims
	flattenCustom[persistentVolumeClaimsPerClassKey] = custom.PersistentvolumeclaimsPerClass
	flattenCustom[requestsCPUKey] = custom.RequestsCPU
	flattenCustom[requestsMemoryKey] = custom.RequestsMemory
	flattenCustom[requestsStorageKey] = custom.RequestsStorage
	flattenCustom[requestsStoragePerClassKey] = custom.RequestsStoragePerClass
	flattenCustom[resourceCountsKey] = custom.ResourceCounts

	return []interface{}{flattenCustom}
}
