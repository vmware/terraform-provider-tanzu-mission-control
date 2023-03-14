/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"testing"

	"github.com/stretchr/testify/require"

	kustomizationclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/cluster"
	kustomizationclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/clustergroup"
)

func TestFlattenSpecForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster kustomization spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster kustomization spec",
			input: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec{
				Interval: "5m",
				Path:     "/",
				Prune:    true,
				Source: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference{
					Name:      "tmc-cd",
					Namespace: "default",
				},
				TargetNamespace: "default",
			},
			expected: []interface{}{
				map[string]interface{}{
					intervalKey: "5m",
					pathKey:     "/",
					pruneKey:    true,
					sourceKey: []interface{}{
						map[string]interface{}{
							nameKey:      "tmc-cd",
							namespaceKey: "default",
						},
					},
					targetNamespaceKey: "default",
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

func TestFlattenSpecForClusterGroupScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster group kustomization spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group kustomization spec",
			input: &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec{
				AtomicSpec: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec{
					Interval: "5m",
					Path:     "/",
					Prune:    true,
					Source: &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationRepositoryReference{
						Name:      "tmc-cd",
						Namespace: "default",
					},
					TargetNamespace: "default",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					intervalKey: "5m",
					pathKey:     "/",
					pruneKey:    true,
					sourceKey: []interface{}{
						map[string]interface{}{
							nameKey:      "tmc-cd",
							namespaceKey: "default",
						},
					},
					targetNamespaceKey: "default",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenSpecForClusterGroupScope(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
