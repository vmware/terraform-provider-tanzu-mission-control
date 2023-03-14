/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
	gitrepositoryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

func TestFlattenStatusForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster git repository status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster git repository status",
			input: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryStatus{
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
				stateKey: fmt.Sprint(statusmodel.VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE),
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
		input       *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster group git repository status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group git repository status",
			input: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryStatus{
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
