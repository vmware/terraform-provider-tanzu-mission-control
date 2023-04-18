/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package repositorycredentialclustermodel

import (
	"github.com/go-openapi/swag"

	credentialmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
)

// VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialStatus Status of the Repository Credential.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.repositorycredential.Status
type VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretStatus struct {

	// Conditions of the Repository Credential resource.
	Status *credentialmodels.VmwareTanzuManageV1alpha1AccountCredentialStatus `json:"status,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
