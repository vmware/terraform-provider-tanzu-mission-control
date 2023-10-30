/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationmodels

import (
	"github.com/go-openapi/swag"

	credentialsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
)

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationSpec The backup location spec.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.Spec.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationSpec struct {

	// List of groups the backup location will be assigned to.
	AssignedGroups []*VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAssignedGroup `json:"assignedGroups"`

	// The bucket to use for object storage.
	Bucket string `json:"bucket,omitempty"`

	// A PEM-encoded certificate bundle to trust while connecting to the storage backend. Optional.
	CaCert string `json:"caCert,omitempty"`

	// Provider-specific storage configuration fields.
	Config *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderSpecificConfig `json:"config"`

	// The name of credential to be used to access the bucket.
	Credential *credentialsmodel.VmwareTanzuManageV1alpha1AccountCredentialFullName `json:"credential,omitempty"`

	// The region of the bucket origin. Optional.
	Region string `json:"region,omitempty"`

	// The target provider of the backup storage.
	TargetProvider *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProvider `json:"targetProvider,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
