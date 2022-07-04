/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iammodel

import (
	"encoding/json"
)

// VmwareTanzuCoreV1alpha1PolicySubjectKind Kind of subject.
//
//  - KIND_UNSPECIFIED: Subject is a undefined.
//  - GROUP: Subject is a group.
//  - SERVICEACCOUNT: Subject is a service.
//  - USER: Subject is a user.
//
// swagger:model vmware.tanzu.core.v1alpha1.policy.Subject.Kind
type VmwareTanzuCoreV1alpha1PolicySubjectKind string

func NewVmwareTanzuCoreV1alpha1PolicySubjectKind(value VmwareTanzuCoreV1alpha1PolicySubjectKind) *VmwareTanzuCoreV1alpha1PolicySubjectKind {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuCoreV1alpha1PolicySubjectKind.
func (m VmwareTanzuCoreV1alpha1PolicySubjectKind) Pointer() *VmwareTanzuCoreV1alpha1PolicySubjectKind {
	return &m
}

const (

	// VmwareTanzuCoreV1alpha1PolicySubjectKindKINDUNSPECIFIED captures enum value "KIND_UNSPECIFIED".
	VmwareTanzuCoreV1alpha1PolicySubjectKindKINDUNSPECIFIED VmwareTanzuCoreV1alpha1PolicySubjectKind = "KIND_UNSPECIFIED"

	// VmwareTanzuCoreV1alpha1PolicySubjectKindGROUP captures enum value "GROUP".
	VmwareTanzuCoreV1alpha1PolicySubjectKindGROUP VmwareTanzuCoreV1alpha1PolicySubjectKind = "GROUP"

	// VmwareTanzuCoreV1alpha1PolicySubjectKindSERVICEACCOUNT captures enum value "SERVICEACCOUNT".
	VmwareTanzuCoreV1alpha1PolicySubjectKindSERVICEACCOUNT VmwareTanzuCoreV1alpha1PolicySubjectKind = "SERVICEACCOUNT"

	// VmwareTanzuCoreV1alpha1PolicySubjectKindUSER captures enum value "USER".
	VmwareTanzuCoreV1alpha1PolicySubjectKindUSER VmwareTanzuCoreV1alpha1PolicySubjectKind = "USER"
)

// for schema.
var vmwareTanzuCoreV1alpha1PolicySubjectKindEnum []interface{}

func init() {
	var res []VmwareTanzuCoreV1alpha1PolicySubjectKind
	if err := json.Unmarshal([]byte(`["KIND_UNSPECIFIED","GROUP","SERVICEACCOUNT","USER"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuCoreV1alpha1PolicySubjectKindEnum = append(vmwareTanzuCoreV1alpha1PolicySubjectKindEnum, v)
	}
}
