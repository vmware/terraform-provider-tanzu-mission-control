resource "tanzu-mission-control_backup_schedule" "sample-full" {
  name                    = "namespaces-hourly"
  scope {
    cluster {
      management_cluster_name = "MGMT_CLS_NAME"
      provisioner_name        = "PROVISIONER_NAME"
      cluster_name            = "CLS_NAME"
    }
  }

  backup_scope                   = "SET_NAMESPACES"

  spec {
    schedule {
      rate = "30 * * * *"
    }

    template {
      included_namespaces = [
        "app-01",
        "app-02",
        "app-03",
        "app-04"
      ]

      excluded_resources = [
        "secrets",
        "configmaps"
      ]

      backup_ttl                = "86400s"
      include_cluster_resources = true
      storage_location          = "TARGET_LOCATION_NAME"

      hooks {
        resource {
          name = "sample-config"

          pre_hook {
            exec {
              command   = ["echo 'hello'"]
              container = "workload"
              on_error  = "CONTINUE"
              timeout   = "10s"
            }
          }

          pre_hook {
            exec {
              command   = ["echo 'hello'"]
              container = "db"
              on_error  = "CONTINUE"
              timeout   = "30s"
            }
          }

          post_hook {
            exec {
              command   = ["echo 'goodbye'"]
              container = "db"
              on_error  = "CONTINUE"
              timeout   = "60s"
            }
          }

          post_hook {
            exec {
              command   = ["echo 'goodbye'"]
              container = "workload"
              on_error  = "FAIL"
              timeout   = "20s"
            }
          }
        }
      }
    }
  }
}

