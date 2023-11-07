/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplatemodels

import (
	"github.com/go-openapi/swag"
)

// K8sIoApimachineryPkgApisMetaV1GroupVersionKind GroupVersionKind unambiguously identifies a kind.  It doesn't anonymously include GroupVersion
// to avoid automatic coersion.  It doesn't use a GroupVersion to avoid custom marshalling
//
// +protobuf.options.(gogoproto.goproto_stringer)=false
//
// swagger:model k8s.io.apimachinery.pkg.apis.meta.v1.GroupVersionKind
type K8sIoApimachineryPkgApisMetaV1GroupVersionKind struct {

	// group
	Group string `json:"group,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation.
func (m *K8sIoApimachineryPkgApisMetaV1GroupVersionKind) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *K8sIoApimachineryPkgApisMetaV1GroupVersionKind) UnmarshalBinary(b []byte) error {
	var res K8sIoApimachineryPkgApisMetaV1GroupVersionKind

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
