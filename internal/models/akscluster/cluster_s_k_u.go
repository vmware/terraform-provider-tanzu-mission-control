/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterClusterSKU The SKU of an AKS cluster
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.ClusterSKU
type VmwareTanzuManageV1alpha1AksclusterClusterSKU struct {

	// Name of the cluster SKU.
	Name *VmwareTanzuManageV1alpha1AksclusterClusterSKUName `json:"name,omitempty"`

	// Tier of the cluster SKU.
	Tier *VmwareTanzuManageV1alpha1AksclusterTier `json:"tier,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterClusterSKU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterClusterSKU) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterClusterSKU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
