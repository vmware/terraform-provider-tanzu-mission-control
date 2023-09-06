/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackage

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterTanzupackageTanzuPackage Represents tanzupackage feature configuration for a cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.tanzupackage.TanzuPackage
type VmwareTanzuManageV1alpha1ClusterTanzupackageTanzuPackage struct {

	// Full name for the TanzuPackage.
	FullName *VmwareTanzuManageV1alpha1ClusterTanzupackageFullName `json:"fullName,omitempty"`

	// Metadata for the TanzuPackage object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Status for the TanzuPackage.
	Status *VmwareTanzuManageV1alpha1ClusterTanzupackageStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterTanzupackageTanzuPackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterTanzupackageTanzuPackage) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterTanzupackageTanzuPackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
