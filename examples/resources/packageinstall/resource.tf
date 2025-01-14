# Create Tanzu Mission Control package install with attached set as default value.
resource "tanzu-mission-control_package_install" "package_install" {
  name = "test-pakage-install-name" # Required

  namespace = "test-namespace-name" # Required

  scope {
    cluster {
      name                    = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }

  spec {
    package_ref {
      package_metadata_name = "test-package-metadata-name" # Required

      version_selection {
        constraints = "test-version-selection" # Required
      }
    }

    path_to_inline_values = "./inline_values.yaml" #<inline-values-file-path>

    inline_values = { "test" : "test" } # Deprecated
  }
}