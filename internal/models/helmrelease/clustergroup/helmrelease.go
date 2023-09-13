/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmreleaseclustergroupmodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease Release is an instance of Helm Chart created at cluster group level.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.helm.release.Release
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease struct {

	// Full name for the Release.
	FullName *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseFullName `json:"fullName,omitempty"`

	// Metadata for the Release object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Release.
	Spec *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseSpec `json:"spec,omitempty"`

	// Status for the Release.
	Status *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
