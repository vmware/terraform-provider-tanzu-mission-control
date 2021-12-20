---
Title: "Provisioning Nodepool Resource"
Description: |-
    Adding the nodepool resource to a cluster.
---

# Provisioning nodepool on a cluster

A nodepool can be added to a workload cluster. For provisioning of a cluster, refer to the `tmc_cluster` in `guides`.

We support Create, read, update and delete operations for node pool.
You could create/manage a `nodepool` for a cluster with the following config:

```terraform
// TMC Cluster Type: Attach. Bring your own k8s cluster and attach it to TMC.
// Operations supported : Read, Create, Update & Delete

# Read TMC cluster nodepool : fetch cluster nodepool details
data "tmc_cluster_node_pool" "read_node_pool" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  cluster_name            = "<cluster_name>"       // Required
  name                    = "<nodepool-namcde>"      // Required
}

# Create TMC nodepool entry
resource "tmc_cluster_node_pool" "create_node_pool" {

  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  cluster_name            = "<cluster_name>"       // Required
  name                    = "<nodepool-name>"      // Required

  spec {
    worker_node_count = "3"

    tkg_service_vsphere {
      class         = "<class>"        // Required
      storage_class = "<storage-class" // Required
    }
  }
}
```