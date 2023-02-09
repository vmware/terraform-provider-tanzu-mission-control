// Tanzu Mission Control Cluster Type: Attach. Bring your own k8s cluster and attach it to Tanzu Mission Control.
// Operations supported : Read, Create, Update & Delete

// Read Tanzu Mission Control cluster : fetch cluster details
data "tanzu_mission_control_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  name                    = "<cluster-name>"       // Required
}

// Create Tanzu Mission Control attach cluster entry
resource "tanzu_mission_control_cluster" "attach_cluster" {
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

  ready_wait_timeout = "15m" // Default: waits until 3 min for the cluster to become ready

  // The deployment link and the command needed to be run to attach this cluster would be provided in the output. status.execution_cmd
}

// Create Tanzu Mission Control attach cluster with k8s cluster kubeconfig provided
// The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tanzu_mission_control_cluster" "attach_cluster_with_kubeconfig" {
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

  ready_wait_timeout = "3m" // Default: waits until 3 min for the cluster to become ready
}

// Create Tanzu Mission Control attach cluster entry with minimal information
resource "tanzu_mission_control_cluster" "attach_cluster_minimal" {
  name = "<cluster-name>" // Required
}
