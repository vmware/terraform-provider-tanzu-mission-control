/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	aksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var ignoredTagsSuffix = "tmc.cloud.vmware.com"

var ClusterSchema = map[string]*schema.Schema{
	CredentialNameKey: {
		Type:        schema.TypeString,
		Description: "Name of the Azure Credential in Tanzu Mission Control",
		Required:    true,
		ForceNew:    true,
	},
	SubscriptionIDKey: {
		Type:        schema.TypeString,
		Description: "Azure Subscription for this cluster",
		Required:    true,
		ForceNew:    true,
	},
	ResourceGroupNameKey: {
		Type:        schema.TypeString,
		Description: "Resource group for this cluster",
		Required:    true,
		ForceNew:    true,
	},
	NameKey: {
		Type:        schema.TypeString,
		Description: "Name of this cluster",
		Required:    true,
		ForceNew:    true,
	},
	common.MetaKey: common.Meta,
	clusterSpecKey: ClusterSpecSchema,
	waitKey: {
		Type:        schema.TypeString,
		Description: "Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero.  The default duration is 30m",
		Default:     "default",
		Optional:    true,
	},
	waitForKubeconfig: {
		Type:        schema.TypeBool,
		Description: "Wait until pinniped extension is ready to provide kubeconfig",
		Default:     false,
		Optional:    true,
	},
	kubeconfigKey: {
		Type:        schema.TypeString,
		Description: "Kubeconfig for connecting to newly created cluster base64 encoded. This will only be returned if you have elected to wait for kubeconfig.",
		Computed:    true,
	},
}

var ClusterSpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			clusterGroupKey: {
				Type:        schema.TypeString,
				Description: "Name of the cluster group to which this cluster belongs",
				Default:     clusterGroupDefaultValue,
				Optional:    true,
			},
			proxyNameKey: {
				Type:        schema.TypeString,
				Description: "Optional proxy name is the name of the Proxy Config to be used for the cluster",
				Optional:    true,
			},
			configKey: {Type: schema.TypeList,
				Description: "AKS config for the cluster control plane",
				Required:    true,
				MaxItems:    1,
				Elem:        ClusterConfig,
			},
			agentNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the cluster in TMC",
				Computed:    true,
				Optional:    true,
			},
			resourceIDKey: {
				Type:        schema.TypeString,
				Description: "Resource ID of the cluster in Azure.",
				Computed:    true,
				Optional:    true,
			},
			nodepoolKey: {
				Type:        schema.TypeList,
				Description: "Nodepool definitions for the cluster",
				Required:    true,
				MinItems:    1,
				Elem:        NodepoolConfig,
			},
		},
	},
}

var ClusterConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		locationKey: {
			Type:        schema.TypeString,
			Description: "The geo-location where the resource lives for the cluster.",
			Required:    true,
			ForceNew:    true,
		},
		kubernetesVersionKey: {
			Type:        schema.TypeString,
			Description: "Kubernetes version of the cluster",
			Required:    true,
		},
		nodeResourceGroupNameKey: {
			Type:        schema.TypeString,
			Description: "Name of the resource group containing nodepools.",
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
		},
		diskEncryptionSetKey: {
			Type:        schema.TypeString,
			Description: "Resource ID of the disk encryption set to use for enabling",
			Optional:    true,
			ForceNew:    true,
		},
		tagsKey: {
			Type:        schema.TypeMap,
			Description: "Metadata to apply to the cluster to assist with categorization and organization",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Computed:    true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return strings.Contains(k, ignoredTagsSuffix)
			},
		},
		skuKey: {
			Type:        schema.TypeList,
			Description: "Azure Kubernetes Service SKU",
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Elem:        SKU,
		},
		accessConfigKey: {
			Type:        schema.TypeList,
			Description: "Access config",
			Optional:    true,
			MaxItems:    1,
			Elem:        AccessConfig,
		},
		apiServerAccessConfigKey: {
			Type:        schema.TypeList,
			Description: "API Server Access Config",
			Optional:    true,
			MaxItems:    1,
			Elem:        APIServerAccessConfig,
		},
		linuxConfigKey: {
			Type:        schema.TypeList,
			Description: "Linux Config",
			Optional:    true,
			ForceNew:    true,
			MaxItems:    1,
			Elem:        LinuxConfig,
		},
		networkConfigKey: {
			Type:        schema.TypeList,
			Description: "Network Config",
			Required:    true,
			MaxItems:    1,
			Elem:        NetworkConfig,
		},
		storageConfigKey: {
			Type:        schema.TypeList,
			Description: "Storage Config",
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Elem:        StorageConfig,
		},
		addonsConfigKey: {
			Type:        schema.TypeList,
			Description: "Addons Config",
			Optional:    true,
			MaxItems:    1,
			Elem:        AddonConfig,
		},
		autoUpgradeConfigKey: {
			Type:        schema.TypeList,
			Description: "Auto Upgrade Config",
			Optional:    true,
			MaxItems:    1,
			Elem:        AutoUpgradeConfig,
		},
		identityConfigKey: {
			Type:        schema.TypeList,
			Description: "Managed Identity Config",
			Optional:    true,
			MaxItems:    1,
			Elem:        ManagedIdentityConfig,
		},
	},
}

var SKU = &schema.Resource{
	Schema: map[string]*schema.Schema{
		skuNameKey: {
			Type:        schema.TypeString,
			Description: "Name of the cluster SKU. Allowed values include: BASIC.",
			Optional:    true,
			Computed:    true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
				string(aksmodel.VmwareTanzuManageV1alpha1AksclusterClusterSKUNameBASIC),
			}, false)),
		},
		skuTierKey: {
			Type:        schema.TypeString,
			Description: "Tier of the cluster SKU. Allowed values include: FREE or PAID.",
			Optional:    true,
			Computed:    true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
				string(aksmodel.VmwareTanzuManageV1alpha1AksclusterTierFREE),
				string(aksmodel.VmwareTanzuManageV1alpha1AksclusterTierPAID),
			}, false)),
		},
	},
}

var AccessConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		enableRbacKey: {
			Type:        schema.TypeBool,
			Description: "Enable kubernetes RBAC",
			Optional:    true,
		},
		disableLocalAccountsKey: {
			Type:        schema.TypeBool,
			Description: "Disable local accounts",
			Optional:    true,
		},
		aadConfigKey: {
			Type:        schema.TypeList,
			Description: "Azure Active Directory config",
			Optional:    true,
			MaxItems:    1,
			Elem:        AADConfig,
		},
	},
}

var AADConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		managedKey: {
			Type:        schema.TypeBool,
			Description: "Enable Managed RBAC",
			Optional:    true,
		},
		tenantIDKey: {
			Type:        schema.TypeString,
			Description: "AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription.",
			Optional:    true,
		},
		adminGroupIDsKey: {
			Type:        schema.TypeList,
			Description: "List of AAD group object IDs that will have admin role of the cluster.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		enableAzureRbacKey: {
			Type:        schema.TypeBool,
			Description: "Enable Azure RBAC for Kubernetes authorization",
			Optional:    true,
		},
	},
}

var APIServerAccessConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		authorizedIPRangesKey: {
			Type:        schema.TypeList,
			Description: "IP ranges authorized to access the Kubernetes API server",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		enablePrivateClusterKey: {
			Type:        schema.TypeBool,
			Description: "Enable Private Cluster",
			Required:    true,
		},
	},
}

var LinuxConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		adminUserNameKey: {
			Type:        schema.TypeString,
			Description: "Administrator username to use for Linux VMs",
			Required:    true,
		},
		sshkeysKey: {
			Type:        schema.TypeList,
			Description: "Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	},
}

var NetworkConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		loadBalancerSkuKey: {
			Type:        schema.TypeString,
			Description: "Load balancer SKU",
			ForceNew:    true,
			Optional:    true,
			Computed:    true,
		},
		networkPluginKey: {
			Type:        schema.TypeString,
			Description: "Network plugin. It is used for building Kubernetes network. Allowed values: azure, kubenet. Specify 'azure' for routable pod IPs from VNET, 'kubenet' for non-routable pod IPs with an overlay network, Defaults to 'kubenet'",
			ForceNew:    true,
			Optional:    true,
			Computed:    true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"",
				aksmodel.VmwareTanzuManageV1alpha1AksClusterNetworkPluginKubenet,
				aksmodel.VmwareTanzuManageV1alpha1AksClusterNetworkPluginAzure,
			}, false)),
		},
		networkPluginModeKey: {
			Type:        schema.TypeString,
			Description: "Network plugin mode. Allowed values: overlay. Used to control the mode the network plugin should operate in. For example, 'overlay' used with networkPlugin=azure will use an overlay network (non-VNET IPs) for pods in the cluster.",
			ForceNew:    true,
			Optional:    true,
			Computed:    true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"",
				aksmodel.VmwareTanzuManageV1alpha1AksClusterNetworkPluginModeOverlay,
			}, false)),
		},
		networkPolicyKey: {
			Type:        schema.TypeString,
			Description: "Network policy",
			ForceNew:    true,
			Optional:    true,
		},
		dnsServiceIPKey: {
			Type:        schema.TypeString,
			Description: "IP address assigned to the Kubernetes DNS service. This key can only be set when the network_config.network_plugin key is set to 'azure'.",
			ForceNew:    true,
			Optional:    true,
			Computed:    true,
		},
		dockerBridgeCidrKey: {
			Type:        schema.TypeString,
			Description: "A CIDR notation IP range assigned to the Docker bridge network",
			ForceNew:    true,
			Optional:    true,
			Computed:    true,
		},
		podCidrKey: {
			Type:        schema.TypeList,
			Description: "CIDR notation IP ranges from which to assign pod IPs",
			ForceNew:    true,
			Optional:    true,
			Computed:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		serviceCidrKey: {
			Type:        schema.TypeList,
			Description: "CIDR notation IP ranges from which to assign service cluster IP. This key can only be set when the network_config.network_plugin key is set to 'azure'.",
			ForceNew:    true,
			Optional:    true,
			Computed:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		dnsPrefixKey: {
			Type:        schema.TypeString,
			Description: "DNS prefix of the cluster",
			ForceNew:    true,
			Required:    true,
		},
	},
}

var StorageConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		enableDiskCsiDriverKey: {
			Type:        schema.TypeBool,
			Description: "Enable the azure disk CSI driver for the storage",
			Optional:    true,
			Computed:    true,
		},
		enableFileCsiDriverKey: {
			Type:        schema.TypeBool,
			Description: "Enable the azure file CSI driver for the storage",
			Optional:    true,
			Computed:    true,
		},
		enableSnapshotControllerKey: {
			Type:        schema.TypeBool,
			Description: "Enable the snapshot controller for the storage",
			Optional:    true,
			Computed:    true,
		},
	},
}

var AddonConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		azureKeyvaultSecretsProviderAddonConfigKey: {
			Type:        schema.TypeList,
			Description: "Keyvault secrets provider addon",
			Optional:    true,
			Elem:        AzureKeyvaulSecretsProviderConfig,
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				return suppressConfig(k, d)
			},
		},
		monitorAddonConfigKey: {
			Type:        schema.TypeList,
			Description: "Monitor addon",
			Optional:    true,
			Elem:        MonitorAddonConfig,
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				return suppressConfig(k, d)
			},
		},
		azurePolicyAddonConfigKey: {
			Type:        schema.TypeList,
			Description: "Azure policy addon",
			Optional:    true,
			Elem:        AzurePolicyAddonConfig,
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				return suppressConfig(k, d)
			},
		},
	},
}

var AzureKeyvaulSecretsProviderConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		enableKey: {
			Type:        schema.TypeBool,
			Description: "Enable Azure Key Vault Secrets Provider",
			Optional:    true,
		},
		enableSecretsRotationKey: {
			Type:        schema.TypeBool,
			Description: "Enable secrets rotation",
			Optional:    true,
		},
		rotationPollIntervalKey: {
			Type:        schema.TypeString,
			Description: "Secret rotation interval",
			Optional:    true,
		},
	},
}

var MonitorAddonConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		enableKey: {
			Type:        schema.TypeBool,
			Description: "Enable monitor",
			Optional:    true,
		},
		logAnalyticsWorkspaceIDKey: {
			Type:        schema.TypeString,
			Description: "Log analytics workspace ID for the monitoring addon",
			Optional:    true,
		},
	},
}

var AzurePolicyAddonConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		enableKey: {
			Type:        schema.TypeBool,
			Description: "Enable policy addon",
			Optional:    true,
		},
	},
}

var AutoUpgradeConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		upgradeChannelKey: {
			Type:        schema.TypeString,
			Description: "Upgrade Channel. Allowed values include: NONE, PATCH, STABLE, RAPID or NODE_IMAGE",
			Optional:    true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
				string(aksmodel.VmwareTanzuManageV1alpha1AksclusterChannelNONE),
				string(aksmodel.VmwareTanzuManageV1alpha1AksclusterChannelPATCH),
				string(aksmodel.VmwareTanzuManageV1alpha1AksclusterChannelSTABLE),
				string(aksmodel.VmwareTanzuManageV1alpha1AksclusterChannelRAPID),
				string(aksmodel.VmwareTanzuManageV1alpha1AksclusterChannelNODEIMAGE),
			}, false)),
		},
	},
}

var ManagedIdentityConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		typeKey: {
			Type:        schema.TypeString,
			Description: "Type of managed identity used by the cluster (default IDENTITY_TYPE_SYSTEM_ASSIGNED). Allowed values include: IDENTITY_TYPE_SYSTEM_ASSIGNED or IDENTITY_TYPE_USER_ASSIGNED",
			Optional:    true,
			Default:     "IDENTITY_TYPE_SYSTEM_ASSIGNED",
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
				string(aksmodel.VmwareTanzuManageV1alpha1AksclusterManagedIdentityTypeSYSTEMASSIGNED),
				string(aksmodel.VmwareTanzuManageV1alpha1AksclusterManagedIdentityTypeUSERASSIGNED),
			}, false)),
		},
		userAssignedKey: {
			Type:        schema.TypeList,
			Description: "User Assigned Managed Identity Config",
			Optional:    true,
			MaxItems:    1,
			Elem:        UserAssignedManagedIdentityConfig,
		},
	},
}

var UserAssignedManagedIdentityConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		resourceIDKey: {
			Type:        schema.TypeString,
			Description: "The ARM resource ID of user assigned identity in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'",
			Required:    true,
		},
	},
}

// NodepoolConfig defines the info and nodepool spec for AKS clusters.
//
// Note: ForceNew is not used in any of the elements because this is a part of
// AKS cluster, and we don't want to replace full clusters because of Nodepool
// changes, these are manually reconciled.
var NodepoolConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		NameKey: {
			Type:        schema.TypeString,
			Description: "Name of the nodepool, immutable",
			Required:    true,
		},
		nodepoolSpecKey: NodepoolSpecSchema,
	},
}

var NodepoolSpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the nodepool",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			modeKey: {
				Type:        schema.TypeString,
				Description: "The mode of the nodepool. Allowed values include: SYSTEM or USER. A cluster must have at least one 'SYSTEM' nodepool at all times.",
				Required:    true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					string(aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolModeUSER),
					string(aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolModeSYSTEM),
				}, false)),
			},
			nodeImageVersionKey: {
				Type:        schema.TypeString,
				Description: "The node image version of the nodepool.",
				Computed:    true,
				Optional:    true,
			},
			typeKey: {
				Type:        schema.TypeString,
				Description: "The Nodepool type. Allowed values include: VIRTUAL_MACHINE_SCALE_SETS or AVAILABILITY_SET.",
				Default:     aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolTypeVIRTUALMACHINESCALESETS,
				Optional:    true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					string(aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolTypeVIRTUALMACHINESCALESETS),
					string(aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolTypeAVAILABILITYSET),
				}, false)),
			},
			availabilityZonesKey: {
				Type:        schema.TypeList,
				Description: "The list of Availability zones to use for nodepool. This can only be specified if the type of the nodepool is AvailabilitySet.",
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			countKey: {
				Type:        schema.TypeInt,
				Description: "Count is the number of nodes",
				Required:    true,
			},
			vmSizeKey: {
				Type:        schema.TypeString,
				Description: "Virtual Machine Size",
				Required:    true,
			},
			scaleSetPriorityKey: {
				Type:        schema.TypeString,
				Description: "Scale set priority. Allowed values include: REGULAR or SPOT.",
				Computed:    true,
				Optional:    true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					string(aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriorityREGULAR),
					string(aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPrioritySPOT),
				}, false)),
			},
			scaleSetEvictionPolicyKey: {
				Type:        schema.TypeString,
				Description: "Scale set eviction policy, Allowed values include: DELETE or DEALLOCATE.",
				Optional:    true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					string(aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicyDELETE),
					string(aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicyDEALLOCATE),
				}, false)),
			},
			maxSpotPriceKey: {
				Type:        schema.TypeFloat,
				Description: "Max spot price",
				Optional:    true,
			},
			osTypeKey: {
				Type:        schema.TypeString,
				Description: "The OS type of the nodepool. Allowed values include: LINUX.",
				Default:     aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolOsTypeLINUX,
				Optional:    true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					string(aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolOsTypeLINUX),
				}, false)),
			},
			osDiskTypeKey: {
				Type:        schema.TypeString,
				Description: "OS Disk Type. Allowed values include: EPHEMERAL or MANAGED.",
				Optional:    true,
				Computed:    true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					string(aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskTypeEPHEMERAL),
					string(aksmodel.VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskTypeMANAGED),
				}, false)),
			},
			osDiskSizeKey: {
				Type:        schema.TypeInt,
				Description: "OS Disk Size in GB to be used to specify the disk size for every machine in the nodepool. If you specify 0, it will apply the default osDisk size according to the vmSize specified",
				Optional:    true,
				Computed:    true,
			},
			maxPodsKey: {
				Type:        schema.TypeInt,
				Description: "The maximum number of pods that can run on a node",
				Optional:    true,
				Computed:    true,
			},
			enableNodePublicIPKey: {
				Type:        schema.TypeBool,
				Description: "Whether each node is allocated its own public IP",
				Optional:    true,
			},
			taintsKey: {
				Type:        schema.TypeList,
				Description: "The taints added to new nodes during nodepool create and scale",
				Optional:    true,
				Elem:        Taints,
			},
			vnetSubnetKey: {
				Type:        schema.TypeString,
				Description: "The ID of a subnet in an existing VNet into which to deploy the cluster. If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes",
				Optional:    true,
			},
			podSubnetKey: {
				Type:        schema.TypeString,
				Description: "The ID of a subnet in an existing VNet into which to assign pods in the cluster. Requires network-plugin to be azure and not compatible with network-plugin-mode overlay",
				Optional:    true,
			},
			nodeLabelsKey: {
				Type:        schema.TypeMap,
				Description: "The node labels to be persisted across all nodes in nodepool",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			tagsKey: {
				Type:        schema.TypeMap,
				Description: "AKS specific node tags",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			autoscalingConfigKey: {
				Type:        schema.TypeList,
				Description: "Auto scaling config.",
				Optional:    true,
				MaxItems:    1,
				Elem:        AutoScaleConfig,
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					return suppressConfig(k, d)
				},
			},
			upgradeConfigKey: {
				Type:        schema.TypeList,
				Description: "upgrade config",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						maxSurgeKey: {
							Type:        schema.TypeString,
							Description: "Max Surge",
							Optional:    true,
						},
					},
				},
			},
		},
	},
}

var Taints = &schema.Resource{
	Schema: map[string]*schema.Schema{
		effectKey: {
			Type:        schema.TypeString,
			Description: "Current effect state of the node pool",
			Optional:    true,
		},
		keyKey: {
			Type:        schema.TypeString,
			Description: "The taint key to be applied to a node",
			Optional:    true,
		},
		valueKey: {
			Type:        schema.TypeString,
			Description: "The taint value corresponding to the taint key",
			Optional:    true,
		},
	},
}

var AutoScaleConfig = &schema.Resource{
	Schema: map[string]*schema.Schema{
		enableKey: {
			Type:        schema.TypeBool,
			Description: "Enable auto scaling",
			Optional:    true,
		},
		minCountKey: {
			Type:             schema.TypeInt,
			Description:      "Minimum node count",
			Optional:         true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
		},
		maxCountKey: {
			Type:             schema.TypeInt,
			Description:      "Maximum node count",
			Optional:         true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
		},
	},
}

func suppressConfig(key string, resourceData *schema.ResourceData) bool {
	lastDotIndex := strings.LastIndex(key, ".")
	if lastDotIndex == -1 {
		return false
	}

	key = key[:lastDotIndex]

	return resourceData.Get(key+".enable") == false
}
