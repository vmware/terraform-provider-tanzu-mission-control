// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

// Code generated by go-swagger; DO NOT EDIT.

package policyrecipecustomcommonmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources TargetKubernetes Resources is a list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources
type VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources struct {

	// APIGroup is a group containing the resource type, for example 'rbac.authorization.k8s.io', 'networking.k8s.io', 'extensions', '' (some resources like Namespace, Pod have empty apiGroup).
	// Required: true
	// Min Items: 1
	APIGroups []string `json:"apiGroups"`

	// Kind is the name of the object schema (resource type), for example 'Namespace', 'Pod', 'Ingress'
	// Required: true
	// Min Items: 1
	Kinds []string `json:"kinds"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
