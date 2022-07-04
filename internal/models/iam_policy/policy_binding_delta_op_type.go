/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/
package iammodel

import (
	"encoding/json"
)

// VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType Type of operation associated with the list of rolebindings.
//
//  - OP_TYPE_UNSPECIFIED: Unspecified operation type.
//  - ADD: Appending rolebindings to the existing policy.
//  - DELETE: Deleting rolebindings from the existing policy.
//
// swagger:model vmware.tanzu.core.v1alpha1.policy.BindingDelta.OpType
type VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType string

func NewVmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType(value VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType) *VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType.
func (m VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType) Pointer() *VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType {
	return &m
}

const (

	// VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeOPTYPEUNSPECIFIED captures enum value "OP_TYPE_UNSPECIFIED".
	VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeOPTYPEUNSPECIFIED VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType = "OP_TYPE_UNSPECIFIED"

	// VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD captures enum value "ADD".
	VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeADD VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType = "ADD"

	// VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeDELETE captures enum value "DELETE".
	VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeDELETE VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType = "DELETE"
)

// for schema.
var vmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeEnum []interface{}

func init() {
	var res []VmwareTanzuCoreV1alpha1PolicyBindingDeltaOpType
	if err := json.Unmarshal([]byte(`["OP_TYPE_UNSPECIFIED","ADD","DELETE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeEnum = append(vmwareTanzuCoreV1alpha1PolicyBindingDeltaOpTypeEnum, v)
	}
}
