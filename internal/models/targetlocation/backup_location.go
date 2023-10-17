/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationmodels

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationBackupLocation A target location for backups.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.BackupLocation.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationBackupLocation struct {

	// Full name for the BackupLocation.
	FullName *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName `json:"fullName,omitempty"`

	// Metadata for the backup location object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the backup location.
	Spec *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationSpec `json:"spec,omitempty"`

	// Status of the backup location.
	Status *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationBackupLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationBackupLocation) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationBackupLocation

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
