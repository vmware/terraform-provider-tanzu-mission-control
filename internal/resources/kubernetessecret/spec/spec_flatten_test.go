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
	secretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup"
)

func TestFlattenClusterScopeSpec(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster secret full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "registry secret test",
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
		{
			description: "opaque secret test",
			input: &secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec{
				SecretType: secretclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEOPAQUE),
				Data: map[string]strfmt.Base64{
					"username": []byte(`myuser`),
					"password": []byte(`somelongpassword`),
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					OpaqueKey: map[string]interface{}{
						"username": "myuser",
						"password": "somelongpassword",
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenSpecForClusterScope(test.input, "somepassword", map[string]interface{}{"username": "myuser", "password": "somelongpassword"})
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenClusterGroupScopeSpec(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster secret full name",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario for registry secret with all values under spec",
			input: &secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec{
				AtomicSpec: &secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec{
					SecretType: secretclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEDOCKERCONFIGJSON),
					Data: map[string]strfmt.Base64{
						".dockerconfigjson": []byte(`{"auths":{"someurl":{"auth":"","password":"","username":"someuname"}}}`),
					},
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
		{
			description: "normal scenario for opaque secret with all values under spec",
			input: &secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec{
				AtomicSpec: &secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec{
					SecretType: secretclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEOPAQUE),
					Data: map[string]strfmt.Base64{
						"username": []byte(`myuser`),
						"password": []byte(`somelongpassword`),
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					OpaqueKey: map[string]interface{}{
						"username": "myuser",
						"password": "somelongpassword",
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenSpecForClusterGroupScope(test.input, "somepassword", map[string]interface{}{"username": "myuser", "password": "somelongpassword"})
			require.Equal(t, test.expected, actual)
		})
	}
}
