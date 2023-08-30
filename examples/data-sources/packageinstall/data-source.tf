# Read Tanzu Mission Control package install with attached set as default value.
data "tanzu-mission-control_package_install" "read_package_install" {
    name = "test-pakage-install-name" # Required

    namespace = "test-namespace-name" # Required

    scope {
        cluster {
        cluster_name            = "testcluster" # Required
        provisioner_name        = "attached"    # Default: attached
        management_cluster_name = "attached"    # Default: attached
        }
    }
}