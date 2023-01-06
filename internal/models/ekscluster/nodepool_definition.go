/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition Definition is the definition of nodepool for cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.Definition
type VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition struct {

	// Info for the nodepool.
	Info *VmwareTanzuManageV1alpha1EksclusterNodepoolInfo `json:"info,omitempty"`

	// Spec for the nodepool.
	Spec *VmwareTanzuManageV1alpha1EksclusterNodepoolSpec `json:"spec,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
