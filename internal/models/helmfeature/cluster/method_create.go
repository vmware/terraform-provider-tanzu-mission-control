// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmfeatureclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterFluxcdHelmRequest Request to create a Helm.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.helm.CreateHelmRequest
type VmwareTanzuManageV1alpha1ClusterFluxcdHelmRequest struct {

	// Helm to create.
	Helm *VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm `json:"helm,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdHelmRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterFluxcdHelmResponse Response from creating a Helm.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.helm.CreateHelmResponse
type VmwareTanzuManageV1alpha1ClusterFluxcdHelmResponse struct {

	// Helm created.
	Helm *VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm `json:"helm,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdHelmResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
