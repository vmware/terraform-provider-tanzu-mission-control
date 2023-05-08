/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetescluster

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	tanzukubernetescluster "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzu_kubernetes_cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var TkcTopology = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Tanzu Kubernetes cluster topology",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			clusterClassKey: {
				Type:        schema.TypeString,
				Description: "The name of the cluster class for the cluster",
				Optional:    true,
				Default:     "tanzukubernetescluster",
			},
			controlPlaneKey: controlPlane,
			coreAddonsKey: {
				Type:        schema.TypeList,
				Description: "The core addons",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						providerKey: {
							Type:        schema.TypeString,
							Description: "Provider of core addon, e.g. 'antrea', 'calico'",
							Optional:    true,
						},
						typeKey: {
							Type:        schema.TypeString,
							Description: "Type of core addon, e.g. 'cni'.",
							Optional:    true,
						},
					},
				},
			},
			networkKey: networkSettings,
			nodePoolsKey: {
				Type:        schema.TypeList,
				Description: "Nodepool definition for the cluster",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						nodePoolInfoKey: tkcNodePoolInfo,
						nodePoolSpecKey: tkcNodePoolSpec,
					},
				},
			},
			variablesKey: {
				Type:        schema.TypeList,
				Description: "Variables configuration for the cluster",
				Optional:    true,
				Elem:        variables,
			},
			kubernetesVersionKey: {
				Type:        schema.TypeString,
				Description: "Kubernetes version of the cluster.",
				Required:    true,
			},
		},
	},
}

func ConstructTKCTopology(data []interface{}) (topology *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology) {
	if len(data) == 0 || data[0] == nil {
		return topology
	}

	topologyData, _ := data[0].(map[string]interface{})
	topology = &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology{}

	if v, ok := topologyData[clusterClassKey]; ok {
		topology.ClusterClass, _ = v.(string)
	}

	if v, ok := topologyData[controlPlaneKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			topology.ControlPlane = expandTKCControlPlane(v1)
		}
	}

	if v, ok := topologyData[coreAddonsKey]; ok {
		coreAddons, _ := v.([]interface{})
		for _, ca := range coreAddons {
			topology.CoreAddons = append(topology.CoreAddons, expandTKCCoreAddons(ca))
		}
	}

	if v, ok := topologyData[networkKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			topology.Network = expandTKCNetwork(v1)
		}
	}

	if v, ok := topologyData[nodePoolsKey]; ok {
		nodepools, _ := v.([]interface{})
		for _, np := range nodepools {
			topology.NodePools = append(topology.NodePools, expandTKCNodePools(np))
		}
	}

	if v, ok := topologyData[variablesKey]; ok {
		variables, _ := v.([]interface{})
		for _, vs := range variables {
			topology.Variables = append(topology.Variables, expandTKCVariables(vs))
		}
	}

	if v, ok := topologyData[kubernetesVersionKey]; ok {
		topology.Version, _ = v.(string)
	}

	return topology
}

func FlattenTKCTopology(topology *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology) (data []interface{}) {
	flattenTopology := make(map[string]interface{})

	if topology == nil {
		return nil
	}

	flattenTopology[clusterClassKey] = topology.ClusterClass
	flattenTopology[controlPlaneKey] = flattenTKCControlPlane(topology.ControlPlane)
	cas := make([]interface{}, 0)

	for _, ca := range topology.CoreAddons {
		cas = append(cas, flattenTKCCoreAddons(ca))
	}

	flattenTopology[coreAddonsKey] = cas

	flattenTopology[networkKey] = flattenTKCNetwork(topology.Network)

	nps := make([]interface{}, 0)

	for _, np := range topology.NodePools {
		nps = append(nps, flattenTKCNodePools(np))
	}

	flattenTopology[nodePoolsKey] = nps

	vbs := make([]interface{}, 0)

	for _, vb := range topology.Variables {
		vbs = append(vbs, flattenTKCVariables(vb))
	}

	flattenTopology[variablesKey] = vbs

	flattenTopology[kubernetesVersionKey] = topology.Version

	return []interface{}{flattenTopology}
}

var controlPlane = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Control plane specific configuration",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			metaDataKey: {
				Type:        schema.TypeList,
				Description: "The metadata of the control plane",
				Computed:    true,
				Optional:    true,
				MaxItems:    1,
				Elem:        metaData,
			},
			osImageKey: {
				Type:        schema.TypeList,
				Description: "The OS image of the control plane",
				Optional:    true,
				MaxItems:    1,
				Elem:        osImage,
			},
			replicasKey: {
				Type:        schema.TypeInt,
				Description: "The replicas of the control plane",
				Optional:    true,
			},
		},
	},
}

func expandTKCControlPlane(data []interface{}) (controlPlane *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterControlPlane) {
	if len(data) == 0 || data[0] == nil {
		return controlPlane
	}

	lookUpControlPlane, _ := data[0].(map[string]interface{})
	controlPlane = &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterControlPlane{
		Metadata: &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata{},
		OsImage:  &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterOSImage{},
	}

	if v, ok := lookUpControlPlane[metaDataKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			controlPlane.Metadata = expandTKCMetadata(v1)
		}
	}

	if v, ok := lookUpControlPlane[osImageKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			controlPlane.OsImage = expandTKCOsImage(v1)
		}
	}

	if v, ok := lookUpControlPlane[replicasKey]; ok {
		replica := v.(int)
		controlPlane.Replicas = int32(replica)
	}

	return controlPlane
}

func flattenTKCControlPlane(controlPlane *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterControlPlane) (data []interface{}) {
	flattenControlPlane := make(map[string]interface{})

	if controlPlane == nil {
		return nil
	}

	if controlPlane.Metadata != nil {
		flattenControlPlane[metaDataKey] = flattenTKCMetadata(controlPlane.Metadata)
	}

	if controlPlane.OsImage != nil {
		flattenControlPlane[osImageKey] = flattenTKCOsImage(controlPlane.OsImage)
	}

	flattenControlPlane[replicasKey] = controlPlane.Replicas

	return []interface{}{flattenControlPlane}
}

var metaData = &schema.Resource{
	Schema: map[string]*schema.Schema{
		annotationsKey: {
			Type:        schema.TypeMap,
			Description: "The annotations configuration",
			Optional:    true,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		labelsKey: {
			Type:        schema.TypeMap,
			Description: "The labels configuration",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	},
}

func expandTKCMetadata(data []interface{}) (meta *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata) {
	meta = &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata{
		Annotations: make(map[string]string),
		Labels:      make(map[string]string),
	}

	if len(data) == 0 || data[0] == nil {
		return meta
	}

	metadata, _ := data[0].(map[string]interface{})

	if v, ok := metadata[annotationsKey]; ok {
		meta.Annotations = common.GetTypeStringMapData(v.(map[string]interface{}))
	}

	if v, ok := metadata[labelsKey]; ok {
		meta.Labels = common.GetTypeStringMapData(v.(map[string]interface{}))
	}

	return meta
}

func flattenTKCMetadata(meta *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata) (data []interface{}) {
	if meta == nil {
		return data
	}

	flattenMetadata := make(map[string]interface{})

	flattenMetadata[annotationsKey] = meta.Annotations
	flattenMetadata[labelsKey] = meta.Labels

	return []interface{}{flattenMetadata}
}

var osImage = &schema.Resource{
	Schema: map[string]*schema.Schema{
		archKey: {
			Type:        schema.TypeString,
			Description: "The arch of the OS image",
			Optional:    true,
			Default:     "",
		},
		nameKey: {
			Type:        schema.TypeString,
			Description: "The name of the OS image",
			Optional:    true,
			Default:     "",
		},
		versionKey: {
			Type:        schema.TypeString,
			Description: "The version of the OS image",
			Optional:    true,
			Default:     "",
		},
	},
}

func expandTKCOsImage(data []interface{}) (osImage *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterOSImage) {
	if len(data) == 0 || data[0] == nil {
		return osImage
	}

	osImageData, _ := data[0].(map[string]interface{})
	osImage = &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterOSImage{}

	if v, ok := osImageData[archKey]; ok {
		osImage.Arch, _ = v.(string)
	}

	if v, ok := osImageData[nameKey]; ok {
		osImage.Name, _ = v.(string)
	}

	if v, ok := osImageData[versionKey]; ok {
		osImage.Version, _ = v.(string)
	}

	return osImage
}

func flattenTKCOsImage(osImage *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterOSImage) (data []interface{}) {
	flattenOsImage := make(map[string]interface{})

	if osImage == nil {
		return nil
	}

	flattenOsImage[archKey] = osImage.Arch
	flattenOsImage[nameKey] = osImage.Name
	flattenOsImage[versionKey] = osImage.Version

	return []interface{}{flattenOsImage}
}

func expandTKCCoreAddons(data interface{}) (coreAddons *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCoreAddon) {
	lookUpAddons, ok := data.(map[string]interface{})
	if !ok {
		return coreAddons
	}

	coreAddons = &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCoreAddon{}

	if v, ok := lookUpAddons[providerKey]; ok {
		coreAddons.Provider, _ = v.(string)
	}

	if v, ok := lookUpAddons[typeKey]; ok {
		coreAddons.Type, _ = v.(string)
	}

	return coreAddons
}

func flattenTKCCoreAddons(coreAddons *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCoreAddon) (data interface{}) {
	flattenCoreAddon := make(map[string]interface{})

	if coreAddons == nil {
		return nil
	}

	flattenCoreAddon[providerKey] = coreAddons.Provider
	flattenCoreAddon[typeKey] = coreAddons.Type

	return flattenCoreAddon
}

var networkSettings = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Network specific configuration",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			serviceDomainKey: {
				Type:        schema.TypeString,
				Description: "Domain name for services",
				Required:    true,
			},
			podsKey: {
				Type:        schema.TypeList,
				Description: "Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cidrBlockKey: cidrBlock,
					},
				},
			},
			servicesKey: {
				Type:        schema.TypeList,
				Description: "Service CIDR for kubernetes services defaults to 10.96.0.0/12",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cidrBlockKey: cidrBlock,
					},
				},
			},
		},
	},
}

var cidrBlock = &schema.Schema{
	Type:        schema.TypeList,
	Description: "CIDRBlocks specifies one or more ranges of IP addresses",
	Required:    true,
	Elem: &schema.Schema{
		Type:         schema.TypeString,
		ValidateFunc: validation.IsCIDR,
	},
}

func expandTKCNetwork(data []interface{}) (network *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings) {
	if len(data) == 0 || data[0] == nil {
		return network
	}

	lookUpNetwork, _ := data[0].(map[string]interface{})
	network = &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings{
		Pods:     &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkRanges{},
		Services: &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkRanges{},
	}

	if v, ok := lookUpNetwork[serviceDomainKey]; ok {
		network.ServiceDomain, _ = v.(string)
	}

	if pods, ok := lookUpNetwork[podsKey]; ok {
		podsData, _ := pods.([]interface{})

		if len(podsData) != 0 || podsData[0] != nil {
			cidrBlockData, _ := podsData[0].(map[string]interface{})

			if cidrBlocks, ok := cidrBlockData[cidrBlockKey]; ok {
				vs, _ := cidrBlocks.([]interface{})

				s := make([]string, 0)

				for _, raw := range vs {
					s = append(s, raw.(string))
				}

				network.Pods.CidrBlocks = s
			}
		}
	}

	if services, ok := lookUpNetwork[servicesKey]; ok {
		servicesData, _ := services.([]interface{})

		if len(servicesData) != 0 || servicesData[0] != nil {
			cidrBlockData, _ := servicesData[0].(map[string]interface{})

			if cidrBlocks, ok := cidrBlockData[cidrBlockKey]; ok {
				vs, _ := cidrBlocks.([]interface{})

				s := make([]string, 0)

				for _, raw := range vs {
					s = append(s, raw.(string))
				}

				network.Services.CidrBlocks = s
			}
		}
	}

	return network
}

func flattenTKCNetwork(network *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings) (data []interface{}) {
	flattenNetworkSettings := make(map[string]interface{})

	if network == nil {
		return nil
	}

	flattenNetworkSettings[serviceDomainKey] = network.ServiceDomain

	if network.Pods != nil {
		flattenNetworkPods := make(map[string]interface{})
		flattenNetworkPods[cidrBlockKey] = network.Pods.CidrBlocks

		flattenNetworkSettings[podsKey] = []interface{}{flattenNetworkPods}
	}

	if network.Services != nil {
		flattenNetworkServices := make(map[string]interface{})
		flattenNetworkServices[cidrBlockKey] = network.Services.CidrBlocks

		flattenNetworkSettings[servicesKey] = []interface{}{flattenNetworkServices}
	}

	return []interface{}{flattenNetworkSettings}
}

var tkcNodePoolInfo = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Info for the nodepool",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			nodePoolNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the nodepool",
				Required:    true,
			},
			nodePoolDescriptionKey: {
				Type:        schema.TypeString,
				Description: "Description for the nodepool",
				Optional:    true,
			},
		},
	},
}

var tkcNodePoolSpec = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the nodepool",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			classKey: {
				Type:        schema.TypeString,
				Description: "The name of the machine deployment class used to create the nodepool",
				Required:    true,
			},
			failureDomainKey: {
				Type:        schema.TypeString,
				Description: "The failure domain the machines will be created in",
				Optional:    true,
			},
			metaDataKey: {
				Type:        schema.TypeList,
				Description: "The metadata of the nodepool",
				Computed:    true,
				Optional:    true,
				MaxItems:    1,
				Elem:        metaData,
			},
			osImageKey: {
				Type:        schema.TypeList,
				Description: "The OS image of the nodepool",
				Optional:    true,
				MaxItems:    1,
				Elem:        osImage,
			},
			overridesKey: {
				Type:        schema.TypeList,
				Description: "Overrides can be used to override cluster level variables",
				Optional:    true,
				Elem:        variables,
			},
			replicasKey: {
				Type:        schema.TypeInt,
				Description: "The replicas of the nodepool",
				Optional:    true,
			},
		},
	},
}

var variables = &schema.Resource{
	Schema: map[string]*schema.Schema{
		variableNameKey: {
			Type:        schema.TypeString,
			Description: "Name of the variable",
			Required:    true,
		},
		valueKey: {
			Type:        schema.TypeString,
			Description: "Value of the variable",
			Optional:    true,
		},
	},
}

func expandTKCNodePools(data interface{}) (nodePools *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition) {
	lookUpNodepool, ok := data.(map[string]interface{})
	if !ok {
		return nodePools
	}

	nodePools = &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition{
		Spec: &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec{},
		Info: &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolInfo{},
	}

	// nolint: nestif
	if v, ok := lookUpNodepool[nodePoolSpecKey]; ok {
		specData, _ := v.([]interface{})

		if len(specData) != 0 || specData[0] != nil {
			spec, _ := specData[0].(map[string]interface{})

			if v1, ok := spec[classKey]; ok {
				nodePools.Spec.Class, _ = v1.(string)
			}

			if v1, ok := spec[failureDomainKey]; ok {
				nodePools.Spec.FailureDomain, _ = v1.(string)
			}

			if v1, ok := spec[metaDataKey]; ok {
				if v2, ok := v1.([]interface{}); ok {
					nodePools.Spec.Metadata = expandTKCMetadata(v2)
				}
			}

			if v1, ok := spec[osImageKey]; ok {
				if v2, ok := v1.([]interface{}); ok {
					nodePools.Spec.OsImage = expandTKCOsImage(v2)
				}
			}

			if v1, ok := spec[overridesKey]; ok {
				overrides, _ := v1.([]interface{})
				for _, ov := range overrides {
					nodePools.Spec.Overrides = append(nodePools.Spec.Overrides, expandTKCVariables(ov))
				}
			}

			if v1, ok := spec[replicasKey]; ok {
				replica := v1.(int)
				nodePools.Spec.Replicas = int32(replica)
			}
		}
	}

	if v, ok := lookUpNodepool[nodePoolInfoKey]; ok {
		infoData, _ := v.([]interface{})

		if len(infoData) != 0 || infoData[0] != nil {
			info, _ := infoData[0].(map[string]interface{})

			if v1, ok := info[nodePoolNameKey]; ok {
				nodePools.Info.Name = v1.(string)
			}

			if v1, ok := info[nodePoolDescriptionKey]; ok {
				nodePools.Info.Description = v1.(string)
			}
		}
	}

	return nodePools
}

func flattenTKCNodePools(nodePools *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition) (data interface{}) {
	flattenNodePool := make(map[string]interface{})

	if nodePools == nil {
		return nil
	}

	if nodePools.Info != nil {
		flattenNodePoolInfo := make(map[string]interface{})

		flattenNodePoolInfo[nodePoolNameKey] = nodePools.Info.Name
		flattenNodePoolInfo[nodePoolDescriptionKey] = nodePools.Info.Description

		flattenNodePool[nodePoolInfoKey] = []interface{}{flattenNodePoolInfo}
	}

	if nodePools.Spec != nil {
		flattenNodePoolSpec := make(map[string]interface{})

		flattenNodePoolSpec[classKey] = nodePools.Spec.Class
		flattenNodePoolSpec[failureDomainKey] = nodePools.Spec.FailureDomain

		if nodePools.Spec.Metadata != nil {
			flattenNodePoolSpec[metaDataKey] = flattenTKCMetadata(nodePools.Spec.Metadata)
		}

		if nodePools.Spec.OsImage != nil {
			flattenNodePoolSpec[osImageKey] = flattenTKCOsImage(nodePools.Spec.OsImage)
		}

		ovs := make([]interface{}, 0)

		for _, ov := range nodePools.Spec.Overrides {
			ovs = append(ovs, flattenTKCVariables(ov))
		}

		flattenNodePoolSpec[overridesKey] = ovs

		flattenNodePoolSpec[replicasKey] = nodePools.Spec.Replicas

		flattenNodePool[nodePoolSpecKey] = []interface{}{flattenNodePoolSpec}
	}

	return flattenNodePool
}

func expandTKCVariables(data interface{}) (variables *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable) {
	lookUpVariables, ok := data.(map[string]interface{})
	if !ok {
		return variables
	}

	variables = &tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable{}

	if v, ok := lookUpVariables[variableNameKey]; ok {
		variables.Name, _ = v.(string)
	}

	if v, ok := lookUpVariables[valueKey]; ok {
		variables.Value, _ = v.(string)
	}

	return variables
}

func flattenTKCVariables(variables *tanzukubernetescluster.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable) (data interface{}) {
	flattenVariables := make(map[string]interface{})

	if variables == nil {
		return nil
	}

	flattenVariables[variableNameKey] = variables.Name
	flattenVariables[valueKey] = variables.Value

	return flattenVariables
}
