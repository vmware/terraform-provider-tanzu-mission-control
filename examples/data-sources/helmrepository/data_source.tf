# Read Tanzu Mission Control helm repository : fetch helm repository details
data "tanzu-mission-control_helm_repository" "get_cluster_helm_repo" {
  name = "test-helm-repository_name"

  metadata_name = "test_namespace_name"

  scope {
    cluster {
      name                    = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }
}