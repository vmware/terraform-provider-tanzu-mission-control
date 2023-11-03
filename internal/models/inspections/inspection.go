/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionsmodels

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterInspectionScan Run an on demand inspection scan on a cluster
//
// Running an inspection scan verifies whether the cluster is certified conformant or security compliant.
// It is a diagnostic tool that helps you understand the state of a cluster by running a set of tests and
// provides clear informative reports about the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.Scan
type VmwareTanzuManageV1alpha1ClusterInspectionScan struct {

	// Full name for the inspection scan.
	FullName *VmwareTanzuManageV1alpha1ClusterInspectionScanFullName `json:"fullName,omitempty"`

	// Metadata for the inspection object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Inspection Scan spec.
	Spec *VmwareTanzuManageV1alpha1ClusterInspectionScanSpec `json:"spec,omitempty"`

	// Status of the Inspection Scan object.
	Status *VmwareTanzuManageV1alpha1ClusterInspectionScanStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScan) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInspectionScan

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
