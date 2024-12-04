// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package gitrepositoryclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest Request to create a GitRepository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.gitrepository.CreateGitRepositoryRequest
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest struct {

	// GitRepository to create.
	GitRepository *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository `json:"gitRepository,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse Response from creating a GitRepository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.gitrepository.CreateGitRepositoryResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse struct {

	// GitRepository created.
	GitRepository *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository `json:"gitRepository,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepositoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
