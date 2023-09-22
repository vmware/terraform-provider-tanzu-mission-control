# Create Tanzu Mission Control cluster scope helm feature with attached set as default value.
resource "tanzu-mission-control_helm_feature" "create_cl_helm_feature" {
  scope {
    cluster {
      name                    = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }
}