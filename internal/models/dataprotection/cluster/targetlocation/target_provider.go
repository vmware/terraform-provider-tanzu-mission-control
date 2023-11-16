/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider Available target provider for back up locations.
//
//   - TARGET_PROVIDER_UNSPECIFIED: Unspecified target provider.
//   - AWS: AWS.
//   - AZURE: AZURE.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.TargetProvider.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider string

func NewVmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider(value VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider) *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider.
func (m VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider) Pointer() *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderTARGETPROVIDERUNSPECIFIED captures enum value "TARGET_PROVIDER_UNSPECIFIED".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderTARGETPROVIDERUNSPECIFIED VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider = "TARGET_PROVIDER_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderAWS captures enum value "AWS".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderAWS VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider = "AWS"

	// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderAZURE captures enum value "AZURE".
	VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderAZURE VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider = "AZURE"
)

// for schema.
var vmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider

	if err := json.Unmarshal([]byte(`["TARGET_PROVIDER_UNSPECIFIED","AWS","AZURE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderEnum = append(vmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderEnum, v)
	}
}
