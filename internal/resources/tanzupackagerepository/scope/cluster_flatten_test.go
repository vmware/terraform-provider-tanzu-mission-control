/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	packagerepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackagerepository"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenClusterPackageRepositoryFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *packagerepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster package repository full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster package repository full name",
			input: &packagerepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName{
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
			actual := FlattenClusterPackageRepositoryFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
