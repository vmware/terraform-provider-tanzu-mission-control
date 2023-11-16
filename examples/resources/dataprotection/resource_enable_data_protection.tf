resource "tanzu-mission-control_enable_data_protection" "demo" {
  scope {
    cluster {
      cluster_name            = "CLS_NAME"
      management_cluster_name = "MGMT_CLS_NAME"
      provisioner_name        = "PROVISIONER_NAME"
    }
  }

  spec {
    disable_restic                       = false
    enable_csi_snapshots                 = false
    enable_all_api_group_versions_backup = false
  }

  deletion_policy {
    delete_backups = false
  }
}
