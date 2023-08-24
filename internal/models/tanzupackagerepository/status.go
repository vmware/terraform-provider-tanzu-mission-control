/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackagerepository

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus Status of Package Repository resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.repository.Status.
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus struct {

	// Conditions of the Package Repository resource.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// If true, the Package Repository is disabled.
	Disabled bool `json:"disabled,omitempty"`

	// If true, the Package Repository is managed by TMC.
	Managed bool `json:"managed,omitempty"`

	// If true, the Package Repository has been subscribed by user organization.
	Subscribed bool `json:"subscribed,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
