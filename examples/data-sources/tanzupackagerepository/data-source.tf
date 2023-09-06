# Read Tanzu Mission Control package repository : fetch cluster package repository details
data "tanzu-mission-control_package_repository" "read_cluster_pkg_repository" {
  name = "tf-pkg-repository-name" # Required

  scope {
    cluster {
      name            = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }
}