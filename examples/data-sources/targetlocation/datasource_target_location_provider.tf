data "tanzu-mission-control_target_location" "demo_provider" {
  scope {
    provider {
      provider_name       = "PROVIDER_NAME"
      credential_name     = "CREDENTIAL_NAME"
      assigned_group_name = "ASSIGNED_GROUP_NAME"
      name                = "TARGET_LOCATION_NAME"
    }
  }

  query         = "QUERY"
  sort_by       = "SORT_BY"
  include_total = true
}
