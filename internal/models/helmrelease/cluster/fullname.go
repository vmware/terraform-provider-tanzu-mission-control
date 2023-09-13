/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmreleaseclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName Full name of the Helm Release.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.release.FullName
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName struct {

	// Name of Cluster.
	ClusterName string `json:"clusterName,omitempty"`

	// Name of management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of the Release.
	Name string `json:"name,omitempty"`

	// Name of Namespace.
	NamespaceName string `json:"namespaceName,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of Provisioner.
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
