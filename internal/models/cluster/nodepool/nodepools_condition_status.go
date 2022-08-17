/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepool

import "encoding/json"

// VmwareTanzuCoreV1alpha1StatusConditionStatus Status describes the state of condition.
/*
  - STATUS_UNSPECIFIED: Controller is actively working to achieve the condition.
  - TRUE: Reconciliation has succeeded. Once all transition conditions have succeeded, the "happy state" condition should be set to True..
  - FALSE: Reconciliation has failed. This should be a terminal failure state until user action occurs.

 swagger:model vmware.tanzu.core.v1alpha1.status.Condition.Status
*/
type VmwareTanzuCoreV1alpha1StatusConditionStatus string

func NewVmwareTanzuCoreV1alpha1StatusConditionStatus(value VmwareTanzuCoreV1alpha1StatusConditionStatus) *VmwareTanzuCoreV1alpha1StatusConditionStatus {
	v := value
	return &v
}

const (

	// VmwareTanzuCoreV1alpha1StatusConditionStatusSTATUSUNSPECIFIED captures enum value "STATUS_UNSPECIFIED".
	VmwareTanzuCoreV1alpha1StatusConditionStatusSTATUSUNSPECIFIED VmwareTanzuCoreV1alpha1StatusConditionStatus = "STATUS_UNSPECIFIED"

	// VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE captures enum value "TRUE".
	VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE VmwareTanzuCoreV1alpha1StatusConditionStatus = "TRUE"

	// VmwareTanzuCoreV1alpha1StatusConditionStatusFALSE captures enum value "FALSE".
	VmwareTanzuCoreV1alpha1StatusConditionStatusFALSE VmwareTanzuCoreV1alpha1StatusConditionStatus = "FALSE"
)

// for schema.
var vmwareTanzuCoreV1alpha1StatusConditionStatusEnum []interface{}

func init() {
	var res []VmwareTanzuCoreV1alpha1StatusConditionStatus
	if err := json.Unmarshal([]byte(`["STATUS_UNSPECIFIED","TRUE","FALSE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuCoreV1alpha1StatusConditionStatusEnum = append(vmwareTanzuCoreV1alpha1StatusConditionStatusEnum, v)
	}
}
