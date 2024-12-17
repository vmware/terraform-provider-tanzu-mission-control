// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	packageinstallclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackageinstall"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenClusterPackageInstallFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *packageinstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster package install full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster package install full name",
			input: &packageinstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName{
				ClusterName:           "c",
				ManagementClusterName: "m",
				ProvisionerName:       "p",
			},
			expected: []interface{}{
				map[string]interface{}{
					commonscope.NameKey:                  "c",
					commonscope.ManagementClusterNameKey: "m",
					commonscope.ProvisionerNameKey:       "p",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenClusterPackageInstallFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
