/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	packageinstallmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackageinstall"
)

func TestFlattenSpecForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster package install spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster package install spec",
			input: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec{
				PackageRef: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef{
					PackageMetadataName: "cert-manager.tanzu.vmware.com",
					VersionSelection: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection{
						Constraints: "1.10.2+vmware.1-tkg.1",
					},
				},
				RoleBindingScope: packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER.Pointer(),
				InlineValues: map[string]interface{}{
					"namespace": "cert-manager",
					"some":      91,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					PackageRefKey: []interface{}{
						map[string]interface{}{
							PackageMetadataNameKey: "cert-manager.tanzu.vmware.com",
							VersionSelectionKey: []interface{}{
								map[string]interface{}{
									ConstraintsKey: "1.10.2+vmware.1-tkg.1",
								},
							},
						},
					},
					RoleBindingScopeKey: fmt.Sprint(packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER),
					InlineValuesKey: map[string]interface{}{
						"namespace": "cert-manager",
						"some":      91,
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenSpecForClusterScope(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
