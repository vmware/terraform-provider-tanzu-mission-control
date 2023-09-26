/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmrepositoryclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.repository.SearchScope
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySearchScope struct {

	// Scope search to the specified cluster_name; supports globbing; default (*).
	ClusterName string `json:"clusterName,omitempty"`

	// Scope search to the specified management_cluster_name; supports globbing; default (*).
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Scope search to the specified namespace_name; supports globbing; default (*).
	NamespaceName string `json:"namespaceName,omitempty"`

	// Scope search to the specified provisioner_name; supports globbing; default (*).
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySearchScope) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
