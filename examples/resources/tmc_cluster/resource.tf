# Create TMC attach cluster entry
resource "tmc_cluster" "attach_cluster_without_apply" {
  management_cluster_name = "attached"         # Default: attached
  provisioner_name        = "attached"         # Default: attached
  name                    = "terraform-attach" # Required

  meta {
    description = "create attach cluster from terraform"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" # Default: default
  }

  wait_until_ready = false
}

# Create TMC attach cluster with k8s cluster kubeconfig provided
# The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tmc_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     # Default: attached
  provisioner_name        = "attached"     # Default: attached
  name                    = "demo-cluster" # Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config path>" # Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" # Default: default
  }

  wait_until_ready = true # Default: false, when set resource waits until 3 min for the cluster to become ready

  # The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}

# Create TMC TKG Service Vsphere workload cluster entry
resource "tmc_cluster" "create_tkgs_workload" {
  management_cluster_name = "test-tkgs"
  provisioner_name        = "test-gc-e2e-demo-ns"
  name                    = "cluster"

  meta {
    labels = { "key" : "test" }
  }

  spec {
    cluster_group = "default"
    tkg_service_vsphere {
      settings {
        network {
          pods {
            cidr_blocks = [
              "172.20.0.0/16", # pods cidr block by default has the value `172.20.0.0/16`
            ]
          }
          services {
            cidr_blocks = [
              "10.96.0.0/16", # services cidr block by default has the value `10.96.0.0/16`
            ]
          }
        }
      }

      distribution {
        version = "v1.20.8+vmware.1-tkg.2"
      }

      topology {
        control_plane {
          class             = "best-effort-xsmall"
          storage_class     = "wcpglobal-storage-profile"
          # storage class is either `wcpglobal-storage-profile` or `gc-storage-profile`
          high_availability = false
        }
        node_pools {
          spec {
            worker_node_count = "1"
            tkg_service_vsphere {
              class         = "best-effort-xsmall"
              storage_class = "wcpglobal-storage-profile"
              # storage class is either `wcpglobal-storage-profile` or `gc-storage-profile`
            }
          }
          info {
            name = "default-nodepool" # default node pool name `default-nodepool`
          }
        }
      }
    }
  }
}

# Create a TKGm Vsphere workload cluster entry
resource "tmc_cluster" "create_tkg_vsphere_cluster" {
  management_cluster_name = "tkg-vsphere-test"
  provisioner_name        = "default"
  name                    = "demo-tkg-vsphere-cluster"

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
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

          control_plane_end_point = "10.185.107.47"
        }

        security {
          ssh_key = "default"
        }
      }

      distribution {
        version = "v1.20.4+vmware.1-tkg.1"

        workspace {
          datacenter        = "/dc0"
          datastore         = "/dc0/datastore/local-0"
          workspace_network = "/dc0/network/Avi Internal"
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
          node_pool_spec {
            worker_node_count = "1"

            tkg_vsphere {
              vm_config {
                cpu       = "2"
                disk_size = "40"
                memory    = "8192"
              }
            }
          }

          node_pool_info {
            node_pool_name        = "default-nodepool"
            node_pool_description = "my nodepool"
          }
        }
      }
    }
  }
}
