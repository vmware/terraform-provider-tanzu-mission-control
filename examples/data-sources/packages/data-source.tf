# Read Tanzu Mission Control packages : fetch cluster packages details
data "tanzu-mission-control_packages" "read_cluster_packages" {
  metadata_name  = "package-metadata-name" # Required

  scope {
    cluster {
      name            = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }
}