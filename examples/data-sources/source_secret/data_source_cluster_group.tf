# Read Tanzu Mission Control source secret : fetch cluster group source secret details
data "tanzu-mission-control_repository_credential" "read_cluster_group_source_secret" {
  name = "tf-source_secret" # Required

  scope {
    cluster_group {
      name = "default" # Required
    }
  }
}
