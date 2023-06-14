/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	sourcesecretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/cluster"
	sourcesecretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/clustergroup"
)

func TestFlattenSpecForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster source secret spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster source secret username/password type spec",
			input: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
				SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeUSERNAMEPASSWORD),
				Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
					Data: map[string]strfmt.Base64{
						"username": []byte("someusername"),
						"password": []byte(""),
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					DataKey: []interface{}{
						map[string]interface{}{
							UsernamePasswordKey: []interface{}{
								map[string]interface{}{
									usernameKey: "someusername",
									PasswordKey: "somevalue",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "normal scenario with complete cluster source secret ssh type spec",
			input: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
				SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeSSH),
				Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
					Data: map[string]strfmt.Base64{
						IdentityKey:   []byte("someidentity"),
						KnownhostsKey: []byte(""),
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					DataKey: []interface{}{
						map[string]interface{}{
							SSHKey: []interface{}{
								map[string]interface{}{
									IdentityKey:   "someidentity",
									KnownhostsKey: "somevalue",
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
			actual := FlattenSpecForClusterScope(test.input, "somevalue")
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenStatusForClusterGroupScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster source secret spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group source secret username/password type spec",
			input: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec{
				AtomicSpec: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
					SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeUSERNAMEPASSWORD),
					Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
						Data: map[string]strfmt.Base64{
							"username": []byte("someusername"),
							"password": []byte(""),
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					DataKey: []interface{}{
						map[string]interface{}{
							UsernamePasswordKey: []interface{}{
								map[string]interface{}{
									usernameKey: "someusername",
									PasswordKey: "somevalue",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "normal scenario with complete cluster group source secret ssh type spec",
			input: &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec{
				AtomicSpec: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{
					SourceSecretType: sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeSSH),
					Data: &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
						Data: map[string]strfmt.Base64{
							IdentityKey:   []byte("someidentity"),
							KnownhostsKey: []byte(""),
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					DataKey: []interface{}{
						map[string]interface{}{
							SSHKey: []interface{}{
								map[string]interface{}{
									IdentityKey:   "someidentity",
									KnownhostsKey: "somevalue",
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
			actual := FlattenSpecForClusterGroupScope(test.input, "somevalue")
			require.Equal(t, test.expected, actual)
		})
	}
}
