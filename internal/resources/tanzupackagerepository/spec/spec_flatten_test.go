/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"testing"

	"github.com/stretchr/testify/require"

	pkgrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackagerepository"
)

func TestFlattenSpecForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster package repository spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with image package repository spec",
			input: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec{
				ImgpkgBundle: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec{
					Image: "sometestimage",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					ImgpkgBundleKey: []interface{}{
						map[string]interface{}{
							ImageKey: "sometestimage",
						},
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
