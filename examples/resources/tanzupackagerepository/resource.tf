# Create Tanzu Mission Control package repository with attached set as default value.
resource "tanzu-mission-control_package_repository" "create_cluster_pkg_repository" {
  name = "tf-pkg-repository-name" # Required

  scope {
    cluster {
      name            = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    imgpkg_bundle {
        image = "testImage" # Required
    }
  }
}