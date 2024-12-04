// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package gitrepositoryclustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest Request to create a GitRepository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.gitrepository.CreateGitRepositoryRequest
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest struct {

	// GitRepository to create.
	GitRepository *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository `json:"gitRepository,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse Response from creating a GitRepository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.gitrepository.CreateGitRepositoryResponse
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse struct {

	// GitRepository created.
	GitRepository *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository `json:"gitRepository,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepositoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
