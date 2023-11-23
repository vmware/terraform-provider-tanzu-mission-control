/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterintegrationmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterIntegrationIndicator Indicator is the value of the abstracted statuses.
//
//   - INDICATOR_UNSPECIFIED: Default indicator.
//   - OK: OK indicates everything is good.
//   - ATTENTION_REQUIRED: ATTENTION_REQUIRED indicates something is bad / requires attention of user.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.integration.Indicator
type VmwareTanzuManageV1alpha1ClusterIntegrationIndicator string

func NewVmwareTanzuManageV1alpha1ClusterIntegrationIndicator(value VmwareTanzuManageV1alpha1ClusterIntegrationIndicator) *VmwareTanzuManageV1alpha1ClusterIntegrationIndicator {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterIntegrationIndicator.
func (m VmwareTanzuManageV1alpha1ClusterIntegrationIndicator) Pointer() *VmwareTanzuManageV1alpha1ClusterIntegrationIndicator {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterIntegrationIndicatorINDICATORUNSPECIFIED captures enum value "INDICATOR_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterIntegrationIndicatorINDICATORUNSPECIFIED VmwareTanzuManageV1alpha1ClusterIntegrationIndicator = "INDICATOR_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterIntegrationIndicatorOK captures enum value "OK".
	VmwareTanzuManageV1alpha1ClusterIntegrationIndicatorOK VmwareTanzuManageV1alpha1ClusterIntegrationIndicator = "OK"

	// VmwareTanzuManageV1alpha1ClusterIntegrationIndicatorATTENTIONREQUIRED captures enum value "ATTENTION_REQUIRED".
	VmwareTanzuManageV1alpha1ClusterIntegrationIndicatorATTENTIONREQUIRED VmwareTanzuManageV1alpha1ClusterIntegrationIndicator = "ATTENTION_REQUIRED"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterIntegrationIndicatorEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterIntegrationIndicator

	if err := json.Unmarshal([]byte(`["INDICATOR_UNSPECIFIED","OK","ATTENTION_REQUIRED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterIntegrationIndicatorEnum = append(vmwareTanzuManageV1alpha1ClusterIntegrationIndicatorEnum, v)
	}
}
