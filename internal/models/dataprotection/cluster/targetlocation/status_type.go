/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType Type of the backup location.
//
//   - TYPE_UNSPECIFIED: Type Unspecified is the default type for a backup location.
//   - MANAGED: Type MANAGED indicates backup location(bucket) is managed by TMC.
//   - UNMANAGED: UNMANAGED indicates backup location(bucket) is not managed by TMC.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.Status.Type.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType string

func NewVmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType(value VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType) *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType.
func (m VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType) Pointer() *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusTypeTYPEUNSPECIFIED captures enum value "TYPE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusTypeTYPEUNSPECIFIED VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType = "TYPE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusTypeMANAGED captures enum value "MANAGED".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusTypeMANAGED VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType = "MANAGED"

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusTypeUNMANAGED captures enum value "UNMANAGED".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusTypeUNMANAGED VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType = "UNMANAGED"
)

// for schema.
var vmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType

	if err := json.Unmarshal([]byte(`["TYPE_UNSPECIFIED","MANAGED","UNMANAGED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusTypeEnum = append(vmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusTypeEnum, v)
	}
}
