/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackage

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackage Represents Tanzu Carvel Package that's available for installation.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.metadata.package.Package
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackage struct {

	// Full name for the Package.
	FullName *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageFullName `json:"fullName,omitempty"`

	// Metadata for the Package object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Package.
	Spec *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec `json:"spec,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackage) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
