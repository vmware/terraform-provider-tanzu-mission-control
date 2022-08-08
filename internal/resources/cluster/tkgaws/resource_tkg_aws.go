/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgaws

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	nodepoolmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	tkgawsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgaws"
)

var TkgAWSClusterSpec = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The Tanzu Kubernetes Grid (TKGm) AWS cluster spec",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			settingsKey:     tkgAWSSettings,
			distributionKey: tkgAWSDistribution,
			topologyKey:     tkgAWSTopology,
		},
	},
}

func ConstructTKGAWSClusterSpec(data []interface{}) (spec *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec) {
	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData, _ := data[0].(map[string]interface{})
	spec = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec{}

	if v, ok := specData[settingsKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.Settings = expandTKGAWSSettings(v1)
		}
	}

	if v, ok := specData[distributionKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.Distribution = expandTKGAWSDistribution(v1)
		}
	}

	if v, ok := specData[topologyKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.Topology = expandTKGAWSTopology(v1)
		}
	}

	return spec
}

func FlattenTKGAWSClusterSpec(spec *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec) (data []interface{}) {
	flattenSpecData := make(map[string]interface{})

	flattenSpecData[settingsKey] = flattenTKGAWSSettings(spec.Settings)
	flattenSpecData[distributionKey] = flattenTKGAWSDistribution(spec.Distribution)
	flattenSpecData[topologyKey] = flattenTKGAWSTopology(spec.Topology)

	return []interface{}{flattenSpecData}
}

var tkgAWSSettings = &schema.Schema{
	Type:        schema.TypeList,
	Description: "AWS related settings for workload cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			networkKey:  tkgAWSNetwork,
			securityKey: tkgAWSSecurity,
		},
	},
}

var tkgAWSSecurity = &schema.Schema{
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

func expandTKGAWSSettings(data []interface{}) (settings *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings) {
	if len(data) == 0 || data[0] == nil {
		return settings
	}

	lookUpSettings, _ := data[0].(map[string]interface{})
	settings = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings{}

	if v, ok := lookUpSettings[networkKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			settings.Network = expandTKGAWSNetworkSettings(v1)
		}
	}

	if v, ok := lookUpSettings[securityKey]; ok {
		security, _ := v.([]interface{})
		if len(security) == 0 || security[0] == nil {
			return settings
		}

		lookUpSecurity, _ := security[0].(map[string]interface{})
		settings.Security = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSecuritySettings{}

		if sshKey, ok := lookUpSecurity[sshKey]; ok {
			settings.Security.SSHKey, _ = sshKey.(string)
		}
	}

	return settings
}

func flattenTKGAWSSettings(settings *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings) (data []interface{}) {
	flattenSettings := make(map[string]interface{})

	if settings == nil {
		return nil
	}

	if settings.Network != nil {
		flattenSettings[networkKey] = flattenTKGAWSNetworkSettings(settings.Network)
	}

	if settings.Security != nil {
		flattenSettingsSecurity := make(map[string]interface{})
		flattenSettingsSecurity[sshKey] = settings.Security.SSHKey

		flattenSettings[securityKey] = []interface{}{flattenSettingsSecurity}
	}

	return []interface{}{flattenSettings}
}

var tkgAWSNetwork = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Network Settings specifies network-related settings for the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			clusterKey:  tkgAWSCluster,
			providerKey: tkgAWSProvider,
		},
	},
}

func expandTKGAWSNetworkSettings(data []interface{}) (network *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkSettings) {
	if len(data) == 0 || data[0] == nil {
		return network
	}

	lookUpNetwork, _ := data[0].(map[string]interface{})
	network = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkSettings{
		Cluster:  &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork{},
		Provider: &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork{},
	}

	if v, ok := lookUpNetwork[clusterKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			network.Cluster = expandTKGAWSClusterNetwork(v1)
		}
	}

	if v, ok := lookUpNetwork[providerKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			network.Provider = expandTKGAWSProviderNetwork(v1)
		}
	}

	return network
}

func flattenTKGAWSNetworkSettings(network *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkSettings) (data []interface{}) {
	flattenNetworkSettings := make(map[string]interface{})

	if network == nil {
		return nil
	}

	if network.Cluster != nil {
		flattenNetworkSettings[clusterKey] = flattenTKGAWSClusterNetwork(network.Cluster)
	}

	if network.Provider != nil {
		flattenNetworkSettings[providerKey] = flattenTKGAWSProviderNetwork(network.Provider)
	}

	return []interface{}{flattenNetworkSettings}
}

var tkgAWSCluster = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Cluster network specifies kubernetes network information for the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			apiServerPortKey: {
				Type:        schema.TypeInt,
				Default:     apiServerPortDefaultValue,
				Description: "APIServerPort specifies the port address for the cluster that defaults to 6443.",
				Optional:    true,
			},
			podsKey: {
				Type:        schema.TypeList,
				Description: "Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cidrBlocksKey: cidrBlock,
					},
				},
			},
			servicesKey: {
				Type:        schema.TypeList,
				Description: "Service CIDR for kubernetes services defaults to 10.96.0.0/12",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cidrBlocksKey: cidrBlock,
					},
				},
			},
		},
	},
}

var cidrBlock = &schema.Schema{
	Type:        schema.TypeString,
	Description: "CIDRBlocks specifies one or more of IP address ranges",
	Required:    true,
}

func expandTKGAWSClusterNetwork(data []interface{}) (cluster *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork) {
	if len(data) == 0 || data[0] == nil {
		return cluster
	}

	lookUpCluster, _ := data[0].(map[string]interface{})
	cluster = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork{
		APIServerPort: apiServerPortDefaultValue,
		Pods:          []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange{},
		Services:      []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange{},
	}

	if v, ok := lookUpCluster[apiServerPortKey]; ok {
		apiServerPort := v.(int)
		cluster.APIServerPort = int32(apiServerPort)
	}

	if v, ok := lookUpCluster[podsKey]; ok {
		pods, _ := v.([]interface{})
		for _, pd := range pods {
			cluster.Pods = append(cluster.Pods, expandTKGAWSNetworkRange(pd))
		}
	}

	if v, ok := lookUpCluster[servicesKey]; ok {
		services, _ := v.([]interface{})
		for _, sv := range services {
			cluster.Services = append(cluster.Services, expandTKGAWSNetworkRange(sv))
		}
	}

	return cluster
}

func expandTKGAWSNetworkRange(data interface{}) (ranges *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange) {
	if data == nil {
		return ranges
	}

	lookUpRanges, _ := data.(map[string]interface{})
	ranges = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange{}

	if v, ok := lookUpRanges[cidrBlocksKey]; ok {
		ranges.CidrBlocks, _ = v.(string)
	}

	return ranges
}

func flattenTKGAWSClusterNetwork(cluster *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork) (data []interface{}) {
	flattenClusterNetwork := make(map[string]interface{})

	if cluster == nil {
		return nil
	}

	flattenClusterNetwork[apiServerPortKey] = cluster.APIServerPort

	pds := make([]interface{}, 0)

	for _, pd := range cluster.Pods {
		pds = append(pds, flattenTKGAWSNetworkRange(pd))
	}

	flattenClusterNetwork[podsKey] = pds

	svs := make([]interface{}, 0)

	for _, sv := range cluster.Services {
		svs = append(svs, flattenTKGAWSNetworkRange(sv))
	}

	flattenClusterNetwork[servicesKey] = svs

	return []interface{}{flattenClusterNetwork}
}

func flattenTKGAWSNetworkRange(ranges *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange) (data interface{}) {
	flattenRanges := make(map[string]interface{})

	if ranges == nil {
		return nil
	}

	flattenRanges[cidrBlocksKey] = ranges.CidrBlocks

	return flattenRanges
}

var tkgAWSProvider = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Provider Network specifies provider specific network information for the cluster",
	Required:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			subnetsKey: tkgAWSSubnets,
			vpcKey:     tkgAWSVPC,
		},
	},
}

func expandTKGAWSProviderNetwork(data []interface{}) (provider *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork) {
	if len(data) == 0 || data[0] == nil {
		return provider
	}

	lookUpProvider, _ := data[0].(map[string]interface{})
	provider = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork{
		Subnets: []*tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet{},
		Vpc:     &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsVPC{},
	}

	if v, ok := lookUpProvider[subnetsKey]; ok {
		subnets, _ := v.([]interface{})
		for _, sn := range subnets {
			provider.Subnets = append(provider.Subnets, expandTKGAWSSubnets(sn))
		}
	}

	if v, ok := lookUpProvider[vpcKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			provider.Vpc = expandTKGAWSVPC(v1)
		}
	}

	return provider
}

func flattenTKGAWSProviderNetwork(provider *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork) (data []interface{}) {
	flattenProviderNetwork := make(map[string]interface{})

	if provider == nil {
		return nil
	}

	sns := make([]interface{}, 0)

	for _, sn := range provider.Subnets {
		sns = append(sns, flattenTKGAWSSubnets(sn))
	}

	flattenProviderNetwork[subnetsKey] = sns

	if provider.Vpc != nil {
		flattenProviderNetwork[vpcKey] = flattenTKGAWSVPC(provider.Vpc)
	}

	return []interface{}{flattenProviderNetwork}
}

var tkgAWSSubnets = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Optional list of subnets used to place the nodes in the cluster",
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			availabilityZoneKey: {
				Type:        schema.TypeString,
				Description: "AWS availability zone e.g. us-west-2a",
				Optional:    true,
			},
			subnetCIDRBlockKey: {
				Type:        schema.TypeString,
				Description: "CIDR for AWS subnet which must be in the range of AWS VPC CIDR block",
				Optional:    true,
			},
			subnetIDKey: {
				Type:        schema.TypeString,
				Description: "This is the subnet ID of AWS. The rest of the fields are ignored if this field is specified",
				Optional:    true,
			},
			isPublicKey: {
				Type:        schema.TypeBool,
				Description: "Describes if it is public subnet or private subnet",
				Default:     false,
				Optional:    true,
			},
		},
	},
}

func expandTKGAWSSubnets(data interface{}) (subnets *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet) {
	if data == nil {
		return subnets
	}

	lookUpSubnets, _ := data.(map[string]interface{})
	subnets = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet{}

	if v, ok := lookUpSubnets[availabilityZoneKey]; ok {
		subnets.AvailabilityZone, _ = v.(string)
	}

	if v, ok := lookUpSubnets[subnetCIDRBlockKey]; ok {
		subnets.CidrBlock, _ = v.(string)
	}

	if v, ok := lookUpSubnets[subnetIDKey]; ok {
		subnets.ID, _ = v.(string)
	}

	if v, ok := lookUpSubnets[isPublicKey]; ok {
		subnets.IsPublic, _ = v.(bool)
	}

	return subnets
}

func flattenTKGAWSSubnets(subnets *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet) (data interface{}) {
	flattenSubnets := make(map[string]interface{})

	if subnets == nil {
		return nil
	}

	flattenSubnets[availabilityZoneKey] = subnets.AvailabilityZone
	flattenSubnets[subnetCIDRBlockKey] = subnets.CidrBlock
	flattenSubnets[subnetIDKey] = subnets.ID
	flattenSubnets[isPublicKey] = subnets.IsPublic

	return flattenSubnets
}

var tkgAWSVPC = &schema.Schema{
	Type:        schema.TypeList,
	Description: "AWS VPC configuration for the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			vpcCIDRBlockKey: {
				Type:        schema.TypeString,
				Description: "CIDR for AWS VPC. A valid example is 10.0.0.0/16",
				Computed:    true,
				Optional:    true,
			},
			vpcIDKey: {
				Type:        schema.TypeString,
				Description: "AWS VPC ID. The rest of the fields are ignored if this field is specified. Kindly add the VPC ID to the terraform script in case of existing VPC.",
				Computed:    true,
				Optional:    true,
			},
		},
	},
}

func expandTKGAWSVPC(data []interface{}) (vpc *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsVPC) {
	if len(data) == 0 || data[0] == nil {
		return vpc
	}

	lookUpVPC, _ := data[0].(map[string]interface{})
	vpc = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsVPC{}

	if v, ok := lookUpVPC[vpcCIDRBlockKey]; ok {
		vpc.CidrBlock, _ = v.(string)
	}

	if v, ok := lookUpVPC[vpcIDKey]; ok {
		vpc.ID, _ = v.(string)
	}

	return vpc
}

func flattenTKGAWSVPC(vpc *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsVPC) (data []interface{}) {
	flattenVPC := make(map[string]interface{})

	if vpc == nil {
		return nil
	}

	flattenVPC[vpcCIDRBlockKey] = vpc.CidrBlock
	flattenVPC[vpcIDKey] = vpc.ID

	return []interface{}{flattenVPC}
}

var tkgAWSDistribution = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Kubernetes version distribution for the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			regionKey: {
				Type:        schema.TypeString,
				Description: "Specifies region of the cluster",
				Required:    true,
			},
			versionKey: {
				Type:        schema.TypeString,
				Description: "Specifies version of the cluster",
				Required:    true,
			},
			provisionerCredentialKey: {
				Type:        schema.TypeString,
				Description: "Specifies name of the account in which to create the cluster",
				Optional:    true,
			},
		},
	},
}

func expandTKGAWSDistribution(data []interface{}) (distribution *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsDistribution) {
	if len(data) == 0 || data[0] == nil {
		return distribution
	}

	lookUpDistribution, _ := data[0].(map[string]interface{})
	distribution = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsDistribution{}

	if v, ok := lookUpDistribution[regionKey]; ok {
		distribution.Region, _ = v.(string)
	}

	if v, ok := lookUpDistribution[versionKey]; ok {
		distribution.Version, _ = v.(string)
	}

	if v, ok := lookUpDistribution[provisionerCredentialKey]; ok {
		distribution.ProvisionerCredentialName, _ = v.(string)
	}

	return distribution
}

func flattenTKGAWSDistribution(distribution *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsDistribution) (data []interface{}) {
	flattenDistribution := make(map[string]interface{})

	if distribution == nil {
		return nil
	}

	flattenDistribution[regionKey] = distribution.Region
	flattenDistribution[versionKey] = distribution.Version
	flattenDistribution[provisionerCredentialKey] = distribution.ProvisionerCredentialName

	return []interface{}{flattenDistribution}
}

var tkgAWSTopology = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Topology configuration of the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			controlPlaneKey: tkgAWSTopologyControlPlane,
			nodePoolsKey: {
				Type:        schema.TypeList,
				Description: "Nodepool specific configuration",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						nodePoolInfoKey: tkgAWSNodePoolInfo,
						nodePoolSpecKey: tkgAWSNodePoolSpec,
					},
				},
			},
		},
	},
}

func expandTKGAWSTopology(data []interface{}) (topology *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology) {
	if len(data) == 0 || data[0] == nil {
		return topology
	}

	lookUpTopology, _ := data[0].(map[string]interface{})
	topology = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology{}

	if v, ok := lookUpTopology[controlPlaneKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			topology.ControlPlane = expandTKGAWSTopologyControlPlane(v1)
		}
	}

	if v, ok := lookUpTopology[nodePoolsKey]; ok {
		nodepools, _ := v.([]interface{})
		for _, np := range nodepools {
			topology.NodePools = append(topology.NodePools, expandTKGAWSTopologyNodePool(np))
		}
	}

	return topology
}

func flattenTKGAWSTopology(topology *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsTopology) (data []interface{}) {
	flattenTopology := make(map[string]interface{})

	if topology == nil {
		return nil
	}

	flattenTopology[controlPlaneKey] = flattenTKGAWSTopologyControlPlane(topology.ControlPlane)

	nps := make([]interface{}, 0)

	for _, np := range topology.NodePools {
		nps = append(nps, flattenTKGAWSTopologyNodePool(np))
	}

	flattenTopology[nodePoolsKey] = nps

	return []interface{}{flattenTopology}
}

var tkgAWSTopologyControlPlane = &schema.Schema{
	Type:        schema.TypeList,
	Description: "AWS specific control plane configuration for workload cluster object",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			availabilityZonesKey: {
				Type:        schema.TypeList,
				Description: "List of availability zones for the control plane nodes",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			instanceTypeKey: {
				Type:        schema.TypeString,
				Description: "Control plane instance type",
				Required:    true,
			},
			highAvailabilityKey: {
				Type:        schema.TypeBool,
				Description: "Flag which controls if the cluster needs to be highly available. HA cluster creates three controlplane machines, and non HA creates just one",
				Default:     false,
				Optional:    true,
			},
		},
	},
}

func expandTKGAWSTopologyControlPlane(data []interface{}) (controlPlane *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsControlPlane) {
	if len(data) == 0 || data[0] == nil {
		return controlPlane
	}

	lookUpControlPlane, _ := data[0].(map[string]interface{})
	controlPlane = &tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsControlPlane{}

	if avzones, ok := lookUpControlPlane[availabilityZonesKey]; ok {
		avz, _ := avzones.([]interface{})

		s := make([]string, 0)

		for _, raw := range avz {
			s = append(s, raw.(string))
		}

		controlPlane.AvailabilityZones = s
	}

	if v, ok := lookUpControlPlane[instanceTypeKey]; ok {
		controlPlane.InstanceType, _ = v.(string)
	}

	if v, ok := lookUpControlPlane[highAvailabilityKey]; ok {
		controlPlane.HighAvailability, _ = v.(bool)
	}

	return controlPlane
}

func flattenTKGAWSTopologyControlPlane(controlPlane *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsControlPlane) (data []interface{}) {
	flattenControlPlane := make(map[string]interface{})

	if controlPlane == nil {
		return nil
	}

	flattenControlPlane[availabilityZonesKey] = controlPlane.AvailabilityZones
	flattenControlPlane[instanceTypeKey] = controlPlane.InstanceType
	flattenControlPlane[highAvailabilityKey] = controlPlane.HighAvailability

	return []interface{}{flattenControlPlane}
}

var tkgAWSNodePoolInfo = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Info is the meta information of nodepool for cluster",
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
				Description: "Description of the nodepool",
				Optional:    true,
			},
		},
	},
}

var tkgAWSNodePoolSpec = &schema.Schema{
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
			tkgAWSKey: {
				Type:        schema.TypeList,
				Description: "Nodepool config for tkg aws",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						nodepoolAvailabilityZoneKey: {
							Type:        schema.TypeString,
							Description: "Availability zone for the nodepool that is to be used when you are creating a nodepool for cluster in TMC hosted AWS solution",
							Optional:    true,
						},
						nodepoolInstanceTypeKey: {
							Type:        schema.TypeString,
							Description: "Nodepool instance type whose potential values could be found using cluster:options api",
							Required:    true,
						},
						nodePlacementKey: {
							Type:        schema.TypeList,
							Description: "List of Availability Zones to place the AWS nodes on. Please use this field to provision a nodepool for workload cluster on an attached TKG AWS management cluster",
							Required:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									awsAvailabilityZoneKey: {
										Type:        schema.TypeString,
										Description: "The Availability Zone where the AWS nodes are placed",
										Required:    true,
									},
								},
							},
						},
						nodePoolSubnetIDKey: {
							Type:        schema.TypeString,
							Description: "Subnet ID of the private subnet in which you want the nodes to be created in",
							Computed:    true,
						},
						nodepoolVersionKey: {
							Type:        schema.TypeString,
							Description: "Kubernetes version of the node pool",
							Required:    true,
						},
					},
				},
			},
		},
	},
}

func expandTKGAWSTopologyNodePool(data interface{}) (nodePools *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition) {
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

			if v1, ok := spec[tkgAWSKey]; ok {
				if v2, ok := v1.([]interface{}); ok {
					nodePools.Spec.TkgAws = expandNodePoolTKGAWS(v2)
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

func flattenTKGAWSTopologyNodePool(nodePool *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolDefinition) (data interface{}) {
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

		if nodePool.Spec.TkgAws != nil {
			flattenNodePoolSpec[tkgAWSKey] = flattenNodePoolTKGAWS(nodePool.Spec.TkgAws)
		}

		flattenNodePool[nodePoolSpecKey] = []interface{}{flattenNodePoolSpec}
	}

	return flattenNodePool
}

func expandNodePoolTKGAWS(data []interface{}) (tkgAWS *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool) {
	if len(data) == 0 || data[0] == nil {
		return tkgAWS
	}

	lookUpTKGAWS, _ := data[0].(map[string]interface{})
	tkgAWS = &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool{}

	if v, ok := lookUpTKGAWS[nodepoolAvailabilityZoneKey]; ok {
		tkgAWS.AvailabilityZone, _ = v.(string)
	}

	if v, ok := lookUpTKGAWS[nodepoolInstanceTypeKey]; ok {
		tkgAWS.InstanceType, _ = v.(string)
	}

	if v, ok := lookUpTKGAWS[nodePlacementKey]; ok {
		nodeplacements, _ := v.([]interface{})
		for _, np := range nodeplacements {
			tkgAWS.NodePlacement = append(tkgAWS.NodePlacement, expandTKGAWSNodePlacement(np))
		}
	}

	if v, ok := lookUpTKGAWS[nodePoolSubnetIDKey]; ok {
		tkgAWS.SubnetID, _ = v.(string)
	}

	if v, ok := lookUpTKGAWS[nodepoolVersionKey]; ok {
		tkgAWS.Version, _ = v.(string)
	}

	return tkgAWS
}

func flattenNodePoolTKGAWS(tkgAWS *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool) (data []interface{}) {
	flattenTKGAWS := make(map[string]interface{})

	if tkgAWS == nil {
		return nil
	}

	flattenTKGAWS[nodepoolAvailabilityZoneKey] = tkgAWS.AvailabilityZone
	flattenTKGAWS[nodepoolInstanceTypeKey] = tkgAWS.InstanceType

	nps := make([]interface{}, 0)

	for _, np := range tkgAWS.NodePlacement {
		nps = append(nps, flattenTKGAWSNodePlacement(np))
	}

	flattenTKGAWS[nodePlacementKey] = nps

	flattenTKGAWS[nodePoolSubnetIDKey] = tkgAWS.SubnetID
	flattenTKGAWS[nodepoolVersionKey] = tkgAWS.Version

	return []interface{}{flattenTKGAWS}
}

func expandTKGAWSNodePlacement(data interface{}) (nodeplacement *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement) {
	if data == nil {
		return nodeplacement
	}

	lookUpNodePlacement, _ := data.(map[string]interface{})
	nodeplacement = &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement{}

	if v, ok := lookUpNodePlacement[awsAvailabilityZoneKey]; ok {
		nodeplacement.AvailabilityZone, _ = v.(string)
	}

	return nodeplacement
}

func flattenTKGAWSNodePlacement(nodeplacement *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement) (data interface{}) {
	flattenNodePlacement := make(map[string]interface{})

	if nodeplacement == nil {
		return nil
	}

	flattenNodePlacement[awsAvailabilityZoneKey] = nodeplacement.AvailabilityZone

	return flattenNodePlacement
}
