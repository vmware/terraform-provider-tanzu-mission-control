/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgvspheremodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1CommonClusterTKGVsphereWorkspace vmware tanzu manage v1alpha1 common cluster t k g vsphere workspace
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.TKGVsphereWorkspace
type VmwareTanzuManageV1alpha1CommonClusterTKGVsphereWorkspace struct {

	// Datacenter in which a VM is created.
	Datacenter string `json:"datacenter,omitempty"`

	// Datastore in which a VM is created.
	Datastore string `json:"datastore,omitempty"`

	// Folder in which a VM is created.
	Folder string `json:"folder,omitempty"`

	// Network used by the VM.
	Network string `json:"network,omitempty"`

	// Resource pool in which a VM is created.
	ResourcePool string `json:"resourcePool,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterTKGVsphereWorkspace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterTKGVsphereWorkspace) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonClusterTKGVsphereWorkspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
