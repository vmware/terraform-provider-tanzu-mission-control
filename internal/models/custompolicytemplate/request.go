/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplatemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1PolicyTemplateCreateTemplateRequest Request to create a Template.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.template.CreateTemplateRequest
type VmwareTanzuManageV1alpha1PolicyTemplateData struct {

	// Template to create.
	Template *VmwareTanzuManageV1alpha1PolicyTemplate `json:"template,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplateData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplateData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1PolicyTemplateData

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
