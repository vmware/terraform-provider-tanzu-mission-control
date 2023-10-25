resource "tanzu-mission-control_backup_schedule" "sample-full" {
  name                    = "label-based-no-storage"
  scope {
    cluster {
      management_cluster_name = "MGMT_CLS_NAME"
      provisioner_name        = "PROVISIONER_NAME"
      cluster_name            = "CLS_NAME"
    }
  }

  backup_scope                   = "LABEL_SELECTOR"


  spec {
    schedule {
      rate = "0 12 * * *"
    }

    template {
      default_volumes_to_fs_backup = false
      include_cluster_resources    = true
      backup_ttl                   = "604800s"
      storage_location             = "TARGET_LOCATION_NAME"

      label_selector {
        match_expression {
          key      = "apps.tanzu.vmware.com/demo"
          operator = "Exists"
        }

        match_expression {
          key      = "apps.tanzu.vmware.com/exclude-from-backup"
          operator = "DoesNotExist"
        }
      }
    }
  }
}

