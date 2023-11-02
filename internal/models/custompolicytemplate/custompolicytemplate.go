/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplatemodels

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1PolicyTemplate A Policy Template wraps a Kubernetes resource that is a pre-requisite/dependency for creating policies.
// An example of a policy template is OPAGatekeeper based ConstraintTemplate.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.template.Template
type VmwareTanzuManageV1alpha1PolicyTemplate struct {

	// Full name for the policy template.
	FullName *VmwareTanzuManageV1alpha1PolicyTemplateFullName `json:"fullName,omitempty"`

	// Metadata for the policy template object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the policy template.
	Spec *VmwareTanzuManageV1alpha1PolicyTemplateSpec `json:"spec,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplate) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1PolicyTemplate

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
