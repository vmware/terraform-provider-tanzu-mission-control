/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyrecipemutationmodel

import (
	"github.com/go-openapi/swag"

	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
	policyrecipemutationcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation/common"
)

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label The input schema for label mutation policy recipe version v1.
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label struct {
	// Label Name and value of the label to be mutated
	Label *policyrecipemutationcommonmodel.KeyValue `json:"label"`

	// Scope Filter the defined target Kubernetes resources by 'Cluster' or 'Namespace' scope. Defaults to '*' (no filter)
	Scope *policyrecipemutationcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope `json:"scope,omitempty"`

	// TargetKubernetesResources List of Kubernetes API resources on which the policy will be enforced, identified using apiGroups and kinds. Use 'kubectl api-resources' to view the list of available API resources
	TargetKubernetesResources []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources `json:"targetKubernetesResources"`
}

func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
