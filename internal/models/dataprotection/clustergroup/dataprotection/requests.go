/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustergroupmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionRequest Request to create a DataProtection.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.CreateDataProtectionRequest.
type VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionRequest struct {

	// DataProtection to create.
	DataProtection *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionDataProtection `json:"dataProtection,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionRequest

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionResponse Response from creating a DataProtection.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.CreateDataProtectionResponse.
type VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionResponse struct {

	// DataProtection created.
	DataProtection *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionDataProtection `json:"dataProtection,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupDataprotectionCreateDataProtectionResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterGroupDataprotectionListDataProtectionsResponse Response from listing DataProtections.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.ListDataProtectionsResponse.
type VmwareTanzuManageV1alpha1ClusterGroupDataprotectionListDataProtectionsResponse struct {

	// List of dataprotections.
	DataProtections []*VmwareTanzuManageV1alpha1ClusterGroupDataprotectionDataProtection `json:"dataProtections"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionListDataProtectionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionListDataProtectionsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupDataprotectionListDataProtectionsResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
