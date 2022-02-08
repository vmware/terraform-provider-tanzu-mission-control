/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgvsphere

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	nodepoolmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	tkgvspheremodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgvsphere"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var TkgVsphereClusterSpec = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The Tanzu Kubernetes Grid (TKGm) VSphere cluster spec",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			settingsKey:     tkgVsphereSettings,
			distributionKey: tkgVsphereDistribution,
			topologyKey:     tkgVsphereTopology,
		},
	},
}

func ConstructTKGVsphereClusterSpec(data []interface{}) (spec *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec) {
	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData, _ := data[0].(map[string]interface{})
	spec = &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec{}

	if v, ok := specData[settingsKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.Settings = expandTKGVsphereSettings(v1)
		}
	}

	if v, ok := specData[distributionKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.Distribution = expandTKGVsphereDistribution(v1)
		}
	}

	if v, ok := specData[topologyKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.Topology = expandTKGVsphereTopology(v1)
		}
	}

	return spec
}

func FlattenTKGVsphereClusterSpec(spec *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec) (data []interface{}) {
	flattenSpecData := make(map[string]interface{})

	flattenSpecData[settingsKey] = flattenTKGVsphereSettings(spec.Settings)
	flattenSpecData[distributionKey] = flattenTKGVsphereDistribution(spec.Distribution)
	flattenSpecData[topologyKey] = flattenTKGVsphereTopology(spec.Topology)

	return []interface{}{flattenSpecData}
}

var tkgVsphereSettings = &schema.Schema{
	Type:        schema.TypeList,
	Description: "VSphere related settings for workload cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			networkKey:  tkgVsphereNetwork,
			securityKey: tkgVsphereSecurity,
		},
	},
}

var tkgVsphereSecurity = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Security Settings specifies security-related settings for the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			sshKey: {
				Type:        schema.TypeString,
				Description: "SSH key for provisioning and accessing the cluster VMs",
				Required:    true,
			},
		},
	},
}

func expandTKGVsphereSettings(data []interface{}) (settings *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSettings) {
	if len(data) == 0 || data[0] == nil {
		return settings
	}

	lookUpSettings, _ := data[0].(map[string]interface{})
	settings = &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSettings{}

	if v, ok := lookUpSettings[networkKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			settings.Network = expandTKGVsphereNetworkSettings(v1)
		}
	}

	if v, ok := lookUpSettings[securityKey]; ok {
		security, _ := v.([]interface{})
		if len(security) == 0 || security[0] == nil {
			return settings
		}

		lookUpSecurity, _ := security[0].(map[string]interface{})
		settings.Security = &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSecuritySettings{}

		if sshKey, ok := lookUpSecurity[sshKey]; ok {
			settings.Security.SSHKey, _ = sshKey.(string)
		}
	}

	return settings
}

func flattenTKGVsphereSettings(settings *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSettings) (data []interface{}) {
	flattenSettings := make(map[string]interface{})

	if settings == nil {
		return nil
	}

	if settings.Network != nil {
		flattenSettings[networkKey] = flattenTKGVsphereNetworkSettings(settings.Network)
	}

	if settings.Security != nil {
		flattenSettingsSecurity := make(map[string]interface{})
		flattenSettingsSecurity[sshKey] = settings.Security.SSHKey

		flattenSettings[securityKey] = []interface{}{flattenSettingsSecurity}
	}

	return []interface{}{flattenSettings}
}

var tkgVsphereNetwork = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Network Settings specifies network-related settings for the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			controlPlaneEndPointKey: {
				Type:        schema.TypeString,
				Description: "ControlPlaneEndpoint specifies the control plane virtual IP address. The value should be unique for every create request, else cluster creation shall fail",
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

func expandTKGVsphereNetworkSettings(data []interface{}) (network *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkSettings) {
	if len(data) == 0 || data[0] == nil {
		return network
	}

	lookUpNetwork, _ := data[0].(map[string]interface{})
	network = &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkSettings{
		ControlPlaneEndpoint: controlPlaneEndpointDefaultValue,
		Pods:                 &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkRanges{},
		Services:             &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkRanges{},
	}

	if v, ok := lookUpNetwork[controlPlaneEndPointKey]; ok {
		network.ControlPlaneEndpoint, _ = v.(string)
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

func flattenTKGVsphereNetworkSettings(network *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkSettings) (data []interface{}) {
	flattenNetworkSettings := make(map[string]interface{})

	if network == nil {
		return nil
	}

	flattenNetworkSettings[controlPlaneEndPointKey] = network.ControlPlaneEndpoint

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

var tkgVsphereDistribution = &schema.Schema{
	Type:        schema.TypeList,
	Description: "VSphere specific distribution",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			versionKey: {
				Type:        schema.TypeString,
				Description: "Version specifies the version of the Kubernetes cluster",
				Required:    true,
			},
			workspaceKey: {
				Type:        schema.TypeList,
				Description: "Workspace defines a workspace configuration for the vSphere cloud provider",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						datacenterKey: {
							Type:     schema.TypeString,
							Required: true,
						},
						datastoreKey: {
							Type:     schema.TypeString,
							Required: true,
						},
						folderKey: {
							Type:     schema.TypeString,
							Required: true,
						},
						workspaceNetworkKey: {
							Type:     schema.TypeString,
							Required: true,
						},
						resourcePoolKey: {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	},
}

func expandTKGVsphereDistribution(data []interface{}) (distribution *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereDistribution) {
	if len(data) == 0 || data[0] == nil {
		return distribution
	}

	lookUpDistribution, _ := data[0].(map[string]interface{})
	distribution = &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereDistribution{}

	if v, ok := lookUpDistribution[versionKey]; ok {
		distribution.Version, _ = v.(string)
	}

	if v, ok := lookUpDistribution[workspaceKey]; ok {
		workspace, _ := v.([]interface{})
		if len(workspace) == 0 || workspace[0] == nil {
			return distribution
		}

		lookUpWorkspace, _ := workspace[0].(map[string]interface{})
		distribution.Workspace = &tkgvspheremodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereWorkspace{}

		if v1, ok := lookUpWorkspace[datacenterKey]; ok {
			distribution.Workspace.Datacenter, _ = v1.(string)
		}

		if v1, ok := lookUpWorkspace[datastoreKey]; ok {
			distribution.Workspace.Datastore, _ = v1.(string)
		}

		if v1, ok := lookUpWorkspace[folderKey]; ok {
			distribution.Workspace.Folder, _ = v1.(string)
		}

		if v1, ok := lookUpWorkspace[workspaceNetworkKey]; ok {
			distribution.Workspace.Network, _ = v1.(string)
		}

		if v1, ok := lookUpWorkspace[resourcePoolKey]; ok {
			distribution.Workspace.ResourcePool, _ = v1.(string)
		}
	}

	return distribution
}

func flattenTKGVsphereDistribution(distribution *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereDistribution) (data []interface{}) {
	flattenDistribution := make(map[string]interface{})

	if distribution == nil {
		return nil
	}

	flattenDistribution[versionKey] = distribution.Version

	if distribution.Workspace != nil {
		flattenDistributionWorkspace := make(map[string]interface{})

		flattenDistributionWorkspace[datacenterKey] = distribution.Workspace.Datacenter
		flattenDistributionWorkspace[datastoreKey] = distribution.Workspace.Datastore
		flattenDistributionWorkspace[folderKey] = distribution.Workspace.Folder
		flattenDistributionWorkspace[workspaceNetworkKey] = distribution.Workspace.Network
		flattenDistributionWorkspace[resourcePoolKey] = distribution.Workspace.ResourcePool

		flattenDistribution[workspaceKey] = []interface{}{flattenDistributionWorkspace}
	}

	return []interface{}{flattenDistribution}
}

var tkgVsphereTopology = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Topology specific configuration",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			controlPlaneKey: tkgVsphereTopologyControlPlane,
			nodePoolsKey: {
				Type:        schema.TypeList,
				Description: "Nodepool specific configuration",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						nodePoolInfoKey: tkgVsphereNodePoolInfo,
						nodePoolSpecKey: tkgVsphereNodePoolSpec,
					},
				},
			},
		},
	},
}

func expandTKGVsphereTopology(data []interface{}) (topology *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology) {
	if len(data) == 0 || data[0] == nil {
		return topology
	}

	lookUpTopology, _ := data[0].(map[string]interface{})
	topology = &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology{}

	if v, ok := lookUpTopology[controlPlaneKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			topology.ControlPlane = expandTKGVsphereTopologyControlPlane(v1)
		}
	}

	if v, ok := lookUpTopology[nodePoolsKey]; ok {
		nodepools, _ := v.([]interface{})
		for _, np := range nodepools {
			topology.NodePools = append(topology.NodePools, expandTKGVsphereTopologyNodePool(np))
		}
	}

	return topology
}

func flattenTKGVsphereTopology(topology *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereTopology) (data []interface{}) {
	flattenTopology := make(map[string]interface{})

	if topology == nil {
		return nil
	}

	flattenTopology[controlPlaneKey] = flattenTKGVsphereTopologyControlPlane(topology.ControlPlane)

	nps := make([]interface{}, 0)

	for _, np := range topology.NodePools {
		nps = append(nps, flattenTKGVsphereTopologyNodePool(np))
	}

	flattenTopology[nodePoolsKey] = nps

	return []interface{}{flattenTopology}
}

var tkgVsphereTopologyControlPlane = &schema.Schema{
	Type:        schema.TypeList,
	Description: "VSphere specific control plane configuration for workload cluster object",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			vmConfigKey: tkgVsphereVMConfig,
			highAvailabilityKey: {
				Type:        schema.TypeBool,
				Description: "High Availability or Non High Availability Cluster. HA cluster creates three controlplane machines, and non HA creates just one",
				Default:     false,
				Optional:    true,
			},
		},
	},
}

var tkgVsphereVMConfig = &schema.Schema{
	Type:        schema.TypeList,
	Description: "VM specific configuration",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			cpuKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			diskKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			memoryKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	},
}

func expandTKGVsphereTopologyControlPlane(data []interface{}) (controlPlane *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereControlPlane) {
	if len(data) == 0 || data[0] == nil {
		return controlPlane
	}

	lookUpControlPlane, _ := data[0].(map[string]interface{})
	controlPlane = &tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereControlPlane{
		VMConfig: &nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig{},
	}

	if v, ok := lookUpControlPlane[vmConfigKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			controlPlane.VMConfig = expandTKGVsphereVMConfig(v1)
		}
	}

	if v, ok := lookUpControlPlane[highAvailabilityKey]; ok {
		controlPlane.HighAvailability, _ = v.(bool)
	}

	return controlPlane
}

func flattenTKGVsphereTopologyControlPlane(controlPlane *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereControlPlane) (data []interface{}) {
	flattenControlPlane := make(map[string]interface{})

	if controlPlane == nil {
		return nil
	}

	if controlPlane.VMConfig != nil {
		flattenControlPlane[vmConfigKey] = flattenTKGVsphereVMConfig(controlPlane.VMConfig)
	}

	flattenControlPlane[highAvailabilityKey] = controlPlane.HighAvailability

	return []interface{}{flattenControlPlane}
}

var tkgVsphereNodePoolInfo = &schema.Schema{
	Type:     schema.TypeList,
	Required: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			nodePoolNameKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			nodePoolDescriptionKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	},
}

var tkgVsphereNodePoolSpec = &schema.Schema{
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
			tkgVsphereKey: {
				Type:        schema.TypeList,
				Description: "Nodepool config for tkgm vsphere",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						vmConfigKey: tkgVsphereVMConfig,
					},
				},
			},
		},
	},
}

func expandTKGVsphereTopologyNodePool(data interface{}) (nodePools *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition) {
	lookUpNodepool := data.(map[string]interface{})
	nodePools = &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition{
		Spec: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{},
		Info: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolInfo{},
	}

	if v, ok := lookUpNodepool[nodePoolSpecKey]; ok {
		specData, _ := v.([]interface{})

		if len(specData) != 0 || specData[0] != nil {
			spec, _ := specData[0].(map[string]interface{})

			if v1, ok := spec[workerNodeCountKey]; ok {
				nodePools.Spec.WorkerNodeCount, _ = v1.(string)
			}

			if v1, ok := spec[nodeLabelKey]; ok {
				nodeLabels, _ := v1.(map[string]interface{})
				nodePools.Spec.NodeLabels = common.GetTypeMapData(nodeLabels)
			}

			if v1, ok := spec[cloudLabelKey]; ok {
				cloudLabels, _ := v1.(map[string]interface{})
				nodePools.Spec.CloudLabels = common.GetTypeMapData(cloudLabels)
			}

			if v1, ok := spec[tkgVsphereKey]; ok {
				if v2, ok := v1.([]interface{}); ok {
					nodePools.Spec.TkgVsphere = expandNodePoolTKGVsphere(v2)
				}
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

func flattenTKGVsphereTopologyNodePool(nodePool *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition) (data interface{}) {
	flattenNodePool := make(map[string]interface{})

	if nodePool == nil {
		return nil
	}

	if nodePool.Info != nil {
		flattenNodePoolInfo := make(map[string]interface{})

		flattenNodePoolInfo[nodePoolNameKey] = nodePool.Info.Name
		flattenNodePoolInfo[nodePoolDescriptionKey] = nodePool.Info.Description

		flattenNodePool[nodePoolInfoKey] = []interface{}{flattenNodePoolInfo}
	}

	if nodePool.Spec != nil {
		flattenNodePoolSpec := make(map[string]interface{})

		flattenNodePoolSpec[workerNodeCountKey] = nodePool.Spec.WorkerNodeCount
		flattenNodePoolSpec[nodeLabelKey] = nodePool.Spec.NodeLabels
		flattenNodePoolSpec[cloudLabelKey] = nodePool.Spec.CloudLabels

		if nodePool.Spec.TkgVsphere != nil {
			flattenNodePoolSpec[tkgVsphereKey] = flattenNodePoolTKGVsphere(nodePool.Spec.TkgVsphere)
		}

		flattenNodePool[nodePoolSpecKey] = []interface{}{flattenNodePoolSpec}
	}

	return flattenNodePool
}

func expandNodePoolTKGVsphere(data []interface{}) (tkgsVsphere *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool) {
	if len(data) == 0 || data[0] == nil {
		return tkgsVsphere
	}

	tkgsVsphereData, _ := data[0].(map[string]interface{})
	tkgsVsphere = &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool{
		VMConfig: &nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig{},
	}

	if v, ok := tkgsVsphereData[vmConfigKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			tkgsVsphere.VMConfig = expandTKGVsphereVMConfig(v1)
		}
	}

	return tkgsVsphere
}

func flattenNodePoolTKGVsphere(tkgVsphere *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool) (data []interface{}) {
	flattenTKGVsphere := make(map[string]interface{})

	if tkgVsphere == nil {
		return nil
	}

	if tkgVsphere.VMConfig != nil {
		flattenTKGVsphere[vmConfigKey] = flattenTKGVsphereVMConfig(tkgVsphere.VMConfig)
	}

	return []interface{}{flattenTKGVsphere}
}

func expandTKGVsphereVMConfig(data []interface{}) (vmConfig *nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig) {
	if len(data) == 0 || data[0] == nil {
		return vmConfig
	}

	lookUpVMConfig, _ := data[0].(map[string]interface{})
	vmConfig = &nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig{}

	if v, ok := lookUpVMConfig[cpuKey]; ok {
		vmConfig.CPU, _ = v.(string)
	}

	if v, ok := lookUpVMConfig[diskKey]; ok {
		vmConfig.DiskGib, _ = v.(string)
	}

	if v, ok := lookUpVMConfig[memoryKey]; ok {
		vmConfig.MemoryMib, _ = v.(string)
	}

	return vmConfig
}

func flattenTKGVsphereVMConfig(vmConfig *nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig) (data []interface{}) {
	flattenVMConfig := make(map[string]interface{})

	if vmConfig == nil {
		return nil
	}

	flattenVMConfig[cpuKey] = vmConfig.CPU
	flattenVMConfig[diskKey] = vmConfig.DiskGib
	flattenVMConfig[memoryKey] = vmConfig.MemoryMib

	return []interface{}{flattenVMConfig}
}
