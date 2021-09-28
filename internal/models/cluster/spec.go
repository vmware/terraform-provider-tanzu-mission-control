/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustermodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterSpec Spec of the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.Spec
type VmwareTanzuManageV1alpha1ClusterSpec struct {

	// Name of the cluster group to which this cluster belongs.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Optional proxy name is the name of the Proxy Config
	// to be used for the cluster.
	ProxyName string `json:"proxyName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
