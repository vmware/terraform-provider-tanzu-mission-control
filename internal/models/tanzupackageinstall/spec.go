// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackageinstall

import (
	"encoding/json"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec Spec of Package Install.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.install.Spec
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec struct {

	// Inline values to configure the Package Install.
	InlineValues interface{} `json:"inlineValues,omitempty"`

	// Reference to the Package which will be installed.
	PackageRef *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef `json:"packageRef,omitempty"`

	// Role binding scope for service account which will be used by Package Install.
	RoleBindingScope *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope `json:"roleBindingScope,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef Reference to Package.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.metadata.package.PackageRef
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef struct {

	// Name of the Package Metadata.
	PackageMetadataName string `json:"packageMetadataName,omitempty"`

	// Version Selection of the Package.
	VersionSelection *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection `json:"versionSelection,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection Version Selection criteria to deploy Package.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.metadata.package.VersionSelection
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection struct {

	// Constraints to select Package. Example: constraints: "v1.2.3", constraints: "<v1.4.0" etc.
	Constraints string `json:"constraints,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope Enumeration defining possible scope of role binding.
//
//   - UNSPECIFIED: Default Role Binding scope. Behaviour is undefined and clients shouldn't use it.
//   - CLUSTER: Role Binding is cluster scoped on the cluster.
//   - NAMESPACE: Role Binding is namespace scoped on the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.install.RoleBindingScope
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope string

func NewVmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope(value VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope) *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope.
func (m VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope) Pointer() *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeUNSPECIFIED captures enum value "UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeUNSPECIFIED VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope = "UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER captures enum value "CLUSTER".
	VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope = "CLUSTER"

	// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeNAMESPACE captures enum value "NAMESPACE".
	VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeNAMESPACE VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope = "NAMESPACE"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScope
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","CLUSTER","NAMESPACE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeEnum = append(vmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeEnum, v)
	}
}
