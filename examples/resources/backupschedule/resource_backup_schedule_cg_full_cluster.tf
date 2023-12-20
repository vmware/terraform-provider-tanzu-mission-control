resource "tanzu-mission-control_backup_schedule" "sample-full" {
  name = "full-weekly"
  scope {
    cluster_group {
      cluster_group_name = "CG_NAME"
    }
  }
  selector {
    names = [
      "cluster1",
      "cluster2"
    ]
  }

  backup_scope = "FULL_CLUSTER"
  spec {
    schedule {
      rate = "0 12 * * 1"
    }

    template {
      backup_ttl = "2592000s"
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

