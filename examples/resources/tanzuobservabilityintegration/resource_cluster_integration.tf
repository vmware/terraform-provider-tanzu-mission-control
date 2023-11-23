resource "tanzu-mission-control_tanzu_observability_integration" "cluster_integration" {
  scope {
    cluster {
      management_cluster_name = "MGMT_CLUSTER_NAME"
      provisioner_name        = "PROVISIONER_NAME"
      name                    = "CLUSTER_NAME"
    }
  }

  spec {
    credentials_name = "CREDENTIALS_NAME"
  }
}
