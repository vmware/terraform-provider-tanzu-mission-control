/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package gitrepositoryclustermodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository Represents a gitrepository source to sync configurations from.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.gitrepository.GitRepository
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository struct {

	// Full name for the Repository.
	FullName *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName `json:"fullName,omitempty"`

	// Metadata for the Repository object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Repository.
	Spec *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec `json:"spec,omitempty"`

	// Status for the Repository.
	Status *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
