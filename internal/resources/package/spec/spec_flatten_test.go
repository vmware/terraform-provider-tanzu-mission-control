/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	packageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/package/cluster"
)

func TestFlattenSpecForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster source secret spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete package spec",
			input: &packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec{
				CapacityRequirementsDescription: "someCapacityRequirementsDescription",
				Licenses: []string{
					"some1",
				},
				ReleaseNotes:   "someReleaseNotes",
				ReleasedAt:     strfmt.DateTime{},
				RepositoryName: "testRepo",
				ValuesSchema: &packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageValuesSchema{
					Template: &packageclustermodel.K8sIoApimachineryPkgRuntimeRawExtension{
						Raw: []byte("somevalue"),
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					CapacityRequirementsDescriptionKey: "someCapacityRequirementsDescription",
					LicensesKey: []string{
						"some1",
					},
					RepositoryNameKey: "testRepo",
					ReleaseNotesKey:   "someReleaseNotes",
					ReleasedAtKey:     strfmt.DateTime{}.String(),
					ValuesSchemaKey: []interface{}{
						map[string]interface{}{
							TemplateKey: []interface{}{
								map[string]interface{}{
									RawKey: "somevalue",
								},
							},
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
