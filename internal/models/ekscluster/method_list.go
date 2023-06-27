/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterListEksClustersResponse Response from listing EksClusters.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.ListEksClustersResponse
type VmwareTanzuManageV1alpha1EksclusterListEksClustersResponse struct {

	// List of eksclusters.
	EksClusters []*VmwareTanzuManageV1alpha1EksclusterEksCluster `json:"eksClusters"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterListEksClustersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterListEksClustersResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterListEksClustersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
