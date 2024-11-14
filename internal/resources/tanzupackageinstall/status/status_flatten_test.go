// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package status

import (
	"testing"

	"github.com/stretchr/testify/require"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
	pkginstallclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackageinstall"
)

func TestFlattenStatusForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus
		expected    interface{}
	}{
		{
			description: "check for nil cluster package install status",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster package install status",
			input: &pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus{
				Managed: false,
				Conditions: map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition{
					conditionReady: {
						Type:     conditionReady,
						Status:   statusmodel.VmwareTanzuCoreV1alpha1StatusConditionStatusTRUE.Pointer(),
						Severity: statusmodel.VmwareTanzuCoreV1alpha1StatusConditionSeverityINFO.Pointer(),
						Reason:   "ReconcileSucceeded",
						Message:  "Reconcile Succeeded",
					},
				},
				ResolvedVersion: "1.9.5+vmware.1-tkg.1",
				GeneratedResources: &pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallGeneratedResources{
					ClusterRoleName:    "testclusterrole",
					RoleBindingName:    "testrolebinding",
					ServiceAccountName: "testserviceaccount",
				},
				ReferredBy: []string{"foo", "bar"},
			},
			expected: []interface{}{
				map[string]interface{}{
					packageInstallPhaseKey: "Reconcile Succeeded",
					managedKey:             false,
					resolvedVersionKey:     "1.9.5+vmware.1-tkg.1",
					referredByKey:          []string{"foo", "bar"},
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
