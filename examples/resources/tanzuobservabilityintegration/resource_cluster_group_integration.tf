resource "tanzu-mission-control_tanzu_observability_integration" "cluster_group_integration" {
  scope {
    cluster_group {
      name = "CLUSTER_GROUP_NAME"
    }
  }

  spec {
    credentials_name = "CREDENTIALS_NAME"
  }
}
