/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package eula

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1TanzupackageTapEulaData vmware tanzu manage v1alpha1 tanzupackage tap eula data
//
// swagger:model vmware.tanzu.manage.v1alpha1.tanzupackage.tap.eula.Data
type VmwareTanzuManageV1alpha1TanzupackageTapEulaData struct {

	// Identifies whether this user has accepted the EULA terms.
	Accepted bool `json:"accepted,omitempty"`

	// URL at which this end user license agreement can be found.
	EulaURL string `json:"eulaUrl,omitempty"`

	// Time when this EULA version was released.
	// Format: date-time
	ReleasedAt strfmt.DateTime `json:"releasedAt,omitempty"`

	// User email identifier.
	User string `json:"user,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1TanzupackageTapEulaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1TanzupackageTapEulaData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1TanzupackageTapEulaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
