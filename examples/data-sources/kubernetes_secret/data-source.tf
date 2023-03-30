# Read Tanzu Mission Control kubernetes secret : fetch namespace details
data "tanzu-mission-control_kubernetes_secret" "read_secret" {
  name                    = "tf-secret" # Required
  namespace_name          = "tf-secret-namespace-name" # Required 

  scope {
    cluster {
        cluster_name            = "testcluster"  # Required
        provisioner_name        = "attached"     # Default: attached
        management_cluster_name = "attached"     # Default: attached
    }
  }
}