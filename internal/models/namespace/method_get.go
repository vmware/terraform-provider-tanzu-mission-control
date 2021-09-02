/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespacemodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse Response from getting a Namespace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.GetNamespaceResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse struct {

	// Namespace returned.
	Namespace *VmwareTanzuManageV1alpha1ClusterNamespaceNamespace `json:"namespace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceGetNamespaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
