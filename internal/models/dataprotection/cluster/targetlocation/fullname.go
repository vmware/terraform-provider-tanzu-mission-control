/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName Full name of the resource. This includes the object name along.
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.FullName.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName struct {

	// Name of the Backup Location.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of the data protection provider name.
	ProviderName string `json:"providerName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationFullName

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
