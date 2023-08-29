/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackageinstall

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus Status of Package Install.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.install.Status
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus struct {

	// Conditions of the Package Install resource.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Kuberenetes RBAC resources and service account created on the cluster by TMC for Package Install.
	GeneratedResources *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGeneratedResources `json:"generatedResources,omitempty"`

	// If true, the Package Install is managed by TMC.
	Managed bool `json:"managed,omitempty"`

	// TMC services/features referencing the package install.
	ReferredBy []string `json:"referredBy"`

	// Resolved version of the Package Install.
	ResolvedVersion string `json:"resolvedVersion,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGeneratedResources Generated Resources for Package Install on the cluster by TMC.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.install.GeneratedResources
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGeneratedResources struct {

	// Name of the cluster role used for Package Install.
	ClusterRoleName string `json:"clusterRoleName,omitempty"`

	// Name of the role binding used for Package Install.
	RoleBindingName string `json:"roleBindingName,omitempty"`

	// Name of the service account used for Package Install.
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGeneratedResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGeneratedResources) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGeneratedResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
