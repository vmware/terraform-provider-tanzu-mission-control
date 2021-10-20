/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgservicevsphere

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	nodepoolmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	tkgservicevspheremodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgservicevsphere"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var TkgServiceVsphere = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The tkg service vsphere cluster spec",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			settingsKey:     settings,
			distributionKey: distribution,
			topologyKey:     topology,
		},
	},
}

func ConstructTKGSSpec(data []interface{}) (spec *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec) {
	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})
	spec = &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec{}

	if v, ok := specData[settingsKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.Settings = expandTKGSSettings(v1)
		}
	}

	if v, ok := specData[distributionKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.Distribution = expandTKGSDistribution(v1)
		}
	}

	if v, ok := specData[topologyKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.Topology = expandTKGSTopology(v1)
		}
	}

	return spec
}

func FlattenTKGSSpec(spec *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec) (data []interface{}) {
	flattenSpec := make(map[string]interface{})

	flattenSpec[settingsKey] = flattenTKGSSettings(spec.Settings)
	flattenSpec[distributionKey] = flattenTKGSDistribution(spec.Distribution)
	flattenSpec[topologyKey] = flattenTKGSTopology(spec.Topology)

	return []interface{}{flattenSpec}
}

var settings = &schema.Schema{
	Type:        schema.TypeList,
	Description: "VSphere related settings for workload cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			networkKey: network,
		},
	},
}

var network = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Network Settings specifies network-related settings for the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			podsKey: {
				Type:        schema.TypeList,
				Description: "Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cidrBlocksKey: cidrBlocks,
					},
				},
			},
			servicesKey: {
				Type:        schema.TypeList,
				Description: "Service CIDR for kubernetes services defaults to 10.96.0.0/12",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cidrBlocksKey: cidrBlocks,
					},
				},
			},
		},
	},
}

func expandTKGSSettings(data []interface{}) (settings *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSettings) {
	if len(data) == 0 || data[0] == nil {
		return settings
	}

	settings = &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSettings{}
	settingData := data[0].(map[string]interface{})

	if v, ok := settingData[networkKey]; ok {
		networks, _ := v.([]interface{})

		if len(networks) == 0 || networks[0] == nil {
			return settings
		}

		networkData := networks[0].(map[string]interface{})
		settings.Network = &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkSettings{
			Pods:     &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkRanges{},
			Services: &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkRanges{},
		}

		if pods, ok := networkData[podsKey]; ok {
			podsData := pods.([]interface{})

			if len(podsData) != 0 || podsData[0] != nil {
				cidrBlockData, _ := podsData[0].(map[string]interface{})

				if cidrBlocks, ok := cidrBlockData[cidrBlocksKey]; ok {
					vs, _ := cidrBlocks.([]interface{})
					s := make([]string, 0)

					for _, raw := range vs {
						s = append(s, raw.(string))
					}

					settings.Network.Pods.CidrBlocks = s
				}
			}
		}

		if services, ok := networkData[servicesKey]; ok {
			servicesData, _ := services.([]interface{})

			if len(servicesData) != 0 || servicesData[0] != nil {
				cidrBlockData, _ := servicesData[0].(map[string]interface{})

				if cidrBlocks, ok := cidrBlockData[cidrBlocksKey]; ok {
					vs, _ := cidrBlocks.([]interface{})
					s := make([]string, 0)

					for _, raw := range vs {
						s = append(s, raw.(string))
					}

					settings.Network.Services.CidrBlocks = s
				}
			}
		}
	}

	return settings
}

func flattenTKGSSettings(settings *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSettings) (data []interface{}) {
	flattenSettings := make(map[string]interface{})

	if settings == nil {
		return nil
	}

	if settings.Network != nil {
		flattenSettingsNetwork := make(map[string]interface{})

		if settings.Network.Pods != nil {
			flattenSettingsNetworkPods := make(map[string]interface{})
			flattenSettingsNetworkPods[cidrBlocksKey] = settings.Network.Pods.CidrBlocks

			flattenSettingsNetwork[podsKey] = []interface{}{flattenSettingsNetworkPods}
		}

		if settings.Network.Services != nil {
			flattenSettingsNetworkServices := make(map[string]interface{})
			flattenSettingsNetworkServices[cidrBlocksKey] = settings.Network.Services.CidrBlocks

			flattenSettingsNetwork[servicesKey] = []interface{}{flattenSettingsNetworkServices}
		}

		flattenSettings[networkKey] = []interface{}{flattenSettingsNetwork}
	}

	return []interface{}{flattenSettings}
}

var distribution = &schema.Schema{
	Type:        schema.TypeList,
	Description: "VSphere specific distribution",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			versionKey: {
				Type:        schema.TypeString,
				Description: "Version of the cluster",
				Required:    true,
			},
		},
	},
}

func expandTKGSDistribution(data []interface{}) (distribution *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereDistribution) {
	if len(data) == 0 || data[0] == nil {
		return distribution
	}

	distribution = &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereDistribution{}
	distributionData, _ := data[0].(map[string]interface{})

	if v, ok := distributionData[versionKey]; ok {
		distribution.Version = v.(string)
	}

	return distribution
}

func flattenTKGSDistribution(distribution *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereDistribution) (data []interface{}) {
	flattenDistribution := make(map[string]interface{})

	if distribution == nil {
		return nil
	}

	flattenDistribution[versionKey] = distribution.Version

	return []interface{}{flattenDistribution}
}

var topology = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Topology specific configuration",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			controlPlaneKey: controlPlane,
			nodePoolsKey: {
				Type:        schema.TypeList,
				Description: "Nodepool specific configuration",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						specKey: nodePoolSpec,
						infoKey: info,
					},
				},
			},
		},
	},
}

func expandTKGSTopology(data []interface{}) (topology *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology) {
	if len(data) == 0 || data[0] == nil {
		return topology
	}

	topology = &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology{}
	topologyData, _ := data[0].(map[string]interface{})

	if v, ok := topologyData[controlPlaneKey]; ok {
		topology.ControlPlane = expandTKGSTopologyControlPlane(v.([]interface{}))
	}

	if v, ok := topologyData[nodePoolsKey]; ok {
		nodePools, _ := v.([]interface{})
		for _, np := range nodePools {
			topology.NodePools = append(topology.NodePools, expandTKGSTopologyNodePool(np))
		}
	}

	return topology
}

func flattenTKGSTopology(topology *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereTopology) (data []interface{}) {
	flattenTopology := make(map[string]interface{})

	if topology == nil {
		return nil
	}

	flattenTopology[controlPlaneKey] = flattenTKGSTopologyControlPlane(topology.ControlPlane)

	nodePools := make([]interface{}, 0)

	for _, nodePool := range topology.NodePools {
		nodePools = append(nodePools, FlattenTKGSTopologyNodePool(nodePool))
	}

	flattenTopology[nodePoolsKey] = nodePools

	return []interface{}{flattenTopology}
}

var cidrBlocks = &schema.Schema{
	Type:        schema.TypeList,
	Description: "CIDRBlocks specifies one or more ranges of IP addresses",
	Required:    true,
	Elem: &schema.Schema{
		Type:         schema.TypeString,
		ValidateFunc: validation.IsCIDR,
	},
}

var controlPlane = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Control plane specific configuration",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			classKey:        class,
			storageClassKey: storageClass,
			highAvailabilityKey: {
				Type:        schema.TypeBool,
				Description: "High Availability or Non High Availability Cluster. HA cluster creates three controlplane machines, and non HA creates just one",
				Default:     false,
				Optional:    true,
			},
		},
	},
}

func expandTKGSTopologyControlPlane(data []interface{}) (controlPlane *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane) {
	if len(data) == 0 || data[0] == nil {
		return controlPlane
	}

	controlPlaneData, _ := data[0].(map[string]interface{})
	controlPlane = &tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane{}

	if v, ok := controlPlaneData[classKey]; ok {
		controlPlane.Class, _ = v.(string)
	}

	if v, ok := controlPlaneData[storageClassKey]; ok {
		controlPlane.StorageClass, _ = v.(string)
	}

	if v, ok := controlPlaneData[highAvailabilityKey]; ok {
		controlPlane.HighAvailability, _ = v.(bool)
	}

	return controlPlane
}

func flattenTKGSTopologyControlPlane(controlPlane *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereControlPlane) (data []interface{}) {
	flattenControlPlane := make(map[string]interface{})

	if controlPlane == nil {
		return nil
	}

	flattenControlPlane[classKey] = controlPlane.Class
	flattenControlPlane[storageClassKey] = controlPlane.StorageClass
	flattenControlPlane[highAvailabilityKey] = controlPlane.HighAvailability

	return []interface{}{flattenControlPlane}
}

var nodePoolSpec = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the cluster nodepool",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			workerNodeCountKey: {
				Type:        schema.TypeString,
				Description: "Count is the number of nodes",
				Optional:    true,
			},
			nodeLabelKey: {
				Type:        schema.TypeMap,
				Description: "Node labels",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			cloudLabelKey: {
				Type:        schema.TypeMap,
				Description: "Cloud labels",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			tkgServiceVsphereKey: {
				Type:        schema.TypeList,
				Description: "Nodepool config for tkg service vsphere",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						classKey:        class,
						storageClassKey: storageClass,
					},
				},
			},
		},
	},
}

var info = &schema.Schema{
	Type:        schema.TypeList,
	Required:    true,
	Description: "Info is the meta information of nodepool for cluster",
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			clusterNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the nodepool",
				Default:     defaultNodePoolName,
				Optional:    true,
			},
			descriptionKey: {
				Type:        schema.TypeString,
				Description: "Description for the nodepool",
				Optional:    true,
			},
		},
	},
}

var class = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Control plane instance type",
	Required:    true,
}

var storageClass = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Storage Class to be used for storage of the disks which store the root filesystems of the nodes",
	Required:    true,
}

func expandTKGSTopologyNodePool(data interface{}) (nodePools *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition) {
	nodePoolData := data.(map[string]interface{})
	nodePools = &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
		Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{},
		Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{},
	}

	if v, ok := nodePoolData[specKey]; ok {
		specData, _ := v.([]interface{})

		if len(specData) != 0 || specData[0] != nil {
			spec := specData[0].(map[string]interface{})

			if v1, k1 := spec[workerNodeCountKey]; k1 {
				nodePools.Spec.WorkerNodeCount = v1.(string)
			}

			if v1, k1 := spec[nodeLabelKey]; k1 {
				nodePools.Spec.NodeLabels = common.GetTypeMapData(v1.(map[string]interface{}))
			}

			if v1, k1 := spec[cloudLabelKey]; k1 {
				nodePools.Spec.CloudLabels = common.GetTypeMapData(v1.(map[string]interface{}))
			}

			if v1, k1 := spec[tkgServiceVsphereKey]; k1 {
				nodePools.Spec.TkgServiceVsphere = expandNodePoolTKGSServiceVsphere(v1.([]interface{}))
			}
		}
	}

	if v1, k1 := nodePoolData[infoKey]; k1 {
		infoData, _ := v1.([]interface{})

		if len(infoData) != 0 || infoData[0] != nil {
			info := infoData[0].(map[string]interface{})

			if v2, k2 := info[clusterNameKey]; k2 {
				nodePools.Info.Name = v2.(string)
			}

			if v2, k2 := info[descriptionKey]; k2 {
				nodePools.Info.Description = v2.(string)
			}
		}
	}

	return nodePools
}

func expandNodePoolTKGSServiceVsphere(data []interface{}) (tkgsServiceVsphere *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool) {
	if len(data) == 0 || data[0] == nil {
		return tkgsServiceVsphere
	}

	tkgsServiceVsphereData, _ := data[0].(map[string]interface{})
	tkgsServiceVsphere = &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{}

	if v, ok := tkgsServiceVsphereData[classKey]; ok {
		tkgsServiceVsphere.Class, _ = v.(string)
	}

	if v, ok := tkgsServiceVsphereData[storageClassKey]; ok {
		tkgsServiceVsphere.StorageClass, _ = v.(string)
	}

	return tkgsServiceVsphere
}

func FlattenTKGSTopologyNodePool(nodePool *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition) (data interface{}) {
	flattenNodePool := make(map[string]interface{})

	if nodePool == nil {
		return nil
	}

	if nodePool.Info != nil {
		flattenNodePoolInfo := make(map[string]interface{})

		flattenNodePoolInfo[clusterNameKey] = nodePool.Info.Name
		flattenNodePoolInfo[descriptionKey] = nodePool.Info.Description

		flattenNodePool[infoKey] = []interface{}{flattenNodePoolInfo}
	}

	if nodePool.Spec != nil {
		flattenNodePoolSpec := make(map[string]interface{})

		flattenNodePoolSpec[workerNodeCountKey] = nodePool.Spec.WorkerNodeCount
		flattenNodePoolSpec[nodeLabelKey] = nodePool.Spec.NodeLabels
		flattenNodePoolSpec[cloudLabelKey] = nodePool.Spec.CloudLabels

		if nodePool.Spec.TkgServiceVsphere != nil {
			flattenNodePoolSpecTKGS := make(map[string]interface{})

			flattenNodePoolSpecTKGS[classKey] = nodePool.Spec.TkgServiceVsphere.Class
			flattenNodePoolSpecTKGS[storageClassKey] = nodePool.Spec.TkgServiceVsphere.StorageClass

			flattenNodePoolSpec[tkgServiceVsphereKey] = []interface{}{flattenNodePoolSpecTKGS}
		}

		flattenNodePool[specKey] = []interface{}{flattenNodePoolSpec}
	}

	return flattenNodePool
}
