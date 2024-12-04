// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzupackage

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataGetPackageResponse Response from getting a Package.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.metadata.package.GetPackageResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataGetPackageResponse struct {

	// Package returned.
	Package *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackage `json:"package,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataGetPackageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataGetPackageResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataGetPackageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
