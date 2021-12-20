/*
  NOTE: Creation of attach cluster depends on cluster-group creation
        Similarly, namespace creation depends on attach cluster and workspace creation
*/

terraform {
  required_providers {
    tmc = {
      source = "vmware/tanzu/tmc"
    }
  }
}

# Create TMC TKG Service Vsphere workload cluster entry
resource "tmc_cluster" "create_tkgs_workload" {
  management_cluster_name = "tkgs-terraform"
  provisioner_name        = "test-gc-e2e-demo-ns"
  name                    = "tkgs-workload-test"

  meta {
    labels      = { "key" : "test"}
  }

  spec {
    cluster_group = "default"
    tkg_service_vsphere {
      settings  {
        network  {
          pods  {
            cidr_blocks = [
              "172.20.0.0/16",
            ]
          }
          services  {
            cidr_blocks = [
              "10.96.0.0/16",
            ]
          }
        }
      }

      distribution  {
        version = "v1.21.2+vmware.1-tkg.1.fea8785"
      }

      topology  {
        control_plane  {
          class = "best-effort-xsmall"
          storage_class = "gc-storage-profile"
          high_availability = false
        }
        node_pools  {
          spec  {
            worker_node_count = "1"
            tkg_service_vsphere  {
              class = "best-effort-xsmall"
              storage_class = "gc-storage-profile"
            }
          }
          info {
            name = "default-nodepool"
          }
        }
      }
    }
  }
}

# Create a TKGm Vsphere workload cluster entry
resource "tmc_cluster" "create_tkg_vsphere_cluster" {
  management_cluster_name = "tkgm-terraform"
  provisioner_name        = "default"
  name                    = "tkgm-workload-test"

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value1" }
  }

  spec {
    cluster_group = "default" # Default: default
    tkg_vsphere {
      settings {
        network {
          pods {
            cidr_blocks = [
              "172.20.0.0/16",
            ]
          }

          services {
            cidr_blocks = [
              "10.96.0.0/16",
            ]
          }

          control_plane_end_point = "10.191.143.100"
        }

        security {
          ssh_key = "default"
        }
      }

      distribution {
        version = "v1.20.5+vmware.2-tkg.1"

        workspace {
          datacenter        = "/dc0"
          datastore         = "/dc0/datastore/local-0"
          workspace_network = "/dc0/network/VM Network"
          folder            = "/dc0/vm"
          resource_pool     = "/dc0/host/cluster0/Resources"
        }
      }

      topology {
        control_plane {
          vm_config {
            cpu       = "2"
            disk_size = "20"
            memory    = "4096"
          }

          high_availability = false
        }

        node_pools {
          spec {
            worker_node_count = "1"

            tkg_vsphere {
              vm_config {
                cpu       = "2"
                disk_size = "40"
                memory    = "4096"
              }
            }
          }

          info {
            name        = "default-nodepool"
            description = "my nodepool"
          }
        }
      }
    }
  }
}

# Create TMC nodepool entry
resource "tmc_cluster_node_pool" "create_node_pool" {

  management_cluster_name = "tkgs-terraform"
  provisioner_name = "test-gc-e2e-demo-ns"
  cluster_name = "tkgs-nodepool"
  name = "terraform-nodepool"

  spec {
    worker_node_count = "3"

    tkg_service_vsphere  {
      class = "best-effort-xsmall"
      storage_class = "gc-storage-profile"
    }
  }
}