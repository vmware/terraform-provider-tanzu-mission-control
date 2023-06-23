/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackage

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec Spec of the Package.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.metadata.package.Spec
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec struct {

	// Minimum capacity requirements to install Package on a cluster.
	CapacityRequirementsDescription string `json:"capacityRequirementsDescription,omitempty"`

	// Licenses under which Package is released.
	Licenses []string `json:"licenses"`

	// Release notes of Package.
	ReleaseNotes string `json:"releaseNotes,omitempty"`

	// Date on which Package is released.
	// Format: date-time
	ReleasedAt strfmt.DateTime `json:"releasedAt,omitempty"`

	// Values schema is used to show template values that can be configured by users.
	ValuesSchema *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageValuesSchema `json:"valuesSchema,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageValuesSchema ValuesSchema is used to show template values that can be configured by users while installing Package.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.metadata.package.ValuesSchema
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageValuesSchema struct {

	// Template values in OpenAPI V3 schema format.
	Template *K8sIoApimachineryPkgRuntimeRawExtension `json:"template,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageValuesSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageValuesSchema) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageValuesSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// swagger:model k8s.io.apimachinery.pkg.runtime.RawExtension
type K8sIoApimachineryPkgRuntimeRawExtension struct {

	// Raw is the underlying serialization of this object.
	//
	// TODO: Determine how to detect ContentType and ContentEncoding of 'Raw' data.
	// Format: byte
	Raw strfmt.Base64 `json:"raw,omitempty"`
}

// MarshalBinary interface implementation.
func (m *K8sIoApimachineryPkgRuntimeRawExtension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *K8sIoApimachineryPkgRuntimeRawExtension) UnmarshalBinary(b []byte) error {
	var res K8sIoApimachineryPkgRuntimeRawExtension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
