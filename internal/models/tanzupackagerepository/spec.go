/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackagerepository

import "github.com/go-openapi/swag"

type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec struct {

	// Pulls imgpkg bundle from Docker/OCI registry.
	ImgpkgBundle *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec `json:"imgpkgBundle,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec Package Repository bundle is an image package bundle that holds Package CRs.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.repository.ImgPkgBundleSpec.
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec struct {

	// Docker image url; unqualified, tagged, or digest references supported.
	Image string `json:"image,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
