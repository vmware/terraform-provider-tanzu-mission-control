package objectmeta

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuCoreV1alpha1ObjectReference Reference references a foreign resource.
//
// swagger:model vmware.tanzu.core.v1alpha1.object.Reference
type VmwareTanzuCoreV1alpha1ObjectReference struct {

	// RID for the object.
	Rid string `json:"rid,omitempty"`

	// UID for the object.
	UID string `json:"uid,omitempty"`
}

// Validate validates this vmware tanzu core v1alpha1 object reference.
func (m *VmwareTanzuCoreV1alpha1ObjectReference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vmware tanzu core v1alpha1 object reference based on context it is used.
func (m *VmwareTanzuCoreV1alpha1ObjectReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1ObjectReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1ObjectReference) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1ObjectReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
