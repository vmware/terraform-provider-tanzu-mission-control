/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package repositorycredentialclustermodel

import (
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Spec of the Source Secret.
type VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialSpec struct {
	// Type of Source Secret(username-password or SSH).
	RepositorycredentialType *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialType `json:"repositorycredentialType,omitempty"`
	// Data is the source credential in the form of key-value.
	Data map[string]strfmt.Base64 `json:"data,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialType RepositoryCredentialType definition - indicates the repository credential type.
//
//   - REPOSITORYCREDENTIAL_TYPE_UNSPECIFIED: REPOSITORYCREDENTIAL_TYPE_UNSPECIFIED, Unspecified credential type (default).
//   - REPOSITORYCREDENTIAL_TYPE_USERNAME_PASSWORD: REPOSITORYCREDENTIAL_TYPE_USERNAME_PASSWORD, basic auth
//   - REPOSITORYCREDENTIAL_TYPE_SSH: REPOSITORYCREDENTIAL_TYPE_SSH, ssh key
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.respoistorycredential.repositorycredentialtype
type VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialType string

func NewVmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialType(value VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialType) *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialType {
	return &value
}

//nolint:all
const (

	// VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialTypeREPOSITORYCREDENTIALTYPEUNSPECIFIED captures enum value "REPOSITORYCREDENTIAL_TYPE_UNSPECIFIED"
	VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialTypeREPOSITORYCREDENTIALTYPEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialType = "REPOSITORYCREDENTIAL_TYPE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialTypeREPOSITORYCREDENTIALTYPEUSERNAMEPASSWORD captures enum value "REPOSITORYCREDENTIAL_TYPE_DOCKERCONFIGJSON"
	VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialTypeREPOSITORYCREDENTIALTYPEUSERNAMEPASSWORD VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialType = "REPOSITORYCREDENTIAL_TYPE_USERNAME_PASSWORD"

	// VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialTypeREPOSITORYCREDENTIALTYPESSH captures enum value "REPOSITORYCREDENTIAL_TYPE_DOCKERCONFIGJSON"
	VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialTypeREPOSITORYCREDENTIALTYPSSH VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialType = "REPOSITORYCREDENTIAL_TYPE_SSH"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredentialType
	if err := json.Unmarshal([]byte(`["REPOSITORYCREDENTIAL_TYPE_UNSPECIFIED","REPOSITORYCREDENTIAL_TYPE_USERNAME_PASSWORD","REPOSITORYCREDENTIAL_TYPE_SSH"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialTypeEnum = append(vmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialTypeEnum, v)
	}
}
