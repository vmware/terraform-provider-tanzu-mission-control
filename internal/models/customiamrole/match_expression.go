/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrolemodels

import (
	"github.com/go-openapi/swag"
)

// K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement A label selector requirement is a selector that contains values, a key, and an operator that
// relates the key and values.
//
// swagger:model k8s.io.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
type K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement struct {

	// key is the label key that the selector applies to.
	// +patchMergeKey=key
	// +patchStrategy=merge
	Key string `json:"key,omitempty"`

	// operator represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator string `json:"operator,omitempty"`

	// values is an array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. This array is replaced during a strategic
	// merge patch.
	// +optional
	Values []string `json:"values"`
}

// MarshalBinary interface implementation.
func (m *K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement) UnmarshalBinary(b []byte) error {
	var res K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
