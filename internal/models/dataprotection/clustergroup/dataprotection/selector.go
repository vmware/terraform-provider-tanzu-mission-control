/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustergroupmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelector The Selector to include/exclude specific targets.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.Spec.Selector
type VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelector struct {
	// List of target names to include.
	Names []string `json:"names,omitempty"`

	// List of target names to exclude.
	ExcludedNames []string `json:"excluded_names,omitempty"`

	// Label based Selector.
	LabelSelector *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelectorLabelSelector `json:"label_selector,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelector) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelector

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
