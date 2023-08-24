/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackage

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterTanzupackageFullName Full name of the TanzuPackage.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.tanzupackage.FullName
type VmwareTanzuManageV1alpha1ClusterTanzupackageFullName struct {

	// Name of Cluster.
	ClusterName string `json:"clusterName,omitempty"`

	// Name of management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of Provisioner.
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterTanzupackageFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterTanzupackageFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterTanzupackageFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
