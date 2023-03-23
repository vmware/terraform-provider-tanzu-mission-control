/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package repositorycredentialclustermodel

import (
	"github.com/go-openapi/swag"

	credmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
)

// VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialStatus Status of the Repository Credential.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.repositorycredential.Status
type VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialStatus struct {

	// Conditions of the Repository Credential resource.
	Status *credmodel.VmwareTanzuManageV1alpha1AccountCredentialStatus `json:"status,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
