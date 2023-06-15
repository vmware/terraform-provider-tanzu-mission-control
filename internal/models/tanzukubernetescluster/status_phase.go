/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcmodels

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// VmwareTanzuManageV1alpha1CommonClusterStatusPhase Phase of the cluster resource.
//
//   - PHASE_UNSPECIFIED: Unspecified phase.
//   - CREATING: Resource is being created.
//   - READY: Resource is in ready state.
//   - DELETING: Resource is being deleted.
//   - ERROR: Error in processing.
//   - UPGRADING: An upgrade is in progress.
//   - UPGRADE_FAILED: An upgrade has failed.
//   - UPDATING: The TanzuKubernetescluster of TKGS is in updating phase.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.Status.Phase
type VmwareTanzuManageV1alpha1CommonClusterStatusPhase string

func NewVmwareTanzuManageV1alpha1CommonClusterStatusPhase(value VmwareTanzuManageV1alpha1CommonClusterStatusPhase) *VmwareTanzuManageV1alpha1CommonClusterStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1CommonClusterStatusPhase.
func (m VmwareTanzuManageV1alpha1CommonClusterStatusPhase) Pointer() *VmwareTanzuManageV1alpha1CommonClusterStatusPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1CommonClusterStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	VmwareTanzuManageV1alpha1CommonClusterStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1CommonClusterStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1CommonClusterStatusPhaseCREATING captures enum value "CREATING"
	VmwareTanzuManageV1alpha1CommonClusterStatusPhaseCREATING VmwareTanzuManageV1alpha1CommonClusterStatusPhase = "CREATING"

	// VmwareTanzuManageV1alpha1CommonClusterStatusPhaseREADY captures enum value "READY"
	VmwareTanzuManageV1alpha1CommonClusterStatusPhaseREADY VmwareTanzuManageV1alpha1CommonClusterStatusPhase = "READY"

	// VmwareTanzuManageV1alpha1CommonClusterStatusPhaseDELETING captures enum value "DELETING"
	VmwareTanzuManageV1alpha1CommonClusterStatusPhaseDELETING VmwareTanzuManageV1alpha1CommonClusterStatusPhase = "DELETING"

	// VmwareTanzuManageV1alpha1CommonClusterStatusPhaseERROR captures enum value "ERROR"
	VmwareTanzuManageV1alpha1CommonClusterStatusPhaseERROR VmwareTanzuManageV1alpha1CommonClusterStatusPhase = "ERROR"

	// VmwareTanzuManageV1alpha1CommonClusterStatusPhaseUPGRADING captures enum value "UPGRADING"
	VmwareTanzuManageV1alpha1CommonClusterStatusPhaseUPGRADING VmwareTanzuManageV1alpha1CommonClusterStatusPhase = "UPGRADING"

	// VmwareTanzuManageV1alpha1CommonClusterStatusPhaseUPGRADEFAILED captures enum value "UPGRADE_FAILED"
	VmwareTanzuManageV1alpha1CommonClusterStatusPhaseUPGRADEFAILED VmwareTanzuManageV1alpha1CommonClusterStatusPhase = "UPGRADE_FAILED"

	// VmwareTanzuManageV1alpha1CommonClusterStatusPhaseUPDATING captures enum value "UPDATING"
	VmwareTanzuManageV1alpha1CommonClusterStatusPhaseUPDATING VmwareTanzuManageV1alpha1CommonClusterStatusPhase = "UPDATING"
)

// for schema
var vmwareTanzuManageV1alpha1CommonClusterStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1CommonClusterStatusPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","CREATING","READY","DELETING","ERROR","UPGRADING","UPGRADE_FAILED","UPDATING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmwareTanzuManageV1alpha1CommonClusterStatusPhaseEnum = append(vmwareTanzuManageV1alpha1CommonClusterStatusPhaseEnum, v)
	}
}

func (m VmwareTanzuManageV1alpha1CommonClusterStatusPhase) validateVmwareTanzuManageV1alpha1CommonClusterStatusPhaseEnum(path, location string, value VmwareTanzuManageV1alpha1CommonClusterStatusPhase) error {
	if err := validate.EnumCase(path, location, value, vmwareTanzuManageV1alpha1CommonClusterStatusPhaseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this vmware tanzu manage v1alpha1 common cluster status phase
func (m VmwareTanzuManageV1alpha1CommonClusterStatusPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVmwareTanzuManageV1alpha1CommonClusterStatusPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this vmware tanzu manage v1alpha1 common cluster status phase based on context it is used
func (m VmwareTanzuManageV1alpha1CommonClusterStatusPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
