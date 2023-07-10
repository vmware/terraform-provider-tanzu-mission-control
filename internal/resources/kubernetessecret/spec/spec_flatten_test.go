/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	secretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
)

func TestFlattenSpec(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster policy full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with all values under spec",
			input: &secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec{
				SecretType: secretclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEDOCKERCONFIGJSON),
				Data: map[string]strfmt.Base64{
					".dockerconfigjson": []byte(`{"auths":{"someurl":{"auth":"","password":"","username":"someuname"}}}`),
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					DockerConfigjsonKey: []interface{}{
						map[string]interface{}{
							UsernameKey:         "someuname",
							ImageRegistryURLKey: "someurl",
							PasswordKey:         "somepassword",
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenSpec(test.input, "somepassword")
			require.Equal(t, test.expected, actual)
		})
	}
}
