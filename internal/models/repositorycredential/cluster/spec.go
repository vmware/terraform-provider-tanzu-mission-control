/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package repositorycredentialclustermodel

import (
	"encoding/json"

	"github.com/go-openapi/swag"
	credentialmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
)

// Spec of the Source Secret.
type VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec struct {
	// Type of Source Secret(username-password or SSH).
	SourceSecretType *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretType `json:"sourceSecretType,omitempty"`
	// Data is the source credential in the form of key-value.
	Data *credentialmodels.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec `json:"data,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

type VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretType string

func NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretType(value VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretType) *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretType {
	return &value
}

//nolint:all
const (
	VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretTypeUNSPECIFIED VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretType = "UNSPECIFIED"

	VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretTypeUSERNAMEPASSWORD VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretType = "USERNAME_PASSWORD"

	VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretTypeSSH VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretType = "SSH"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretType
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","USERNAME_PASSWORD","SSH"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretTypeEnum = append(vmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretTypeEnum, v)
	}
}
