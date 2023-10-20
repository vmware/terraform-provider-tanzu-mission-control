---
Title: "Provisioning of a Azure AKS cluster"
Description: |-
    An example of provisioning Azure AKS clusters.
---
# AKS Cluster

The `tanzu-mission-control_akscluster` resource can directly perform cluster lifecycle management operations on AKS clusters
(and associated node groups) including create, update, upgrade, and delete through Tanzu Mission Control.

## Prerequisites

To manage the lifecycle of AKS clusters, you need the following prerequisites.

- Set up a credential that allows VMware Tanzu Mission Control to connect to your Azure subscription and manage resources in your Azure account. Please refer [connecting an Azure account for AKS cluster lifecycle management][azure-account] guide for detailed steps.
You can also use `tanzu-mission-control_credential` Terraform resource for this purpose. The name of the Azure AKS credential in Tanzu Mission Control will be referred to as `credential_name` in this guide.

- Create a Service Principal with Contributor role on each Azure subscription it has access to. Select either Azure CLI or Azure Portal UI and follow the instructions for the selected method. Please refer to Tanzu documentation on how to [create a Service Principal for AKS cluster lifecycle management][tanzu-aks-credential] .

- Ensure the CSP token used in initialising the terraform provider has the right set of permissions to create a workload cluster.

[tanzu-aks-credential]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-90ED8C73-8A40-46FF-85AE-A8DAA9048AA9.html?hWord=N4IghgNiBcIMoFEBKA1AkgYQQAgApLQDkM1cBBAGRAF8g
[azure-account]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-2CA6A21A-1D33-4852-B8F2-86BB3A1337E4.html

## Provisioning the cluster

You can use the following template as reference to write your own `tanzu-mission-control_akscluster` resource in the terraform scripts.

```terraform
// Tanzu Mission Control AKS Cluster Type: Azure AKS clusters.
// Operations supported : Read, Create, Update & Delete

// Read Tanzu Mission Control Azure AKS cluster : fetch cluster details
data "tanzu-mission-control_akscluster" "tf_aks_cluster" {
  credential_name = "<aws-credential-name>" // Required
  subscription    = "<subscription>"        // Required
  resource_group  = "<resource-group>"      // Required
  name            = "<cluster-name>"        // Required
}

resource "tanzu-mission-control_akscluster" "tf_aks_cluster" {
  credential_name = "<credential-name>" // Required
  subscription    = "<subscription>"    // Required
  resource_group  = "<resource-group>"  // Required
  name            = "<cluster-name>"    // Required

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "<cluster-group>" // Default: default
    proxy         = "<proxy>"

    azure_AKS {
      location                 = "<location>"                 // Required     // Force Recreate
      version                  = "<version>"                  // Required
      node_resource_group_name = "<node-resource-group-name>" // Force Recreate
      disk_encryption_set      = "<disk-encryption-set>"      // Force Recreate
      tags                     = { "key" : "value" }

      sku {
        name = "<name>"
        tier = "<tier>" // Required
      }

      access_config {
        enable_rbac            = true
        disable_local_accounts = true
        aad_config {
          managed  = true
          tenantId = "<tenant-id>"
          admin_group_ids = [
            "<admin-group-id-1>",
            "<admin-group-id-2>",
          ]
          enable_azure_rbac = true
        }
      }

      api_server_access_config {
        authorized_ip_ranges = [
          "<ip-range-1>",
          "<ip-range-2>",
        ]
        enable_private_cluster = true // Forces Recreate
      }

      linux_config {
        // Force Recreate
        admin_username = "<admin-username>"
        ssh_keys = [
          "<ssh-key-1>",
          "<ssh-key-2>",
        ]
      }

      network_config {
        // Required
        load_balancer_sku   = "<load-balancer-sku>"   // Forces Recreate
        network_plugin      = "<network-plugin>"      // Forces Recreate
        network_plugin_mode = "<network-plugin-mode>" // Forces Recreate
        network_policy      = "<network-policy>"      // Forces Recreate
        dns_service_ip      = "<dns-service-ip>"      // Forces Recreate
        docker_bridge_cidr  = "<docker-bridge-cidr>"  // Forces Recreate
        pod_cidr = [
          // Forces Recreate
          "<pod-cidr-1>",
          "<pod-cidr-2>",
        ]
        service_cidr = [
          // Forces Recreate
          "<service-cidr-1>",
          "<service-cidr-2>",
        ]
        dns_prefix                      = "<dns-prefix>" // Required
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
            rotation_poll_interval = "<rotation_poll_interval>"
          }
        }

        monitor_addon_config {
          enable                     = true
          log_analytics_workspace_id = "<log-analytics-workspace-id>"
        }

        azure_policy_addon_config {
          enable = true
        }
      }

      auto_upgrade_config {
        upgrade_channel = "<upgrade_channel>"
      }

      nodepools = [
        {
          info = {
            name = "<nodepool-name>"
          }

          spec = {
            mode = "<mode>" // Required
            type = "<type>"
            availabilityZones = [
              "<availability-zone-1>",
              "<availability-zone-2>",
            ]
            count                     = 1                             // Required
            vm_size                   = "<vm-size>"                   // Required // Force Recreate
            scale_set_priority        = "<scale-set-priority>"        // Force Recreate
            scale_set_eviction_policy = "<scale-set-eviction-policy>" // Force Recreate
            spot_max_price            = 1.00
            os_type                   = "<os-type>"
            os_disk_type              = "<os-disk-type>" // Force Recreate
            os_disk_size_gb           = 60               // Force Recreate
            max_pods                  = 10               // Force Recreate
            enable_node_public_ip     = true
            node_taints = [
              {
                effect = "<effect>"
                key    = "<key>"
                value  = "<value>"
              }
            ]
            vnet_subnet_id = "<vnet-subnet-id>" // Force Recreate
            pod_subnet_id  = "<pod-subnet-id>"  // Force Recreate
            node_labels    = { "key" : "value" }
            tags           = { "key" : "value" }

            auto_scaling_config = {
              enable    = true // Force Recreate
              min_count = 1
              max_count = 5
            }

            upgrade_config = {
              max_surge = "<max-surge>"
            }
          }
        }
      ]
    }
  }
}
```