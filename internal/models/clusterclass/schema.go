/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclass

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassVariableSchema VariableSchema defines the schema of a variable.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.clusterclass.VariableSchema
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassVariableSchema struct {

	// Template values in OpenAPI V3 schema format.
	Template *K8sIoApimachineryPkgRuntimeRawExtension `json:"template,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassVariableSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassVariableSchema) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassVariableSchema

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// swagger:model k8s.io.apimachinery.pkg.runtime.RawExtension
type K8sIoApimachineryPkgRuntimeRawExtension struct {

	// Raw is the underlying serialization of this object.
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
