data "tanzu-mission-control_backup_schedule" "demo" {
  name = "BACKUP_SCHEDULE_NAME"
  scope {
    cluster_group {
      cluster_group_name = "CG_NAME"
    }
  }

  query         = "QUERY"
  sort_by       = "SORT_BY"
  include_total = true
}
