// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package gitrepositoryclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGetGitRepositoryResponse Response from getting a GitRepository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.gitrepository.GetGitRepositoryResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGetGitRepositoryResponse struct {

	// GitRepository returned.
	GitRepository *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository `json:"gitRepository,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGetGitRepositoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGetGitRepositoryResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGetGitRepositoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
