// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackageinstall

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest Request to create an Install.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.install.CreateInstallRequest
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest struct {

	// Install to create.
	Install *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall `json:"install,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse Response from creating an Install.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.install.CreateInstallResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse struct {

	// Install created.
	Install *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall `json:"install,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstallResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
