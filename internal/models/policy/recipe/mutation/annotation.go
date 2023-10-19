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

// VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation The input schema for annotation mutation policy recipe version v1.
type VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation struct {
	Annotation *policyrecipemutationcommonmodel.KeyValue `json:"annotation"`

	Scope *policyrecipemutationcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope `json:"scope,omitempty"`

	TargetKubernetesResources []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources `json:"targetKubernetesResources"`
}

func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

func (m *VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
