// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package gitrepositoryclustermodel

import (
	"encoding/json"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec Spec of the Repository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.gitrepository.Spec
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec struct {

	// GitImplementation specifies which client library implementation to use.
	GitImplementation *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation `json:"gitImplementation,omitempty"`

	// Interval at which to check gitrepository for updates.
	Interval string `json:"interval,omitempty"`

	// Reference specifies git reference to resolve.
	Ref *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference `json:"ref,omitempty"`

	// Reference to the secret.
	SecretRef string `json:"secretRef,omitempty"`

	// URL of the git repository.
	URL string `json:"url,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation GitImplementation specifies which client library implementation to use.
//
//   - GO_GIT: GO_GIT specifies go-git library to use.
//   - LIB_GIT2: LIB_GIT2 specifies libgit2 library to use which supports git v2 protocol.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.gitrepository.GitImplementation
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation string

func NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(value VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation) *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation.
func (m VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation) Pointer() *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT captures enum value "GO_GIT".
	VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation = "GO_GIT"

	// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2 captures enum value "LIB_GIT2".
	VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2 VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation = "LIB_GIT2"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation
	if err := json.Unmarshal([]byte(`["GO_GIT","LIB_GIT2"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationEnum = append(vmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationEnum, v)
	}
}

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference Reference specifies git reference to resolve and checkout.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.gitrepository.Reference
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference struct {

	// Branch from git to checkout.
	Branch string `json:"branch,omitempty"`

	// Commit SHA to checkout. Takes precedence over all other reference fields.
	// When GitRepository.spec.git_implementation is `go-git`, this can be combined
	// with branch to shallow clone branch in which the commit is expected to exist.
	Commit string `json:"commit,omitempty"`

	// SemVer expression to checkout from git tags. Takes precedence over tag.
	Semver string `json:"semver,omitempty"`

	// Tag from git to checkout. Takes precedence over branch.
	Tag string `json:"tag,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
