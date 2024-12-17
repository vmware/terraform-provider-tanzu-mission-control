// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package status

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	releaseclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/cluster"
	releaseclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/clustergroup"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

func TestFlattenStatusForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster helm release status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster helm release status",
			input: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus{
				Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
					conditionReady: {
						Type:     conditionReady,
						Status:   statusmodel.VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE.Pointer(),
						Severity: statusmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO.Pointer(),
						Reason:   conditionEnabled,
						Message:  "Feature is enabled on the cluster",
					},
				},
				GeneratedResources: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseGeneratedResources{
					ClusterRoleName:    "testclusterrole",
					RoleBindingName:    "testrolebinding",
					ServiceAccountName: "testserviceaccount",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					phaseKey: conditionEnabled,
					generatedResourcesKey: []interface{}{
						map[string]interface{}{
							clusterRoleNameKey:    "testclusterrole",
							roleBindingNameKey:    "testrolebinding",
							serviceAccountNameKey: "testserviceaccount",
						},
					},
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

func TestFlattenStatusForClusterGroupScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster group helm release status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group helm release status",
			input: &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus{
				Phase: statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED.Pointer(),
				Details: &statusmodel.VmwareTanzuManageV1alpha1CommonBatchDetails{
					Overridden:       0,
					Applied:          3,
					AvailableTargets: 3,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					phaseKey: fmt.Sprint(statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhaseAPPLIED),
				},
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
