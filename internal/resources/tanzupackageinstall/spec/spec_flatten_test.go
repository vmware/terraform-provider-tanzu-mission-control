// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	packageinstallmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackageinstall"
)

const (
	test1102vmware1tkg1           = "1.10.2+vmware.1-tkg.1"
	testCertManager               = "cert-manager"
	testCertmanagertanzuvmwarecom = "cert-manager.tanzu.vmware.com"
	testSome                      = "some"

	testNamespace = "namespace"
)

func TestFlattenSpecForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec
		expected    []interface{}
		filePath    string
	}{
		{
			description: "check for nil cluster package install spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster package install spec with inline values file",
			input: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec{
				PackageRef: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef{
					PackageMetadataName: testCertmanagertanzuvmwarecom,
					VersionSelection: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection{
						Constraints: test1102vmware1tkg1,
					},
				},
				RoleBindingScope: packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER.Pointer(),
				InlineValues: map[string]interface{}{
					testNamespace: testCertManager,
					testSome:      91,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					PackageRefKey: []interface{}{
						map[string]interface{}{
							PackageMetadataNameKey: testCertmanagertanzuvmwarecom,
							VersionSelectionKey: []interface{}{
								map[string]interface{}{
									ConstraintsKey: test1102vmware1tkg1,
								},
							},
						},
					},
					RoleBindingScopeKey:   fmt.Sprint(packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER),
					PathToInlineValuesKey: "test.yaml",
				},
			},
			filePath: "test.yaml",
		},
		{
			description: "normal scenario with complete cluster package install spec without inline values file",
			input: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec{
				PackageRef: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackageRef{
					PackageMetadataName: testCertmanagertanzuvmwarecom,
					VersionSelection: &packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageVersionSelection{
						Constraints: test1102vmware1tkg1,
					},
				},
				RoleBindingScope: packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER.Pointer(),
				InlineValues: map[string]interface{}{
					testNamespace: testCertManager,
					testSome:      91,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					PackageRefKey: []interface{}{
						map[string]interface{}{
							PackageMetadataNameKey: testCertmanagertanzuvmwarecom,
							VersionSelectionKey: []interface{}{
								map[string]interface{}{
									ConstraintsKey: test1102vmware1tkg1,
								},
							},
						},
					},
					RoleBindingScopeKey: fmt.Sprint(packageinstallmodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallRoleBindingScopeCLUSTER),
					InlineValuesKey: map[string]string{
						testNamespace: testCertManager,
						testSome:      "91",
					},
				},
			},
			filePath: "",
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual, _ := FlattenSpecForClusterScope(test.input, test.filePath)
			require.Equal(t, test.expected, actual)
		})
	}
}
