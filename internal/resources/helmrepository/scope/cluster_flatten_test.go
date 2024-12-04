// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrepository"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenClusterHelmFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *helmclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster helm repository full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster helm repository full name",
			input: &helmclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryFullName{
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
			actual := FlattenClusterHelmFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
