/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	releaseclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenClusterHelmReleaseFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster git repository full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster git repository full name",
			input: &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseFullName{
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
			actual := FlattenClusterReleaseFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
