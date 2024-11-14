// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmfeatureclustergroupmodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm Represents helm feature for a cluster group.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.helm.Helm
type VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm struct {

	// Full name for the helm.
	FullName *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmFullName `json:"fullName,omitempty"`

	// Metadata for the helm object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Status for the helm.
	Status *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
