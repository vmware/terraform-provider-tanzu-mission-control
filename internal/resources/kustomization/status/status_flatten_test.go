// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package status

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	kustomizationclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/cluster"
	kustomizationclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/clustergroup"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

func TestFlattenStatusForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster kustomization status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster kustomization status",
			input: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationStatus{
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
		input       *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster group kustomization status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group kustomization status",
			input: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationStatus{
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
