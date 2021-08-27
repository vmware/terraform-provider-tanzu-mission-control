/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespacemodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceCreateNamespaceRequest Request to create a Namespace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.CreateNamespaceRequest
type VmwareTanzuManageV1alpha1ClusterNamespaceCreateNamespaceRequest struct {

	// Namespace to create.
	Namespace *VmwareTanzuManageV1alpha1ClusterNamespaceNamespace `json:"namespace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceCreateNamespaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceCreateNamespaceRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceCreateNamespaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceCreateNamespaceResponse Response from creating a Namespace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.CreateNamespaceResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceCreateNamespaceResponse struct {

	// Namespace created.
	Namespace *VmwareTanzuManageV1alpha1ClusterNamespaceNamespace `json:"namespace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceCreateNamespaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceCreateNamespaceResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceCreateNamespaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
