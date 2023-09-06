/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackage

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterTanzupackageListTanzuPackagesResponse Response from listing TanzuPackages.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.tanzupackage.ListTanzuPackagesResponse
type VmwareTanzuManageV1alpha1ClusterTanzupackageListTanzuPackagesResponse struct {

	// List of tanzupackages.
	TanzuPackages []*VmwareTanzuManageV1alpha1ClusterTanzupackageTanzuPackage `json:"tanzuPackages"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterTanzupackageListTanzuPackagesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterTanzupackageListTanzuPackagesResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterTanzupackageListTanzuPackagesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
