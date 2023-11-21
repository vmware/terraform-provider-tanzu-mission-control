/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package managementcluster

const (
	ResourceName = "tanzu-mission-control_management_cluster"

	NameKey  = "name"
	OrgIDKey = "org_id"

	specKey            = "spec"
	registerClusterKey = "register_management_cluster"

	registerClusterKubeConfigPathForTKGmKey = "tkgm_kubeconfig_file"
	registerClusterDescriptionForTKGmKey    = "tkgm_description"
	registerClusterKubeConfigRawForTKGmKey  = "tkgm_kubeconfig_raw"
	StatusKey                               = "status"
	waitKey                                 = "ready_wait_timeout"
	clusterGroupKey                         = "cluster_group"
	kubernetesProviderTypeKey               = "kubernetes_provider_type"
	imageRegistryKey                        = "image_registry"
	managedWorkloadClusterImageRegistryKey  = "managed_workload_cluster_image_registry"
	managementClusterProxyNameKey           = "management_proxy_name"
	managedWorkloadClusterProxyNameKey      = "managed_workload_cluster_proxy_name"
)
