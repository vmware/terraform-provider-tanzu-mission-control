/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package eula

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1TanzupackageTapEulaValidateEulaResponse Response from validating an Eula.
//
// swagger:model vmware.tanzu.manage.v1alpha1.tanzupackage.tap.eula.ValidateEulaResponse
type VmwareTanzuManageV1alpha1TanzupackageTapEulaValidateEulaResponse struct {

	// Eula returned.
	Eula *VmwareTanzuManageV1alpha1TanzupackageTapEulaEula `json:"eula,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1TanzupackageTapEulaValidateEulaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1TanzupackageTapEulaValidateEulaResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1TanzupackageTapEulaValidateEulaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
