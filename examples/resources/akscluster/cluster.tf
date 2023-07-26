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
