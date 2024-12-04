// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmreleaseclustermodel

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus Status of the Helm Release.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.release.Status
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus struct {

	// Conditions of the Helm Release resource.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Kuberenetes RBAC resources and service account created on the cluster by TMC for Helm Release.
	GeneratedResources *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGeneratedResources `json:"generatedResources,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGeneratedResources Generated Resources for Helm Release on the cluster by TMC.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.release.GeneratedResources
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGeneratedResources struct {

	// Name of the cluster role used for Helm Release.
	ClusterRoleName string `json:"clusterRoleName,omitempty"`

	// Name of the role binding used for Helm Release.
	RoleBindingName string `json:"roleBindingName,omitempty"`

	// Name of the service account used for Helm Release.
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGeneratedResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGeneratedResources) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGeneratedResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
