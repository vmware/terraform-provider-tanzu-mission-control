/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmreleaseclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest Request to create a Release.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.release.CreateReleaseRequest
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest struct {

	// Release to create.
	Release *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease `json:"release,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse Response from creating a Release.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.release.CreateReleaseResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse struct {

	// Release created.
	Release *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease `json:"release,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
