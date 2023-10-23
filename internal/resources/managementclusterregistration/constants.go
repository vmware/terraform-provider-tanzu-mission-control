/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package managementclusterregistration

const (
	NameKey = "name"

	specKey          = "spec"
	attachClusterKey = "attach_management_cluster"

	attachClusterKubeConfigPathKey         = "kubeconfig_file"
	attachClusterDescriptionKey            = "description"
	attachClusterKubeConfigRawKey          = "kubeconfig_raw"
	StatusKey                              = "status"
	waitKey                                = "ready_wait_timeout"
	defaultClusterGroupKey                 = "default_cluster_group"
	kubernetesProviderTypeKey              = "kubernetes_provider_type"
	imageRegistryKey                       = "image_registry"
	defaultWorkloadClusterImageRegistryKey = "default_workload_cluster_image_registry"
	proxyNameKey                           = "proxy_name"
	defaultWorkloadClusterProxyNameKey     = "default_workload_cluster_proxy_name"
)
