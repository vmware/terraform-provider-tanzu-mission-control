// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmreleaseclustermodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease Instance of Helm Chart.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.release.Release
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease struct {

	// Full name for the Release.
	FullName *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName `json:"fullName,omitempty"`

	// Metadata for the Release object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Release.
	Spec *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec `json:"spec,omitempty"`

	// Status for the Release.
	Status *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
