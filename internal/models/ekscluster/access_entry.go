/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterAccessEntry EKS cluster IAM Access Entry.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.AccessEntry
type VmwareTanzuManageV1alpha1EksclusterAccessEntry struct {

	// The AccessEntry will be primary if the value is set to true.
	IsPrimary bool `json:"isPrimary,omitempty"`

	// The user name to which the access entry is associated.
	UserName string `json:"userName,omitempty"`

	// The group names to which the access entry is associated.
	GroupNames []string `json:"groupNames,omitempty"`

	// Tags applied to AccessEntry.
	Tags map[string]string `json:"tags,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterAccessEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterAccessEntry) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterAccessEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
