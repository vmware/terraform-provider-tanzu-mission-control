/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import (
	"testing"

	"github.com/stretchr/testify/require"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
	pkgrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackagerepository"
)

func TestFlattenStatusForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster source secret status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster source secret status",
			input: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus{
				Disabled:   false,
				Managed:    false,
				Subscribed: true,
				Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
					conditionReady: {
						Type:     conditionReady,
						Status:   statusmodel.VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE.Pointer(),
						Severity: statusmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO.Pointer(),
						Reason:   "ReconcileSucceeded",
						Message:  "Reconcile Succeeded",
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					packageRepositoryPhaseKey: "Reconcile Succeeded",
					disabledKey:               false,
					managedKey:                false,
					subscribedKey:             true,
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenStatusForClusterScope(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
