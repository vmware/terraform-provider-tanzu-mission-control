// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgservicevsphere

import (
	"testing"

	"github.com/stretchr/testify/require"

	tkgservicevspheremodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgservicevsphere"
)

func TestFlattenDistribution(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereDistribution
		expected []interface{}
	}{
		{
			name:     "check for nil distribution data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with all fields of distribution data",
			input: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereDistribution{
				Version:   "v1.20",
				OsArch:    "amd64",
				OsName:    "photon",
				OsVersion: "3",
			},
			expected: []interface{}{
				map[string]interface{}{
					versionKey:   "v1.20",
					osArchKey:    "amd64",
					osNameKey:    "photon",
					osVersionKey: "3",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := flattenTKGSDistribution(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
