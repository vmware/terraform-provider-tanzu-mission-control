// Tanzu Mission Control AKS Cluster Type: Azure AKS clusters.
// Operations supported : Read, Create, Update & Delete

// Read Tanzu Mission Control Azure AKS cluster : fetch cluster details
data "tanzu-mission-control_akscluster" "tf_aks_cluster" {
  credential_name = "<credential-name>" // Required
  subscription    = "<subscription>"    // Required
  resource_group = "<resource-group>"  // Required
  name            = "<cluster-name>"    // Required
}

// Create Tanzu Mission Control Azure AKS workload cluster entry
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
      location                 = "<location>" // Required     // Force Recreate
      version                  = "<version>"  // Required
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
          managed         = true
          tenantId        = "<tenant-id>"
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
        ssh_keys       = [
          "<ssh-key-1>",
          "<ssh-key-2>",
        ]
      }

      network_config {
        // Required
        load_balancer_sku  = "<load-balancer-sku>"  // Forces Recreate
        network_plugin     = "<network-plugin>"     // Forces Recreate
        network_policy     = "<network-policy>"     // Forces Recreate
        dns_service_ip     = "<dns-service-ip>"     // Forces Recreate
        docker_bridge_cidr = "<docker-bridge-cidr>" // Forces Recreate
        pod_cidr           = [
          // Forces Recreate
          "<pod-cidr-1>",
          "<pod-cidr-2>",
        ]
        service_cidr = [
          // Forces Recreate
          "<service-cidr-1>",
          "<service-cidr-2>",
        ]
        dns_prefix                      = "<dns-prefix>" // Required, Forces Recreate
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

      nodepools {
        name = "<node-pool-name>"   // Required // Forces Recreate

        spec {
          mode              = "<mode>" // Required
          type              = "<type>"
          availabilityZones = [
            "<availability-zone-1>",
            "<availability-zone-2>",
          ]
          count                 = 1 // Required
          vm_size               = "<vm-size>" // Required // Force Recreate
          os_type               = "<os-type>"
          os_disk_type          = "<os-disk-type>"        // Force Recreate
          os_disk_size_gb       = 60                      // Force Recreate
          max_pods              = 10                      // Force Recreate
          enable_node_public_ip = true
          node_taints           = [
            {
              effect = "<effect>"
              key    = "<key>"
              value  = "<value>"
            }
          ]
          vnet_subnet_id = "<vnet-subnet-id>" // Force Recreate
          node_labels    = { "key" : "value" }
          tags           = { "key" : "value" }

          auto_scaling_config {
            enable                    = true // Force Recreate
            min_count                 = 1
            max_count                 = 5
            scale_set_priority        = "<scale-set-priority>" // Force Recreate
            scale_set_eviction_policy = "<scale-set-eviction-policy>"
            spot_max_price            = 1.00
          }

          upgrade_config {
            max_surge = "<max-surge>"
          }
        }
      }
    }
  }
}