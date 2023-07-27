/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterStatus Status of the AKS cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.Status
type VmwareTanzuManageV1alpha1AksclusterStatus struct {

	// Conditions of the cluster resource.
	Conditions map[string]VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Phase of the cluster resource.
	Phase *VmwareTanzuManageV1alpha1AksclusterPhase `json:"phase,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
