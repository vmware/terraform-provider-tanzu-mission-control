---
Title: "AKS Cluster Resource"
Description: |-
    Create an Azure AKS cluster resource managed by Tanzu Mission Control.
---

# AKS Cluster

The `tanzu-mission-control_akscluster` resource allows you to provision and manage [Azure AKS](https://azure.microsoft.com/en-us/products/kubernetes-service) through Tanzu Mission Control.
It allows users to connect Tanzu Mission Control to their Microsoft Azure account and create, update/upgrade, and delete AKS clusters and node groups (called node pools in Tanzu).

## Provisioning a AKS Cluster

To use the **Tanzu Mission Control** for creating a new cluster, you must first log into Azure and set up an Azure AKS
credential that allows VMware Tanzu Mission Control to connect to your Azure subscription and manage resources in your
Azure account. For more information, see [connecting an Azure account for AKS cluster lifecycle management][azure-account]
and [create an AKS Cluster][create-cluster].

You must also have the appropriate permissions in Tanzu Mission Control:

- To provision a cluster, you must have `cluster.admin` permissions.
- You must also have `clustergroup.edit` permissions on the cluster group to detach a cluster.

[azure-account]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-2CA6A21A-1D33-4852-B8F2-86BB3A1337E4.html
[create-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-90ED8C73-8A40-46FF-85AE-A8DAA9048AA9.html

__Note__: Fields under the [nested Schema for `spec.nodepool`](#nestedblock--spec--nodepool) which are marked as "immutable" can't be changed. To update those fields, you need to create a new node pool or rename the node pool (which will have the same effect).

## Example Usage

```terraform
// Create Tanzu Mission Control Azure AKS workload cluster entry
resource "tanzu-mission-control_akscluster" "tf_aks_cluster" {
  credential_name = "aks-test-credential" // Required
  subscription    = "azure-test-subscription"    // Required
  resource_group  = "azure-test-resource-group"  // Required
  name            = "tf-aks-cluster"    // Required

  meta {
    description = "aks test cluster"
    labels      = { "key1" : "value1" }
  }

  spec {
    cluster_group = "test-cluster-group" // Default: default
    proxy         = "tmc-proxy"

    azure_AKS {
      location                 = "westus2" // Required     // Force Recreate
      version                  = "1.23"  // Required
      node_resource_group_name = "test-aks-resource-group>" // Force Recreate
      disk_encryption_set      = "test-aks-disk-encryption-set-name"      // Force Recreate
      tags                     = { "tagkey" : "tagvalue" }

      sku {
        name = "test-aks-sku-name"
        tier = "Premium" // Required
      }

      access_config {
        enable_rbac            = true
        disable_local_accounts = true
        aad_config {
          managed         = true
          tenantId        = "1325478f-9gg7-1376-e473-6kek13986365"
          admin_group_ids = [
            "5d241325-8rr4-9104-j598-9d14afa27aed",
            "2k907631-t454-w335-p132-7e25nmz98brf",
          ]
          enable_azure_rbac = true
        }
      }

      api_server_access_config {
        authorized_ip_ranges = [
          "73.140.245.0/24",
          "71.952.241.0/32",
        ]
        enable_private_cluster = true // Forces Recreate
      }

      linux_config {
        // Force Recreate
        admin_username = "test-admin-username"
        ssh_keys       = [
          "test-ssh-key-1",
          "test-ssh-key-2",
        ]
      }

      network_config {
        // Required
        load_balancer_sku  = "standard"  // Forces Recreate
        network_plugin     = "azure"     // Forces Recreate
        network_policy     = "azure"     // Forces Recreate
        dns_service_ip     = "10.2.0.10"     // Forces Recreate
        docker_bridge_cidr = "172.17.0.1/16" // Forces Recreate
        pod_cidr           = [
          // Forces Recreate
          "10.244.0.0/16",
          "10.246.0.0/16",
        ]
        service_cidr = [
          // Forces Recreate
          "10.100.0.0/24",
          "10.101.0.0/24",
        ]
        dns_prefix                      = "testdnsprefix" // Required
        enable_http_application_routing = true
      }

      storage_config {
        enable_disk_csi_driver     = true
        enable_file_csi_driver     = true
        enable_snapshot_controller = true
      }

      addons_config {
        azure_keyvault_secrets_provider_addon_config {
          enable = true
          keyvault_secrets_provider_config {
            enable_secret_rotation = true
            rotation_poll_interval = "5m"
          }
        }

        monitor_addon_config {
          enable                     = true
          log_analytics_workspace_id = "test-log-analytics-workspace-id"
        }

        azure_policy_addon_config {
          enable = true
        }
      }

      auto_upgrade_config {
        upgrade_channel = "stable"
      }

      nodepools = [
        {
          info = {
            name = "third-np"
          }

          spec = {
            mode              = "System" // Required
            type              = "Microsoft.ContainerService/managedClusters/agentPools"
            availabilityZones = [
              "West US 2",
              "West US 3",
            ]
            count                     = 1 // Required
            vm_size                   = "Standard_DS2_v2" // Required // Force Recreate
            scale_set_priority        = "Regular"// Force Recreate
            scale_set_eviction_policy = "Delete" // Force Recreate
            spot_max_price            = 1.00
            os_type                   = "Linux"
            os_disk_type              = "Ephemeral"        // Force Recreate
            os_disk_size_gb           = 60                      // Force Recreate
            max_pods                  = 10                      // Force Recreate
            enable_node_public_ip     = true
            node_taints               = [
              {
                effect = "NoSchedule"
                key    = "randomkey"
                value  = "randomvalue"
              }
            ]
            vnet_subnet_id = "test-vnet-subnet-id" // Force Recreate
            node_labels    = { "nplabelkey" : "nplabelvalue" }
            tags           = { "nptagkey" : "nptagvalue3" }

            auto_scaling_config = {
              enable    = true // Force Recreate
              min_count = 1
              max_count = 5
            }

            upgrade_config = {
              max_surge = "30%"
            }
          }
        }
      ]
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `credential_name` (String) Name of the Azure Credential in Tanzu Mission Control
- `name` (String) Name of this cluster
- `resource_group` (String) Resource group for this cluster
- `spec` (Block List, Min: 1, Max: 1) Spec for the cluster (see [below for nested schema](#nestedblock--spec))
- `subscription_id` (String) Azure Subscription for this cluster

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))
- `ready_wait_timeout` (String) Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero.  The default duration is 30m

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Required:

- `config` (Block List, Min: 1, Max: 1) AKS config for the cluster control plane (see [below for nested schema](#nestedblock--spec--config))
- `nodepool` (Block List, Min: 1) Nodepool definitions for the cluster (see [below for nested schema](#nestedblock--spec--nodepool))

Optional:

- `agent_name` (String) Name of the cluster in TMC
- `cluster_group` (String) Name of the cluster group to which this cluster belongs
- `proxy` (String) Optional proxy name is the name of the Proxy Config to be used for the cluster
- `resource_id` (String) Resource ID of the cluster in Azure.

<a id="nestedblock--spec--config"></a>
### Nested Schema for `spec.config`

Required:

- `kubernetes_version` (String) Kubernetes version of the cluster
- `location` (String) The geo-location where the resource lives for the cluster.
- `network_config` (Block List, Min: 1, Max: 1) Network Config (see [below for nested schema](#nestedblock--spec--config--network_config))

Optional:

- `access_config` (Block List, Max: 1) Access config (see [below for nested schema](#nestedblock--spec--config--access_config))
- `addon_config` (Block List, Max: 1) Addons Config (see [below for nested schema](#nestedblock--spec--config--addon_config))
- `api_server_access_config` (Block List, Max: 1) API Server Access Config (see [below for nested schema](#nestedblock--spec--config--api_server_access_config))
- `auto_upgrade_config` (Block List, Max: 1) Auto Upgrade Config (see [below for nested schema](#nestedblock--spec--config--auto_upgrade_config))
- `disk_encryption_set` (String) Resource ID of the disk encryption set to use for enabling
- `linux_config` (Block List, Max: 1) Linux Config (see [below for nested schema](#nestedblock--spec--config--linux_config))
- `node_resource_group_name` (String) Name of the resource group containing nodepools.
- `sku` (Block List, Max: 1) Azure Kubernetes Service SKU (see [below for nested schema](#nestedblock--spec--config--sku))
- `storage_config` (Block List, Max: 1) Storage Config (see [below for nested schema](#nestedblock--spec--config--storage_config))
- `tags` (Map of String) Metadata to apply to the cluster to assist with categorization and organization

<a id="nestedblock--spec--config--network_config"></a>
### Nested Schema for `spec.config.network_config`

Required:

- `dns_prefix` (String) DNS prefix of the cluster

Optional:

- `dns_service_ip` (String) IP address assigned to the Kubernetes DNS service
- `docker_bridge_cidr` (String) A CIDR notation IP range assigned to the Docker bridge network
- `load_balancer_sku` (String) Load balancer SKU
- `network_plugin` (String) Network plugin
- `network_policy` (String) Network policy
- `pod_cidr` (List of String) CIDR notation IP ranges from which to assign pod IPs
- `service_cidr` (List of String) CIDR notation IP ranges from which to assign service cluster IP


<a id="nestedblock--spec--config--access_config"></a>
### Nested Schema for `spec.config.access_config`

Optional:

- `aad_config` (Block List, Max: 1) Azure Active Directory config (see [below for nested schema](#nestedblock--spec--config--access_config--aad_config))
- `disable_local_accounts` (Boolean) Disable local accounts
- `enable_rbac` (Boolean) Enable kubernetes RBAC

<a id="nestedblock--spec--config--access_config--aad_config"></a>
### Nested Schema for `spec.config.access_config.aad_config`

Optional:

- `admin_group_ids` (List of String) List of AAD group object IDs that will have admin role of the cluster.
- `enable_azure_rbac` (Boolean) Enable Azure RBAC for Kubernetes authorization
- `managed` (Boolean) Enable Managed RBAC
- `tenant_id` (String) AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription.



<a id="nestedblock--spec--config--addon_config"></a>
### Nested Schema for `spec.config.addon_config`

Optional:

- `azure_keyvault_secrets_provider_addon_config` (Block List) Keyvault secrets provider addon (see [below for nested schema](#nestedblock--spec--config--addon_config--azure_keyvault_secrets_provider_addon_config))
- `azure_policy_addon_config` (Block List) Azure policy addon (see [below for nested schema](#nestedblock--spec--config--addon_config--azure_policy_addon_config))
- `monitor_addon_config` (Block List) Monitor addon (see [below for nested schema](#nestedblock--spec--config--addon_config--monitor_addon_config))

<a id="nestedblock--spec--config--addon_config--azure_keyvault_secrets_provider_addon_config"></a>
### Nested Schema for `spec.config.addon_config.azure_keyvault_secrets_provider_addon_config`

Optional:

- `enable` (Boolean) Enable Azure Key Vault Secrets Provider
- `enable_secret_rotation` (Boolean) Enable secrets rotation
- `rotation_poll_interval` (String) Secret rotation interval


<a id="nestedblock--spec--config--addon_config--azure_policy_addon_config"></a>
### Nested Schema for `spec.config.addon_config.azure_policy_addon_config`

Optional:

- `enable` (Boolean) Enable policy addon


<a id="nestedblock--spec--config--addon_config--monitor_addon_config"></a>
### Nested Schema for `spec.config.addon_config.monitor_addon_config`

Optional:

- `enable` (Boolean) Enable monitor
- `log_analytics_workspace_id` (String) Log analytics workspace ID for the monitoring addon



<a id="nestedblock--spec--config--api_server_access_config"></a>
### Nested Schema for `spec.config.api_server_access_config`

Required:

- `enable_private_cluster` (Boolean) Enable Private Cluster

Optional:

- `authorized_ip_ranges` (List of String) IP ranges authorized to access the Kubernetes API server


<a id="nestedblock--spec--config--auto_upgrade_config"></a>
### Nested Schema for `spec.config.auto_upgrade_config`

Optional:

- `upgrade_channel` (String) Upgrade Channel


<a id="nestedblock--spec--config--linux_config"></a>
### Nested Schema for `spec.config.linux_config`

Required:

- `admin_username` (String) Administrator username to use for Linux VMs

Optional:

- `ssh_keys` (List of String) Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers


<a id="nestedblock--spec--config--sku"></a>
### Nested Schema for `spec.config.sku`

Optional:

- `name` (String) Name of the cluster SKU
- `tier` (String) Tier of the cluster SKU


<a id="nestedblock--spec--config--storage_config"></a>
### Nested Schema for `spec.config.storage_config`

Optional:

- `enable_disk_csi_driver` (Boolean) Enable the azure disk CSI driver for the storage
- `enable_file_csi_driver` (Boolean) Enable the azure file CSI driver for the storage
- `enable_snapshot_controller` (Boolean) Enable the snapshot controller for the storage



<a id="nestedblock--spec--nodepool"></a>
### Nested Schema for `spec.nodepool`

Required:

- `name` (String) Name of the nodepool, immutable
- `spec` (Block List, Min: 1, Max: 1) Spec for the nodepool (see [below for nested schema](#nestedblock--spec--nodepool--spec))

<a id="nestedblock--spec--nodepool--spec"></a>
### Nested Schema for `spec.nodepool.spec`

Required:

- `count` (Number) Count is the number of nodes
- `mode` (String) The mode of the nodepool SYSTEM or USER. A cluster must have at least one 'SYSTEM' nodepool at all times.
- `vm_size` (String) Virtual Machine Size

Optional:

- `auto_scaling_config` (Block List, Max: 1) Auto scaling config. (see [below for nested schema](#nestedblock--spec--nodepool--spec--auto_scaling_config))
- `availability_zones` (List of String) The list of Availability zones to use for nodepool. This can only be specified if the type of the nodepool is AvailabilitySet.
- `enable_node_public_ip` (Boolean) Whether each node is allocated its own public IP
- `max_pods` (Number) The maximum number of pods that can run on a node
- `node_labels` (Map of String) The node labels to be persisted across all nodes in nodepool
- `os_disk_size_gb` (Number) OS Disk Size in GB to be used to specify the disk size for every machine in the nodepool. If you specify 0, it will apply the default osDisk size according to the vmSize specified
- `os_disk_type` (String) OS Disk Type
- `os_type` (String) The OS type of the nodepool
- `scale_set_eviction_policy` (String) Scale set eviction policy
- `scale_set_priority` (String) Scale set priority
- `spot_max_price` (Number) Max spot price
- `tags` (Map of String) AKS specific node tags
- `taints` (Block List) The taints added to new nodes during nodepool create and scale (see [below for nested schema](#nestedblock--spec--nodepool--spec--taints))
- `type` (String) Nodepool type
- `upgrade_config` (Block List, Max: 1) upgrade config (see [below for nested schema](#nestedblock--spec--nodepool--spec--upgrade_config))
- `vnet_subnet_id` (String) If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes

<a id="nestedblock--spec--nodepool--spec--auto_scaling_config"></a>
### Nested Schema for `spec.nodepool.spec.auto_scaling_config`

Optional:

- `enable` (Boolean) Enable auto scaling
- `max_count` (Number) Maximum node count
- `min_count` (Number) Minimum node count


<a id="nestedblock--spec--nodepool--spec--taints"></a>
### Nested Schema for `spec.nodepool.spec.taints`

Optional:

- `effect` (String) Current effect state of the node pool
- `key` (String) The taint key to be applied to a node
- `value` (String) The taint value corresponding to the taint key


<a id="nestedblock--spec--nodepool--spec--upgrade_config"></a>
### Nested Schema for `spec.nodepool.spec.upgrade_config`

Optional:

- `max_surge` (String) Max Surge





<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource