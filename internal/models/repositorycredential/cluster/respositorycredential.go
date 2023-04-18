/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package repositorycredentialclustermodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// Repositorycredential represents a credential used to authenticate to a fluxcd source such as GitRepository.
type VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecret struct {
	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
	// Full name for the Source Secret.
	FullName *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName `json:"fullName,omitempty"`
	// Metadata for the Source Secret object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`
	// Spec for the Source Secret.
	Spec *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec `json:"spec,omitempty"`
	// Status for the Source Secret.
	Status *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretStatus `json:"status,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecret) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
