/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyrecipecustommodel

import (
	"github.com/go-openapi/swag"

	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
)

// VmwareTanzuManageV1alpha1CommonPolicySpecCustom tmc-external-ips recipe schema.
//
// The input schema for tmc-external-ips recipe.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPS
type VmwareTanzuManageV1alpha1CommonPolicySpecCustom struct {

	// Audit (dry-run).
	// Creates this policy for dry-run. Violations will be logged but not denied. Defaults to false (deny).
	Audit bool `json:"audit,omitempty"`

	// Parameters.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// TargetKubernetesResources is a list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. You can use 'kubectl api-resources' to view the list of available api resources on your cluster.
	// Required: true
	// Min Items: 1
	TargetKubernetesResources []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources `json:"targetKubernetesResources"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecCustom) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecCustom) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecCustom

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
