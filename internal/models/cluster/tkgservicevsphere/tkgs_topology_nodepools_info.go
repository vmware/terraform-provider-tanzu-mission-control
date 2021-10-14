/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgservicevspheremodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNodepoolInfo Info is the meta information of nodepool for cluster
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.Info
type VmwareTanzuManageV1alpha1ClusterNodepoolInfo struct {

	// Description for the nodepool.
	Description string `json:"description,omitempty"`

	// Name of the nodepool.
	Name string `json:"name,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolInfo) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
