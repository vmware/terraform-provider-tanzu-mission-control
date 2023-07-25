/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"fmt"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterFullName Full name of the cluster. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.FullName
type VmwareTanzuManageV1alpha1AksclusterFullName struct {

	// Name of the credential.
	CredentialName string `json:"credentialName,omitempty"`

	// Name of this cluster.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of the resource group.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ID of Azure subscription of the cluster.
	SubscriptionID string `json:"subscriptionId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

func (m *VmwareTanzuManageV1alpha1AksclusterFullName) ToString() string {
	if m == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s:%s:%s:%s", m.OrgID, m.CredentialName, m.SubscriptionID, m.ResourceGroupName, m.Name)
}
