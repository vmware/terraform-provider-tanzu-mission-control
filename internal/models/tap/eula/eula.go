/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package eula

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1TanzupackageTapEulaEula Represents Tanzu Application Platform (TAP) EULA Acceptance.
//
// swagger:model vmware.tanzu.manage.v1alpha1.tanzupackage.tap.eula.Eula
type VmwareTanzuManageV1alpha1TanzupackageTapEulaEula struct {

	// data
	Data *VmwareTanzuManageV1alpha1TanzupackageTapEulaData `json:"data,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// tap version
	TapVersion string `json:"tapVersion,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1TanzupackageTapEulaEula) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1TanzupackageTapEulaEula) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1TanzupackageTapEulaEula
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
