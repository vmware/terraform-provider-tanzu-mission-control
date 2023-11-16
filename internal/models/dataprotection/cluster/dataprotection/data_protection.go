/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustermodels

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection Backup, restore, or migrate cluster data.
//
// Protect Kubernetes cluster data with the DataProtection resource. Backup, restore, or.
// migrate cluster objects and volumes.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.DataProtection.
type VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection struct {

	// Full name for the DataProtection.
	FullName *VmwareTanzuManageV1alpha1ClusterDataprotectionFullName `json:"fullName,omitempty"`

	// Metadata for the DataProtection object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec field for DataProtection.
	Spec *VmwareTanzuManageV1alpha1ClusterDataprotectionSpec `json:"spec"`

	// Status field.
	Status *VmwareTanzuManageV1alpha1ClusterDataprotectionStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
