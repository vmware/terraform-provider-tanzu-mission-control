/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package manifestmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterClusterManifestGetResponse The response type for getting the attach manifest for a cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.ClusterManifestGetResponse
type VmwareTanzuManageV1alpha1ClusterClusterManifestGetResponse struct {

	// Attach manifest for the cluster resource.
	Manifest string `json:"manifest,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterClusterManifestGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterClusterManifestGetResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterClusterManifestGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
