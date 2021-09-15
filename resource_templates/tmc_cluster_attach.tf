// TMC Cluster Type: Attach. Bring your own k8s cluster and attach it to TMC.
// Operations supported : Read, Create, Update & Delete

// Read TMC cluster : fetch cluster details
data "tmc_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  name                    = "<cluster-name>"       // Required
}

// Create TMC attach cluster entry
resource "tmc_cluster" "attach_cluster" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  name                    = "<cluster-name>"       // Required

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "<cluster-group>" // Default: default
  }

  wait_until_ready = false // Default: false, when set resource waits until 3 min for the cluster to become ready

  // The deployment link and the command needed to be run to attach this cluster would be provided in the output. status.execution_cmd
}

// Create TMC attach cluster with k8s cluster kubeconfig provided
// The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tmc_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  name                    = "<cluster-name>"       // Required

  attach_k8s_cluster {
    kubeconfig_file = "<k8s kubeconfig file path>" // Required
    description     = "optional description about the kubeconfig provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "<cluster-group>" // Default: default
  }

  wait_until_ready = true // Default: false, when set resource waits until 3 min for the cluster to become ready

  // The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}

// Create TMC attach cluster entry with minimal information
resource "tmc_cluster" "attach_cluster_minimal" {
  name = "<cluster-name>" // Required
}