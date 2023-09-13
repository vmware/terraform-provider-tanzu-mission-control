/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmfeatureclustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmRequest Request to create a Helm.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.helm.CreateHelmRequest
type VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmRequest struct {

	// Helm to create.
	Helm *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm `json:"helm,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmResponse Response from creating a Helm.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.helm.CreateHelmResponse
type VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmResponse struct {

	// Helm created.
	Helm *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm `json:"helm,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
