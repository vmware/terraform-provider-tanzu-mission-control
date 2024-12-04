// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgservicevspheremodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereStorageSettings Storage related settings for workload cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgservicevsphere.StorageSettings
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereStorageSettings struct {

	// Classes is a list of storage classes from the supervisor namespace to expose within a cluster.
	// If omitted, all storage classes from the supervisor namespace will be exposed within the cluster.
	Classes []string `json:"classes"`

	// DefaultClass is the valid storage class name which is treated as the default storage class within a cluster.
	// If omitted, no default storage class is set.
	DefaultClass string `json:"defaultClass,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereStorageSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereStorageSettings) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereStorageSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
