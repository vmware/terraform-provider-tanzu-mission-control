// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/stretchr/testify/require"

	kustomizationclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func TestFlattenClusterGroupKustomizationFullname(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName
		expected    []interface{}
	}{
		{
			description: "check for nil cluster group kustomization full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group kustomization full name",
			input: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName{
				ClusterGroupName: "c",
			},
			expected: []interface{}{
				map[string]interface{}{
					commonscope.NameKey: "c",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenClusterGroupKustomizationFullname(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
