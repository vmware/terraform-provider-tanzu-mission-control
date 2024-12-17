// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackageinstall

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGetInstallResponse Response from getting an Install.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.install.GetInstallResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGetInstallResponse struct {

	// Install returned.
	Install *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall `json:"install,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGetInstallResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGetInstallResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGetInstallResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
