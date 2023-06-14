/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	sourcesecretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenClusterSourcesecretFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster source secret full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster source secret full name",
			input: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName{
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
			actual := FlattenClusterSourcesecretFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
