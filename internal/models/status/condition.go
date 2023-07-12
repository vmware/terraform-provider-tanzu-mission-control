/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package statusmodel

import (
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuCoreV1alpha1StatusCondition Condition describes the status of resource.
// Each resource should provide meaningful set of conditions.
<<<<<<< HEAD
// For Tanzu, each resource must support 'Ready' and 'Scheduled' conditions
// Here is meaning of base conditions and their states:
// Condition 'Ready' with Status 'True' means user action has reached the desired state
// Condition 'Ready' with Status 'False' means user action failed to reach desired state.
// Condition 'Scheduled' with Status 'False' means user action can not be scheduled due to some reason
// Condition 'Scheduled' with Status 'True', Ready unknown means job is scheduled and system is working/will work on reaching to desires state
// Condition 'Scheduled' with Status 'Unknown' means system does not know the status of the action
//
// swagger:model vmware.tanzu.core.v1alpha1.status.Condition
=======
// For Tanzu, each resource must support 'Ready' and 'Scheduled' conditions.
// Here is meaning of base conditions and their states:
// Condition 'Ready' with Status 'True' means user action has reached the desired state.
// Condition 'Ready' with Status 'False' means user action failed to reach desired state.
// Condition 'Scheduled' with Status 'False' means user action can not be scheduled due to some reason.
// Condition 'Scheduled' with Status 'True', Ready unknown means job is scheduled and system is working/will work on reaching to desires state.
// Condition 'Scheduled' with Status 'Unknown' means system does not know the status of the action.
//
// swagger:model vmware.tanzu.core.v1alpha1.status.Condition.
>>>>>>> 2dc0bf6 (Add models and clients for Package repository resource and define schema)
type VmwareTanzuCoreV1alpha1StatusCondition struct {

	// Last time the condition transit from one status to another.
	// Format: date-time
	LastTransitionTime strfmt.DateTime `json:"lastTransitionTime,omitempty"`

	// Human readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// One-word reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Severity of condition, one of Error, Warning, Info.
	// Default is Error.
	Severity *VmwareTanzuCoreV1alpha1StatusConditionSeverity `json:"severity,omitempty"`

	// Status of the condition, one of True, False, Unknown.
	// Default is Unknown.
	Status *VmwareTanzuCoreV1alpha1StatusConditionStatus `json:"status,omitempty"`

	// Type of condition.
	Type string `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1StatusCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1StatusCondition) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1StatusCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuCoreV1alpha1StatusConditionSeverity Severity expresses the severity of a Condition Type failing.
//
//   - SEVERITY_UNSPECIFIED: Unspecified severity.
//   - ERROR: Failure of a condition type should be viewed as an error.
//   - WARNING: Failure of a condition type should be viewed as a warning, but that things could still work.
//   - INFO: Failure of a condition type should be viewed as purely informational, and that things could still work.
//
<<<<<<< HEAD
// swagger:model vmware.tanzu.core.v1alpha1.status.Condition.Severity
=======
// swagger:model vmware.tanzu.core.v1alpha1.status.Condition.Severity.
>>>>>>> 2dc0bf6 (Add models and clients for Package repository resource and define schema)
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

// VmwareTanzuCoreV1alpha1StatusConditionStatus Status describes the state of condition.
//
//   - STATUS_UNSPECIFIED: Controller is actively working to achieve the condition.
//   - TRUE: Reconciliation has succeeded. Once all transition conditions have succeeded, the "happy state" condition should be set to True..
//   - FALSE: Reconciliation has failed. This should be a terminal failure state until user action occurs.
//
<<<<<<< HEAD
// swagger:model vmware.tanzu.core.v1alpha1.status.Condition.Status
=======
// swagger:model vmware.tanzu.core.v1alpha1.status.Condition.Status.
>>>>>>> 2dc0bf6 (Add models and clients for Package repository resource and define schema)
type VmwareTanzuCoreV1alpha1StatusConditionStatus string

func NewVmwareTanzuCoreV1alpha1StatusConditionStatus(value VmwareTanzuCoreV1alpha1StatusConditionStatus) *VmwareTanzuCoreV1alpha1StatusConditionStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuCoreV1alpha1StatusConditionStatus.
func (m VmwareTanzuCoreV1alpha1StatusConditionStatus) Pointer() *VmwareTanzuCoreV1alpha1StatusConditionStatus {
	return &m
}

const (

	// VmwareTanzuCoreV1alpha1StatusConditionStatusSTATUSUNSPECIFIED captures enum value "STATUS_UNSPECIFIED".
	VmwareTanzuCoreV1alpha1StatusConditionStatusSTATUSUNSPECIFIED VmwareTanzuCoreV1alpha1StatusConditionStatus = "STATUS_UNSPECIFIED"

	// VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE captures enum value "TRUE".
	VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE VmwareTanzuCoreV1alpha1StatusConditionStatus = "TRUE"

	// VmwareTanzuCoreV1alpha1StatusConditionStatusFALSE captures enum value "FALSE".
	VmwareTanzuCoreV1alpha1StatusConditionStatusFALSE VmwareTanzuCoreV1alpha1StatusConditionStatus = "FALSE"
)

// for schema.
var vmwareTanzuCoreV1alpha1StatusConditionStatusEnum []interface{}

func init() {
	var res []VmwareTanzuCoreV1alpha1StatusConditionStatus
	if err := json.Unmarshal([]byte(`["STATUS_UNSPECIFIED","TRUE","FALSE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuCoreV1alpha1StatusConditionStatusEnum = append(vmwareTanzuCoreV1alpha1StatusConditionStatusEnum, v)
	}
}
