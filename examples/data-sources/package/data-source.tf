# Read Tanzu Mission Control package : fetch cluster package details
data "tanzu-mission-control_package" "get_cluster_package" {
  name           = "test-package-version" # Required

  metadata_name  = "package-metadata-name" # Required

  scope {
    cluster {
      cluster_name            = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }
}