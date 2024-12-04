// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package status

import (
	"testing"

	"github.com/stretchr/testify/require"

	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrepository"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

func TestFlattenStatusForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *helmclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster helm repository status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster helm repository status",
			input: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryStatus{
				Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
					conditionReady: {
						Type:     conditionReady,
						Status:   statusmodel.VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE.Pointer(),
						Severity: statusmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO.Pointer(),
						Reason:   conditionEnabled,
						Message:  "Feature is enabled on the cluster",
					},
				},
			},
			expected: map[string]interface{}{
				phaseKey: conditionEnabled,
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
