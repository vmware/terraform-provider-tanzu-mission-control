# Create Tanzu Mission Control attach cluster with k8s cluster kubeconfig provided
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