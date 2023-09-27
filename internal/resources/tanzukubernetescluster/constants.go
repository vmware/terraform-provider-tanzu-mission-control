/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkc

const (
	ResourceName = "tanzu-mission-control_tanzu_kubernetes_cluster"

	ManagementClusterNameKey = "management_cluster_name"
	ProvisionerNameKey       = "provisioner_name"
	NameKey                  = "name"
	specKey                  = "spec"
	StatusKey                = "status"
	waitKey                  = "ready_wait_timeout"
	clusterGroupKey          = "cluster_group_name"
	proxyNameKey             = "proxy_name"
	imageRegistryKey         = "image_registry"
	tmcManagedKey            = "tmc_managed"
	topologyKey              = "topology"
	clusterClassKey          = "cluster_class"
	controlPlaneKey          = "control_plane"
	nodepoolsKey             = "node_pools"
	clusterVariablesKey      = "cluster_variables"
	networkKey               = "network"
	coreAddonsKey            = "core_addons"
	replicasKey              = "replicas"
	metaKey                  = "meta"
	osImageKey               = "os_image"
	nameKey                  = "name"
	versionKey               = "version"
	archKey                  = "arch"
	metadataKey              = "meta"
	labelsKey                = "labels"
	annotationsKey           = "annotations"
	variablesKey             = "variables"
	tkgVsphereV100Key        = "tkg_vsphere_v100"
	vsphereTanzuV100Key      = "vsphere_tanzu_v100"
	podCidrBlocksKey         = "pod_cidr_blocks"
	serviceCidrBlocksKey     = "service_cidr_blocks"
	serviceDomainKey         = "service_domain"
	typeKey                  = "type"
	providerKey              = "provider"
	clusterGroupDefaultValue = "default"
	overridesKey             = "overrides"
	valueKey                 = "value"

	classKey         = "class"
	failureDomainKey = "failure_domain"
	tkgVsphereKey    = "tkg_vsphere"
	vsphereTanzuKey  = "vsphere_tanzu"

	readyCondition = "Ready"
	errorSeverity  = "ERROR"
)
