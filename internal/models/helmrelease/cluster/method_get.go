/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmreleaseclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGetResponse Response from getting a Release.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.release.GetReleaseResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGetResponse struct {

	// Release returned.
	Release *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRelease `json:"release,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGetResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
