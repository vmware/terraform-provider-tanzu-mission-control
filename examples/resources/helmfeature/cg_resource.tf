# Create Tanzu Mission Control cluster group scope helm feature with attached set as default value.
resource "tanzu-mission-control_helm_feature" "cg_helm_feature" {
  scope {
    cluster_group {
      name = "default" # Required
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }
}