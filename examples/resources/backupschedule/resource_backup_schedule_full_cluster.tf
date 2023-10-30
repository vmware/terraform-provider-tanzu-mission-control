resource "tanzu-mission-control_backup_schedule" "sample-full" {
  name                    = "full-weekly"
  management_cluster_name = "MGMT_CLS_NAME"
  provisioner_name        = "PROVISIONER_NAME"
  cluster_name            = "CLS_NAME"

  scope = "FULL_CLUSTER"

  spec {
    schedule {
      rate = "0 12 * * 1"
    }

    template {
      backup_ttl          = "2592000s"
      excluded_namespaces = [
        "app-01",
        "app-02",
        "app-03",
        "app-04"
      ]
      excluded_resources = [
        "secrets",
        "configmaps"
      ]

      storage_location = "TARGET_LOCATION_NAME"
    }
  }
}

