/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclass

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassStatus Status of the cluster class.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.clusterclass.Status
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassStatus struct {

	// Conditions of the resource.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Schema of cluster class variables for UI to render.
	VariablesSchema *K8sIoApimachineryPkgRuntimeRawExtension `json:"variablesSchema,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
