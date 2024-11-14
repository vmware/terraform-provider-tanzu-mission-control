// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package namespacemodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceRequest Request to create a Namespace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.CreateNamespaceRequest
type VmwareTanzuManageV1alpha1ClusterNamespaceRequest struct {

	// Namespace to create.
	Namespace *VmwareTanzuManageV1alpha1ClusterNamespaceNamespace `json:"namespace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceResponse Response from creating a Namespace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.CreateNamespaceResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceResponse struct {

	// Namespace created.
	Namespace *VmwareTanzuManageV1alpha1ClusterNamespaceNamespace `json:"namespace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
