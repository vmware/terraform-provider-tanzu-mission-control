/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"github.com/go-openapi/swag"

	policymodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
)

// K8sIoApimachineryPkgApisMetaV1LabelSelector A label selector is a label query over a set of resources. The result of matchLabels and.
// matchExpressions are ANDed. An empty label selector matches all objects. A null,
// label selector matches no objects.
//
// swagger:model k8s.io.apimachinery.pkg.apis.meta.v1.LabelSelector.
type K8sIoApimachineryPkgApisMetaV1LabelSelector struct {

	// matchExpressions is a list of label selector requirements. The requirements are ANDed.
	// +optional.
	MatchExpressions []*policymodels.K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement `json:"matchExpressions"`

	// matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels.
	// map is equivalent to an element of matchExpressions, whose key field is "key", the.
	// operator is "In", and the values array contains only "value". The requirements are ANDed.
	// +optional.
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

// MarshalBinary interface implementation.
func (m *K8sIoApimachineryPkgApisMetaV1LabelSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *K8sIoApimachineryPkgApisMetaV1LabelSelector) UnmarshalBinary(b []byte) error {
	var res K8sIoApimachineryPkgApisMetaV1LabelSelector

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
