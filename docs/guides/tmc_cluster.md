---
Title: "Provisioning of a workload cluster"
Description: |-
    An example of provisioning a TKGs and a TKGm workload cluster.
---

# TKGs Workload Cluster

The TKGs workload cluster can be provisioned through the terraform provider using the following example. For the
provisioning, it is expected that the user already has a management cluster registered of the kind `vSphere with Tanzu`
on their TMC instance. The following example demonstrates the resource of a TKG Service Vsphere workload cluster:

```terraform
// TMC Cluster Type: TKGServiceVsphere workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read TMC TKG Service Vsphere workload cluster : fetch cluster details
data "tmc_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<management-cluster>" // Required
  provisioner_name        = "<prov-name>"          // Required
  name                    = "<cluster-name>"       // Required
}

# Create TMC TKG Service Vsphere workload cluster entry
resource "tmc_cluster" "create_tkgs_workload" {
  management_cluster_name = "<management-cluster>"
  provisioner_name        = "<prov-name>"
  name                    = "<cluster-name>"

  meta {
    labels = { "key" : "test" }
  }

  spec {
    cluster_group = "<cluster-group>" // Default: default
    tkg_service_vsphere {
      settings {
        network {
          pods {
            cidr_blocks = [
              "<pods-cidr-blocks>", // Required
            ]
          }
          services {
            cidr_blocks = [
              "<services-cidr-blocks>", // Required
            ]
          }
        }
      }

      distribution {
        version = "<version>" // Required
      }

      topology {
        control_plane {
          class             = "<class>"        // Required
          storage_class     = "<storage-class" // Required
          high_availability = false            // Default: false
        }
        node_pools {
          spec {
            worker_node_count = "<worker-node-count>" // Required
            cloud_label = {
              "<key>" : "<val>"
            }
            node_label = {
              "<key>" : "<val>"
            }
            tkg_service_vsphere {
              class         = "<class>"         // Required
              storage_class = "<storage-class>" // Required
            }
          }
          info {
            name = "<node-pool-name>" // Required
          }
        }
      }
    }
  }
}
```

# TKGm Vsphere Workload Cluster

The TKGm Vsphere workload cluster can be provisioned through the terraform provider using the following example. For the
provisioning, it is expected that the user already has a management cluster registered of the kind `Tanzu Kubernetes Grid`
on their TMC instance. The following example demonstrates the resource of a TKGm Vsphere workload cluster:

```terraform
// TMC Cluster Type: TKGVsphere workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read TMC TKG Vsphere workload cluster : fetch cluster details
data "tmc_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<management-cluster>" // Required
  provisioner_name        = "<prov-name>"          // Required
  name                    = "<cluster-name>"       // Required
}

// Create TMC TKG Vsphere workload cluster entry
resource "tmc_cluster" "create_tkg_vsphere_cluster" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  name                    = "<cluster-name>"       // Required

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "<cluster-group>" // Default: default
    tkg_vsphere {
      settings {
        network {
          pods {
            cidr_blocks = [
              "<pods-cidr-blocks>", // Required
            ]
          }

          services {
            cidr_blocks = [
              "<services-cidr-blocks>", // Required
            ]
          }

          control_plane_end_point = "<end-point>" // Required
        }

        security {
          ssh_key = "<ssh-key>" // Required
        }
      }

      distribution {
        version = "<version>" // Required

        workspace {
          datacenter        = "<datacenter>"        // Required
          datastore         = "<datastore>"         // Required
          workspace_network = "<workspace_network>" // Required
          folder            = "<folder>"            // Required
          resource_pool     = "<resource_pool>"     // Required
        }
      }

      topology {
        control_plane {
          vm_config {
            cpu       = "<cpu>"       // Required
            disk_size = "<disk_size>" // Required
            memory    = "<memory>"    // Required
          }

          high_availability = false // Default: false
        }

        node_pools {
          spec {
            worker_node_count = "<worker-node-count>" // Required
            cloud_label = {
              "<key>" : "<val>"
            }
            node_label = {
              "<key>" : "<val>"
            }

            tkg_vsphere {
              vm_config {
                cpu       = "<cpu>"       // Required
                disk_size = "<disk_size>" // Required
                memory    = "<memory>"    // Required
              }
            }
          }

          info {
            name        = "<node-pool-name>" // Required
            description = "<node-pool-description>"
          }
        }
      }
    }
  }
}
```

When you apply these configurations it will create the respective workload cluster on the chosen management cluster.
If you need to update the cluster, you simply make an update to the rule definition and Terraform will
apply/update it across all the sites. If you add / or remove a site from the list, Terraform will also
handle creating or removing the rule on the subsequent `terraform apply`.