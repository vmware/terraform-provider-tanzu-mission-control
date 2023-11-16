/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustergroupmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName Full name of the namespace. This includes the object name along.
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.FullName.
type VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName struct {
	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of Cluster group.
	ClusterGroupName string `json:"cluster_group_name,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
