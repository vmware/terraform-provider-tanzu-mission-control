// Copyright © 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: MPL-2.0

package policyrecipemutationmodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity The input schema for pod-security mutation policy recipe version v1.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity struct {

	// allow privilege escalation
	AllowPrivilegeEscalation *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation `json:"allowPrivilegeEscalation,omitempty"`

	// capabilities add
	CapabilitiesAdd *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd `json:"capabilitiesAdd,omitempty"`

	// capabilities drop
	CapabilitiesDrop *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop `json:"capabilitiesDrop,omitempty"`

	// fs group
	FsGroup *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup `json:"fsGroup,omitempty"`

	// privileged
	Privileged *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged `json:"privileged,omitempty"`

	// read only root filesystem
	ReadOnlyRootFilesystem *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem `json:"readOnlyRootFilesystem,omitempty"`

	// run as group
	RunAsGroup *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup `json:"runAsGroup,omitempty"`

	// run as non root
	RunAsNonRoot *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot `json:"runAsNonRoot,omitempty"`

	// run as user
	RunAsUser *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser `json:"runAsUser,omitempty"`

	// se linux options
	SeLinuxOptions *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions `json:"seLinuxOptions,omitempty"`

	// supplemental groups
	SupplementalGroups *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups `json:"supplementalGroups,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation Set allowPrivilegeEscalation flag in container security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation struct {

	// Condition specifies whether to always mutate/set this value or only if container security context contains or does not contain this field
	// Required: true
	// Enum: [Always IfFieldExists IfFieldDoesNotExist]
	Condition *string `json:"condition"`

	// Value to set for allowPrivilegeEscalation field in container security context
	// Required: true
	Value *bool `json:"value"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityAllowPrivilegeEscalation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd Set linux capabilities.add field in container security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd struct {

	// Option to either override the list or merge values into the list or prune values from the list
	// Required: true
	// Enum: [override merge prune]
	Operation *string `json:"operation"`

	// List of values to override/merge/prune in capabilities.add field in container security context
	// Required: true
	// Min Items: 1
	Values []string `json:"values"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesAdd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop Set linux capabilities.drop field in container security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop struct {

	// Option to either override the list or merge values into the list or prune values from the list
	// Required: true
	// Enum: [override merge prune]
	Operation *string `json:"operation"`

	// List of values to override/merge/prune in capabilities.drop field in container security context.
	// Required: true
	// Min Items: 1
	Values []string `json:"values"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityCapabilitiesDrop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup Set numerical supplemental group ID in fsGroup flag in pod security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup struct {

	// Condition specifies whether to always mutate/set this value or only if pod security context contains or does not contain this field
	// Required: true
	// Enum: [Always IfFieldExists IfFieldDoesNotExist]
	Condition *string `json:"condition"`

	// Value to set for fsGroup field in pod security context
	// Required: true
	// Maximum: 65535
	// Minimum: 0
	Value *float64 `json:"value"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityFsGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged Set privileged flag in container security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged struct {

	// Condition specifies whether to always mutate/set this value or only if container security context contains or does not contain this field
	// Required: true
	// Enum: [Always IfFieldExists IfFieldDoesNotExist]
	Condition *string `json:"condition"`

	// Value to set for privileged field in container security context
	// Required: true
	Value *bool `json:"value"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityPrivileged
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem Set readOnlyRootFilesystem flag in container security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem struct {

	// Condition specifies whether to always mutate/set this value or only if container security context contains or does not contain this field
	// Required: true
	// Enum: [Always IfFieldExists IfFieldDoesNotExist]
	Condition *string `json:"condition"`

	// Value to set for readOnlyRootFilesystem field in container security context
	// Required: true
	Value *bool `json:"value"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityReadOnlyRootFilesystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup Set numerical group ID in runAsGroup flag in container security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup struct {

	// Condition specifies whether to always mutate/set this value or only if container security context contains or does not contain this field
	// Required: true
	// Enum: [Always IfFieldExists IfFieldDoesNotExist]
	Condition *string `json:"condition"`

	// Value to set for runAsGroup field in container security context
	// Required: true
	// Maximum: 65535
	// Minimum: 0
	Value *float64 `json:"value"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot Set runAsNonRoot flag in container security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot struct {

	// Condition specifies whether to always mutate/set this value or only if container security context contains or does not contain this field
	// Required: true
	// Enum: [Always IfFieldExists IfFieldDoesNotExist]
	Condition *string `json:"condition"`

	// Value to set for runAsNonRoot field in container security context
	// Required: true
	Value *bool `json:"value"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsNonRoot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser Set numerical user ID in runAsUser flag in container security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser struct {

	// Condition specifies whether to always mutate/set this value or only if container security context contains or does not contain this field
	// Required: true
	// Enum: [Always IfFieldExists IfFieldDoesNotExist]
	Condition *string `json:"condition"`

	// Value to set for runAsUser field in container security context
	// Required: true
	// Maximum: 65535
	// Minimum: 0
	Value *float64 `json:"value"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurityRunAsUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions Set seLinuxOptions in container security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions struct {

	// Condition specifies whether to always mutate/set this value or only if container security context contains or does not contain this field
	// Required: true
	// Enum: [Always IfFieldExists IfFieldDoesNotExist]
	Condition *string `json:"condition"`

	// value
	// Required: true
	Value *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptionsValue `json:"value"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptionsValue Value to set for seLinuxOptions field in container security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptionsValue
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptionsValue struct {

	// Value to set for level field in container security context seLinuxOptions
	Level string `json:"level,omitempty"`

	// Value to set for role field in container security context seLinuxOptions
	Role string `json:"role,omitempty"`

	// Value to set for type field in container security context seLinuxOptions
	Type string `json:"type,omitempty"`

	// Value to set for user field in container security context seLinuxOptions
	User string `json:"user,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptionsValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptionsValue) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySeLinuxOptionsValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups Set list of supplemental group IDs in supplementalGroups flag in pod security context.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups struct {

	// Condition specifies whether to always mutate/set this value or only if pod security context contains or does not contain this field
	// Required: true
	// Enum: [Always IfFieldExists IfFieldDoesNotExist]
	Condition *string `json:"condition"`

	// List of values to set for supplementalGroups field in pod security context
	// Required: true
	// Min Items: 1
	Values []*float64 `json:"values"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecuritySupplementalGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
