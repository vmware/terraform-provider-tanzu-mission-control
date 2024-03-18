/*
Copyright 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

const ()

// VmwareTanzuManageV1alpha1EksclusterAuthenticationMode The EKS authentication mode.
//
//   - API: Controlled only by API.
//   - API_AND_CONFIG_MAP: Controlled by API and by config map.
//   - CONFIG_MAP: Controlled only by config map.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.AuthenticationMode
type VmwareTanzuManageV1alpha1EksclusterAuthenticationMode string

func NewVmwareTanzuManageV1alpha1EksAuthenticationMode(value VmwareTanzuManageV1alpha1EksclusterPhase) *VmwareTanzuManageV1alpha1EksclusterPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1EksclusterAuthenticationMode.
func (m VmwareTanzuManageV1alpha1EksclusterAuthenticationMode) Pointer() *VmwareTanzuManageV1alpha1EksclusterAuthenticationMode {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1EksclusterAuthenticationModeAPI captures enum value "API".
	VmwareTanzuManageV1alpha1EksclusterAuthenticationModeAPI VmwareTanzuManageV1alpha1EksclusterPhase = "API"

	// VmwareTanzuManageV1alpha1EksclusterAuthenticationModeAPIANDCONFIGMAP captures enum value "API_AND_CONFIG_MAP".
	VmwareTanzuManageV1alpha1EksclusterAuthenticationModeAPIANDCONFIGMAP VmwareTanzuManageV1alpha1EksclusterPhase = "API_AND_CONFIG_MAP"

	// VmwareTanzuManageV1alpha1EksclusterAuthenticationModeCONFIGMAP captures enum value "CONFIG_MAP".
	VmwareTanzuManageV1alpha1EksclusterAuthenticationModeCONFIGMAP VmwareTanzuManageV1alpha1EksclusterPhase = "CONFIG_MAP"
)
