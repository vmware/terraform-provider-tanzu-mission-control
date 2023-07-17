package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterTier Tier options of cluster SKU.
//
//   - TIER_UNSPECIFIED: Unspecified tier.
//   - FREE: No guaranteed SLA, no additional charges. Free tier clusters have an SLO of 99.5%.
//   - PAID: Guarantees 99.95% availability of the Kubernetes API server endpoint for clusters that use
//
// Availability Zones and 99.9% of availability for clusters that don't use Availability Zones.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.Tier
type VmwareTanzuManageV1alpha1AksclusterTier string

func NewVmwareTanzuManageV1alpha1AksclusterTier(value VmwareTanzuManageV1alpha1AksclusterTier) *VmwareTanzuManageV1alpha1AksclusterTier {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterTier.
func (m VmwareTanzuManageV1alpha1AksclusterTier) Pointer() *VmwareTanzuManageV1alpha1AksclusterTier {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterTierTIERUNSPECIFIED captures enum value "TIER_UNSPECIFIED".
	VmwareTanzuManageV1alpha1AksclusterTierTIERUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterTier = "TIER_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterTierFREE captures enum value "FREE".
	VmwareTanzuManageV1alpha1AksclusterTierFREE VmwareTanzuManageV1alpha1AksclusterTier = "FREE"

	// VmwareTanzuManageV1alpha1AksclusterTierPAID captures enum value "PAID".
	VmwareTanzuManageV1alpha1AksclusterTierPAID VmwareTanzuManageV1alpha1AksclusterTier = "PAID"
)

// for schema.
var vmwareTanzuManageV1alpha1AksclusterTierEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterTier
	if err := json.Unmarshal([]byte(`["TIER_UNSPECIFIED","FREE","PAID"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterTierEnum = append(vmwareTanzuManageV1alpha1AksclusterTierEnum, v)
	}
}
