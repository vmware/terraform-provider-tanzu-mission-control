/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package objectmetamodel

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuCoreV1alpha1ObjectMeta Holds general shared object metadatas.
//
// swagger:model vmware.tanzu.core.v1alpha1.object.Meta
type VmwareTanzuCoreV1alpha1ObjectMeta struct {

	// Annotations for the object. Annotations hold system level information provisioned by controllers.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Creation time of the object.
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creationTime,omitempty"`

	// Description of the resource.
	Description string `json:"description,omitempty"`

	// Generation of the resource as specified by the user, increments on changes.
	Generation string `json:"generation,omitempty"`

	// Labels to apply to the object.
	Labels map[string]string `json:"labels,omitempty"`

	// Hard object references to parents of this resource.
	ParentReferences []*VmwareTanzuCoreV1alpha1ObjectReference `json:"parentReferences"`

	// A string that identifies the internal version of this object that can be used by clients to
	// determine when objects have changed. This value MUST be treated as opaque by clients and
	// passed unmodified back to the server.
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// UID for the object.
	UID string `json:"uid,omitempty"`

	// Update time of the object.
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"updateTime,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1ObjectMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1ObjectMeta) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1ObjectMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
