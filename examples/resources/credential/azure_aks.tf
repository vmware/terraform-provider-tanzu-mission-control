# Create AZURE_AKS credential resource
resource "tanzu-mission-control_credential" "azure_aks_cred" {
  name = "tf-azure-aks-cred-test"

  meta {
    description = "Azure AKS credential for AKS cluster lifecycle management"
    labels = {
      "key1" : "value1",
    }
  }

  spec {
    capability = "MANAGED_K8S_PROVIDER"
    provider   = "AZURE_AKS"
    data {
      azure_credential {
        service_principal_with_certificate {
          subscription_id    = "some_subscription_id"
          tenant_id          = "some_tenant_id"
          client_id          = "some_client_id"
          client_certificate = "-----BEGIN PRIVATE KEY-----\nMIvIEFk\nv+FiTAfd5MYtJYjkuU7MVA==\n-----END PRIVATE KEY-----\n-----BEGIN CERTIFICATE-----\nMIICoTCCAYkCAgPoMA0GCSqGSIb3DQEBBQUAMXyNh2KI=\n-----END CERTIFICATE-----\n"
        }
      }
    }
  }
  ready_wait_timeout = "5m" // Wait time for credential create operations to finish (default: 3m).
}
