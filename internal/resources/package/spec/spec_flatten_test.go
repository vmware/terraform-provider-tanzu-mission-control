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
			description: "check for nil cluster package spec",
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
				ReleaseNotes:   "cert-manager 1.1.0 https://github.com/jetstack/cert-manager/1.1.0",
				ReleasedAt:     strfmt.DateTime{},
				RepositoryName: "testRepo",
				ValuesSchema: &packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageValuesSchema{
					Template: &packageclustermodel.K8sIoApimachineryPkgRuntimeRawExtension{
						Raw: []byte("{\"examples\":[{\"namespace\":\"cert-manager\"}],\"properties\":{\"namespace\":{\"default\":\"cert-manager\",\"description\":\"The namespace in which to deploy cert-manager.\",\"type\":\"string\"}},\"title\":\"cert-manager.tanzu.vmware.com.1.1.0+vmware.1-tkg.2 values schema\"}"),
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
					ReleaseNotesKey: []interface{}{
						map[string]interface{}{
							metadataNameKey: "cert-manager",
							versionKey:      "1.1.0",
							urlKey:          "https://github.com/jetstack/cert-manager/1.1.0",
						},
					},
					ReleasedAtKey: strfmt.DateTime{}.String(),
					ValuesSchemaKey: []interface{}{
						map[string]interface{}{
							TemplateKey: []interface{}{
								map[string]interface{}{
									RawKey: []interface{}{
										map[string]interface{}{
											examplesKey: []interface{}{
												map[string]interface{}{
													namespaceKey: "cert-manager",
												},
											},
											propertiesKey: []interface{}{
												map[string]interface{}{
													namespaceKey: []interface{}{
														map[string]interface{}{
															defaultKey:     "cert-manager",
															descriptionKey: "The namespace in which to deploy cert-manager.",
															typeKey:        "string",
														},
													},
												},
											},
											titleKey: "cert-manager.tanzu.vmware.com.1.1.0+vmware.1-tkg.2 values schema",
										},
									},
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
