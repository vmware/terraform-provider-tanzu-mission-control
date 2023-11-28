resource "tanzu-mission-control_target_location" "demo_tmc_managed" {
  name = "TARGET_LOCATION_NAME"

  spec {
    target_provider = "TARGET_PROVIDER_NAME"
    credential = {
      name = "CREDENTIAL_NAME"
    }

    assigned_groups {
      cluster {
        management_cluster_name = "MGMT_CLS_NAME"
        provisioner_name        = "PROVISIONER_NAME"
        name                    = "CLS_NAME"
      }

      cluster {
        management_cluster_name = "MGMT_CLS_NAME"
        provisioner_name        = "PROVISIONER_NAME"
        name                    = "CLS_NAME"
      }

      cluster_groups = ["CLS_GROUP_NAME", "CLS_GROUP_NAME"]
    }
  }
}