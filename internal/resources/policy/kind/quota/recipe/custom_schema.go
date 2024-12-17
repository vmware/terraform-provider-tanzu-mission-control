// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

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
			LimitsCPUKey: {
				Type:        schema.TypeString,
				Description: "The sum of CPU limits across all pods in a non-terminal state cannot exceed this value",
				Optional:    true,
			},
			LimitsMemoryKey: {
				Type:        schema.TypeString,
				Description: "The sum of memory limits across all pods in a non-terminal state cannot exceed this value",
				Optional:    true,
			},
			PersistentVolumeClaimsKey: {
				Type:        schema.TypeInt,
				Description: "The total number of PersistentVolumeClaims that can exist in a namespace",
				Optional:    true,
			},
			PersistentVolumeClaimsPerClassKey: {
				Type:        schema.TypeMap,
				Description: "Across all persistent volume claims associated with each storage class, the total number of persistent volume claims that can exist in the namespace",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			RequestsCPUKey: {
				Type:        schema.TypeString,
				Description: "The sum of CPU requests across all pods in a non-terminal state cannot exceed this value",
				Optional:    true,
			},
			RequestsMemoryKey: {
				Type:        schema.TypeString,
				Description: "The sum of memory requests across all pods in a non-terminal state cannot exceed this value",
				Optional:    true,
			},
			RequestsStorageKey: {
				Type:        schema.TypeString,
				Description: "The sum of storage requests across all persistent volume claims cannot exceed this value",
				Optional:    true,
			},
			RequestsStoragePerClassKey: {
				Type:        schema.TypeMap,
				Description: "Across all persistent volume claims associated with each storage class, the sum of storage requests cannot exceed this value",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			ResourceCountsKey: {
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

	if v, ok := customData[LimitsCPUKey]; ok {
		helper.SetPrimitiveValue(v, &custom.LimitsCPU, LimitsCPUKey)
	}

	if v, ok := customData[LimitsMemoryKey]; ok {
		helper.SetPrimitiveValue(v, &custom.LimitsMemory, LimitsMemoryKey)
	}

	if v, ok := customData[PersistentVolumeClaimsKey]; ok {
		helper.SetPrimitiveValue(v, &custom.Persistentvolumeclaims, PersistentVolumeClaimsKey)
	}

	if v, ok := customData[PersistentVolumeClaimsPerClassKey]; ok {
		custom.PersistentvolumeclaimsPerClass = common.GetTypeIntMapData(v.(map[string]interface{}))
	}

	if v, ok := customData[RequestsCPUKey]; ok {
		helper.SetPrimitiveValue(v, &custom.RequestsCPU, RequestsCPUKey)
	}

	if v, ok := customData[RequestsMemoryKey]; ok {
		helper.SetPrimitiveValue(v, &custom.RequestsMemory, RequestsMemoryKey)
	}

	if v, ok := customData[RequestsStorageKey]; ok {
		helper.SetPrimitiveValue(v, &custom.RequestsStorage, RequestsStorageKey)
	}

	if v, ok := customData[RequestsStoragePerClassKey]; ok {
		custom.RequestsStoragePerClass = common.GetTypeStringMapData(v.(map[string]interface{}))
	}

	if v, ok := customData[ResourceCountsKey]; ok {
		custom.ResourceCounts = common.GetTypeIntMapData(v.(map[string]interface{}))
	}

	return custom
}

func FlattenCustom(custom *policyrecipequotamodel.VmwareTanzuManageV1alpha1CommonPolicySpecQuotaV1Custom) (data []interface{}) {
	if custom == nil {
		return data
	}

	flattenCustom := make(map[string]interface{})

	flattenCustom[LimitsCPUKey] = custom.LimitsCPU
	flattenCustom[LimitsMemoryKey] = custom.LimitsMemory
	flattenCustom[PersistentVolumeClaimsKey] = custom.Persistentvolumeclaims
	flattenCustom[PersistentVolumeClaimsPerClassKey] = custom.PersistentvolumeclaimsPerClass
	flattenCustom[RequestsCPUKey] = custom.RequestsCPU
	flattenCustom[RequestsMemoryKey] = custom.RequestsMemory
	flattenCustom[RequestsStorageKey] = custom.RequestsStorage
	flattenCustom[RequestsStoragePerClassKey] = custom.RequestsStoragePerClass
	flattenCustom[ResourceCountsKey] = custom.ResourceCounts

	return []interface{}{flattenCustom}
}
