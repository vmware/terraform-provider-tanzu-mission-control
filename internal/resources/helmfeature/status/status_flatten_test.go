/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/cluster"
	helmclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

func TestFlattenStatusForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster git repository status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster git repository status",
			input: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus{
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

func TestFlattenStatusForClusterGroupScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster group git repository status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group git repository status",
			input: &helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmStatus{
				Phase: statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED.Pointer(),
				Details: &statusmodel.VmwareTanzuManageV1alpha1CommonBatchDetails{
					Overridden:       0,
					Applied:          3,
					AvailableTargets: 3,
				},
			},
			expected: map[string]interface{}{
				phaseKey: fmt.Sprint(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenStatusForClusterGroupScope(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
