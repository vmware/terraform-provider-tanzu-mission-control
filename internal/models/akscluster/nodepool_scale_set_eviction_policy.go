package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy ScaleSetEvictionPolicy of the auto scaling config.
//
//   - SCALE_SET_EVICTION_POLICY_UNSPECIFIED: Unspecified scale set eviction policy.
//   - DELETE: Nodes in the underlying Scale Set of the nodepool are deleted when they're evicted.
//   - DEALLOCATE: Nodes in the underlying Scale Set of the nodepool are set to the stopped-deallocated state upon eviction.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.ScaleSetEvictionPolicy
type VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy string

func NewVmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy(value VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy) *VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy.
func (m VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy) Pointer() *VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicySCALESETEVICTIONPOLICYUNSPECIFIED captures enum value "SCALE_SET_EVICTION_POLICY_UNSPECIFIED"
	VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicySCALESETEVICTIONPOLICYUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy = "SCALE_SET_EVICTION_POLICY_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicyDELETE captures enum value "DELETE"
	VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicyDELETE VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy = "DELETE"

	// VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicyDEALLOCATE captures enum value "DEALLOCATE"
	VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicyDEALLOCATE VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy = "DEALLOCATE"
)

// for schema
var vmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicyEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy
	if err := json.Unmarshal([]byte(`["SCALE_SET_EVICTION_POLICY_UNSPECIFIED","DELETE","DEALLOCATE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicyEnum = append(vmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicyEnum, v)
	}
}
