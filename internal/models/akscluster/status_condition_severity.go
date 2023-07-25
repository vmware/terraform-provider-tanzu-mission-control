/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"encoding/json"
)

// VmwareTanzuCoreV1alpha1StatusConditionSeverity Severity expresses the severity of a Condition Type failing.
//
//   - SEVERITY_UNSPECIFIED: Unspecified severity.
//   - ERROR: Failure of a condition type should be viewed as an error.
//   - WARNING: Failure of a condition type should be viewed as a warning, but that things could still work.
//   - INFO: Failure of a condition type should be viewed as purely informational, and that things could still work.
//
// swagger:model vmware.tanzu.core.v1alpha1.status.Condition.Severity
type VmwareTanzuCoreV1alpha1StatusConditionSeverity string

func NewVmwareTanzuCoreV1alpha1StatusConditionSeverity(value VmwareTanzuCoreV1alpha1StatusConditionSeverity) *VmwareTanzuCoreV1alpha1StatusConditionSeverity {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuCoreV1alpha1StatusConditionSeverity.
func (m VmwareTanzuCoreV1alpha1StatusConditionSeverity) Pointer() *VmwareTanzuCoreV1alpha1StatusConditionSeverity {
	return &m
}

const (

	// VmwareTanzuCoreV1alpha1StatusConditionSeveritySEVERITYUNSPECIFIED captures enum value "SEVERITY_UNSPECIFIED".
	VmwareTanzuCoreV1alpha1StatusConditionSeveritySEVERITYUNSPECIFIED VmwareTanzuCoreV1alpha1StatusConditionSeverity = "SEVERITY_UNSPECIFIED"

	// VmwareTanzuCoreV1alpha1StatusConditionSeverityERROR captures enum value "ERROR".
	VmwareTanzuCoreV1alpha1StatusConditionSeverityERROR VmwareTanzuCoreV1alpha1StatusConditionSeverity = "ERROR"

	// VmwareTanzuCoreV1alpha1StatusConditionSeverityWARNING captures enum value "WARNING".
	VmwareTanzuCoreV1alpha1StatusConditionSeverityWARNING VmwareTanzuCoreV1alpha1StatusConditionSeverity = "WARNING"

	// VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO captures enum value "INFO".
	VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO VmwareTanzuCoreV1alpha1StatusConditionSeverity = "INFO"
)

// for schema.
var vmwareTanzuCoreV1alpha1StatusConditionSeverityEnum []interface{}

func init() {
	var res []VmwareTanzuCoreV1alpha1StatusConditionSeverity
	if err := json.Unmarshal([]byte(`["SEVERITY_UNSPECIFIED","ERROR","WARNING","INFO"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuCoreV1alpha1StatusConditionSeverityEnum = append(vmwareTanzuCoreV1alpha1StatusConditionSeverityEnum, v)
	}
}
