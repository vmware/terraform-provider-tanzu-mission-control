/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustermodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest Request to create a DataProtection.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.CreateDataProtectionRequest.
type VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest struct {

	// DataProtection to create.
	DataProtection *VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection `json:"dataProtection,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionRequest

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionResponse Response from creating a DataProtection.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.CreateDataProtectionResponse.
type VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionResponse struct {

	// DataProtection created.
	DataProtection *VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection `json:"dataProtection,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionCreateDataProtectionResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse Response from listing DataProtections.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.ListDataProtectionsResponse.
type VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse struct {

	// List of dataprotections.
	DataProtections []*VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection `json:"dataProtections"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionListDataProtectionsResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
