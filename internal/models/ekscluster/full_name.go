// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"fmt"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterFullName Full name of the cluster. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.FullName
type VmwareTanzuManageV1alpha1EksclusterFullName struct {

	// Name of the credential.
	CredentialName string `json:"credentialName,omitempty"`

	// Name of this cluster.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of the region.
	Region string `json:"region,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

func (m *VmwareTanzuManageV1alpha1EksclusterFullName) ToString() string {
	if m == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s:%s:%s", m.OrgID, m.CredentialName, m.Region, m.Name)
}
