package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuCoreV1alpha1StatusCondition Condition describes the status of resource.
// Each resource should provide meaningful set of conditions.
// For Tanzu, each resource must support 'Ready' and 'Scheduled' conditions
// Here is meaning of base conditions and their states:
// Condition 'Ready' with Status 'True' means user action has reached the desired state
// Condition 'Ready' with Status 'False' means user action failed to reach desired state.
// Condition 'Scheduled' with Status 'False' means user action can not be scheduled due to some reason
// Condition 'Scheduled' with Status 'True', Ready unknown means job is scheduled and system is working/will work on reaching to desires state
// Condition 'Scheduled' with Status 'Unknown' means system does not know the status of the action
//
// swagger:model vmware.tanzu.core.v1alpha1.status.Condition
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
