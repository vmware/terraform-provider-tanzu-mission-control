data "tanzu-mission-control_target_location" "demo_cluster" {
  scope {
    cluster {
      management_cluster_name = "MGMT_CLS_NAME"
      provisioner_name        = "PROVISIONER_NAME"
      cluster_name            = "CLS_NAME"
      name                    = "TARGET_LOCATION_NAME"
    }
  }

  query         = "QUERY"
  sort_by       = "SORT_BY"
  include_total = true
}
