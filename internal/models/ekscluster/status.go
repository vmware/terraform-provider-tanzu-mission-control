/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1EksclusterStatus Status of the EKS cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.Status
type VmwareTanzuManageV1alpha1EksclusterStatus struct {

	// Conditions of the cluster resource.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Phase of the cluster resource.
	Phase *VmwareTanzuManageV1alpha1EksclusterPhase `json:"phase,omitempty"`

	// AWS EKS platform version that this cluster uses.
	// https://docs.aws.amazon.com/eks/latest/userguide/platform-versions.html
	PlatformVersion string `json:"platformVersion,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
