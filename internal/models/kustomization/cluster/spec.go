// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kustomizationclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec Spec of the Kustomization.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.kustomization.Spec
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec struct {

	// Interval defines the interval at which to reconcile kustomization.
	Interval string `json:"interval,omitempty"`

	// Path within the source from which configurations will be applied.
	Path string `json:"path,omitempty"`

	// If true, the workloads will be deleted when the kustomization CR is deleted.
	Prune bool `json:"prune,omitempty"`

	// Reference to the source from which the configurations will be applied.
	Source *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference `json:"source,omitempty"`

	// TargetNamespace sets or overrides the namespaces of resources/kustomization yaml while applying on cluster.
	// Namespace specified here must exist on cluster. It won't be created as a result of specifying here.
	TargetNamespace string `json:"targetNamespace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference Reference to the repository in same or different namespace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.kustomization.RepositoryReference
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference struct {

	// Name of the repository.
	Name string `json:"name,omitempty"`

	// Namespace of the repository.
	Namespace string `json:"namespace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
