---
Title: "Provisioning Node pool Resource"
Description: |-
    Adding the node pool resource to a cluster.
---

# Provisioning node pool on a cluster

For clusters that you create in Tanzu Mission Control, you can define a pool of worker nodes on which your workloads can run.
By default, each workload cluster that you create through Tanzu Mission Control has a node pool.
The `tanzu-mission-control_cluster_node_pool` resource allows you to create additional node pools, as well as read, update, and delete existing node pools in your clusters.
Because Tanzu Mission Control cannot provision additional resources in a cluster that is created elsewhere and subsequently attached, you cannot create a node pool in an attached cluster.

For provisioning of a cluster, refer to the `tmc_cluster` in guides.

You could create/manage a `node pool` for a cluster with the following config:

```terraform
// Tanzu Mission Control Nodepool
// Operations supported : Read, Create, Update & Delete

# Read Tanzu Mission Control cluster nodepool : fetch cluster nodepool details
data "tanzu-mission-control_cluster_node_pool" "read_node_pool" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  cluster_name            = "<cluster_name>"       // Required
  name                    = "<node_pool-name>"     // Required
}

# Create Tanzu Mission Control cluster nodepool entry
resource "tanzu-mission-control_cluster_node_pool" "create_node_pool" {

  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  cluster_name            = "<cluster_name>"       // Required
  name                    = "<node_pool-name>"     // Required

  spec {
    worker_node_count = "<worker-nodes>"
    cloud_labels = {
      "<key>" : "<val>"
    }
    node_labels = {
      "<key>" : "<val>"
    }

    tkg_service_vsphere {
      class         = "<class>"        // Required
      storage_class = "<storage-class" // Required
    }
  }
}
```