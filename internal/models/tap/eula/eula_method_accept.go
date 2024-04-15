/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package eula

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaRequest Request to accept an Eula.
//
// swagger:model vmware.tanzu.manage.v1alpha1.tanzupackage.tap.eula.AcceptEulaRequest
type VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaRequest struct {

	// Eula to accept.
	Eula *VmwareTanzuManageV1alpha1TanzupackageTapEulaEula `json:"eula,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaResponse Response from accepting an Eula.
//
// swagger:model vmware.tanzu.manage.v1alpha1.tanzupackage.tap.eula.AcceptEulaResponse
type VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaResponse struct {

	// Eula accepted.
	Eula *VmwareTanzuManageV1alpha1TanzupackageTapEulaEula `json:"eula,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1TanzupackageTapEulaAcceptEulaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
