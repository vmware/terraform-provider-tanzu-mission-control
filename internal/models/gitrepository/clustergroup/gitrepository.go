// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package gitrepositoryclustergroupmodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository Represents a gitrepository source to sync configurations from.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.gitrepository.GitRepository
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository struct {

	// Full name for the Repository.
	FullName *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName `json:"fullName,omitempty"`

	// Metadata for the Repository object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Repository.
	Spec *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositorySpec `json:"spec,omitempty"`

	// Status for the Repository.
	Status *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryGitRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
