/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package managementclusterregistration

import (
	"testing"

	"github.com/stretchr/testify/require"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	managementclusterregistrationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/managementclusterregistration"
)

func TestFlattenSpec(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterSpec
		expected    []interface{}
	}{
		{
			description: "check for nil data in management cluster registration spec",
		},
		{
			description: "normal scenario with management cluster registration",
			input: &managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterSpec{
				DefaultClusterGroup:                 "cluster_group_value",
				DefaultWorkloadClusterImageRegistry: "workload_cluster_image_registry_value",
				DefaultWorkloadClusterProxyName:     "workload_cluster_proxy_name_value",
				ImageRegistry:                       "image_registry_value",
				KubernetesProviderType:              clustermodel.NewVmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType("VMWARE_TANZU_KUBERNETES_GRID"),
				ProxyName:                           "proxy_name_value",
			},
			expected: []interface{}{
				map[string]interface{}{
					defaultClusterGroupKey:                 "cluster_group_value",
					defaultWorkloadClusterImageRegistryKey: "workload_cluster_image_registry_value",
					defaultWorkloadClusterProxyNameKey:     "workload_cluster_proxy_name_value",
					imageRegistryKey:                       "image_registry_value",
					kubernetesProviderTypeKey:              "VMWARE_TANZU_KUBERNETES_GRID",
					proxyNameKey:                           "proxy_name_value",
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenSpec(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
