/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package repositorycredentialclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterFluxcdGetRepositorycredentialResponse Response from getting a Credential.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.repositorycredential.GetRepositorycredentialResponse
type VmwareTanzuManageV1alpha1ClusterFluxcdGetRepositorycredentialResponse struct {

	// credential returned.
	Repositorycredential *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredential `json:"respositorycredential,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdGetRepositorycredentialResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdGetRepositorycredentialResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdGetRepositorycredentialResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
