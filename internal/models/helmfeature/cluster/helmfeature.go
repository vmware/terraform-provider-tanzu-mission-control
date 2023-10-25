/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmfeatureclustermodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm Represents helm feature for a cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.helm.Helm
type VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm struct {

	// Full name for the helm.
	FullName *VmwareTanzuManageV1alpha1ClusterFluxcdHelmFullName `json:"fullName,omitempty"`

	// Metadata for the helm object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Status for the helm.
	Status *VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
