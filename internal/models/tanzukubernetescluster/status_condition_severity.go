/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcmodels

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
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

	// VmwareTanzuCoreV1alpha1StatusConditionSeveritySEVERITYUNSPECIFIED captures enum value "SEVERITY_UNSPECIFIED"
	VmwareTanzuCoreV1alpha1StatusConditionSeveritySEVERITYUNSPECIFIED VmwareTanzuCoreV1alpha1StatusConditionSeverity = "SEVERITY_UNSPECIFIED"

	// VmwareTanzuCoreV1alpha1StatusConditionSeverityERROR captures enum value "ERROR"
	VmwareTanzuCoreV1alpha1StatusConditionSeverityERROR VmwareTanzuCoreV1alpha1StatusConditionSeverity = "ERROR"

	// VmwareTanzuCoreV1alpha1StatusConditionSeverityWARNING captures enum value "WARNING"
	VmwareTanzuCoreV1alpha1StatusConditionSeverityWARNING VmwareTanzuCoreV1alpha1StatusConditionSeverity = "WARNING"

	// VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO captures enum value "INFO"
	VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO VmwareTanzuCoreV1alpha1StatusConditionSeverity = "INFO"
)

// for schema
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

func (m VmwareTanzuCoreV1alpha1StatusConditionSeverity) validateVmwareTanzuCoreV1alpha1StatusConditionSeverityEnum(path, location string, value VmwareTanzuCoreV1alpha1StatusConditionSeverity) error {
	if err := validate.EnumCase(path, location, value, vmwareTanzuCoreV1alpha1StatusConditionSeverityEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this vmware tanzu core v1alpha1 status condition severity
func (m VmwareTanzuCoreV1alpha1StatusConditionSeverity) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVmwareTanzuCoreV1alpha1StatusConditionSeverityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this vmware tanzu core v1alpha1 status condition severity based on context it is used
func (m VmwareTanzuCoreV1alpha1StatusConditionSeverity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
