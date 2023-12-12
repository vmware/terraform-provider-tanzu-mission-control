// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClustergroupDataprotectionCreateDataProtectionRequest Request to create a DataProtection.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.CreateDataProtectionRequest
type VmwareTanzuManageV1alpha1ClustergroupDataprotectionCreateDataProtectionRequest struct {

	// DataProtection to create.
	DataProtection *VmwareTanzuManageV1alpha1ClustergroupDataprotectionDataProtection `json:"dataProtection,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionCreateDataProtectionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionCreateDataProtectionRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupDataprotectionCreateDataProtectionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupDataprotectionCreateDataProtectionResponse Response from creating a DataProtection.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.CreateDataProtectionResponse
type VmwareTanzuManageV1alpha1ClustergroupDataprotectionCreateDataProtectionResponse struct {

	// DataProtection created.
	DataProtection *VmwareTanzuManageV1alpha1ClustergroupDataprotectionDataProtection `json:"dataProtection,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionCreateDataProtectionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionCreateDataProtectionResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupDataprotectionCreateDataProtectionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupDataprotectionDeleteDataProtectionResponse Response from deleting a DataProtection.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.DeleteDataProtectionResponse
type VmwareTanzuManageV1alpha1ClustergroupDataprotectionDeleteDataProtectionResponse struct {

	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionDeleteDataProtectionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionDeleteDataProtectionResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupDataprotectionDeleteDataProtectionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupDataprotectionListDataProtectionsResponse Response from listing DataProtections.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.ListDataProtectionsResponse
type VmwareTanzuManageV1alpha1ClustergroupDataprotectionListDataProtectionsResponse struct {

	// List of dataprotections.
	DataProtections []*VmwareTanzuManageV1alpha1ClustergroupDataprotectionDataProtection `json:"dataProtections"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionListDataProtectionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionListDataProtectionsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupDataprotectionListDataProtectionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}