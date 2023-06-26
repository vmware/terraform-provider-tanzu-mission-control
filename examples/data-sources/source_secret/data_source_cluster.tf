# Read Tanzu Mission Control source secret : fetch cluster source secret details
data "tanzu-mission-control_repository_credential" "read_cluster_source_secret" {
  name = "tf-source_secret" # Required

  scope {
    cluster {
      cluster_name            = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }
}
