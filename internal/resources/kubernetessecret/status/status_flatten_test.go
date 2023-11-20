/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	secretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
	secretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

func TestFlattenStatusForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster secret status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster secret status",
			input: &secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretStatus{
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
		input       *secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster group secret status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group secret status",
			input: &secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretStatus{
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
