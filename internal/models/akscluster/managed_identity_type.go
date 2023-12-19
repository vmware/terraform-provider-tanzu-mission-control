/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterManagedIdentityType Managed identity type options of identity config.
//
//   - IDENTITY_TYPE_SYSTEM_ASSIGNED: Indicates that a system assigned managed identity should be used by the cluster.
//   - IDENTITY_TYPE_USER_ASSIGNED: Indicates that a user assigned managed identity should be used by the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.ManagedIdentityType
type VmwareTanzuManageV1alpha1AksclusterManagedIdentityType string

func NewVmwareTanzuManageV1alpha1AksclusterManagedIdentityType(value VmwareTanzuManageV1alpha1AksclusterManagedIdentityType) *VmwareTanzuManageV1alpha1AksclusterManagedIdentityType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterManagedIdentityType.
func (m VmwareTanzuManageV1alpha1AksclusterManagedIdentityType) Pointer() *VmwareTanzuManageV1alpha1AksclusterManagedIdentityType {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterManagedIdentityTypeSYSTEMASSIGNED captures enum value "IDENTITY_TYPE_SYSTEM_ASSIGNED".
	VmwareTanzuManageV1alpha1AksclusterManagedIdentityTypeSYSTEMASSIGNED VmwareTanzuManageV1alpha1AksclusterManagedIdentityType = "IDENTITY_TYPE_SYSTEM_ASSIGNED"

	// VmwareTanzuManageV1alpha1AksclusterManagedIdentityTypeUSERASSIGNED captures enum value "IDENTITY_TYPE_USER_ASSIGNED".
	VmwareTanzuManageV1alpha1AksclusterManagedIdentityTypeUSERASSIGNED VmwareTanzuManageV1alpha1AksclusterManagedIdentityType = "IDENTITY_TYPE_USER_ASSIGNED"
)

// for schema.
var vmwareTanzuManageV1alpha1AksclusterManagedIdentityTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterManagedIdentityType
	if err := json.Unmarshal([]byte(`["IDENTITY_TYPE_SYSTEM_ASSIGNED","IDENTITY_TYPE_USER_ASSIGNED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterManagedIdentityTypeEnum = append(vmwareTanzuManageV1alpha1AksclusterManagedIdentityTypeEnum, v)
	}
}
