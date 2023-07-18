package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig Auto scaling config for the nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.AutoScalingConfig
type VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig struct {

	// Whether to enable auto-scaler.
	Enabled bool `json:"enabled,omitempty"`

	// The maximum number of nodes for auto-scaling.
	MaxCount int32 `json:"maxCount,omitempty"`

	// The minimum number of nodes for auto-scaling.
	MinCount int32 `json:"minCount,omitempty"`

	// The Virtual Machine Scale Set eviction policy to use.
	// This cannot be specified unless the scaleSetPriority is 'Spot'.
	ScaleSetEvictionPolicy *VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy `json:"scaleSetEvictionPolicy,omitempty"`

	// The Virtual Machine Scale Set priority.
	ScaleSetPriority *VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority `json:"scaleSetPriority,omitempty"`

	// The max price (in US Dollars) you are willing to pay for spot instances.
	// Possible values are any decimal value greater than zero or -1 which indicates default price to be up-to on-demand.
	SpotMaxPrice float32 `json:"spotMaxPrice,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
