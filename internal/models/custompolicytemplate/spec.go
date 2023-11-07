/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplatemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1PolicyTemplateSpec Spec of policy template.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.template.Spec
type VmwareTanzuManageV1alpha1PolicyTemplateSpec struct {

	// DataInventory is a list of Kubernetes api-resource kinds that need to be synced/replicated
	// in Gatekeeper in order to enforce policy rules on those resources.
	// Note: This is used for OPAGatekeeper based templates, and should be used if the policy
	// enforcement logic in Rego code uses cached data using "data.inventory" fields.
	DataInventory []*K8sIoApimachineryPkgApisMetaV1GroupVersionKind `json:"dataInventory"`

	// Deprecated specifies whether this version (latest version) of the policy template is deprecated.
	// Updating a policy template deprecates the previous versions. To view all versions, use Versions API.
	Deprecated bool `json:"deprecated"`

	// Object is a yaml-formatted Kubernetes resource.
	// The Kubernetes object has to be of the type defined in ObjectType ('ConstraintTemplate').
	// The object name must match the name of the wrapping policy template.
	// This will be applied on the cluster after a policy is created using this version of the template.
	// This contains the latest version of the object. For previous versions, check Versions API.
	Object string `json:"object,omitempty"`

	// ObjectType is the type of Kubernetes resource encoded in Object.
	// Currently, we only support OPAGatekeeper based 'ConstraintTemplate' object.
	ObjectType string `json:"objectType,omitempty"`

	// PolicyUpdateStrategy on how to handle policies after a policy template update.
	PolicyUpdateStrategy *VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategy `json:"policyUpdateStrategy,omitempty"`

	// TemplateType is the type of policy template.
	// Currently, we only support 'OPAGatekeeper' based policy templates.
	TemplateType string `json:"templateType,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplateSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplateSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1PolicyTemplateSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
