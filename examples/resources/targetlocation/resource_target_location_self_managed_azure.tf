resource "tanzu-mission-control_target_location" "demo_azure_self_provisioned" {
  name          = "TARGET_LOCATION_NAME"

  spec {
    target_provider = "AZURE"
    credential      = {
      name = "AZURE_CREDENTIAL_NAME"
    }

    bucket = "BUCKET_NAME"

    config {
      azure {
        resource_group  = "AZURE_RESOURCE_GROUP_NAME"
        storage_account = "AZURE_STORAGE_ACCOUNT_NAME"
        subscription_id = "AZURE_SUBSCRIPTION_ID_NAME"
      }
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