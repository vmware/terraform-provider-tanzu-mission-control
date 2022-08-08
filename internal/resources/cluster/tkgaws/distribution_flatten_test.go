/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgaws

import (
	"testing"

	"github.com/stretchr/testify/require"

	tkgawsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgaws"
)

func TestFlattenDistribution(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsDistribution
		expected []interface{}
	}{
		{
			name:     "check for nil distribution data",
			input:    nil,
			expected: nil,
		},
		{
			name: "normal scenario with all fields of distribution data",
			input: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsDistribution{
				Region:                    "us-west-2",
				Version:                   "v1.21.2+vmware.1-tkg.2",
				ProvisionerCredentialName: "default",
			},
			expected: []interface{}{
				map[string]interface{}{
					regionKey:                "us-west-2",
					versionKey:               "v1.21.2+vmware.1-tkg.2",
					provisionerCredentialKey: "default",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.name, func(t *testing.T) {
			actual := flattenTKGAWSDistribution(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
