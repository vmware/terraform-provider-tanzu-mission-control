/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustergroupmodels

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterGroupDataprotectionDataProtection Backup, restore, or migrate cluster group data.
//
// Protect Kubernetes cluster group data with the DataProtection resource. Backup, restore, or.
// migrate cluster objects and volumes.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.DataProtection.
type VmwareTanzuManageV1alpha1ClusterGroupDataprotectionDataProtection struct {

	// Full name for the DataProtection.
	FullName *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName `json:"fullName,omitempty"`

	// Metadata for the DataProtection object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec field for DataProtection.
	Spec *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpec `json:"spec"`

	// Status field.
	Status *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionDataProtection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionDataProtection) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupDataprotectionDataProtection

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
