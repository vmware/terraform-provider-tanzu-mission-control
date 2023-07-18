package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolTaint The node this Taint is attached to has the "effect" on
// any pod that does not tolerate the Taint.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.Taint
type VmwareTanzuManageV1alpha1AksclusterNodepoolTaint struct {

	// Current effect state of the nodepool.
	Effect *VmwareTanzuManageV1alpha1AksclusterNodepoolTaintEffect `json:"effect,omitempty"`

	// The taint key to be applied to a node.
	Key string `json:"key,omitempty"`

	// The taint value corresponding to the taint key.
	Value string `json:"value,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolTaint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolTaint) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolTaint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
