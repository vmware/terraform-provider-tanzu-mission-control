/*
  NOTE: Creation of attach cluster depends on cluster-group creation
        Similarly, namespace creation depends on attach cluster and workspace creation
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/dev/tanzu-mission-control"
    }
  }
}

resource "tanzu-mission-control_enable_data_protection" "demo" {
  scope {
    cluster {
      cluster_name            = "aks-raven-20231116-qwax"
      management_cluster_name = "aks"
      provisioner_name        = "aks"
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

