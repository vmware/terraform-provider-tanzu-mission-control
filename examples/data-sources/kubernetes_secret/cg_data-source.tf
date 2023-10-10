# Read Tanzu Mission Control kubernetes secret : fetch namespace details
data "tanzu-mission-control_kubernetes_secret" "read_secret" {
  name           = "tf-secret"                # Required
  namespace_name = "tf-secret-namespace-name" # Required 

  scope {
    cluster_group {
      name = "default" # Required
    }
  }
}