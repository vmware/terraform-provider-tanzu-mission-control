/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
Code generated by go-swagger; DO NOT EDIT.
*/

package policyrecipenetworkmodel

import (
	"github.com/go-openapi/swag"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

// V1alpha1CommonPolicySpecNetworkV1CustomIngress The input schema for network policy custom ingress recipe version v1.
type V1alpha1CommonPolicySpecNetworkV1CustomIngress struct {

	// This specifies list of ingress rules to be applied to the selected pods.
	Rules []policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRules `json:"rules"`

	// Use a label selector to identify the pods to which the policy applies.
	ToPodLabels *[]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels `json:"toPodLabels"`
}

// MarshalBinary interface implementation
func (m *V1alpha1CommonPolicySpecNetworkV1CustomIngress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1CommonPolicySpecNetworkV1CustomIngress) UnmarshalBinary(b []byte) error {
	var res V1alpha1CommonPolicySpecNetworkV1CustomIngress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
