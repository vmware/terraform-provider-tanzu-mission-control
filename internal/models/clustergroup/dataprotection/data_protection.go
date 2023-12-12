// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClustergroupDataprotectionDataProtection Backup, restore, or migrate cluster data.
//
// Protect Kubernetes cluster data with the DataProtection resource. Backup, restore, or
// migrate cluster objects and volumes.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.DataProtection
type VmwareTanzuManageV1alpha1ClustergroupDataprotectionDataProtection struct {

	// Full name for the DataProtection.
	FullName *VmwareTanzuManageV1alpha1ClustergroupDataprotectionFullName `json:"fullName,omitempty"`

	// Metadata for the DataProtection object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec field for DataProtection.
	Spec *VmwareTanzuManageV1alpha1ClustergroupDataprotectionSpec `json:"spec,omitempty"`

	// Status field.
	Status *VmwareTanzuManageV1alpha1ClustergroupDataprotectionStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionDataProtection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionDataProtection) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupDataprotectionDataProtection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
