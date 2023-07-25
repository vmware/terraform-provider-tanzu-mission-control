/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package gitrepositoryclustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName Full name of the Repository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.gitrepository.FullName
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName struct {

	// Name of Cluster Group.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Name of the Repository.
	Name string `json:"name,omitempty"`

	// Name of Namespace.
	NamespaceName string `json:"namespaceName,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
