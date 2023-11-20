# Create self provisioned storage Azure blob for Data protection.
resource "tanzu-mission-control_credential" "azure_ad_cred" {
  name = "tf-azure-ad-self-dp-test"

  meta {
    description = "Self provisioned storage Azure blob for Data protection"
    labels = {
      "key1" : "value1",
    }
  }

  spec {
    capability = "DATA_PROTECTION"
    provider   = "AZURE_AD"
    data {
      azure_credential {
        service_principal {
          subscription_id  = "some_subscription_id"
          tenant_id        = "some_tenant_id"
          resource_group   = "dp-backup-rg"
          client_id        = "some_client_id"
          client_secret    = "some_client_id"
          azure_cloud_name = "AzurePublicCloud"
        }
      }
    }
  }
  ready_wait_timeout = "5m" // Wait time for credential create operations to finish (default: 3m).
}
