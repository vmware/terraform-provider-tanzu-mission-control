// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolFullName Full name of the nodepool. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.FullName
type VmwareTanzuManageV1alpha1AksclusterNodepoolFullName struct {

	// Name of the AKS cluster.
	AksClusterName string `json:"aksClusterName,omitempty"`

	// Name of the credential.
	CredentialName string `json:"credentialName,omitempty"`

	// Name of the nodepool.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Resource group name of the cluster.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ID of Azure subscription of the cluster.
	SubscriptionID string `json:"subscriptionId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
