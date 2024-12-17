// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmreleaseclustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest Request to create a Release.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.helm.release.CreateReleaseRequest
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest struct {

	// Release to create.
	Release *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease `json:"release,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse Response from updating a Release.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.helm.release.UpdateReleaseResponse
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse struct {

	// Release updated.
	Release *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseRelease `json:"release,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
